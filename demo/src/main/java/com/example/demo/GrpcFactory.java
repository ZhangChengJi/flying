//package com.example.demo;
//
//
//import com.github.zhangchengji.proto.client.Client;
//import com.github.zhangchengji.proto.client.ClientServiceGrpc;
//import com.github.zhangchengji.proto.client.FlyingConfig;
//import com.github.zhangchengji.proto.common.Response;
//import com.google.protobuf.InvalidProtocolBufferException;
//import io.grpc.ManagedChannel;
//import io.grpc.Server;
//import io.grpc.netty.shaded.io.grpc.netty.GrpcSslContexts;
//import io.grpc.netty.shaded.io.grpc.netty.NettyChannelBuilder;
//import io.grpc.netty.shaded.io.netty.handler.ssl.SslContext;
//import io.grpc.netty.shaded.io.netty.handler.ssl.SslContextBuilder;
//import org.slf4j.Logger;
//import org.slf4j.LoggerFactory;
//import org.springframework.core.io.ClassPathResource;
//
//import javax.net.ssl.SSLException;
//import java.io.IOException;
//import java.io.InputStream;
//import java.util.concurrent.TimeUnit;
//
//public class GrpcFactory implements Grpc {
//
//    private static final Logger log = LoggerFactory
//            .getLogger(GrpcFactory.class);
//
//    private  ClientServiceGrpc.ClientServiceBlockingStub blockingStub;
//    private ManagedChannel channel;
//    private Server server;
//    private  Client client;
//    private  String appId;
//    private String address;
//    public GrpcFactory(String address,String appId){
//        this.appId = appId;
//        this.address=address;
//        initGrpcConnection(address);
//    }
//
//
//    public  void initGrpcConnection(String address) {
//        ClassPathResource classPathResource = new ClassPathResource("ca.pem");
//        InputStream stream = null;
//        try {
//            stream = classPathResource.getInputStream();
//        } catch (IOException e) {
//            e.printStackTrace();
//        }
//        SslContextBuilder builder = GrpcSslContexts.forClient();
//        SslContext sslContext = null;
//        try {
//            sslContext = builder.trustManager(stream).build();
//        } catch (SSLException e) {
//            e.printStackTrace();
//        }
////        ChannelCredentials channelCredentials=null;
////        try {
////            assert stream != null;
////            channelCredentials= TlsChannelCredentials.newBuilder().trustManager(stream).build();
////        } catch (IOException e) {
////            e.printStackTrace();
////        }
//
//        // 构建 Channel
//         this.channel = NettyChannelBuilder.forTarget(address).overrideAuthority("flying.com").sslContext(sslContext).build();
//
//        this.blockingStub = ClientServiceGrpc.newBlockingStub(channel);
//
//    }
//
//
//
//
//
//    public String getAppId() {
//        return appId;
//    }
//
//    public void setAppId(String appId) {
//        this.appId = appId;
//    }
//
//    private Client createClient(String namespace){
//        client= Client.newBuilder().setAppId(this.getAppId()).setNamespace(namespace).build();
//          return client;
//    }
//
//    @Override
//    public FlyingConfig config(String namespace) {
//        FlyingConfig flyingConfig = null;
//        Response resp= this.blockingStub.config(createClient(namespace));
//        if (resp.getCode() == 200) {
//            try {
//                flyingConfig = resp.getData().unpack(FlyingConfig.class);
//                if (flyingConfig==null){
//                    throw new RuntimeException("获取配置信息为null，请检查配置中心是否存在此项目");
//                }
//            } catch (InvalidProtocolBufferException e) {
//                e.printStackTrace();
//            }
//        }else if (resp.getCode() == 500) {
//            log.error(resp.getMessage());
//        }
//
//        return flyingConfig;
//    }
//
//    @Override
//    public boolean listener() {
//        Response resp = this.blockingStub.listener(client);
//       if (resp.getCode() ==304){
//           log.debug(resp.getMessage());
//           return false;
//       }else return resp.getCode() == 200;
//    }
//    @Override
//    public  void shutdown(){
//        channel.shutdown();
//        // fail the test if cannot gracefully shutdown
//        try {
//            assert channel.awaitTermination(5, TimeUnit.SECONDS) : "channel cannot be gracefully shutdown";
//        } catch (InterruptedException e) {
//            e.printStackTrace();
//        } finally {
//            channel.shutdownNow();
//        }
//
//    }
//}
