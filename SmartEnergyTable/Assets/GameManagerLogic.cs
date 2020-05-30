using Network;
using System;
using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class GameManagerLogic : MonoBehaviour
{
    private NetworkManager _netMan;
    private string _id = Guid.NewGuid().ToString();

    // Start is called before the first frame update
    void Start()
    {
        try
        {
            _netMan = GameObject.Find("GameManager").GetComponent<NetworkManager>();
            _netMan.ObserveMaster(_id, (state) =>
            {
                Debug.Log("Steet: " + state);
            });

            _netMan.ObserveEnergyData(_id, (ener) =>
            {
                Debug.Log("Got this: " + ener.EnergyUsers[1].Pv);
            });

            EnergyData en = _netMan.GetEnergyData();

            _netMan.AddToken("Windmill", 90, new Vector3(50, 0, 0), 1);
        } catch
        {

        }



    }

    // Update is called once per frame
    void Update()
    {
        
    }

    void Destroy()
    {

    }


}
