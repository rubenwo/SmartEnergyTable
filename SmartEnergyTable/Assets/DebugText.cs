using Network;
using System;
using System.Collections;
using System.Collections.Generic;
using System.Linq;
using TMPro;
using UnityEngine;
using UnityEngine.SceneManagement;
using UnityEngine.UI;

public class DebugText : MonoBehaviour
{

    public Button _clicker;
    public static Button Clicker;
    private NetworkManager _networkManager;

    public ARVRSwitcher _switcher;

    // Start is called before the first frame update
    void Start()
    {
        if (ARVRSwitcher.ARVRSwitch == null)
            new ARVRSwitcher();

        _switcher = ARVRSwitcher.ARVRSwitch;

        Clicker = this._clicker;

        _networkManager = GameObject.Find("GameManager").GetComponent<NetworkManager>();
        _clicker.onClick.AddListener(() =>
        {
            string a = "";

            try
            {
                foreach(var stuff in SceneManager.GetActiveScene().GetRootGameObjects())
                {
                    a += stuff.name +"("+ stuff.transform.position +")";
                }

                a += SceneManager.GetActiveScene().GetRootGameObjects().Length.ToString(); // 7 -> 6
            } catch (Exception e)
            {
                a += e.Message;
            }

            try
            {
                a += "\n"+ SceneManager.GetActiveScene().name + "("+ SceneManager.GetActiveScene().buildIndex +")";
                var obj = GameObject.Find("Camera Rig");
                a += "\nPos: "+ "("+ obj.transform.position + ")";
            }
            catch (Exception e)
            {
                a += e.Message;
            }

            try
            {
                a += "\n" + _switcher.ArEnabled + "\n";
            }
            catch (Exception e)
            {
                a += e.Message;
            }

            try
            {
                
                a += "\n" + _networkManager._currentScene.Count + "\n";
            }
            catch (Exception e)
            {
                a += e.Message;
            }
            try
            {
                var camerapos = GameObject.Find("Camera Rig").transform;
                camerapos.position = SceneManager.GetActiveScene().GetRootGameObjects().Last(o => o.name.Contains("Windmill")).transform.position;
                GameObject.Find("Camera Rig").transform.position.Set(camerapos.position.x, camerapos.position.y, camerapos.position.z - 5);
            } catch (Exception e)
            {
                a += "\n"+ e.Message;
            }



            GameObject.Find("Textos").GetComponent<TextMeshProUGUI>().text = a + " and  "+ GameObject.FindObjectsOfType(typeof(MonoBehaviour)).Length.ToString();
        });
    }

    // Update is called once per frame
    void Update()
    {
        
    }
}
