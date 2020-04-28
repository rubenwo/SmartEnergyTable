using System.Threading;
using System.Threading.Tasks;
using UnityEngine;
using UnityEngine.UI;

namespace UI
{
    public class OverviewMenu : MonoBehaviour
    {
        #region MainButtons

        public Button addTokenButton;
        public Button removeTokenButton;
        public Button moveTokenButton;
        public Button saveSessionButton;
        public Button leaveSessionButton;
        public Button shareSessionButton;
        public Button stopSessionButton;
        public Button clearButton;
        public Button moveUsersButton;
        public Button changeSceneButton;

        #endregion

        #region MyRegion

        public GameObject tokenSelectionPanel;
        public Button cubeButton;
        public Button sphereButton;
        public Button capsuleButton;

        #endregion


        private NetworkManager _networkManager;

        private void Start()
        {
            _networkManager = GameObject.Find("GameManager").GetComponent<NetworkManager>();

            if (!_networkManager.IsMaster)
            {
                gameObject.SetActive(false);
            }

            addTokenButton.onClick.AddListener(() => { tokenSelectionPanel.SetActive(true); });

            cubeButton.onClick.AddListener(() =>
            {
                var r = new System.Random();
                _networkManager.AddToken("Cube",
                    new UnityEngine.Vector3(r.Next(-5, 5), r.Next(-5, 5), r.Next(0, 5)));
                tokenSelectionPanel.SetActive(false);
            });
            sphereButton.onClick.AddListener(() =>
            {
                var r = new System.Random();
                _networkManager.AddToken("Sphere",
                    new UnityEngine.Vector3(r.Next(-5, 5), r.Next(-5, 5), r.Next(0, 5)));
                tokenSelectionPanel.SetActive(false);
            });
            capsuleButton.onClick.AddListener(() =>
            {
                var r = new System.Random();
                _networkManager.AddToken("Capsule",
                    new UnityEngine.Vector3(r.Next(-5, 5), r.Next(-5, 5), r.Next(0, 5)));
                tokenSelectionPanel.SetActive(false);
            });
            removeTokenButton.onClick.AddListener(() =>
            {
                //TODO: Implement remove token button.
                //_networkManager.RemoveToken();
                Debug.Log("Remove Token");
            });
            moveTokenButton.onClick.AddListener(() =>
            {
                //TODO: Implement Move token button
                //_networkManager.MoveToken();
                Debug.Log("Move Token");
            });

            saveSessionButton.onClick.AddListener(() => { _networkManager.SaveRoom(); });
            leaveSessionButton.onClick.AddListener(() => { _networkManager.LeaveRoom(); });
            shareSessionButton.onClick.AddListener(() => { Debug.Log("Share Session"); });
            stopSessionButton.onClick.AddListener(() => { Debug.Log("Stop Session"); });
            clearButton.onClick.AddListener(() => { _networkManager.ClearScene(); });
            moveUsersButton.onClick.AddListener(() => { Debug.Log("Move"); });
            changeSceneButton.onClick.AddListener(() => { _networkManager.LoadScene(2); });
        }
    }
}