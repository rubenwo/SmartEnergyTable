using Mapbox.Json;
using Network;
using System;
using System.Collections;
using System.Collections.Generic;
using System.Linq;
using System.Reflection;
using System.Threading;
using UnityEngine;
using UnityEngine.Networking;

public class AddPointsToLineRenderer : MonoBehaviour
{
    public Color c1 = Color.yellow;
    public Color c2 = Color.red;

    public Color TextColor = Color.black;

    // Which Json value do you want to show? (Must exist else won't show)
    public string GraphPropertyName = "TotalDemand";
    private string DisplayPropertyName { get => GraphPropertyName + "FieldNumber"; }

    // Represents all our raw data
    private List<Vector3> _points = new List<Vector3>();
    //private List<float> _values = new List<float>(new float[] { 25, 50, 45, 18, 29, 12 });
    //private List<string> _names = new List<string>(new string[] { "Lab", "Helix", "Auditorium", "HVAC", "ICT", "Laplace" });

    private EnergyDemandHourly[] _energy;

    

    IEnumerator GetHourly()
    {
        UnityWebRequest www = UnityWebRequest.Get("http://localhost:8080/Hourly");
        yield return www.Send();

        if (www.isError)
        {
            Debug.Log(www.error);
        }
        else
        {
            _energy = JsonConvert.DeserializeObject<EnergyDemandHourly[]>(www.downloadHandler.text);
            Debug.Log(_energy[1]);

            //var hourlyData = UnityWebRequest.Get("http://localhost:8080/Hourly");
            //var monthlyData = UnityWebRequest.Get("http://localhost:8080/Monthly");
            //Type type = _energy[1].GetType();
            //FieldInfo field = type.GetField("TotalDemand");
            //object value = field.GetValue(_energy[1]);
            //float levelValue = (float)value;

            //foreach (var f in _energy[1].GetType().GetFields())
            //    Debug.Log(f.Name);


            float relX, relY, relZ;

            RectTransform b = gameObject.GetComponent<RectTransform>();

            relX = gameObject.transform.position.x - b.rect.width / 2;
            relY = gameObject.transform.position.y - b.rect.height / 2;
            relZ = 50;

            // Calculate Graph size
            _energy = _energy.Take(10).ToArray();
            Debug.Log(_energy[1].TotalDemand);
            Debug.Log(_energy[1].GetType().GetField(DisplayPropertyName).GetValue(_energy[1]).GetType());
            float maxY = _energy.Max(eData => (int)eData.GetType().GetField(DisplayPropertyName).GetValue(eData));
            Debug.Log("Max: " + maxY);
            float diffYPerX = b.rect.height / maxY * 0.9f;
            float diffX = b.rect.width / _energy.Length;

            short counter = 0;
            // Generate 4 points for each raw value
            foreach (var values in _energy)
            {
                var val = (int)values.GetType().GetField(DisplayPropertyName).GetValue(values);

                float startX = relX + counter * diffX;
                float endX = relX + counter * diffX + diffX;
                float upperY = relY + val * diffYPerX;

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
        }
    }

    // Start is called before the first frame update
    void Start()
    {
        StartCoroutine(GetHourly());



    }
    private void AddText(string text, string value, Vector3 start, Vector3 end)
    {
        // Create the Text GameObject.
        GameObject textGO = new GameObject("infolabel"+text);
        textGO.transform.parent = GameObject.Find("GraphCanvas").transform;
        var textMesh = textGO.AddComponent<TextMesh>();
        textMesh.fontSize = 30;
        textMesh.color = TextColor;
        textMesh.alignment = TextAlignment.Center;
        textMesh.font = Resources.GetBuiltinResource(typeof(Font), "Arial.ttf") as Font;

        var rTransf = (RectTransform)gameObject.transform;

        textMesh.transform.position = start;
        rTransf.sizeDelta = new Vector2(Math.Abs(end.x - start.x), Math.Abs(end.y - start.y));
        textMesh.text = text + ": " + value;

        //label.transform.position = start;
    }

    // Update is called once per frame
    void Update()
    {
        LineRenderer lineRenderer = GetComponent<LineRenderer>();
        int counter = 0;
        foreach (var point in _points)
        {
            lineRenderer.SetPosition(counter++, point);
        }

    }


}

//class EnergyUser {

//    public string Time { get; set; }
//    public string Label { get; set; }
//    public string Name { get;set; }

//	public string SourceId {get;set; }
//	public string TotalDemand { get;set; }
//	public string Lighting { get;set; }
//	public string HVAC { get;set; }
//	public string Appliances  { get;set; }
//	public string Lab   { get;set; }
//	public string PV   { get;set; }
//	public string Unit   { get;set; }
//}

//class EnergyDemandHourly {

//    public string Id { get; set; }
//    public string Date { get; set; }
//    public string Year { get;set; }
//    public string Month { get;set; }
//	public string Day { get;set; }
//	public string Hour { get;set; }
//	public string Minutes { get;set; }
//	public string SourceId { get;set; }	
//	public string ChannelId { get;set; }
//	public string Unit { get;set; }
//	public string TotalDemand { get;set; }
//	public string DeltaValue { get;set; }
//	public string SourceTag { get; set; }
//    public string ChannelTag { get; set; }
//    public string Label { get; set; }
//    public string Name { get; set; }
//    public string Height { get; set; }
//    public string Area { get; set; }
//    public string WindSpeed { get; set; }
//    public string Temperature { get; set; }
//    public string SolarRad { get; set; }
//    public string ElectricityPrice { get; set; }
//    public string supply { get; set; }
//    public string renewables { get; set; }
//}