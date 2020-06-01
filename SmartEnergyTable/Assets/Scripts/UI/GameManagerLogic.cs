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

            
            // Cheats
            _netMan.AddToken("Windmill", 90, new Vector3(50, 0, 0), 2);

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
