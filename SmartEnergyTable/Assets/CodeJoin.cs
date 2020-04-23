using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.UI;

public class CodeJoin : MonoBehaviour
{
    public InputField codeField;

    public NetworkManager Manager;

    // Start is called before the first frame update
    void Start()
    {
    }

    // Update is called once per frame
    void Update()
    {
    }

    public void onClick()
    {
        Debug.Log(codeField.text);
        Manager.JoinRoom(codeField.text);
    }
}