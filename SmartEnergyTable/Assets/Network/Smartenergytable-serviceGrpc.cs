// <auto-generated>
//     Generated by the protocol buffer compiler.  DO NOT EDIT!
//     source: smartenergytable-service.proto
// </auto-generated>
#pragma warning disable 0414, 1591
#region Designer generated code

using grpc = global::Grpc.Core;

public static partial class SmartEnergyTableService
{
  static readonly string __ServiceName = "SmartEnergyTableService";

  static readonly grpc::Marshaller<global::Empty> __Marshaller_Empty = grpc::Marshallers.Create((arg) => global::Google.Protobuf.MessageExtensions.ToByteArray(arg), global::Empty.Parser.ParseFrom);
  static readonly grpc::Marshaller<global::Room> __Marshaller_Room = grpc::Marshallers.Create((arg) => global::Google.Protobuf.MessageExtensions.ToByteArray(arg), global::Room.Parser.ParseFrom);
  static readonly grpc::Marshaller<global::RoomId> __Marshaller_RoomId = grpc::Marshallers.Create((arg) => global::Google.Protobuf.MessageExtensions.ToByteArray(arg), global::RoomId.Parser.ParseFrom);
  static readonly grpc::Marshaller<global::Update> __Marshaller_Update = grpc::Marshallers.Create((arg) => global::Google.Protobuf.MessageExtensions.ToByteArray(arg), global::Update.Parser.ParseFrom);
  static readonly grpc::Marshaller<global::GameObject> __Marshaller_GameObject = grpc::Marshallers.Create((arg) => global::Google.Protobuf.MessageExtensions.ToByteArray(arg), global::GameObject.Parser.ParseFrom);

  static readonly grpc::Method<global::Empty, global::Room> __Method_CreateRoom = new grpc::Method<global::Empty, global::Room>(
      grpc::MethodType.Unary,
      __ServiceName,
      "CreateRoom",
      __Marshaller_Empty,
      __Marshaller_Room);

  static readonly grpc::Method<global::RoomId, global::Update> __Method_JoinRoom = new grpc::Method<global::RoomId, global::Update>(
      grpc::MethodType.ServerStreaming,
      __ServiceName,
      "JoinRoom",
      __Marshaller_RoomId,
      __Marshaller_Update);

  static readonly grpc::Method<global::GameObject, global::Empty> __Method_AddGameObject = new grpc::Method<global::GameObject, global::Empty>(
      grpc::MethodType.Unary,
      __ServiceName,
      "AddGameObject",
      __Marshaller_GameObject,
      __Marshaller_Empty);

  /// <summary>Service descriptor</summary>
  public static global::Google.Protobuf.Reflection.ServiceDescriptor Descriptor
  {
    get { return global::SmartenergytableServiceReflection.Descriptor.Services[0]; }
  }

  /// <summary>Base class for server-side implementations of SmartEnergyTableService</summary>
  [grpc::BindServiceMethod(typeof(SmartEnergyTableService), "BindService")]
  public abstract partial class SmartEnergyTableServiceBase
  {
    public virtual global::System.Threading.Tasks.Task<global::Room> CreateRoom(global::Empty request, grpc::ServerCallContext context)
    {
      throw new grpc::RpcException(new grpc::Status(grpc::StatusCode.Unimplemented, ""));
    }

    public virtual global::System.Threading.Tasks.Task JoinRoom(global::RoomId request, grpc::IServerStreamWriter<global::Update> responseStream, grpc::ServerCallContext context)
    {
      throw new grpc::RpcException(new grpc::Status(grpc::StatusCode.Unimplemented, ""));
    }

    public virtual global::System.Threading.Tasks.Task<global::Empty> AddGameObject(global::GameObject request, grpc::ServerCallContext context)
    {
      throw new grpc::RpcException(new grpc::Status(grpc::StatusCode.Unimplemented, ""));
    }

  }

  /// <summary>Client for SmartEnergyTableService</summary>
  public partial class SmartEnergyTableServiceClient : grpc::ClientBase<SmartEnergyTableServiceClient>
  {
    /// <summary>Creates a new client for SmartEnergyTableService</summary>
    /// <param name="channel">The channel to use to make remote calls.</param>
    public SmartEnergyTableServiceClient(grpc::ChannelBase channel) : base(channel)
    {
    }
    /// <summary>Creates a new client for SmartEnergyTableService that uses a custom <c>CallInvoker</c>.</summary>
    /// <param name="callInvoker">The callInvoker to use to make remote calls.</param>
    public SmartEnergyTableServiceClient(grpc::CallInvoker callInvoker) : base(callInvoker)
    {
    }
    /// <summary>Protected parameterless constructor to allow creation of test doubles.</summary>
    protected SmartEnergyTableServiceClient() : base()
    {
    }
    /// <summary>Protected constructor to allow creation of configured clients.</summary>
    /// <param name="configuration">The client configuration.</param>
    protected SmartEnergyTableServiceClient(ClientBaseConfiguration configuration) : base(configuration)
    {
    }

