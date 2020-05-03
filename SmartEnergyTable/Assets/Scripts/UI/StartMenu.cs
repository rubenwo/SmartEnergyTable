using Network;
using UnityEngine;
using UnityEngine.Android;
using UnityEngine.UI;
using ZXing;

namespace UI
{
    public class StartMenu : MonoBehaviour
    {
        public Button create;
        public Button join;
        public RawImage qrImage;


        private NetworkManager _networkManager;


        private WebCamTexture _camTexture;
        private bool drawQr;

        void Start()
        {
            if (!Permission.HasUserAuthorizedPermission(Permission.Camera))
                Permission.RequestUserPermission(Permission.Camera);

            _networkManager = GameObject.Find("GameManager").GetComponent<NetworkManager>();
            create.onClick.AddListener(() => _networkManager.CreateRoom());
            join.onClick.AddListener(() =>
            {
                qrImage.gameObject.SetActive(true);

                drawQr = true;
                Debug.Log("Join");
            });
            if (Application.platform == RuntimePlatform.Android)
            {
                Debug.Log(WebCamTexture.devices);
                _camTexture = new WebCamTexture();
                _camTexture.deviceName = WebCamTexture.devices[0].name;
                if (_camTexture != null)
                    _camTexture.Play();
                qrImage.texture = _camTexture;
            }
        }

        private void Update()
        {
            if (drawQr)
            {
                IBarcodeReader reader = new BarcodeReader();

                var result = reader.Decode(_camTexture.GetPixels32(), _camTexture.width, _camTexture.height);
                if (result != null)
                {
                    Debug.Log(result.Text);
                    drawQr = false;
                    _camTexture.Stop();
                    _networkManager.JoinRoom(result.Text);
                }
            }
        }
    }
}