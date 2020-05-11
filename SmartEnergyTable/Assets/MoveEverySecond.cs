using System.Collections;
using System.Collections.Generic;
using UnityEngine;

using UnityEngine.SceneManagement;

public class MoveEverySecond : MonoBehaviour
{

    public int Seconds { get; set; } 


    // Start is called before the first frame update
    void Start()
    {
        
    }

    // Update is called once per frame
    void Update()
    {
        

        this.Seconds++;

        if (this.Seconds > 30)
        {
            gameObject.transform.Translate(1, 0, 0);
            this.Seconds = 0;

        }
        //SceneManager.LoadScene("Launcher");
    }
}
