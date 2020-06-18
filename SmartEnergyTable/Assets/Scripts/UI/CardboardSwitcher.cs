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

    public static CardboardSwitcher cardboard;

    // Start is called before the first frame update
    void Start()
    {
        cardboard = this;

    }

    public void Update()
    {


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
        StartCoroutine(LoadDevice("cardboard"));
    }

    void cardboardOff()
    {
        StartCoroutine(LoadDevice("none"));
    }

    IEnumerator LoadDevice(string newDevice)
    {
        XRSettings.LoadDeviceByName(newDevice);
        yield return null;
        XRSettings.enabled = true;
    }
}
