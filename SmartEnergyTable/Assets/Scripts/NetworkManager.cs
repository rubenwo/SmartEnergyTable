using System;
using System.Collections;
using System.Collections.Generic;
using System.Threading;
using System.Threading.Tasks;
using Grpc.Core;
using UnityEditor;
using UnityEditor.PackageManager;
using UnityEngine;
using UnityEngine.Events;
using UnityEngine.SceneManagement;
using Client = Network.Client;

public class NetworkManager : MonoBehaviour
{
    public List<UnityEngine.GameObject> objectLibrary = new List<UnityEngine.GameObject>();
    private Channel _channel;
    private Client _client;
    private Queue<Action> obj = new Queue<Action>();
    public static NetworkManager Instance { get; private set; }


    private static NetworkManager s_Instance = null;

    void Awake()
    {
        if (s_Instance == null)
        {
            s_Instance = this;
            DontDestroyOnLoad(gameObject);

            //Initialization code goes here[/INDENT]
        }
        else
        {
            Destroy(gameObject);
        }
    }

    private string roomId = "";
    private string userId = Guid.NewGuid().ToString();

    // Start is called before the first frame update
    private void Start()
    {
        _channel = new Channel("192.168.2.14:8080", ChannelCredentials.Insecure);
        _client = new Client(new SmartEnergyTableService.SmartEnergyTableServiceClient(_channel));
    }

    public void CreateRoom()
    {
        if (roomId != "")
            return;
        var room = _client.CreateRoom();
        Debug.Log(room.Id);
        roomId = room.Id;
        JoinRoom(roomId);
    }

    public void JoinRoom(string id)
    {
        Task.Run(() => _client.JoinRoom(id, userId, update =>
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

    public void AddGameObject(string prefab, float posX, float posY, float posZ)
    {
        var empty = _client.AddGameObject(roomId, userId, prefab, posX, posY, posZ);
    }

    public void LoadScene(uint buildIndex)
    {
        if (buildIndex > SceneManager.sceneCountInBuildSettings - 1)
            throw new IndexOutOfRangeException("buildIndex is out of bounds. Check build settings for valid indices.");
        SceneManager.LoadScene(Convert.ToInt32(buildIndex));
    }

    private void OnDisable()
    {
        Debug.Log("OnDisable()");
        _client.LeaveRoom(roomId, userId);
        _channel.ShutdownAsync().Wait();
    }

    // Update is called once per frame
    private void Update()
    {
        if (obj.Count > 0)
        {
            obj.Dequeue()();
        }
    }
}