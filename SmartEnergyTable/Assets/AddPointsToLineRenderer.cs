using System;
using System.Collections.Generic;
using System.Linq;
using UnityEngine;

public class AddPointsToLineRenderer : MonoBehaviour
{
    public Color c1 = Color.yellow;
    public Color c2 = Color.red;

    public Color TextColor = Color.black;

    // Represents all our raw data
    private List<Vector3> _points = new List<Vector3>();
    private List<float> _values = new List<float>(new float[] { 25, 50, 45, 18, 29, 12 });
    private List<string> _names = new List<string>(new string[] { "Lab", "Helix", "Auditorium", "HVAC", "ICT", "Laplace" });


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
        float diffYPerX = b.rect.height / maxY * 0.9f;
        float diffX = b.rect.width / _values.Count;

        short counter = 0;
        // Generate 4 points for each raw value
        foreach (int val in _values)
        {
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

            //AddGraphBar(new Vector3(startX, relY, relZ), new Vector3(endX, upperY, relZ));


            //Add text above our graph bar here
            AddText(val.ToString(), new Vector3(relX + counter * diffX, relY + val * diffYPerX + 10, relZ),
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

    //private void AddGraphBar(Vector3 leftBottom, Vector3 rightTop)
    //{
    //    // Create the Text GameObject.

    //    GameObject plane = GameObject.CreatePrimitive(PrimitiveType.Plane);

    //    // Position and rotation plane
    //    //plane.transform.position = leftBottom;
    //    var a = plane.AddComponent<RectTransform>();

    //    //GameObject textGO = new GameObject("bars");
    //    plane.transform.parent = GameObject.Find("Plane (1)").transform;
    //    var mesh = plane.GetComponent<MeshRenderer>();
    //    //mesh.material = new Material(Shader.Find("Specular"));
    //    mesh.material.color = Color.blue;
    //    //textMesh.mater
    //    //textMesh.color = TextColor;
    //    //textMesh.alignment = TextAlignment.Center;
    //    //textMesh.font = Resources.GetBuiltinResource(typeof(Font), "Arial.ttf") as Font;

    //    //var rTransf = (RectTransform)plane.transform;
    //    a.transform.position = leftBottom;
    //    a.transform.rotation = GameObject.Find("Plane (1)").transform.rotation;
    //    //start.y += rTransf.rect.height / 10;
    //    //textMesh.transform.position = start;
    //    a.sizeDelta = new Vector2(rightTop.x - leftBottom.x, rightTop.y - leftBottom.y);
    //    //textMesh.text = _names.ElementAt(_values.IndexOf(int.Parse(text))) + "\n " + text;
        

    //}

    private void AddText(string text, Vector3 start, Vector3 end)
    {
        // Create the Text GameObject.
        GameObject textGO = new GameObject("infolabel"+text);
        textGO.transform.parent = GameObject.Find("Plane (1)").transform;
        var textMesh = textGO.AddComponent<TextMesh>();
        textMesh.fontSize = 30;
        textMesh.color = TextColor;
        textMesh.alignment = TextAlignment.Center;
        textMesh.font = Resources.GetBuiltinResource(typeof(Font), "Arial.ttf") as Font;

        var rTransf = (RectTransform)gameObject.transform;

        //start.y += rTransf.rect.height/10;
        textMesh.transform.position = start;
        rTransf.sizeDelta = new Vector2(Math.Abs(end.x - start.x), Math.Abs(end.y - start.y));
        textMesh.text = _names.ElementAt(_values.IndexOf(int.Parse(text))) + "\n "+ text;

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
