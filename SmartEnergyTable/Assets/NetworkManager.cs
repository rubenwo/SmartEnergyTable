using System;
using System.Collections;
using System.Collections.Generic;
using System.Threading;
using System.Threading.Tasks;
using Grpc.Core;
using Network;
using UnityEngine;

public class NetworkManager : MonoBehaviour
{
    private static NetworkManager _instance;
    private Network.Client _client;
    private Channel _channel;

    public static NetworkManager Instance
    {
        get { return _instance; }
    }


    private void Awake()
    {
        if (_instance != null && _instance != this)
        {
            Destroy(this.gameObject);
        }
        else
        {
            _instance = this;
        }
    }

    // Start is called before the first frame update
    void Start()
    {
        _channel = new Channel("127.0.0.1:8080", ChannelCredentials.Insecure);
        _client = new Client(new SmartEnergyTableService.SmartEnergyTableServiceClient(_channel));
        var room = _client.CreateRoom();
        Debug.Log(room.Id);
        Task.Run(() => _client.JoinRoom(room.Id, update => Debug.Log(update.Id)));
    }

    private void OnDisable()
    {
        _channel.ShutdownAsync().Wait();
    }

    // Update is called once per frame
    void Update()
    {
    }
}