using System;
using System.Collections.Generic;
using System.Linq;
using UnityEngine;

public class AddPointsToLineRenderer : MonoBehaviour
{
    public Color c1 = Color.yellow;
    public Color c2 = Color.red;

    public Color TextColor = Color.blue;

    // Represents all our raw data
    private List<Vector3> _points = new List<Vector3>();
    private List<float> _values = new List<float>(new float[] { 25, 50, 45, 18, 29, 12 });
    private List<string> _names = new List<string>(new string[] { "Lab", "Helix", "Auditorium", "HVAC", "ICT", "Laplace" });


    // Start is called before the first frame update
    void Start()
    {
        LineRenderer lineRenderer = gameObject.GetComponent<LineRenderer>();
        lineRenderer.material = new Material(Shader.Find("Sprites/Default"));
        lineRenderer.widthMultiplier = 1f;
        lineRenderer.positionCount = _points.Count;
        lineRenderer.transform.rotation = gameObject.transform.rotation;

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
            float startX = relX + counter * diffX;
            float endX = relX + counter * diffX + diffX;

            for (float c = startX; c < endX; c++)
            {
                // Draw graph
                _points.Add(new Vector3(c, relY, relZ));
                _points.Add(new Vector3(c, relY + val * diffYPerX, relZ));
                _points.Add(new Vector3(c+1, relY + val * diffYPerX, relZ));
                _points.Add(new Vector3(c+1, relY, relZ));
            }

            AddText(val.ToString(), new Vector3(relX + counter * diffX, relY + val * diffYPerX, relZ),
                                           new Vector3(relX + counter * diffX + diffX, relY + val * diffYPerX, relZ));
            
            // Draw info panel

            counter++;
        }



        // A simple 2 color gradient with a fixed alpha of 1.0f.
        float alpha = 1.0f;
        Gradient gradient = new Gradient();
        gradient.SetKeys(
            new GradientColorKey[] { new GradientColorKey(c1, 0.0f), new GradientColorKey(c2, 1.0f) },
            new GradientAlphaKey[] { new GradientAlphaKey(alpha, 0.0f), new GradientAlphaKey(alpha, 1.0f) }
        );
        lineRenderer.colorGradient = gradient;


    }

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

        start.y += rTransf.rect.height/10;
        textMesh.transform.position = start;
        ((RectTransform)gameObject.transform).sizeDelta = new Vector2(end.x - start.x, end.y - start.y);
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
