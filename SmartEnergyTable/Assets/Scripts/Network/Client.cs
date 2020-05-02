using System.Threading.Tasks;
using Grpc.Core;
using UnityEngine;

namespace Network
{
    // this function is used to return the current update as soon as it's available.
    public delegate void UpdateCallback(Patch patch);

    public class Client
    {
        private readonly SmartEnergyTableService.SmartEnergyTableServiceClient _client;

        public Client(SmartEnergyTableService.SmartEnergyTableServiceClient client)
        {
            _client = client;
        }

        // RPC to create a new room
        internal RoomUser CreateRoom()
        {
            return _client.CreateRoom(new Empty());
        }

        /*
         * RPC to join an existing room.
         * @param roomId: the ID corresponding to the room.
         * @param userId: a GUID generated by the client to indicate the user in the room.
         * @param callback: a delegate function used to callback an update.
         */
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
                        callback.Invoke(call.ResponseStream.Current);
                    }
                }
            }
            catch (RpcException e)
            {
                Debug.Log("RPC failed" + e);
                throw;
            }
        }

        /*
         * RPC to save the room in the server.
         * @param room: the room contains the ID of the room that should be saved.
         */
        internal Empty SaveRoom(RoomUser user)
        {
            var empty = _client.SaveRoom(user);
            return empty;
        }

        /*
         * RPC to add a token to the current room.
         * @param roomId: the ID of the room where this function should be executed.
         * @param userId: the GUID generated by the client should be corresponding to the master in the room.
         * @param index: index in the objectLibrary that contains the prefabs.
         * @param position: UnityEngine version of the Vector3. This contains the x,y,z coordinates where the token should be placed
         */
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

        /*
         * RPC to remove a token from the current room.
         * @param roomId: the ID of the room where this function should be executed.
         * @param userId: the GUID generated by the client should be corresponding to the master in the room.
         * @param uuid: the uuid from the token. This uuid is generated server-side when the AddToken RPC is called.
         */
        internal Empty RemoveToken(string roomId, string userId, string uuid)
        {
            var empty = _client.RemoveToken(new Token
            {
                RoomUser = new RoomUser {Id = roomId, UserId = userId},
                ObjectId = uuid
            });
            return empty;
        }

        /*
         * RPC to move a token in the current room.
         * @param roomId: the ID of the room where this function should be executed.
         * @param userId: the GUID generated by the client should be corresponding to the master in the room.
         * @param uuid: the uuid from the token. This uuid is generated server-side when the AddToken RPC is called.
         * @param position: UnityEngine version of the Vector3. This contains the x,y,z coordinates where the token should be moved to.
         */
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

        /*
         * RPC to clear a room of all tokens.
         * @param roomId: the ID of the room where this function should be executed.
         * @param userId: the GUID generated by the client should be corresponding to the master in the room.
         */
        internal Empty ClearRoom(string roomId, string userId)
        {
            var empty = _client.ClearRoom(new RoomUser
            {
                Id = roomId,
                UserId = userId
            });
            return empty;
        }

        /*
         * RPC to change the scene in all clients
         * @param roomId: the ID of the room where this function should be executed.
         * @param userId: the GUID generated by the client should be corresponding to the master in the room.
         * @param sceneId: an integer indicated the sceneBuildIndex from Unity.
         */
        internal Empty ChangeScene(string roomId, string userId, int sceneId)
        {
            var empty = _client.ChangeScene(new Scene
            {
                RoomUser = new RoomUser {Id = roomId, UserId = userId},
                SceneId = sceneId
            });
            return empty;
        }

        /*
         * RPC to move all users to a position
         * @param roomId: the ID of the room where this function should be executed.
         * @param userId: the GUID generated by the client should be corresponding to the master in the room.
         * @param position: UnityEngine version of the Vector3. This contains the x,y,z coordinates where the users should be moved to.
         */
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

        /*
         * RPC to change the master of the room.
         * @param roomId: the ID of the room where this function should be executed.
         * @param userId: the GUID generated by the client should be corresponding to the master in the room.
         * @param newMasterId: the id of the new master.
         */
        internal Empty ChangeMaster(string roomId, string userId, string newMasterId)
        {
            var empty = _client.ChangeMaster(new MasterSwitch
                {Id = roomId, MasterId = userId, NewMasterId = newMasterId});
            return empty;
        }

        /*
         * RPC to leave the room
         * @param roomId: the ID corresponding to the room.
         * @param userId: a GUID generated by the client to indicate the user in the room.
         */
        internal Empty LeaveRoom(string roomId, string userId)
        {
            var empty = _client.LeaveRoom(new RoomUser {Id = roomId, UserId = userId});
            return empty;
        }
    }
}