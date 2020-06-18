using Mapbox.Json;
using Network;
using System;
using System.Collections;
using System.Collections.Generic;
using System.Linq;
using System.Linq.Expressions;
using System.Reflection;
using System.Threading;
using TMPro;
using UnityEngine;
using UnityEngine.Networking;
using Microsoft.CSharp;
using NetworkManager = Network.NetworkManager;

public class AddPointsToLineRenderer : MonoBehaviour
{
    public Color c1 = Color.yellow;
    public Color c2 = Color.red;

    public Color TextColor = Color.black;

    private Vector3 _lastRootRotation;

    public enum GraphType
    {
        DAILY,
        MONTHLY,
        POWER_UNIT
    };

    public GraphType GraphTypeToDisplay = GraphType.POWER_UNIT;

    private NetworkManager _networkManager;

    // Which Json value do you want to show? (Must exist else won't show)
    public string GraphPropertyName = "TotalDemand";

    // Represents all our raw data
    private List<Vector3> _points = new List<Vector3>();

    private EnergyData EnergyDataStore;

    private Token Tok
    {
        get => this.GetComponent<TokenData>().Tok;
        set => this.GetComponent<TokenData>().Tok = value;
    }

    void GetHourly()
    {
        this.EnergyDataStore = _networkManager.GetEnergyData();

        // First entries are invalid and cause errors.
        //EnergyDataStore.removeWrongEntries();
        //EnergyDataStore.limitBy(10);

        List<object> data = new List<object>();

        switch (GraphTypeToDisplay)
        {
            case GraphType.DAILY:
                int c = 0;
                foreach (var a in EnergyDataStore.EnergyUsers)
                {
                    data.Add(a);

                    if (c++ > 10)
                        break;
                }

                ;
                break;
            case GraphType.MONTHLY:
                int d = 0;
                foreach (var a in EnergyDataStore.EnergyDemandHourly)
                {
                    data.Add(a);

                    if (d++ > 10)
                        break;
                }

                ;
                break;
            case GraphType.POWER_UNIT:
                {
                    foreach (var energy in _networkManager.GeneratedEnergy.Data)
                    {
                        if (energy.Token.ObjectId == Tok.ObjectId)
                        {
                            var name = gameObject.name;

                            if (name.Contains("Windmill")
                            ) // 365 * 24 * 1,500(kW) * .25 = 3,285,000 (Yearly output windmill per year) 
                                energy.Energy = (float)(3285000 * ((double)new System.Random().Next(1, 10) / 10));
                            else if (name.Contains("SPV")) // 500-550 kWh (Yearly output windmill per year) 
                                energy.Energy = (float)(550000 * ((double)new System.Random().Next(1, 10) / 10));
                            else if (name.Contains("BAT")
                            ) // 365 * 24 * 1,500(kW) * .25 = 3,285,000 (Yearly output windmill per year) 
                                energy.Energy = (float)(500000 * ((double)new System.Random().Next(1, 10) / 10));

                            data.Add(energy);
                        }
                    }

                    GraphPropertyName = "Energy";

                    break;
                }
        }

        // Set Title bar
        GameObject.Find("TitleBar").GetComponent<TextMeshPro>().text = GraphPropertyName;

        // Set current rotation to a variable so we can compare changes
        _lastRootRotation = gameObject.transform.eulerAngles;

        RectTransform b = gameObject.GetComponent<RectTransform>();

        float maxY = 0;

        foreach (var a in data)
        {
            try
            {
                float num = (float)Convert.ToDouble(a.GetType().GetProperty(GraphPropertyName).GetValue(a));
                // Get value and see if it's higher. Then make it our new highest number, if higher.
                if (num > maxY)
                    maxY = num;
            } catch (Exception e)
            {

            }

        }

        // Calculate desired sizes
        float diffYPerX = b.rect.height / maxY * 0.1f;
        float diffX = b.rect.width / data.Count;

        if (GraphTypeToDisplay == GraphType.POWER_UNIT)
        {
            dynamic thingy = data[0];

            string tokenName = _networkManager.getTokenNameById(thingy.Token.ObjectId);

            //_points.Add(new Vector3(0, 0, 0));
            //_points.Add(new Vector3(1, 0, diffYPerX * thingy.Energy));

            AddTextToPowerUnit(tokenName, thingy.Energy.ToString());
        }
        else
        {
            short counter = 0;
            // Generate 4 points for each raw value
            foreach (var values in data)
            {
                try
                {


                    var val = (float)Convert.ToDouble(values.GetType().GetProperty(GraphPropertyName).GetValue(values));

                    float startX = counter * diffX;
                    float endX = counter * diffX + diffX;
                    float upperZ = val * diffYPerX;

                    for (float c = startX; c < endX; c++)
                    {
                        // Draw graph
                        _points.Add(new Vector3(c, 0, 0));
                        _points.Add(new Vector3(c, 0, diffYPerX * val));
                        _points.Add(new Vector3(c + diffX, 0, diffYPerX * val));
                        _points.Add(new Vector3(c + diffX, 0, 0));
                    }


                    AddText(values.GetType().GetProperty("Name").GetValue(values).ToString(),
                        val.ToString(), new Vector3(0, 0, 0), new Vector3(0, 0, 0));


                    counter++;
                } catch (Exception e)
                {

                }
            }
        }

        LineRenderer lineRenderer = gameObject.GetComponent<LineRenderer>();
        lineRenderer.material = new Material(Shader.Find("Sprites/Default"));
        lineRenderer.widthMultiplier = 1f;
        lineRenderer.positionCount = _points.Count;
        //lineRenderer.transform.rotation = gameObject.transform.rotation;

        // A simple 2 color gradient with a fixed alpha of 1.0f.
        float alpha = 1.0f;
        Gradient gradient = new Gradient();
        gradient.SetKeys(
            new GradientColorKey[] {new GradientColorKey(c1, 0.0f), new GradientColorKey(c2, 1.0f)},
            new GradientAlphaKey[] {new GradientAlphaKey(alpha, 0.0f), new GradientAlphaKey(alpha, 1.0f)}
        );
        lineRenderer.colorGradient = gradient;

        //LineRenderer lineRenderer = GetComponent<LineRenderer>();
        int count = 0;
        foreach (var point in _points)
        {
            lineRenderer.SetPosition(count++, point);
        }
    }

