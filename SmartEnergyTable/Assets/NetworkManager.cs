using System.Collections;
using System.Collections.Generic;
using Network;
using UnityEngine;

public class NetworkManager : MonoBehaviour
{
    private static NetworkManager _instance;
    private Network.Client _client;

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
        _client = new Client();
        _client.GetRandomColor("0xffffff");
    }

    // Update is called once per frame
    void Update()
    {
    }
}