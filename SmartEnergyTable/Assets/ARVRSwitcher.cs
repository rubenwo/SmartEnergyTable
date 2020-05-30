using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.UI;
using UnityEngine.VR;
using UnityEngine.XR;

public class ARVRSwitcher : MonoBehaviour
{

    public Sprite OffSprite;
    public Sprite OnSprite;

    private Button Source { get => gameObject.GetComponent<Button>(); }

    public static bool ArEnabled;

    // Start is called before the first frame update
    void Start()
    {
        gameObject.GetComponent<Button>().onClick.AddListener(() => SwitchARVR());
    }

    // Update is called once per frame
    void Update()
    {
        
    }

    public void SwitchARVR()
    {
        ArEnabled = !ArEnabled;

        if (ArEnabled) {
            Source.image.sprite = OffSprite;
            unSetVRComponents();
            setARComponents();

            if (XRSettings.loadedDeviceName == "None")
            {
                StartCoroutine(LoadDevice("cardboard"));
            }
        }
        else
        {
            if (XRSettings.loadedDeviceName == "cardboard")
            {
                StartCoroutine(LoadDevice("None"));
            }

            Source.image.sprite = OnSprite;
            unsetARComponents();
            setVRComponents();
        }


    }

    void setARComponents()
    {
        GameObject.Find("Map").SetActive(true);
        GameObject.Find("PlaneDiscovery").SetActive(true);

    }

    void unsetARComponents()
    {
        GameObject.Find("Map").SetActive(false);
        GameObject.Find("PlaneDiscovery").SetActive(false);
    }

    void setVRComponents()
    {
        GameObject.Find("CitySimulatorMap").SetActive(true);
        GameObject.Find("").SetActive(true);

    }

    void unSetVRComponents()
    {
        GameObject.Find("CitySimulatorMap").SetActive(false);
        GameObject.Find("").SetActive(false);

    }

    IEnumerator LoadDevice(string newDevice)
    {
        XRSettings.LoadDeviceByName(newDevice);
        yield return null;
        XRSettings.enabled = true;
    }
}
