using UnityEngine;
using UnityEngine.UI;

namespace UI
{
    public class OverviewMenu : MonoBehaviour
    {
        public Button addTokenButton;

        private NetworkManager _networkManager;

        private string[] prefabs = {"Cube", "Sphere", "Capsule"};

        void Start()
        {
            _networkManager = GameObject.Find("GameManager").GetComponent<NetworkManager>();

            if (!_networkManager.IsMaster)
            {
                gameObject.SetActive(false);
            }

            addTokenButton.onClick.AddListener(() =>
            {
                var r = new System.Random();
                _networkManager.AddToken(prefabs[r.Next(0, 2)],
                    new UnityEngine.Vector3(r.Next(-5, 5), r.Next(-5, 5), r.Next(0, 5)));
            });
        }
    }
}