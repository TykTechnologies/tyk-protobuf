package coprocess;

import static io.grpc.MethodDescriptor.generateFullMethodName;
import static io.grpc.stub.ClientCalls.asyncBidiStreamingCall;
import static io.grpc.stub.ClientCalls.asyncClientStreamingCall;
import static io.grpc.stub.ClientCalls.asyncServerStreamingCall;
import static io.grpc.stub.ClientCalls.asyncUnaryCall;
import static io.grpc.stub.ClientCalls.blockingServerStreamingCall;
import static io.grpc.stub.ClientCalls.blockingUnaryCall;
import static io.grpc.stub.ClientCalls.futureUnaryCall;
import static io.grpc.stub.ServerCalls.asyncBidiStreamingCall;
import static io.grpc.stub.ServerCalls.asyncClientStreamingCall;
import static io.grpc.stub.ServerCalls.asyncServerStreamingCall;
import static io.grpc.stub.ServerCalls.asyncUnaryCall;
import static io.grpc.stub.ServerCalls.asyncUnimplementedStreamingCall;
import static io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall;

/**
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.16.1)",
    comments = "Source: coprocess_object.proto")
public final class DispatcherGrpc {

  private DispatcherGrpc() {}

  public static final String SERVICE_NAME = "coprocess.Dispatcher";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<coprocess.CoprocessObject.Object,
      coprocess.CoprocessObject.Object> getDispatchMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "Dispatch",
      requestType = coprocess.CoprocessObject.Object.class,
      responseType = coprocess.CoprocessObject.Object.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<coprocess.CoprocessObject.Object,
      coprocess.CoprocessObject.Object> getDispatchMethod() {
    io.grpc.MethodDescriptor<coprocess.CoprocessObject.Object, coprocess.CoprocessObject.Object> getDispatchMethod;
    if ((getDispatchMethod = DispatcherGrpc.getDispatchMethod) == null) {
      synchronized (DispatcherGrpc.class) {
        if ((getDispatchMethod = DispatcherGrpc.getDispatchMethod) == null) {
          DispatcherGrpc.getDispatchMethod = getDispatchMethod = 
              io.grpc.MethodDescriptor.<coprocess.CoprocessObject.Object, coprocess.CoprocessObject.Object>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(
                  "coprocess.Dispatcher", "Dispatch"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  coprocess.CoprocessObject.Object.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  coprocess.CoprocessObject.Object.getDefaultInstance()))
                  .setSchemaDescriptor(new DispatcherMethodDescriptorSupplier("Dispatch"))
                  .build();
          }
        }
     }
     return getDispatchMethod;
  }

  private static volatile io.grpc.MethodDescriptor<coprocess.CoprocessObject.Event,
      coprocess.CoprocessObject.EventReply> getDispatchEventMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "DispatchEvent",
      requestType = coprocess.CoprocessObject.Event.class,
      responseType = coprocess.CoprocessObject.EventReply.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<coprocess.CoprocessObject.Event,
      coprocess.CoprocessObject.EventReply> getDispatchEventMethod() {
    io.grpc.MethodDescriptor<coprocess.CoprocessObject.Event, coprocess.CoprocessObject.EventReply> getDispatchEventMethod;
    if ((getDispatchEventMethod = DispatcherGrpc.getDispatchEventMethod) == null) {
      synchronized (DispatcherGrpc.class) {
        if ((getDispatchEventMethod = DispatcherGrpc.getDispatchEventMethod) == null) {
          DispatcherGrpc.getDispatchEventMethod = getDispatchEventMethod = 
              io.grpc.MethodDescriptor.<coprocess.CoprocessObject.Event, coprocess.CoprocessObject.EventReply>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(
                  "coprocess.Dispatcher", "DispatchEvent"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  coprocess.CoprocessObject.Event.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  coprocess.CoprocessObject.EventReply.getDefaultInstance()))
                  .setSchemaDescriptor(new DispatcherMethodDescriptorSupplier("DispatchEvent"))
                  .build();
          }
        }
     }
     return getDispatchEventMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static DispatcherStub newStub(io.grpc.Channel channel) {
    return new DispatcherStub(channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static DispatcherBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    return new DispatcherBlockingStub(channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static DispatcherFutureStub newFutureStub(
      io.grpc.Channel channel) {
    return new DispatcherFutureStub(channel);
  }

  /**
   */
  public static abstract class DispatcherImplBase implements io.grpc.BindableService {

    /**
     */
    public void dispatch(coprocess.CoprocessObject.Object request,
        io.grpc.stub.StreamObserver<coprocess.CoprocessObject.Object> responseObserver) {
      asyncUnimplementedUnaryCall(getDispatchMethod(), responseObserver);
    }

    /**
     */
    public void dispatchEvent(coprocess.CoprocessObject.Event request,
        io.grpc.stub.StreamObserver<coprocess.CoprocessObject.EventReply> responseObserver) {
      asyncUnimplementedUnaryCall(getDispatchEventMethod(), responseObserver);
    }

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
          .addMethod(
            getDispatchMethod(),
            asyncUnaryCall(
              new MethodHandlers<
                coprocess.CoprocessObject.Object,
                coprocess.CoprocessObject.Object>(
                  this, METHODID_DISPATCH)))
          .addMethod(
            getDispatchEventMethod(),
            asyncUnaryCall(
              new MethodHandlers<
                coprocess.CoprocessObject.Event,
                coprocess.CoprocessObject.EventReply>(
                  this, METHODID_DISPATCH_EVENT)))
          .build();
    }
  }

  /**
   */
  public static final class DispatcherStub extends io.grpc.stub.AbstractStub<DispatcherStub> {
    private DispatcherStub(io.grpc.Channel channel) {
      super(channel);
    }

    private DispatcherStub(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected DispatcherStub build(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      return new DispatcherStub(channel, callOptions);
    }

    /**
     */
    public void dispatch(coprocess.CoprocessObject.Object request,
        io.grpc.stub.StreamObserver<coprocess.CoprocessObject.Object> responseObserver) {
      asyncUnaryCall(
          getChannel().newCall(getDispatchMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void dispatchEvent(coprocess.CoprocessObject.Event request,
        io.grpc.stub.StreamObserver<coprocess.CoprocessObject.EventReply> responseObserver) {
      asyncUnaryCall(
          getChannel().newCall(getDispatchEventMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   */
  public static final class DispatcherBlockingStub extends io.grpc.stub.AbstractStub<DispatcherBlockingStub> {
    private DispatcherBlockingStub(io.grpc.Channel channel) {
      super(channel);
    }

    private DispatcherBlockingStub(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected DispatcherBlockingStub build(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      return new DispatcherBlockingStub(channel, callOptions);
    }

    /**
     */
    public coprocess.CoprocessObject.Object dispatch(coprocess.CoprocessObject.Object request) {
      return blockingUnaryCall(
          getChannel(), getDispatchMethod(), getCallOptions(), request);
    }

    /**
     */
    public coprocess.CoprocessObject.EventReply dispatchEvent(coprocess.CoprocessObject.Event request) {
      return blockingUnaryCall(
          getChannel(), getDispatchEventMethod(), getCallOptions(), request);
    }
  }

  /**
   */
  public static final class DispatcherFutureStub extends io.grpc.stub.AbstractStub<DispatcherFutureStub> {
    private DispatcherFutureStub(io.grpc.Channel channel) {
      super(channel);
    }

    private DispatcherFutureStub(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected DispatcherFutureStub build(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      return new DispatcherFutureStub(channel, callOptions);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<coprocess.CoprocessObject.Object> dispatch(
        coprocess.CoprocessObject.Object request) {
      return futureUnaryCall(
          getChannel().newCall(getDispatchMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<coprocess.CoprocessObject.EventReply> dispatchEvent(
        coprocess.CoprocessObject.Event request) {
      return futureUnaryCall(
          getChannel().newCall(getDispatchEventMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_DISPATCH = 0;
  private static final int METHODID_DISPATCH_EVENT = 1;

  private static final class MethodHandlers<Req, Resp> implements
      io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
    private final DispatcherImplBase serviceImpl;
    private final int methodId;

    MethodHandlers(DispatcherImplBase serviceImpl, int methodId) {
      this.serviceImpl = serviceImpl;
      this.methodId = methodId;
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        case METHODID_DISPATCH:
          serviceImpl.dispatch((coprocess.CoprocessObject.Object) request,
              (io.grpc.stub.StreamObserver<coprocess.CoprocessObject.Object>) responseObserver);
          break;
        case METHODID_DISPATCH_EVENT:
          serviceImpl.dispatchEvent((coprocess.CoprocessObject.Event) request,
              (io.grpc.stub.StreamObserver<coprocess.CoprocessObject.EventReply>) responseObserver);
          break;
        default:
          throw new AssertionError();
      }
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public io.grpc.stub.StreamObserver<Req> invoke(
        io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        default:
          throw new AssertionError();
      }
    }
  }

  private static abstract class DispatcherBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    DispatcherBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return coprocess.CoprocessObject.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("Dispatcher");
    }
  }

  private static final class DispatcherFileDescriptorSupplier
      extends DispatcherBaseDescriptorSupplier {
    DispatcherFileDescriptorSupplier() {}
  }

  private static final class DispatcherMethodDescriptorSupplier
      extends DispatcherBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final String methodName;

    DispatcherMethodDescriptorSupplier(String methodName) {
      this.methodName = methodName;
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.MethodDescriptor getMethodDescriptor() {
      return getServiceDescriptor().findMethodByName(methodName);
    }
  }

  private static volatile io.grpc.ServiceDescriptor serviceDescriptor;

  public static io.grpc.ServiceDescriptor getServiceDescriptor() {
    io.grpc.ServiceDescriptor result = serviceDescriptor;
    if (result == null) {
      synchronized (DispatcherGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new DispatcherFileDescriptorSupplier())
              .addMethod(getDispatchMethod())
              .addMethod(getDispatchEventMethod())
              .build();
        }
      }
    }
    return result;
  }
}
