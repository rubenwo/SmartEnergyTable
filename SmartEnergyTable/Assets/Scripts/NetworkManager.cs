﻿using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Grpc.Core;
using Network;
using UnityEngine;
using UnityEngine.SceneManagement;

public sealed class NetworkManager : MonoBehaviour
{
    //This is the server address for the backend.
    public string serverAddr;

    //The objectLibrary is a list of prefabs. As we're not using many different tokens it's more efficient to store them in an array.
    //The library is filled in the editor. Rather than loading them from Resources when we need them.
    public List<GameObject> objectLibrary = new List<GameObject>();

    //To make the life of developers easier the _prefabLookup returns the index for the objectLibrary based on the name of the prefab.
    private readonly Dictionary<string, int> _prefabLookUp = new Dictionary<string, int>();

    //Static instance so there is only 1 NetworkManager
    private static NetworkManager _instance;
    private Channel _channel;
    private Client _client;
    private readonly Queue<Action> _actionQueue = new Queue<Action>();


    private readonly string _userId = Guid.NewGuid().ToString(); //Generate a new GUID as userId.
    private string _roomId = "";
    private bool _master;

    //This is a representation of the current scene according to the server. This manages all the tokens.
    //Key (string) is the uuid generated by the server on AddToken. Value (GameObject) is the token.
    private readonly Dictionary<string, GameObject> _currentScene = new Dictionary<string, GameObject>();

    private void Awake()
    {
        if (_instance == null)
        {
            _instance = this;
            //We want the network manager to exist in every scene, so we need to call DontDestroyOnLoad on this gameObject.
            DontDestroyOnLoad(gameObject);

            //Fill the prefab
            for (var i = 0; i < objectLibrary.Count; i++)
            {
                _prefabLookUp[objectLibrary[i].name] = i;
            }

            //Create the channel and client
            _channel = new Channel(serverAddr, ChannelCredentials.Insecure);
            _client = new Client(new SmartEnergyTableService.SmartEnergyTableServiceClient(_channel));
        }
        else
        {
            Destroy(gameObject);
        }
    }

    private void OnApplicationPause(bool pauseStatus)
    {
        if (Application.platform == RuntimePlatform.Android && pauseStatus)
            OnApplicationQuit();
    }

    private void OnApplicationQuit()
    {
        Debug.Log("OnApplicationQuit()");
        //Leave the room before we shutdown.
        _client.LeaveRoom(_roomId, _userId);
        //Shutdown the channel synchronously to avoid bugs.
        _channel.ShutdownAsync().Wait();
    }


    // Update is called once per frame and runs on the main thread.
    private void Update()
    {
        //The update from the JoinRoom adds actions to the queue. We perform these actions in Update() to run them on the main thread.
        //Unity doesn't support multithreaded calls to the engine like 'Instantiate'.
        while (_actionQueue.Count > 0)
            _actionQueue.Dequeue()();
    }


    //IsMaster can be used in UI elements to show a "master" only view.
    public bool IsMaster => _master;

    #region RPCs

    /*
     * CreateRoom is an abstraction over the RPC. This function creates a room, sets the roomId and master.
     * Then joins the newly created room.
     */
    public void CreateRoom()
    {
        if (_roomId != "")
            return;
        var room = _client.CreateRoom();
        _roomId = room.Id;
        _master = true;
        JoinRoom(_roomId);
    }