    private void trk(Action ac)
    {
        try
        {
            ac.Invoke();
        }
        catch (Exception e)
        {
            Debug.Log(e.Message);
        }
    }

    // Start is called before the first frame update
    void Start()
    {
        _networkManager = GameObject.Find("GameManager").GetComponent<NetworkManager>();

        GetHourly();
    }

    private void AddText(string text, string value, Vector3 start, Vector3 end)
    {
        // Create the Text GameObject.
        GameObject textGO = new GameObject("infolabel" + text);
        textGO.transform.parent = gameObject.transform;
        var textMesh = textGO.AddComponent<TextMesh>();
        textMesh.fontSize = 20;
        textMesh.color = TextColor;
        textMesh.alignment = TextAlignment.Center;
        textMesh.font = Resources.GetBuiltinResource(typeof(Font), "Arial.ttf") as Font;

        var rTransf = (RectTransform) gameObject.transform;

        if (_networkManager.IsMaster)
            textGO.transform.position = new Vector3(0, 0, 20);
        else
            textGO.transform.position = new Vector3(0, 0, 0.2f);

        rTransf.sizeDelta = new Vector2(1f, 1f);
        textMesh.text = text + "\n" + value;
    }

    private void AddTextToPowerUnit(string text, string value)
    {
        // Create the Text GameObject.
        GameObject textGO = new GameObject("infolabel" + text);
        textGO.transform.parent = gameObject.transform;
        var textMesh = textGO.AddComponent<TextMesh>();
        textMesh.fontSize = 20;
        textMesh.color = TextColor;
        textMesh.alignment = TextAlignment.Center;
        textMesh.font = Resources.GetBuiltinResource(typeof(Font), "Arial.ttf") as Font;

        var rTransf = (RectTransform)gameObject.transform;
        var a = textGO.transform.position;

        var b = GameObject.Find("infolabel" + text);

        if (_networkManager.IsMaster)
        {
            textGO.transform.position = new Vector3(0, 0, 20);
            b.transform.localPosition = new Vector3(0, 0, 20);
            b.transform.localScale = new Vector3(1f, 1f, 1f);
        }
            
        else
        {
            textGO.transform.position = new Vector3(0, 0.2f, 0);
            b.transform.localPosition = new Vector3(0, 0, 1f);
            b.transform.localScale = new Vector3(1f, 1f, 1f);
        }
            




        //rTransf.sizeDelta = new Vector2(0.1f, 0.1f);
        textMesh.text = text + "\n" + value + " kW/y";
    }

    // Update is called once per frame
    void Update()
    {
        //Rotate();
    }

    void Rotate()
    {
        // Don't rotate if rotation wasnt changed
        if (_lastRootRotation == gameObject.transform.eulerAngles)
            return;

        var pivot = gameObject.GetComponent<Transform>();

        List<Vector3> newPoints = new List<Vector3>();
        foreach (var point in _points)
        {
            // Calculate new point
            var dir = point - pivot.position; // get point direction relative to pivot
            dir = Quaternion.Euler(pivot.eulerAngles) * dir; // rotate it
            var newPoint = dir + pivot.position; // calculate rotated point
            // Add to list of points
            newPoints.Add(newPoint);
        }

        _points = newPoints;

        LineRenderer lineRenderer = GetComponent<LineRenderer>();

        short counter = 0;
        lineRenderer.positionCount = 0;
        lineRenderer.positionCount = newPoints.Count;

        foreach (var point in newPoints)
            lineRenderer.SetPosition(counter++, point);

        _lastRootRotation = GameObject.Find("GraphCanvas").transform.eulerAngles;
    }
}