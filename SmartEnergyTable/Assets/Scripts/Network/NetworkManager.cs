﻿using System;
using System.Collections;
using System.Collections.Generic;
using System.IO;
using System.Net.Http;
using System.Security.Cryptography.X509Certificates;
using System.Threading.Tasks;
using Grpc.Core;
using Grpc.Core.Api;
using UnityEngine;
using UnityEngine.SceneManagement;

namespace Network
{
    /// <summary>
    /// NetworkManager is a singleton that manages the state of the game through a networked server.
    /// </summary>
    public sealed class NetworkManager : MonoBehaviour
    {
        ///This is the server address for the backend.
        public string serverAddr;

        ///<summary>The objectLibrary is a list of prefabs. As we're not using many different tokens it's more efficient to store them in an array.
        ///The library is filled in the editor. Rather than loading them from Resources when we need them.</summary>
        public List<GameObject> objectLibrary = new List<GameObject>();

        ///To make the life of developers easier the _prefabLookup returns the index for the objectLibrary based on the name of the prefab.
        private readonly Dictionary<string, int> _prefabLookUp = new Dictionary<string, int>();

        public readonly List<string> Prefabs = new List<string>();

        ///Static instance so there is only 1 NetworkManager
        private static NetworkManager _instance;

        private Channel _channel;
        private Client _client;
        private readonly Queue<Action> _actionQueue = new Queue<Action>();


        private readonly string _userId = Guid.NewGuid().ToString(); //Generate a new GUID as userId.
        private string _roomId = "";
        private bool _master;
        private EnergyData _energyData = null;
        private bool _connected;
        private bool _sceneLoaded = true;

        private Transform _parentTransformForTokens = null;
        private Vector3 _userPosition = Vector3.zero;
        private GeneratedEnergy _generatedEnergy;

        ///<summary>This is a representation of the current scene according to the server. This manages all the tokens.
        ///Key (string) is the uuid generated by the server on AddToken. Value (GameObject) is the token.</summary>
        private readonly Dictionary<string, GameObject> _currentScene = new Dictionary<string, GameObject>();

        ///<summary>_uuidLookUp is the same as the _currentScene dictionary except using GameObjects as value and uuid as key.
        ///When using collision systems like RayCasts this dictionary can be used to (re)move objects in the scene.</summary>
        private readonly Dictionary<GameObject, string> _uuidLookUp = new Dictionary<GameObject, string>();

        private readonly Dictionary<string, Action<bool>> _masterChangeListeners =
            new Dictionary<string, Action<bool>>();

        private readonly Dictionary<string, Action<EnergyData>> _energyDataListeners =
            new Dictionary<string, Action<EnergyData>>();

        private readonly Dictionary<string, Action<Vector3>> _userPositionListeners =
            new Dictionary<string, Action<Vector3>>();


