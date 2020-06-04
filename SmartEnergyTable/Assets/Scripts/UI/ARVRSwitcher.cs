using Mapbox.Unity.Utilities;
using System.Collections;
using System.Collections.Generic;
using System.Diagnostics;
using UnityEngine;
using UnityEngine.UI;
using UnityEngine.VR;
using UnityEngine.XR;

public class ARVRSwitcher : MonoBehaviour
{

    public Sprite OffSprite;
    public Sprite OnSprite;

    // All AR/VR Objects present in the scene
    private List<GameObject> ARObjects;
    private List<GameObject> VRObjects;

    private Button Source { get => gameObject.GetComponent<Button>(); }

    public static bool ArEnabled;

    // Start is called before the first frame update
    void Start()
    {
        // Insert all our objects into the right lists (Other ojects are rendered in both scenes)
        ARObjects = new List<GameObject>() {
            GameObject.Find("PlaneDiscovery")
        };
        VRObjects = new List<GameObject>() {
            GameObject.Find("Camera Rig")
        };


        unSetVRComponents();
    }

    // Update is called once per frame
    void Update()
    {
        ArEnabled = true;
        SwitchARVR();
    }

    public void SwitchARVR()
    {
        ArEnabled = !ArEnabled;

        if (ArEnabled) {
            Source.image.sprite = OffSprite;
            unSetVRComponents();
            setARComponents();

            if (XRSettings.loadedDeviceName == "cardboard")
                StartCoroutine(LoadDevice("None"));
        }
        else
        {
            if (XRSettings.loadedDeviceName == "None")
                StartCoroutine(LoadDevice("cardboard"));

            Source.image.sprite = OnSprite;
            unsetARComponents();
            setVRComponents();
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

    IEnumerator LoadDevice(string newDevice)
    {
        XRSettings.LoadDeviceByName(newDevice);
        yield return null;
        XRSettings.enabled = true;
    }

}
