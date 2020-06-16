using Network;
using System;
using System.Collections;
using System.Collections.Generic;
using TMPro;
using UnityEngine;
using UnityEngine.SceneManagement;
using UnityEngine.UI;

public class DebugText : MonoBehaviour
{

    public Button _clicker;
    public static Button Clicker;
    private NetworkManager _networkManager;

    // Start is called before the first frame update
    void Start()
    {
        Clicker = this._clicker;
        _networkManager = GameObject.Find("GameManager").GetComponent<NetworkManager>();
        _clicker.onClick.AddListener(() =>
        {
            string a = "";

            try
            {
                foreach(var stuff in SceneManager.GetActiveScene().GetRootGameObjects())
                {
                    a += stuff.name +" ";
                }
                a += SceneManager.GetActiveScene().GetRootGameObjects().Length.ToString(); // 7 -> 6
            } catch (Exception e)
            {
                a += e.Message;
            }

            try
            {
                a += "\n"+ SceneManager.GetActiveScene().name + "("+ SceneManager.GetActiveScene().buildIndex +")";
                a += "\n"+ SceneManager.GetActiveScene().GetRootGameObjects().Length;
            }
            catch (Exception e)
            {
                a += e.Message;
            }

            a += "\n" + ARVRSwitcher.ArEnabled + "\n";
            a += "\n" + _networkManager._currentScene.Count + "\n";

            GameObject.Find("Textos").GetComponent<TextMeshProUGUI>().text = a + " and  "+ GameObject.FindObjectsOfType(typeof(MonoBehaviour)).Length.ToString();
        });
    }

    // Update is called once per frame
    void Update()
    {
        
    }
}
