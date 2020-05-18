using System;
using System.Collections.Generic;
using System.Linq;
using UnityEngine;

public class AddPointsToLineRenderer : MonoBehaviour
{
    public Color c1 = Color.yellow;
    public Color c2 = Color.red;

    // Represents all our raw data
    List<Vector3> _points = new List<Vector3>();
    List<float> _values = new List<float>(new float[] { 25, 50, 45, 18, 29, 12 });

    List<Vector3> infoLabelPoints = new List<Vector3>();

    private int _counter = 0;

    public Material material;
    private Mesh tileMesh;

    // Start is called before the first frame update
    void Start()
    {
        float relX, relY, relZ;

        RectTransform b = gameObject.GetComponent<RectTransform>();

        relX = gameObject.transform.position.x - b.rect.width / 2;
        relY = gameObject.transform.position.y - b.rect.height / 2;
        relZ = 50;


        // Calculate Graph size
        float maxY = _values.Max(points => points);
        float diffYPerX = b.rect.height / maxY;
        float diffX = b.rect.width / _values.Count;

        short counter = 0;
        // Generate 4 points for each raw value
        foreach (int val in _values)
        {
            // Draw graph
            _points.Add(new Vector3(relX + counter * diffX, relY, relZ));
            _points.Add(new Vector3(relX + counter * diffX, relY + val * diffYPerX, relZ));
            _points.Add(new Vector3(relX + counter * diffX + diffX, relY + val * diffYPerX, relZ));
            _points.Add(new Vector3(relX + counter * diffX + diffX, relY, relZ));

            // Add right bottom and left top corners to the list
           // AddRect("Lab", (new Vector3(relX + counter * diffX, relY, relZ), new Vector3(relX + counter * diffX + diffX, relY + val * diffYPerX, relZ)));

            AddText(val.ToString(), new Vector3(relX + counter * diffX, relY + val * diffYPerX, relZ),
                                           new Vector3(relX + counter * diffX + diffX, relY + val * diffYPerX, relZ));
            
            // Draw info panel

            counter++;
        }

        LineRenderer lineRenderer = gameObject.GetComponent<LineRenderer>();
        lineRenderer.material = new Material(Shader.Find("Sprites/Default"));
        lineRenderer.widthMultiplier = 1f;
        lineRenderer.positionCount = _points.Count;

        // A simple 2 color gradient with a fixed alpha of 1.0f.
        float alpha = 1.0f;
        Gradient gradient = new Gradient();
        gradient.SetKeys(
            new GradientColorKey[] { new GradientColorKey(c1, 0.0f), new GradientColorKey(c2, 1.0f) },
            new GradientAlphaKey[] { new GradientAlphaKey(alpha, 0.0f), new GradientAlphaKey(alpha, 1.0f) }
        );
        lineRenderer.colorGradient = gradient;


    }

    //private void AddRect(string text, (Vector3, Vector3) p)
    //{
    //    GameObject textGO = new GameObject("colormesh-" + text);
    //    textGO.transform.parent = GameObject.Find("Plane (1)").transform;
    //    //var canvas = textGO.AddComponent<Canvas>();
    //    var render = textGO.AddComponent<MeshRenderer>();
    //    //render.transform.position = 
    //    render.material.color = Color.gray;

    //    //((RectTransform)textGO.transform).sizeDelta = new Vector2(p.Item2.x - p.Item1.x, p.Item2.y - p.Item1.y);
    //    textGO.transform.position = p.Item1;


    //}

    private void AddText(string text, Vector3 start, Vector3 end)
    {
        //Canvas myTextCanvas = gameObject.AddComponent<Canvas>();
        //myTextCanvas.name = text;
        //myTextCanvas.transform.rotation = myTextCanvas.transform.rotation;

        //Text myText = myTextCanvas.gameObject.transform.parent.gameObject.AddComponent<Text>();
        //myText.text = text;
        //Debug.Log(gameObject.transform.parent.gameObject.name);
        //gameObject.transform.parent.gameObject.name = "TEST";

        // Create the Text GameObject.
        GameObject textGO = new GameObject("infolabel"+text);
        textGO.transform.parent = GameObject.Find("Plane (1)").transform;
        var textMesh = textGO.AddComponent<TextMesh>();
        textMesh.fontSize = 30;
        textMesh.color = Color.red;
        textMesh.alignment = TextAlignment.Center;
        textMesh.font = Resources.GetBuiltinResource(typeof(Font), "Arial.ttf") as Font;

        var rTransf = (RectTransform)gameObject.transform;

        start.y += rTransf.rect.height/15;
        textMesh.transform.position = start;
        ((RectTransform)gameObject.transform).sizeDelta = new Vector2(end.x - start.x, end.y - start.y);
        textMesh.text = text;


        //label.transform.position = start;
    }

    // Update is called once per frame
    void Update()
    {
        LineRenderer lineRenderer = GetComponent<LineRenderer>();
        var t = Time.time;
        int counter = 0;
        foreach (var point in _points)
        {
            lineRenderer.SetPosition(counter++, point);
        }

    }

}