    /*
     * JoinRoom is an abstraction over the RPC for easy use.
     * This function starts a long-running task that enqueues actions on every update from the server.
     */
    public void JoinRoom(string id)
    {
        //If we're not the master, the roomId won't be set yet.
        if (_roomId == "")
            _roomId = id;
        Task.Run(() => _client.JoinRoom(id, _userId, patch =>
        {
            _actionQueue.Enqueue(() =>
            {
                _master = patch.IsMaster; // This is true if the master switched.

                //Load the scene if it is not the currentScene, meaning the scene has changed.
                if (patch.SceneId != SceneManager.GetActiveScene().buildIndex)
                    SceneManager.LoadScene(patch.SceneId);

                if (_currentScene.Count == 0)
                {
                    Debug.Log("Destroy/Instantiatie everything");
                    //TODO: create a comparator algorithm to avoid destroying the scene. While functional at 'normal' use, this can be broken rather quickly.
                    //Destroy all tokens from the scene, then clear the currentScene as we get the entire scene from the server.
                    foreach (var keyValuePair in _currentScene)
                    {
                        Destroy(keyValuePair.Value);
                    }

                    _currentScene.Clear();


                    //Instantiate the tokens and add the to the currentScene, where the uuid generated by the server is the key.
                    foreach (var roomObject in patch.Objects)
                    {
                        var obj = Instantiate(objectLibrary[roomObject.ObjectIndex], new UnityEngine.Vector3
                        {
                            x = roomObject.Position.X,
                            y = roomObject.Position.Y,
                            z = roomObject.Position.Z
                        }, Quaternion.identity);
                        _currentScene.Add(roomObject.ObjectId, obj);
                    }
                }
                else
                {
                    Debug.Log("Applying patches...");
                    foreach (var diff in patch.Diffs)
                    {
                        switch (diff.Action)
                        {
                            case Diff.Types.Action.Add:
                                var obj = Instantiate(objectLibrary[diff.Token.ObjectIndex], new UnityEngine.Vector3
                                {
                                    x = diff.Token.Position.X,
                                    y = diff.Token.Position.Y,
                                    z = diff.Token.Position.Z
                                }, Quaternion.identity);
                                _currentScene.Add(diff.Token.ObjectId, obj);
                                break;
                            case Diff.Types.Action.Delete:
                                Destroy(_currentScene[diff.Token.ObjectId]);
                                _currentScene.Remove(diff.Token.ObjectId);
                                break;
                            case Diff.Types.Action.Move:
                                var vec3 = new UnityEngine.Vector3(diff.Token.Position.X, diff.Token.Position.Y,
                                    diff.Token.Position.Z);
                                _currentScene[diff.Token.ObjectId].transform.position = vec3;
                                break;
                        }
                    }
                }
            });
        }));
    }

    /*
     * SaveRoom is an abstraction to make saving a room as simple as possible.
     */
    public void SaveRoom()
    {
        _client.SaveRoom(new RoomUser
        {
            Id = _roomId,
            UserId = _userId
        });
    }

    /*
     * AddToken is an abstraction of the AddToken RPC.
     * @param prefab: the name of the prefab of the token. This is case sensitive.
     * @param position: UnityEngine version of the Vector3 class. This is the position of a newly placed token.
     */
    public void AddToken(string prefab, UnityEngine.Vector3 position)
    {
        _client.AddToken(_roomId, _userId, _prefabLookUp[prefab], position);
    }

    /*
     * RemoveToken is used to delete a token from the scene.
     * @param uuid: the ObjectID that can be found in the _currentScene.
     */
    public void RemoveToken(string uuid)
    {
        _client.RemoveToken(_roomId, _userId, uuid);
    }

    /*
     * MoveToken is used to move a token in the scene.
     * @param uuid: the ObjectID that can be found in the _currentScene.
     * @param position: UnityEngine version of the Vector3 class. This is the new position of an existing token.
     */
    public void MoveToken(string uuid, UnityEngine.Vector3 position)
    {
        _client.MoveToken(_roomId, _userId, uuid, position);
    }


    /*
     * LoadScene is an abstraction over the RPC. Before sending the rpc is checks if the requests index is a valid index.
     * @param buildIndex: the sceneBuildIndex found in the Unity Editor. The value can not be negative.
     */
    public void LoadScene(int buildIndex)
    {
        if (buildIndex > SceneManager.sceneCountInBuildSettings - 1)
            throw new IndexOutOfRangeException("buildIndex is out of bounds. Check build settings for valid indices.");
        if (buildIndex < 0)
            throw new IndexOutOfRangeException("buildIndex can not be less than 0");

        _client.ChangeScene(_roomId, _userId, buildIndex);
    }

    /*
     * MoveUsers is used to move all users except the master to the new position.
     * @param newPosition: UnityEngine version of the Vector3 class. This is the new position of all users except the master.
     */
    public void MoveUsers(UnityEngine.Vector3 newPosition)
    {
        _client.MoveUsers(_roomId, _userId, newPosition);
    }

    /*
     * ChangeMaster changes the master from the current master to a new master.
     * @param newMasterId: this is the GUID of the new master.
     */
    public void ChangeMaster(string newMasterId)
    {
        _client.ChangeMaster(_roomId, _userId, newMasterId);
    }

    /*
     * LeaveRoom is used to leave the current room.
     * When called from the master, everyone is kicked from the room. When called from a client, only that client leaves the room.
     * This function should always be called when a session ends otherwise Unity might introduce bugs where Tasks keep running indefinitely.
     */
    public void LeaveRoom()
    {
        _client.LeaveRoom(_roomId, _userId);
    }

    #endregion
}