        private void Awake()
        {
            if (_instance == null)
            {
                _instance = this;
                //We want the network manager to exist in every scene, so we need to call DontDestroyOnLoad on this gameObject.
                DontDestroyOnLoad(gameObject);

                //Fill the prefab lookup 
                for (var i = 0; i < objectLibrary.Count; i++)
                {
                    _prefabLookUp[objectLibrary[i].name] = i;
                    Prefabs.Add(objectLibrary[i].name);
                }


                try
                {
                    var secureCredentials = new SslCredentials();

                    //Create the gRPC channel and client
                    _channel = new Channel(serverAddr, secureCredentials);
                    _client = new Client(new SmartEnergyTableService.SmartEnergyTableServiceClient(_channel));
                }
                catch (Exception e)
                {
                    Debug.Log(e);
                    throw;
                }
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
            if (SceneManager.GetActiveScene().buildIndex != 0
            ) //If we are on the Launcher menu we have no room to leave.
                _client.LeaveRoom(_roomId, _userId); //Leave the room before we shutdown.
            //Shutdown the channel synchronously to avoid bugs.
            _channel.ShutdownAsync().Wait();
        }


        // Update is called once per frame and runs on the main thread.
        private void Update()
        {
            //The update from the JoinRoom adds actions to the queue. We perform these actions in Update() to run them on the main thread.
            //Unity doesn't support multithreaded calls to the engine like 'Instantiate'.
            lock (_actionQueue)
            {
                while (_actionQueue.Count > 0)
                {
                    _actionQueue.Dequeue()();
                }
            }
        }


        /// <summary>
        /// IsMaster can be used in UI elements to show a "master" only view.
        /// </summary>
        public bool IsMaster => _master;

        /// <summary>
        /// SessionID returns the current room ID.
        /// </summary>
        public string SessionID => _roomId;


        /// <summary>
        /// ObserveMaster adds the callback action to a list. When a patch changes the master role we update these listeners.
        /// </summary>
        /// <param name="callback">Action with bool as param, this is called when a patch comes in</param>
        /// <param name="uuid">This is an identifier for the listener</param>
        public void ObserveMaster(string uuid, Action<bool> callback)
        {
            _masterChangeListeners.Add(uuid, callback);
            foreach (var masterChangeListener in _masterChangeListeners)
            {
                masterChangeListener.Value.Invoke(_master);
            }
        }

        /// <summary>
        /// ObserveEnergyData adds a callback to the internal list. When a patch updates the energy data, these callbacks
        /// are invoked.
        /// </summary>
        /// <param name="callback">Action(EnergyData), an action that is called when the EnergyData has changed</param>
        /// <param name="uuid">This is an identifier for the listener</param>
        public void ObserveEnergyData(string uuid, Action<EnergyData> callback)
        {
            _energyDataListeners.Add(uuid, callback);
            foreach (var energyDataListener in _energyDataListeners)
            {
                energyDataListener.Value.Invoke(_energyData);
            }
        }


        /// <summary>
        /// UnObserveMaster: When a listener no longer needs to listen they should unsubscribe.
        /// </summary>
        /// <param name="uuid">This is an identifier for the listener</param>
        public void UnObserveMaster(string uuid)
        {
            _masterChangeListeners.Remove(uuid);
        }

        /// <summary>
        /// UnObserveEnergyData: When a listener no longer needs to listen they should unsubscribe.
        /// </summary>
        /// <param name="uuid">This is an identifier for the listener</param>
        public void UnObserveEnergyData(string uuid)
        {
            _energyDataListeners.Remove(uuid);
        }

        /// <summary>
        /// SetTransformForTokens sets the transform of the parent for the tokens. useful when the tokens need to be in local space.
        /// </summary>
        /// <param name="t">Transform of the parent object</param>
        public void SetTransformForTokens(Transform t)
        {
            _parentTransformForTokens = t;
        }


        public GeneratedEnergy GeneratedEnergy => _generatedEnergy;

        #region RPCs

        /// <summary>
        /// CreateRoom is an abstraction over the RPC. This function creates a room, sets the roomId and master.
        /// Then joins the newly created room.
        /// </summary>
        public void CreateRoom()
        {
            if (_roomId != "")
                return;
            var room = _client.CreateRoom();
            _roomId = room.Id;
            _master = true;
            JoinRoom(_roomId);
        }


        /// <summary>
        /// LoadSceneAsync is a used as a coroutine to load a scene asynchronously.
        /// </summary>
        /// <param name="sceneBuildIndex">this is the build index created by the Unity Editor in 'File->Build Settings'.</param>
        /// <returns>IEnumerator for the StartCoroutine</returns>
        private IEnumerator LoadSceneAsync(int sceneBuildIndex)
        {
            _sceneLoaded = false;

            var asyncLoad = SceneManager.LoadSceneAsync(sceneBuildIndex);

            // Wait until the asynchronous scene fully loads
            while (!asyncLoad.isDone)
            {
                yield return null;
            }

            _sceneLoaded = true;
        }

        /// <summary>
        /// ProcessDiffs is a coroutine that processes the diffs from a patch
        /// </summary>
        /// <param name="diffs">this is a list of the diffs. This coroutine will loop through this and execute the action specified by the diff</param>
        /// <returns>IEnumerator for the StartCoroutine</returns>
        private IEnumerator ProcessDiffs(IEnumerable<Diff> diffs)
        {
            // If the scene is not fully loaded we need te wait before processing diffs or we risk instantiating/moving/destroying
            // objects in the wrong scene.
            while (!_sceneLoaded)
                yield return null;
            foreach (var diff in diffs)
            {
                switch (diff.Action)
                {
                    case Diff.Types.Action.Add:
                        var t = _parentTransformForTokens != null ? _parentTransformForTokens : transform;
                        var obj = Instantiate(objectLibrary[diff.Token.ObjectIndex],
                            t.position + new Vector3
                            {
                                x = diff.Token.Position.X,
                                y = diff.Token.Position.Y,
                                z = diff.Token.Position.Z
                            }, Quaternion.identity);
                        obj.transform.localScale *= diff.Token.Scale;
                        _currentScene.Add(diff.Token.ObjectId, obj);
                        break;
                    case Diff.Types.Action.Delete:
                        Destroy(_currentScene[diff.Token.ObjectId]);
                        _currentScene.Remove(diff.Token.ObjectId);
                        break;
                    case Diff.Types.Action.Move:
                        _currentScene[diff.Token.ObjectId].transform.position = new Vector3(
                            diff.Token.Position.X,
                            diff.Token.Position.Y,
                            diff.Token.Position.Z
                        );
                        break;
                    default:
                        Debug.Log("ProcessDiffs() => ERROR: Unknown diff action!");
                        //We should never come here, but ReShaper won't shut up about 'missing default case'.
                        break;
                }
            }
        }

        /// <summary>
        /// JoinRoom uses the RPC callback to manage the game state managed by the server.
        /// This function starts a long-running task that enqueues actions on every update from the server.
        /// </summary>
        /// <param name="id">String with the ID of the room</param>
        public void JoinRoom(string id)
        {
            //If we're not the master, the roomId won't be set yet.
            if (_roomId == "")
                _roomId = id;
            Task.Run(async () =>
            {
                _connected = true;
                await _client.JoinRoom(id, _userId, patch =>
                {
                    lock (_actionQueue)
                    {
                        GetEnergyData();
                        _actionQueue.Enqueue(() =>
                        {
                            _master = patch.IsMaster;
                            foreach (var masterChangeListener in _masterChangeListeners)
                            {
                                masterChangeListener.Value.Invoke(_master);
                            }

                            foreach (var energyDataListener in _energyDataListeners)
                            {
                                energyDataListener.Value.Invoke(_energyData);
                            }

                            _userPosition = new Vector3(patch.UserPosition.X, patch.UserPosition.Y,
                                patch.UserPosition.Z);
                            foreach (var userPositionListener in _userPositionListeners)
                            {
                                userPositionListener.Value.Invoke(_userPosition);
                            }

                            _generatedEnergy = patch.Energy;
                            //Load the scene if it is not the currentScene, meaning the scene has changed.
                            if (patch.SceneId != SceneManager.GetActiveScene().buildIndex)
                            {
                                //The master should only be able to change between scene 0 and 1 (Launcher and Overview).
                                //The other clients should be able to change to any scene.
                                if (!_master)
                                    StartCoroutine(LoadSceneAsync(patch.SceneId));
                                else if (!(patch.SceneId > 1))
                                    StartCoroutine(
                                        LoadSceneAsync(patch.SceneId));
                            }

                            //If the _currentScene is empty we want to process the entire history as this might mean we joined
                            //later on and some object might not be send through the 'normal' diffs.
                            if (_currentScene.Count == 0)
                            {
                                Debug.Log("Process patch history...");
                                StartCoroutine(ProcessDiffs(patch.History));
                            }
                            else
                            {
                                Debug.Log("Applying patches...");
                                StartCoroutine(ProcessDiffs(patch.Diffs));
                            }

                            _uuidLookUp.Clear();
                            foreach (var keyValuePair in _currentScene)
                            {
                                _uuidLookUp.Add(keyValuePair.Value, keyValuePair.Key);
                            }

                            Debug.Log(_currentScene.Count);
                        });
                    }
                });
                _connected = false;
                LeaveRoom(); // Leave the room when the patches stop (causes are: RPC failure or master left).
            });
        }


        /// <summary>
        /// SaveRoom calls the server to save the session to the database.
        /// </summary>
        public void SaveRoom()
        {
            _client.SaveRoom(new RoomUser
            {
                Id = _roomId,
                UserId = _userId
            });
        }

        /// <summary>
        /// AddToken calls the server to add a new Token to the room.
        /// </summary>
        /// <param name="prefab">the name of the prefab of the token. This is case sensitive.</param>
        /// <param name="efficiency">The efficiency of the token. This value should be between 0 & 100</param>
        /// <param name="position">UnityEngine version of the Vector3 class. This is the position of a newly placed token.</param>
        public void AddToken(string prefab, int efficiency, Vector3 position, float scale = 1)
        {
            if (efficiency < 0) efficiency = 0;
            if (efficiency > 100) efficiency = 100;
            _client.AddToken(_roomId, _userId, _prefabLookUp[prefab], efficiency, position, scale);
        }


        /// <summary>
        /// RemoveToken is used to delete a token from the scene it's uuid.
        /// </summary>
        /// <param name="uuid">the ObjectID that can be found in the _currentScene.</param>
        public void RemoveToken(string uuid)
        {
            _client.RemoveToken(_roomId, _userId, uuid);
        }


        /// <summary>
        /// RemoveToken is used to delete a token from the scene by the GameObject. Useful when using collision systems like RayCasts.
        /// </summary>
        /// <param name="obj">the GameObject that should be removed.</param>
        public void RemoveToken(GameObject obj)
        {
            _client.RemoveToken(_roomId, _userId, _uuidLookUp[obj]);
        }

        /// <summary>
        /// MoveToken is used to move a token in the scene.
        /// </summary>
        /// <param name="uuid">the ObjectID that can be found in the _currentScene.</param>
        /// <param name="position">UnityEngine version of the Vector3 class. This is the new position of an existing token.</param>
        public void MoveToken(string uuid, UnityEngine.Vector3 position)
        {
            _client.MoveToken(_roomId, _userId, uuid, position);
        }

        /// <summary>
        /// MoveToken is used to move a token in the scene using the GameObject. Useful when using collision systems like RayCasts.
        /// </summary>
        /// <param name="obj"> the GameObject that should be removed.</param>
        /// <param name="position">UnityEngine version of the Vector3 class. This is the new position of an existing token.</param>
        public void MoveToken(GameObject obj, UnityEngine.Vector3 position)
        {
            _client.MoveToken(_roomId, _userId, _uuidLookUp[obj], position);
        }

        /// <summary>
        /// ClearScene removes all the tokens from the current scene.
        /// </summary>
        public void ClearScene()
        {
            _client.ClearRoom(_roomId, _userId);
        }


        /// <summary>
        /// LoadScene checks if the requests index is valid before sending a request to the server to change the scene for all users.
        /// </summary>
        /// <param name="buildIndex">the sceneBuildIndex found in the Unity Editor. The value can not be negative.</param>
        /// <exception cref="IndexOutOfRangeException">If the request buildIndex is negative or larger than the last index
        /// of the scenes defined in the build settings this exception is thrown.</exception>
        public void LoadScene(int buildIndex)
        {
            if (buildIndex > SceneManager.sceneCountInBuildSettings - 1)
                throw new IndexOutOfRangeException(
                    "buildIndex is out of bounds. Check build settings for valid indices.");
            if (buildIndex < 0)
                throw new IndexOutOfRangeException("buildIndex can not be less than 0");

            _client.ChangeScene(_roomId, _userId, buildIndex);
        }


        /// <summary>
        ///  MoveUsers is used to move all users except the master to the new position.
        /// </summary>
        /// <param name="newPosition">UnityEngine version of the Vector3 class. This is the new position of all users except the master.</param>
        public void MoveUsers(UnityEngine.Vector3 newPosition)
        {
            _client.MoveUsers(_roomId, _userId, newPosition);
        }

        /// <summary>
        /// ChangeMaster changes the master from the current master to a new master.
        /// </summary>
        /// <param name="newMasterId">this is the GUID of the new master.</param>
        public void ChangeMaster(string newMasterId)
        {
            _client.ChangeMaster(_roomId, _userId, newMasterId);
        }


        /// <summary>
        /// LeaveRoom is used to leave the current room.
        /// When called from the master, everyone is kicked from the room. When called from a client, only that client leaves the room.
        /// This function should always be called when a session ends otherwise Unity keeps Tasks running indefinitely.
        /// </summary>
        public void LeaveRoom()
        {
            if (_connected)
                _client.LeaveRoom(_roomId,
                    _userId); //Leave the room on the server only if we're connected to the server.
            //Clear the network manager after leaving the room.
            SceneManager.LoadScene(0);
            _master = false;
            _roomId = "";
            _currentScene.Clear();
            lock (_actionQueue)
                _actionQueue.Clear();
            _uuidLookUp.Clear();
        }

        public EnergyData GetEnergyData()
        {
            if (_energyData == null)
                _energyData = _client.GetEnergyData(_roomId, _userId);
            return _energyData;
        }

        #endregion
    }
}