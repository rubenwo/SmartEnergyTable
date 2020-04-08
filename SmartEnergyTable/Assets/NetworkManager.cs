using System.Threading.Tasks;
using Grpc.Core;
using Network;
using UnityEngine;

public class NetworkManager : MonoBehaviour
{
    private Channel _channel;
    private Client _client;

    public static NetworkManager Instance { get; private set; }


    private void Awake()
    {
        if (Instance != null && Instance != this)
            Destroy(gameObject);
        else
            Instance = this;
    }

    // Start is called before the first frame update
    private void Start()
    {
        _channel = new Channel("192.168.2.14:8080", ChannelCredentials.Insecure);
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
    private void Update()
    {
    }
}