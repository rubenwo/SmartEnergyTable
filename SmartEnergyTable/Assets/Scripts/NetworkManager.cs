using System;
using System.Collections;
using System.Collections.Generic;
using System.Threading;
using System.Threading.Tasks;
using Grpc.Core;
using Network;
using UnityEditor;
using UnityEngine;
using UnityEngine.Events;
using UnityEngine.SceneManagement;

public class NetworkManager : MonoBehaviour
{
    public List<UnityEngine.GameObject> GameObjects = new List<UnityEngine.GameObject>();
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

    private string roomId;
    private string userId = "test123";

    // Start is called before the first frame update
    private void Start()
    {
        _channel = new Channel("192.168.2.14:8080", ChannelCredentials.Insecure);
        _client = new Client(new SmartEnergyTableService.SmartEnergyTableServiceClient(_channel));

        var room = _client.CreateRoom();
        Debug.Log(room.Id);
        roomId = room.Id;

        Task.Run(() => _client.JoinRoom(roomId, userId, update =>
        {
            obj.Enqueue(() => Instantiate(GameObjects[0]));

            Debug.Log(update.Id);
        }));
        SceneManager.LoadScene(1);
        Task.Run(() =>
        {
            Thread.Sleep(5000);
            _client.AddGameObject(roomId, userId, "Cube", 0, 0, 0);
        });
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