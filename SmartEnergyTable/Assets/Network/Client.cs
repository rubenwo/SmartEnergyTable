using System.Threading.Tasks;
using Grpc.Core;
using UnityEngine;

namespace Network
{
    public delegate void UpdateCallback(Update update);

    public class Client
    {
        private readonly SmartEnergyTableService.SmartEnergyTableServiceClient _client;

        public Client(SmartEnergyTableService.SmartEnergyTableServiceClient client)
        {
            _client = client;
        }

        internal Room CreateRoom()
        {
            var room = _client.CreateRoom(new Empty());
            return room;
        }

        internal async Task JoinRoom(string roomId, UpdateCallback callback)
        {
            try
            {
                using (var call = _client.JoinRoom(new RoomId {Id = roomId}))
                {
                    while (await call.ResponseStream.MoveNext())
                    {
                        var s = call.ResponseStream.Current;
                        callback.Invoke(s);
                    }
                }
            }
            catch (RpcException e)
            {
                Debug.Log("RPC failed" + e);
                throw;
            }
        }

        internal Empty AddGameObject(string name, float posX, float posY, float posZ)
        {
            var empty = _client.AddGameObject(new GameObject {Name = name, PosX = posX, PosY = posY, PosZ = posZ});
            return empty;
        }
    }
}