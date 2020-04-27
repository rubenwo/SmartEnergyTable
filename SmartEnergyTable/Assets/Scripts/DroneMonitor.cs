using System.Collections;
using System.Collections.Generic;
using System.IO.Ports;
using System.Threading.Tasks;
using UnityEngine;

public class DroneMonitor : MonoBehaviour
{
    private SerialPort _serialPort;


    // Start is called before the first frame update
    void Start()
    {
        try
        {
            _serialPort = new SerialPort("COM4", 9600, Parity.None);
        }

        catch
        {
            Debug.Log("Please select a COM port first.");
        }

        _serialPort.Open();
        Task.Run(() =>
        {
            while (true)
            {
                Debug.Log(_serialPort.ReadLine());
            }
        });
    }

    // Update is called once per frame
    void Update()
    {
    }
}