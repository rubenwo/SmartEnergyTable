using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using Network;
using System;

public class OverviewManager : MonoBehaviour
{
    private NetworkManager _networkManager;
    private bool isRendering;

    public GameObject GameObjectMapPrefab;
    public Camera MasterCamera;
    public Camera FPVCamera;

    // Start is called before the first frame update
    void Start()
    {
        _networkManager = GameObject.Find("GameManager").GetComponent<NetworkManager>();
        _networkManager.ObserveMaster(Guid.NewGuid().ToString(), isMaster =>
        {
            if (!isMaster)
            {
                FPVCamera.tag = "MainCamera";
                //Activate AR
                var g = GameObject.Find("ARCore Device");
                g.SetActive(true);
                g = GameObject.Find("ARController");
                g.SetActive(true);
                g = GameObject.Find("Environmental Light");
                g.SetActive(true);
                g = GameObject.Find("PlaneDiscovery");
                g.SetActive(true);
                MasterCamera.enabled = false;
                FPVCamera.enabled = true;
            }
            else
            {
                MasterCamera.tag = "MainCamera";
                MasterCamera.enabled = true;
                FPVCamera.enabled = false;
                //create mapbox at gameObject.transform.position
                if (!isRendering)
                {
                    var map = Instantiate(GameObjectMapPrefab, gameObject.transform.position, Quaternion.identity);
                    map.transform.localScale = new Vector3(0.005f, 0.005f, 0.005f);
                    isRendering = true;
                }
            }
        });
    }

    // Update is called once per frame
    void Update()
    {
        
    }
}
