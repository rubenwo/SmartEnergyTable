using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Grpc.Core;
using Network;
using UnityEngine;
using UnityEngine.SceneManagement;

public class NetworkManager : MonoBehaviour
{
    private static NetworkManager s_Instance;
    private readonly string _userId = Guid.NewGuid().ToString();
    private Channel _channel;
    private Client _client;

    private string _roomId = "";
    private readonly Queue<Action> obj = new Queue<Action>();
    public List<UnityEngine.GameObject> objectLibrary = new List<UnityEngine.GameObject>();
    public static NetworkManager Instance { get; private set; }

    private void Awake()
    {
        if (s_Instance == null)
        {
            s_Instance = this;
            DontDestroyOnLoad(gameObject);

            _channel = new Channel("192.168.2.14:8080", ChannelCredentials.Insecure);
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
        if (obj.Count > 0) obj.Dequeue()();
    }

    #region RPCs

    public void CreateRoom()
    {
        if (_roomId != "")
            return;
        var room = _client.CreateRoom();
        _roomId = room.Id;
        JoinRoom(_roomId);
    }

    public void JoinRoom(string id)
    {
        Task.Run(() => _client.JoinRoom(id, _userId, update =>
        {
            obj.Enqueue(() =>
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

    public void AddGameObject(string prefab, Vector3 position)
    {
        _client.AddGameObject(_roomId, _userId, prefab, position);
    }

    public void RemoveGameObject(string prefab, Vector3 position)
    {
        _client.RemoveGameObject(_roomId, _userId, prefab, position);
    }

    public void MoveGameObject(string prefab, Vector3 position)
    {
        _client.MoveGameObject(_roomId, _userId, prefab, position);
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