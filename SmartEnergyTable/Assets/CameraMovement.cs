using Network;
using System;
using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.Networking;
using NetworkManager = Network.NetworkManager;

public class CameraMovement : MonoBehaviour
{
    private NetworkManager _networkManager;

    private GameObject _camera;

    // Start is called before the first frame update
    void Start()
    {
        _networkManager = GameObject.Find("GameManager").GetComponent<NetworkManager>();
        _camera = GameObject.Find("Camera Rig");

        // Controls Here
        _networkManager.ObserveUserPosition(Guid.NewGuid().ToString(), (vec3) =>
        {
            this._camera.transform.position = vec3;
        });

    }

    // Update is called once per frame
    void Update()
    {

    }
}
