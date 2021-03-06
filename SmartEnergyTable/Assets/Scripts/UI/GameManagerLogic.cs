﻿using Network;
using System;
using System.Linq;
using UnityEngine;
using UnityEngine.SceneManagement;
using UnityEngine.UI;

public class GameManagerLogic : MonoBehaviour
{
    private NetworkManager _netMan;
    private string _id = Guid.NewGuid().ToString();

    public Button graphButton;

    private bool _graphsActive;
    public GameObject gameObjectMapPrefab;

    private Token Tok
    {
        get => this.GetComponent<TokenData>().Tok;
        set => this.GetComponent<TokenData>().Tok = value;
    }

    public GameObject prefab;

    private ARVRSwitcher _switcher;

    // Start is called before the first frame update
    void Start()
    {
        ARVRSwitcher.CreateIfNeeded();

        try
        {
            _netMan = GameObject.Find("GameManager").GetComponent<NetworkManager>();

            _netMan.ObserveMaster(_id, (state) => { Debug.Log("Steet: " + state); });

            _netMan.ObserveEnergyData(_id, (ener) =>
            {
                //Debug.Log("Got this: " + ener.EnergyUsers[1].Pv);
            });

            _netMan.ObserveGeneratedEnergy(_id, (ener) =>
            {
                EnergyData en = _netMan.GetEnergyData();
            });
            if (!_netMan.IsMaster)
            {
//                GameObject ob = GameObject.Find("CitySimulatorMap");
//                //ob.transform.localScale = new Vector3(0.005f, 0.005f, 0.005f);
//                _netMan.SetTransformForTokens(ob.transform);

                var map = Instantiate(gameObjectMapPrefab, gameObject.transform.position, Quaternion.identity);
                map.transform.localScale = new Vector3(0.005f, 0.005f, 0.005f);
                _netMan.SetTransformForTokens(map.transform);
            }


            _netMan.ObserveViewMode(_id, (view) =>
            {
                if (_netMan.IsMaster)
                    return;

                ARVRSwitcher.ARVRSwitch.switchClientMode(view);
            });
        }
        catch (Exception e)
        {
            Debug.Log("Got err: " + e.Message);
            Debug.Log(e.StackTrace);
            Debug.Log(e.Data);
        }

        graphButton.onClick.AddListener(() => showGraphs());
    }

    // Update is called once per frame
    void Update()
    {
    }

    void Destroy()
    {
    }

    public void showGraphs()
    {
        _graphsActive = !_graphsActive;


        if (_graphsActive)
        {
            var data = _netMan.GetEnergyData();

            foreach (var token in SceneManager.GetActiveScene().GetRootGameObjects().Where(ob =>
                ob.name.Contains("Windmill") || ob.name.Contains("SPV") || name.Contains("BAT")))
            {
                addGraphToScene(token);
            }
        }
        else
        {
            foreach (var token in SceneManager.GetActiveScene().GetRootGameObjects()
                .Where(ob => ob.name.Contains("GenGraph")))
            {
                token.Destroy();
            }
        }
    }


    void addGraphToScene(GameObject ob)
    {
        // Get our already existing graph
        var graphCanvas = UnityEngine.Object.Instantiate(prefab, ob.transform.position, ob.transform.rotation);
        graphCanvas.GetComponent<TokenData>().Tok = ob.GetComponent<TokenData>().Tok;

        graphCanvas.name = "GenGraph" + ob.name;
        graphCanvas.SetActive(true);

        var graphScript = graphCanvas.GetComponent<AddPointsToLineRenderer>();
        graphScript.GraphTypeToDisplay = AddPointsToLineRenderer.GraphType.POWER_UNIT;

        graphCanvas.transform.parent = ob.transform;

        RectTransform obPos = gameObject.GetComponent<RectTransform>();
        //obPos.sizeDelta = new Vector2(0.1f, 0.1f);
        graphCanvas.transform.localPosition = new Vector3(0, 0, 0);
        graphCanvas.transform.localScale *= ob.GetComponent<TokenData>().Tok.Scale;
        graphCanvas.transform.rotation = Quaternion.Euler(-90, 0, 0);
    }

    void moveToGrapPosition()
    {
        // Choose a position
        var position = new Vector3(0, 0, 0);

        // 
        //GameObject prefab = UnityEngine.Object.Instantiate(prefab, position, Quaternion.identity);
        //gameObject.
    }
}