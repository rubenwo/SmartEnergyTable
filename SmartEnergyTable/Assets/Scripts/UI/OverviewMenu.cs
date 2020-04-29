using System.Threading.Tasks;
using UnityEngine;
using UnityEngine.UI;
using ZXing;
using ZXing.QrCode;

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

        #region TokenSelection

        public GameObject tokenSelectionPanel;
        public Button cubeButton;
        public Button sphereButton;
        public Button capsuleButton;

        #endregion

        #region QRCode

        public GameObject qrCodePanel;
        private bool _showQrCode;

        #endregion


        private NetworkManager _networkManager;

        private void Start()
        {
            _networkManager = GameObject.Find("GameManager").GetComponent<NetworkManager>();

            if (!_networkManager.IsMaster)
            {
                gameObject.SetActive(false);
            }

            addTokenButton.onClick.AddListener(() =>
            {
//                Task.Run(() =>
//                {
//                    for (int i = 0; i < 1000; i++)
//                    {
//                        var r = new System.Random();
//                        _networkManager.AddToken("Cube",
//                            new UnityEngine.Vector3(r.Next(-5, 5), r.Next(-5, 5), r.Next(0, 5)));
//                    }
//                });
                tokenSelectionPanel.SetActive(true);
            });

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
            shareSessionButton.onClick.AddListener(() =>
            {
                _showQrCode = !_showQrCode;
                qrCodePanel.SetActive(_showQrCode);
                qrCodePanel.GetComponent<RawImage>().texture = GenerateQrCode(_networkManager.SessionID);
            });
            stopSessionButton.onClick.AddListener(() => { Debug.Log("Stop Session"); });
            clearButton.onClick.AddListener(() => { _networkManager.ClearScene(); });
            moveUsersButton.onClick.AddListener(() => { Debug.Log("Move"); });
            changeSceneButton.onClick.AddListener(() => { _networkManager.LoadScene(2); });
        }

        private static Color32[] Encode(string textForEncoding,
            int width, int height)
        {
            var writer = new BarcodeWriter
            {
                Format = BarcodeFormat.QR_CODE,
                Options = new QrCodeEncodingOptions
                {
                    Height = height,
                    Width = width
                }
            };
            return writer.Write(textForEncoding);
        }

        private Texture2D GenerateQrCode(string text)
        {
            var encoded = new Texture2D(256, 256);
            var color32 = Encode(text, encoded.width, encoded.height);
            encoded.SetPixels32(color32);
            encoded.Apply();
            return encoded;
        }
    }
}