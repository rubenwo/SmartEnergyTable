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

    private Token Tok { get => this.GetComponent<TokenData>().Tok; set => this.GetComponent<TokenData>().Tok = value; }

    public GameObject prefab;
     
    // Start is called before the first frame update
    void Start()
    {
        try
        {
            _netMan = GameObject.Find("GameManager").GetComponent<NetworkManager>();
            _netMan.ObserveMaster(_id, (state) =>
            {
                //Debug.Log("Steet: " + state);
            });

            _netMan.ObserveEnergyData(_id, (ener) =>
            {
                //Debug.Log("Got this: " + ener.EnergyUsers[1].Pv);
            });

            _netMan.ObserveGeneratedEnergy(_id, (ener) =>
            {
                EnergyData en = _netMan.GetEnergyData();
                //Debug.Log("Got " + en.EnergyDemandHourly.Count + " items");
            });

        } catch
        {

        }

        _netMan.AddToken("Windmill", 1, new Vector3(10, 0, 0), 0.05f);

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

            foreach (var token in SceneManager.GetActiveScene().GetRootGameObjects().Where(ob => ob.name.Contains("Windmill") || ob.name.Contains("Solar Panel") || name.Contains("Battery")))
            {
                addGraphToScene((GameObject)token);
            }
        } else
        {
            foreach (var token in SceneManager.GetActiveScene().GetRootGameObjects().Where(ob => ob.name.Contains("GenGraph")))
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
        var obPos = ob.transform.position;
        ob.transform.localScale *= ob.GetComponent<TokenData>().Tok.Scale;
        obPos.y += 10;
        obPos.z += 10;

        graphCanvas.name = "GenGraph" + ob.name;
        graphCanvas.SetActive(true);

        var graphScript = graphCanvas.GetComponent<AddPointsToLineRenderer>();
        graphScript.GraphTypeToDisplay = AddPointsToLineRenderer.GraphType.POWER_UNIT;

        graphCanvas.transform.parent = this.gameObject.transform;

    }


}