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

public class AddPointsToLineRenderer : MonoBehaviour
{
    public Color c1 = Color.yellow;
    public Color c2 = Color.red;

    public Color TextColor = Color.black;

    public enum GraphType { DAILY, MONTHLY };
    public GraphType GraphTypeToDisplay = GraphType.DAILY;

    // Which Json value do you want to show? (Must exist else won't show)
    public string GraphPropertyName = "TotalDemand";

    // Represents all our raw data
    private List<Vector3> _points = new List<Vector3>();

    private EnergyDataContainer EnergyDataStore;

    IEnumerator GetHourly()
    {
        UnityWebRequest www = UnityWebRequest.Get("https://smartenergytable.rubenwoldhuis.nl/energydata");
        www.certificateHandler = null;

        yield return www.Send();

        if (www.isNetworkError)
        {
            Debug.Log(www.error);
        }
        else
        {

            EnergyDataStore = JsonConvert.DeserializeObject<EnergyDataContainer>(www.downloadHandler.text);
            Debug.Log(EnergyDataStore.EnergyUser[0].Name);
            // First entries are invalid and cause errors.
            EnergyDataStore.removeWrongEntries();
            EnergyDataStore.limitBy(10);

            dynamic data = new List<object>();

            switch (GraphTypeToDisplay)
            {
                case GraphType.DAILY: data = EnergyDataStore.EnergyUser; break;
                case GraphType.MONTHLY: data = EnergyDataStore.EnergyDemand; break;
            }

            // Set Title bar
            GameObject.Find("TitleBar").GetComponent<TextMeshPro>().text = GraphPropertyName;

            float relX, relY, relZ;

            RectTransform b = gameObject.GetComponent<RectTransform>();

            relX = gameObject.transform.position.x - b.rect.width / 2;
            relY = gameObject.transform.position.y - b.rect.height / 2;
            relZ = 50;

            int maxY = 0;

            foreach (var a in data)
            {
                int num = (int)Double.Parse(a.GetType().GetProperty(GraphPropertyName).GetValue(a));
                // Get value and see if it's higher. Then make it our new highest number, if higher.
                Debug.Log(num);

                if (num > maxY)
                    maxY = num;
            }

            // Calculate desired sizes
            float diffYPerX = b.rect.height / maxY * 0.9f;
            float diffX = b.rect.width / data.Count;

            short counter = 0;
            // Generate 4 points for each raw value
            foreach (var values in data)
            {
                var val = (float)Convert.ToDouble(values.GetType().GetProperty(GraphPropertyName).GetValue(values));

                float startX = relX + counter * diffX;
                float endX = relX + counter * diffX + diffX;
                float upperY = relY + val * diffYPerX;

                //Debug.Log("Start");
                //Debug.Log(diffYPerX);
                //Debug.Log(diffX);
                //Debug.Log(startX);
                //Debug.Log(endX);
                //Debug.Log(upperY);

                for (float c = startX; c < endX; c++)
                {
                    // Draw graph
                    _points.Add(new Vector3(c, relY, relZ));
                    _points.Add(new Vector3(c, upperY, relZ));
                    _points.Add(new Vector3(c + 1, upperY, relZ));
                    _points.Add(new Vector3(c + 1, relY, relZ));
                }

                //Add text above our graph bar here
                AddText(values.Name, val.ToString(), new Vector3(relX + counter * diffX, relY + val * diffYPerX + 10, relZ),
                                                    new Vector3(relX + counter * diffX + diffX, relY + val * diffYPerX + b.rect.height / 10, relZ));

                counter++;
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
                new GradientColorKey[] { new GradientColorKey(c1, 0.0f), new GradientColorKey(c2, 1.0f) },
                new GradientAlphaKey[] { new GradientAlphaKey(alpha, 0.0f), new GradientAlphaKey(alpha, 1.0f) }
            );
            lineRenderer.colorGradient = gradient;

            //LineRenderer lineRenderer = GetComponent<LineRenderer>();
            counter = 0;
            foreach (var point in _points)
            {
                lineRenderer.SetPosition(counter++, point);
            }
        }


    }

    // Start is called before the first frame update
    void Start()
    {
        StartCoroutine(GetHourly());
    }

    void calculateAverage(string unitName, string property)
    {

    }

    private void AddText(string text, string value, Vector3 start, Vector3 end)
    {
        // Create the Text GameObject.
        GameObject textGO = new GameObject("infolabel"+text);
        textGO.transform.parent = GameObject.Find("GraphCanvas").transform;
        var textMesh = textGO.AddComponent<TextMesh>();
        textMesh.fontSize = 20;
        textMesh.color = TextColor;
        textMesh.alignment = TextAlignment.Center;
        textMesh.font = Resources.GetBuiltinResource(typeof(Font), "Arial.ttf") as Font;

        var rTransf = (RectTransform)gameObject.transform;

        textMesh.transform.position = start;
        rTransf.sizeDelta = new Vector2(Math.Abs(end.x - start.x), Math.Abs(end.y - start.y));
        textMesh.text = text + "\n" + value;

    }

    // Update is called once per frame
    void Update()
    {


    }


}
// Class used to deserialize complex json array
public class EnergyDataContainer
{
    public List<EnergyUser2> EnergyUser { get; set; }
    public List<EnergyDemandHourly2> EnergyDemand { get; set; }

    public EnergyDataContainer()
    {
        EnergyUser = new List<EnergyUser2>();
        EnergyDemand = new List<EnergyDemandHourly2>();
    }

    public void removeWrongEntries()
    {
        //Debug.Log(EnergyUser.FindIndex(0, EnergyUser.Count, (user) => user.Name == "Laplace"));
        EnergyUser = EnergyUser.Where(user => user.Name != "Name").ToList();
        EnergyDemand = EnergyDemand.Where(user => user.Name != "Name").ToList();

        limitBy(10);
    }

    internal void limitBy(int v)
    {
        EnergyUser = EnergyUser.Take(v).ToList();
        EnergyDemand = EnergyDemand.Take(v).ToList();
    }
}

// Classes are also loaded in gRPC but serializing to these classes seems to give our objects other variablenames for some reason
// Quick-fix: Use 2 new classes and serialize them into these.
public class EnergyUser2
{
    public string Time { get; set; }
    public string Label { get; set; }
    public string Name { get; set; }

    public string SourceId { get; set; }
    public string TotalDemand { get; set; }
    public string Lighting { get; set; }
    public string HVAC { get; set; }
    public string Appliances { get; set; }
    public string Lab { get; set; }
    public string PV { get; set; }
    public string Unit { get; set; }

    public override string ToString()
    {
        return TotalDemand.ToString();
    }
}

public class EnergyDemandHourly2
{

    public string Id { get; set; }
    public string Date { get; set; }
    public string Year { get; set; }
    public string Month { get; set; }
    public string Day { get; set; }
    public string Hour { get; set; }
    public string Minutes { get; set; }
    public string SourceId { get; set; }
    public string ChannelId { get; set; }
    public string Unit { get; set; }
    public string TotalDemand { get; set; }
    public string DeltaValue { get; set; }
    public string SourceTag { get; set; }
    public string ChannelTag { get; set; }
    public string Label { get; set; }
    public string Name { get; set; }
    public string Height { get; set; }
    public string Area { get; set; }
    public string WindSpeed { get; set; }
    public string Temperature { get; set; }
    public string SolarRad { get; set; }
    public string ElectricityPrice { get; set; }
    public string supply { get; set; }
    public string renewables { get; set; }

    public override string ToString()
    {
        return TotalDemand.ToString();
    }
}