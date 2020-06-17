using Mapbox.Unity.Utilities;
using System.Collections;
using System.Collections.Generic;
using System.Diagnostics;
using UnityEngine;
using UnityEngine.UI;
using UnityEngine.VR;
using UnityEngine.XR;
using Network;
using System.Security.Policy;
using System;
using UnityEngine.SceneManagement;

public class ARVRSwitcher : MonoBehaviour
{

    public Sprite OffSprite;
    public Sprite OnSprite;

    // All AR/VR Objects present in the scene
    private NetworkManager _networkManager;

    private Button Source { get => GameObject.Find("SwitchARVR").GetComponent<Button>(); }

    public bool ArEnabled { get; set; } = true;

    public static ARVRSwitcher ARVRSwitch;

    // Start is called before the first frame update
    void Start()
    {
        ARVRSwitch = this;

        _networkManager = GameObject.Find("GameManager").GetComponent<NetworkManager>();

        Source.onClick.AddListener(() => SwitchARVR());

    }

    // Update is called once per frame
    void Update()
    {
    }

    public void switchClientMode(ViewMode view)
    {
        //if (_networkManager.IsMaster)
        //    return;

        if (view == ViewMode.Overview)
        {
            ArEnabled = true;
            SceneManager.LoadScene(1);
        }
        else
        {
            ArEnabled = false;
            SceneManager.LoadScene(2);
        }
    }

    public void SwitchARVR()
    {
        ArEnabled = !ArEnabled;

        if (ArEnabled) {
            Source.image.sprite = OffSprite;

            // send to server: swap all clients to AR
            _networkManager.LoadScene(1);

        }
        else
        {
            Source.image.sprite = OnSprite;

            ////send to server: swap all clients to VR
            _networkManager.LoadScene(2);

        }

        UnityEngine.Debug.Log("Sent");
    }

}