    public virtual global::Room CreateRoom(global::Empty request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
    {
      return CreateRoom(request, new grpc::CallOptions(headers, deadline, cancellationToken));
    }
    public virtual global::Room CreateRoom(global::Empty request, grpc::CallOptions options)
    {
      return CallInvoker.BlockingUnaryCall(__Method_CreateRoom, null, options, request);
    }
    public virtual grpc::AsyncUnaryCall<global::Room> CreateRoomAsync(global::Empty request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
    {
      return CreateRoomAsync(request, new grpc::CallOptions(headers, deadline, cancellationToken));
    }
    public virtual grpc::AsyncUnaryCall<global::Room> CreateRoomAsync(global::Empty request, grpc::CallOptions options)
    {
      return CallInvoker.AsyncUnaryCall(__Method_CreateRoom, null, options, request);
    }
    public virtual grpc::AsyncServerStreamingCall<global::Update> JoinRoom(global::RoomId request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
    {
      return JoinRoom(request, new grpc::CallOptions(headers, deadline, cancellationToken));
    }
    public virtual grpc::AsyncServerStreamingCall<global::Update> JoinRoom(global::RoomId request, grpc::CallOptions options)
    {
      return CallInvoker.AsyncServerStreamingCall(__Method_JoinRoom, null, options, request);
    }
    public virtual global::Empty AddGameObject(global::GameObject request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
    {
      return AddGameObject(request, new grpc::CallOptions(headers, deadline, cancellationToken));
    }
    public virtual global::Empty AddGameObject(global::GameObject request, grpc::CallOptions options)
    {
      return CallInvoker.BlockingUnaryCall(__Method_AddGameObject, null, options, request);
    }
    public virtual grpc::AsyncUnaryCall<global::Empty> AddGameObjectAsync(global::GameObject request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
    {
      return AddGameObjectAsync(request, new grpc::CallOptions(headers, deadline, cancellationToken));
    }
    public virtual grpc::AsyncUnaryCall<global::Empty> AddGameObjectAsync(global::GameObject request, grpc::CallOptions options)
    {
      return CallInvoker.AsyncUnaryCall(__Method_AddGameObject, null, options, request);
    }
    /// <summary>Creates a new instance of client from given <c>ClientBaseConfiguration</c>.</summary>
    protected override SmartEnergyTableServiceClient NewInstance(ClientBaseConfiguration configuration)
    {
      return new SmartEnergyTableServiceClient(configuration);
    }
  }

  /// <summary>Creates service definition that can be registered with a server</summary>
  /// <param name="serviceImpl">An object implementing the server-side handling logic.</param>
  public static grpc::ServerServiceDefinition BindService(SmartEnergyTableServiceBase serviceImpl)
  {
    return grpc::ServerServiceDefinition.CreateBuilder()
        .AddMethod(__Method_CreateRoom, serviceImpl.CreateRoom)
        .AddMethod(__Method_JoinRoom, serviceImpl.JoinRoom)
        .AddMethod(__Method_AddGameObject, serviceImpl.AddGameObject).Build();
  }

  /// <summary>Register service method with a service binder with or without implementation. Useful when customizing the  service binding logic.
  /// Note: this method is part of an experimental API that can change or be removed without any prior notice.</summary>
  /// <param name="serviceBinder">Service methods will be bound by calling <c>AddMethod</c> on this object.</param>
  /// <param name="serviceImpl">An object implementing the server-side handling logic.</param>
  public static void BindService(grpc::ServiceBinderBase serviceBinder, SmartEnergyTableServiceBase serviceImpl)
  {
    serviceBinder.AddMethod(__Method_CreateRoom, serviceImpl == null ? null : new grpc::UnaryServerMethod<global::Empty, global::Room>(serviceImpl.CreateRoom));
    serviceBinder.AddMethod(__Method_JoinRoom, serviceImpl == null ? null : new grpc::ServerStreamingServerMethod<global::RoomId, global::Update>(serviceImpl.JoinRoom));
    serviceBinder.AddMethod(__Method_AddGameObject, serviceImpl == null ? null : new grpc::UnaryServerMethod<global::GameObject, global::Empty>(serviceImpl.AddGameObject));
  }

}
#endregion
