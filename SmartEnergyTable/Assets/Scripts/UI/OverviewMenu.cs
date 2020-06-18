using System;
using System.Collections.Generic;
using Mapbox.Json.Bson;
using Network;
using TMPro;
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

        #endregion

        #region QRCode

        public GameObject qrCodePanel;
        private bool _showQrCode;

        #endregion


        public Button PrefabButton;

        private enum State
        {
            Idle,
            PlacingToken,
            RemovingToken,
            MovingToken,
            SelectedTokenForMoving,
            MovingUser
        }

        private string _prefab = "Cube";

        private int _efficiency = 0;

        //private GameObject test;
        private readonly string _uuid = Guid.NewGuid().ToString();

        private NetworkManager _networkManager;
        private State _state = State.Idle;
        private RaycastHit _selectedToken;

        private readonly List<Button> _buttons = new List<Button>();
        public InputField efficiencyInputField;

        private void Start()
        {
            _networkManager = GameObject.Find("GameManager").GetComponent<NetworkManager>();
            _networkManager.ObserveMaster(_uuid, isMaster => gameObject.SetActive(isMaster));

            var pos = tokenSelectionPanel.transform.position;
            pos.x -= _networkManager.Prefabs.Count / 2 * 300;

            for (var i = 0; i < _networkManager.Prefabs.Count; i++)
            {
                var button = Instantiate(PrefabButton, pos, Quaternion.identity,
                    tokenSelectionPanel.transform) as Button;
                button.GetComponentInChildren<TextMeshProUGUI>().text = _networkManager.Prefabs[i];
                button.onClick.AddListener(() =>
                {
                    _prefab = button.GetComponentInChildren<TextMeshProUGUI>().text;
                    var e = efficiencyInputField.text;
                    _efficiency = e != "" ? int.Parse(e) : 0;
                    tokenSelectionPanel.SetActive(false);
                });
                _buttons.Add(button);
                pos.x += 300;
            }

            addTokenButton.onClick.AddListener(() =>
            {
                _state = State.PlacingToken;
                tokenSelectionPanel.SetActive(true);
            });
            foreach (var data in _networkManager.GeneratedEnergy.Data)
            {
                Debug.Log(data.Energy);
            }

            removeTokenButton.onClick.AddListener(() => { _state = State.RemovingToken; });
            moveTokenButton.onClick.AddListener(() => { _state = State.MovingToken; });

            saveSessionButton.onClick.AddListener(() => { _networkManager.SaveRoom(); });
            leaveSessionButton.onClick.AddListener(() => { _networkManager.LeaveRoom(); });
            shareSessionButton.onClick.AddListener(() =>
            {
                _showQrCode = !_showQrCode;
                qrCodePanel.SetActive(_showQrCode);
                if (_showQrCode)
                    qrCodePanel.GetComponent<RawImage>().texture = GenerateQrCode(_networkManager.SessionID);
            });
            stopSessionButton.onClick.AddListener(() =>
            {
                _networkManager.SaveRoom();
                _networkManager.LeaveRoom();
            });
            clearButton.onClick.AddListener(() => { _networkManager.ClearScene(); });
            moveUsersButton.onClick.AddListener(() => { _state = State.MovingUser; });


            //changeSceneButton.onClick.AddListener(() => { _networkManager.LoadScene(2); });
        }


        private void Update()
        {
            if (!_networkManager.IsMaster || tokenSelectionPanel.activeSelf)
                return;
            if (!Input.GetMouseButtonDown(0))
                return;


            RaycastHit hit;
            bool ok;
            switch (_state)
            {
                case State.Idle:
                    break;
                case State.PlacingToken:
                    (hit, ok) = Select();
                    if (ok)
                    {
                        _networkManager.AddToken(_prefab, _efficiency, hit.point, 0.005f);
                        _state = State.Idle;
                    }

                    break;
                case State.RemovingToken:
                    (hit, ok) = Select();
                    if (ok)
                    {
                        _networkManager.RemoveToken(hit.transform.gameObject);
                        _state = State.Idle;
                    }

                    break;
                case State.MovingToken:
                    (hit, ok) = Select();
                    if (ok)
                    {
                        _selectedToken = hit;
                        _state = State.SelectedTokenForMoving;
                    }

                    break;
                case State.SelectedTokenForMoving:
                    (hit, ok) = Select();
                    if (ok)
                    {
                        _networkManager.MoveToken(_selectedToken.transform.gameObject, hit.point);
                        _state = State.Idle;
                    }

                    break;
                case State.MovingUser:
                    (hit, ok) = Select();
                    if (ok)
                    {
                        _networkManager.MoveUsers(hit.point);
                        _state = State.Idle;
                    }

                    break;
            }
        }

        private void OnDestroy()
        {
            _buttons.ForEach(button => Destroy(button));
            _networkManager.UnObserveMaster(_uuid);
        }

        private (RaycastHit, bool) Select()
        {
            if (Camera.main == null)
                return (new RaycastHit(), false);
            var ray = Camera.main.ScreenPointToRay(Input.mousePosition);
            if (Physics.Raycast(ray, out var hit, 100.0f))
            {
                return (hit, true);
            }

            return (new RaycastHit(), false);
        }

        #region QrEncoder

        /*
         * Encode is a function that returns a Color32[] containing the data for the QR code.
         * @param textForEncoding: is a string containing the text that needs to be encoded to a QR Code.
         * @param width: integer for the width of the texture.
         * @param height: integer for the height of the texture.
         */
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

        /*
         * GenerateQrCode is a function to create a Texture2D as a QR code from a string.
         * @param text: the string that is to be encoded to a QR Code.
         */
        private static Texture2D GenerateQrCode(string text)
        {
            var encoded = new Texture2D(256, 256);
            var color32 = Encode(text, encoded.width, encoded.height);
            encoded.SetPixels32(color32);
            encoded.Apply();
            return encoded;
        }

        #endregion
    }
}