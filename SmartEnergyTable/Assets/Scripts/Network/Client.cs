using System.Threading.Tasks;
using Grpc.Core;
using UnityEditor;
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

        internal async Task JoinRoom(string roomId, string userId, UpdateCallback callback)
        {
            try
            {
                using (var call = _client.JoinRoom(new RoomUser {Id = roomId, UserId = userId}))
                {
                    while (true)
                    {
                        await call.ResponseStream.MoveNext();

                        var s = call.ResponseStream.Current;
                        if (s.Id != "-1") callback.Invoke(s);
                    }
                }
            }
            catch (RpcException e)
            {
                Debug.Log("RPC failed" + e);
                throw;
            }
        }

        internal Empty SaveRoom(Room room)
        {
            var empty = _client.SaveRoom(room);
            return empty;
        }

        internal Empty AddToken(string roomId, string userId, int index, UnityEngine.Vector3 position)
        {
            var empty = _client.AddToken(new Token
            {
                RoomUser = new RoomUser {Id = roomId, UserId = userId},
                ObjectIndex = index,
                Position = new Vector3
                {
                    X = position.x, Y = position.y, Z = position.z
                }
            });
            return empty;
        }

        internal Empty RemoveToken(string roomId, string userId, string uuid)
        {
            var empty = _client.RemoveToken(new Token
            {
                RoomUser = new RoomUser {Id = roomId, UserId = userId},
                ObjectId = uuid
            });
            return empty;
        }

        internal Empty MoveToken(string roomId, string userId, string uuid, UnityEngine.Vector3 position)
        {
            var empty = _client.MoveToken(new Token
            {
                RoomUser = new RoomUser {Id = roomId, UserId = userId},
                ObjectId = uuid,
                Position = new Vector3
                {
                    X = position.x, Y = position.y, Z = position.z
                }
            });
            return empty;
        }

        internal Empty ChangeScene(string roomId, string userId, int sceneId)
        {
            var empty = _client.ChangeScene(new Scene
            {
                RoomUser = new RoomUser {Id = roomId, UserId = userId},
                SceneId = sceneId
            });
            return empty;
        }

        internal Empty MoveUsers(string roomId, string userId, UnityEngine.Vector3 newPosition)
        {
            var empty = _client.MoveUsers(new UserPosition
            {
                RoomUser = new RoomUser
                {
                    Id = roomId,
                    UserId = userId
                },
                NewPosition = new Vector3
                {
                    X = newPosition.x, Y = newPosition.y, Z = newPosition.z
                }
            });
            return empty;
        }

        internal Empty LeaveRoom(string roomId, string userId)
        {
            var empty = _client.LeaveRoom(new RoomUser {Id = roomId, UserId = userId});
            return empty;
        }
    }
}