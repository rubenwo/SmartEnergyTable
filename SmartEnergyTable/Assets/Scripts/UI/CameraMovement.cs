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
        // Won't start when not started from Launcher, so this is a bypass.
        // We don't need Camera movement from server wehn we're in editor mode anyway
        try
        {
            _networkManager = GameObject.Find("GameManager").GetComponent<NetworkManager>();
            _camera = GameObject.Find("Camera Rig");

            // Controls Here
            //_networkManager.ObserveUserPosition(Guid.NewGuid().ToString(), (vec3) =>
            //{
            //    // Elevate to correct map height.
            //    Vector3 newVec = new Vector3(vec3.x, 20, vec3.z);
            //    this._camera.transform.position = vec3;
            //});


        } catch
        {

        }

    }

    // Update is called once per frame
    void Update()
    {
        
    }
    
    public void MovePlayer(Vector3 pos, bool IsAbsolutePosition)
    {
        if (IsAbsolutePosition)
            gameObject.transform.position = (pos);
        else
        {
            var oldPos = this._camera.transform.position;
            var newPos = new Vector3(oldPos.x + pos.x, oldPos.y + pos.y, oldPos.z + pos.z);
            gameObject.transform.position = this._camera.transform.position;
        }
    }


}
