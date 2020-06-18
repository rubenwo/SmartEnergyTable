using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using Network;
using System;
using UnityEngine.Experimental.GlobalIllumination;

public class OverviewManager : MonoBehaviour
{
    private NetworkManager _networkManager;
    private bool _isRendering;

    public GameObject gameObjectMapPrefab;
    public Camera masterCamera;
    public Camera fpvCamera;

    public Light sceneLight;

    public GameObject arController;
    public GameObject arCoreDevice;
    public GameObject envLight;
    public GameObject planeDiscovery;

    // Start is called before the first frame update
    void Start()
    {
        _networkManager = GameObject.Find("GameManager").GetComponent<NetworkManager>();
        _networkManager.ObserveMaster(Guid.NewGuid().ToString(), isMaster =>
        {
            if (!isMaster)
            {
                Debug.Log("Overview manager as non-master");
                arController.SetActive(true);
                arCoreDevice.SetActive(true);
                envLight.SetActive(true);
                planeDiscovery.SetActive(true);
                sceneLight.gameObject.SetActive(false);
                Debug.Log("Set AR go's to active");

                fpvCamera.tag = "MainCamera";

                masterCamera.enabled = false;
                fpvCamera.enabled = true;

                GameObject.Find("UserCanvas").SetActive(false);
            }
            else
            {
                masterCamera.tag = "MainCamera";
                masterCamera.enabled = true;
                fpvCamera.enabled = false;
                if (!_isRendering)
                {
                    var map = Instantiate(gameObjectMapPrefab, gameObject.transform.position, Quaternion.identity);
                    Debug.Log(map.transform.position);
                    map.transform.localScale = new Vector3(0.005f, 0.005f, 0.005f);
                    _networkManager.SetTransformForTokens(map.transform);
                    _isRendering = true;
                }
            }
        });
    }
}