using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.UI;

namespace UI
{
    public class StartMenu : MonoBehaviour
    {
        public Button create;
        public Button join;

        void Start()
        {
            create.onClick.AddListener(() =>
                GameObject.Find("GameManager").GetComponent<NetworkManager>().CreateRoom());
            join.onClick.AddListener(() => Debug.Log("Join"));
        }
    }
}