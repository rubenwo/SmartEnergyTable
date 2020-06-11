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
    private List<GameObject> ARObjects;
    private List<GameObject> VRObjects;
    private NetworkManager _networkManager;

    private Button Source { get => gameObject.GetComponent<Button>(); }

    public static bool ArEnabled { get; set; } = true;

    public static ARVRSwitcher ARVRSwitch;

    // Start is called before the first frame update
    void Start()
    {
        ARVRSwitch = this;

        // Insert all our objects into the right lists (Other ojects are rendered in both scenes)
        //ARObjects = new List<GameObject>() {
        //    GameObject.Find("PlaneDiscovery")
        //};
        //VRObjects = new List<GameObject>() {
        //    GameObject.Find("Camera Rig")
        //};
        ARObjects = new List<GameObject>() {
        };
        VRObjects = new List<GameObject>() {
        };

        _networkManager = GameObject.Find("GameManager").GetComponent<NetworkManager>();

        Source.onClick.AddListener(() => SwitchARVR());

        //unSetVRComponents();
    }

    // Update is called once per frame
    void Update()
    {
        //ArEnabled = true;
        //SwitchARVR();
    }

    internal static void switchClientMode(ViewMode view)
    {
        if (view == ViewMode.Overview)
        {
            ArEnabled = true;
            SceneManager.LoadScene(1);
        } else
        {
            ArEnabled = false;
            SceneManager.LoadScene(2);
        }
    }

    public void SwitchARVR()
    {

        if (ArEnabled) {
            Source.image.sprite = OffSprite;

            // send to server: swap all clients to AR
            _networkManager.LoadScene(1);

            //unSetVRComponents();
            //setARComponents();
            UnityEngine.Debug.Log("Hallo VR");
        }
        else
        {
            Source.image.sprite = OnSprite;

            ////send to server: swap all clients to VR
            _networkManager.LoadScene(2);
            UnityEngine.Debug.Log("Hallo AR");

            //unsetARComponents();
            //setVRComponents();
        }
    }

    

    void setARComponents()
    {
        ARObjects.ForEach(ob => ob.SetActive(true));
    }

    void unsetARComponents()
    {
        ARObjects.ForEach(ob => ob.SetActive(false));
    }

    void setVRComponents()
    {
        VRObjects.ForEach(ob => ob.SetActive(true));
    }

    void unSetVRComponents()
    {
        VRObjects.ForEach(ob => ob.SetActive(false));

    }

}
