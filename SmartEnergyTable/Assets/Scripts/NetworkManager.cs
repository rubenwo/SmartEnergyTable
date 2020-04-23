using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Grpc.Core;
using Network;
using UnityEngine;
using UnityEngine.SceneManagement;

public class NetworkManager : MonoBehaviour
{
    public string serverAddr;
    public List<GameObject> objectLibrary = new List<GameObject>();


    private static NetworkManager s_Instance;
    private Channel _channel;
    private Client _client;
    private readonly Queue<Action> _actionQueue = new Queue<Action>();


    private readonly string _userId = Guid.NewGuid().ToString();
    private string _roomId = "";
    private bool _master;

    private void Awake()
    {
        if (s_Instance == null)
        {
            s_Instance = this;
            DontDestroyOnLoad(gameObject);

            _channel = new Channel(serverAddr, ChannelCredentials.Insecure);
            _client = new Client(new SmartEnergyTableService.SmartEnergyTableServiceClient(_channel));
        }
        else
        {
            Destroy(gameObject);
        }
    }

    private void OnDisable()
    {
        Debug.Log("OnDisable()");
        _client.LeaveRoom(_roomId, _userId);
        _channel.ShutdownAsync().Wait();
    }

    // Update is called once per frame
    private void Update()
    {
        while (_actionQueue.Count > 0)
            _actionQueue.Dequeue()();
    }


    public bool IsMaster => _master;

    #region RPCs

    public void CreateRoom()
    {
        if (_roomId != "")
            return;
        var room = _client.CreateRoom();
        _roomId = room.Id;
        _master = true;
        JoinRoom(_roomId);
    }

    public void JoinRoom(string id)
    {
        Task.Run(() => _client.JoinRoom(id, _userId, update =>
        {
            _actionQueue.Enqueue(() =>
            {
                if (update.Room.SceneId != SceneManager.GetActiveScene().buildIndex)
                    SceneManager.LoadScene(update.Room.SceneId);
                Instantiate(objectLibrary[0]);
            });

            Debug.Log(update.Id);
        }));
    }

    public void SaveRoom()
    {
        _client.SaveRoom(new Room
        {
            Id = _roomId,
            SceneId = SceneManager.GetActiveScene().buildIndex
        });
    }

    public void AddToken(string prefab, Vector3 position)
    {
        _client.AddToken(_roomId, _userId, prefab, position);
    }

    public void RemoveToken(string prefab, Vector3 position)
    {
        _client.RemoveToken(_roomId, _userId, prefab, position);
    }

    public void MoveToken(string prefab, Vector3 position)
    {
        _client.MoveToken(_roomId, _userId, prefab, position);
    }


    public void LoadScene(uint buildIndex)
    {
        if (buildIndex > SceneManager.sceneCountInBuildSettings - 1)
            throw new IndexOutOfRangeException("buildIndex is out of bounds. Check build settings for valid indices.");
        //SceneManager.LoadScene(Convert.ToInt32(buildIndex));
        _client.ChangeScene(_roomId, _userId, Convert.ToInt32(buildIndex));
    }

    public void MoveUsers(Vector3 newPosition)
    {
        _client.MoveUsers(_roomId, _userId, newPosition);
    }

    public void LeaveRoom()
    {
        _client.LeaveRoom(_roomId, _userId);
    }

    #endregion
}