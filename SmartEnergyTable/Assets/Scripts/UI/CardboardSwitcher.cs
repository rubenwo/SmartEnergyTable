using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.UI;
using UnityEngine.XR;

public class CardboardSwitcher : MonoBehaviour
{

    public Button CardboardButton;

    private bool _cardboardActive;
    public bool CardboardActive { get => _cardboardActive; set { _cardboardActive = value; updateCardboard(); } }

    // Start is called before the first frame update
    void Start()
    {
        CardboardButton.onClick.AddListener(() => CardboardActive = !CardboardActive);
    }


    void updateCardboard()
    {
        if (_cardboardActive)
            cardboardOn();
        else
            cardboardOff();
    }

    void cardboardOn()
    {
        if (XRSettings.loadedDeviceName == "cardboard")
            StartCoroutine(LoadDevice("None"));
    }

    void cardboardOff()
    {
        if (XRSettings.loadedDeviceName == "cardboard")
            StartCoroutine(LoadDevice("None"));
    }

    IEnumerator LoadDevice(string newDevice)
    {
        XRSettings.LoadDeviceByName(newDevice);
        yield return null;
        XRSettings.enabled = true;
    }
}
