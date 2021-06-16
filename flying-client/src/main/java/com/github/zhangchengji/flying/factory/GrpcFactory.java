package com.github.zhangchengji.flying.factory;

import com.github.zhangchengji.flying.constants.Constants;

import com.github.zhangchengji.flying.client.Client;
import com.github.zhangchengji.flying.common.Response;
import com.github.zhangchengji.flying.client.FlyingConfig;
import com.google.protobuf.InvalidProtocolBufferException;

import io.grpc.ChannelCredentials;
import io.grpc.ManagedChannel;
import io.grpc.Server;
import io.grpc.TlsChannelCredentials;
import io.grpc.netty.shaded.io.grpc.netty.GrpcSslContexts;
import io.grpc.netty.shaded.io.grpc.netty.NettyChannelBuilder;
import io.grpc.netty.shaded.io.netty.handler.ssl.SslContext;
import io.grpc.netty.shaded.io.netty.handler.ssl.SslContextBuilder;
import okhttp3.OkHttpClient;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.cloud.bootstrap.encrypt.KeyProperties;
import org.springframework.core.io.ClassPathResource;
import retrofit2.Call;
import retrofit2.Retrofit;
import retrofit2.adapter.rxjava.RxJavaCallAdapterFactory;
import retrofit2.converter.gson.GsonConverterFactory;

import javax.net.ssl.*;
import java.io.*;
import java.security.KeyStore;
import java.security.SecureRandom;
import java.security.cert.CertificateFactory;
import java.util.concurrent.TimeUnit;

public class GrpcFactory implements Grpc {

    private static final Logger log = LoggerFactory
            .getLogger(GrpcFactory.class);

  // private  ClientServiceGrpc.ClientServiceBlockingStub blockingStub;
    private ManagedChannel channel;
    private Server server;
    private  Client client;
    private  String appId;
    private String address;
    private Grpc grpc;
    public GrpcFactory(String address,String appId){
        this.appId = appId;
        this.address=address;
        initGrpcConnection(address);
    }


    public static Grpc initGrpcConnection(String address) {

        OkHttpClient okHttpClient = new OkHttpClient.Builder()
                .connectTimeout(5, TimeUnit.SECONDS)
                .readTimeout(10, TimeUnit.SECONDS)
                .sslSocketFactory(getSSLSocketFactory(), new CustomTrustManager())
                .hostnameVerifier(getHostnameVerifier())
                .build();

        Retrofit retrofit = new Retrofit.Builder()
                .baseUrl("https://localhost:8881")
                .client(okHttpClient)
                .addConverterFactory(new NullOnEmptyConverterFactory())
                .addConverterFactory(GsonConverterFactory.create())
                .addCallAdapterFactory(RxJavaCallAdapterFactory.create())
                .build();

        Grpc  grpc = retrofit.create(Grpc.class);
         return  grpc;

//
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
//        ChannelCredentials channelCredentials=null;
//        try {
//            assert stream != null;
//            channelCredentials= TlsChannelCredentials.newBuilder().trustManager(stream).build();
//        } catch (IOException e) {
//            e.printStackTrace();
//        }

        // 构建 Channel
//         this.channel = NettyChannelBuilder.forTarget(address).overrideAuthority(Constants.SERVER_NAME).sslContext(sslContext).build();
//
//        this.blockingStub = ClientServiceGrpc.newBlockingStub(channel);

    }





    public String getAppId() {
        return appId;
    }

    public void setAppId(String appId) {
        this.appId = appId;
    }

    private Client createClient(String namespace){
        client= Client.newBuilder().setAppId(this.getAppId()).setNamespace(namespace).build();
          return client;
    }
    public static HostnameVerifier getHostnameVerifier() {
        HostnameVerifier hostnameVerifier = new HostnameVerifier() {
            public boolean verify(String hostname, SSLSession session) {
                return true;
            }
        };
        return hostnameVerifier;
    }
    private static SSLSocketFactory getSSLSocketFactory() {
        SSLContext sslContext = null;
        try {
            KeyStore keyStore = KeyStore.getInstance(KeyStore.getDefaultType());
            ClassPathResource classPathResource = new ClassPathResource("ca.pem");
        InputStream stream = null;
        try {
            stream = classPathResource.getInputStream();
        } catch (IOException e) {
            e.printStackTrace();
        }
            CertificateFactory certificateFactory;
            certificateFactory = CertificateFactory.getInstance("X.509");
            keyStore.load(null, null);
            keyStore.setCertificateEntry("flying.com",certificateFactory.generateCertificate(stream));
             sslContext = SSLContext.getInstance("TLS");
            TrustManagerFactory trustManagerFactory = TrustManagerFactory.getInstance(TrustManagerFactory.getDefaultAlgorithm());
            trustManagerFactory.init(keyStore);
            sslContext.init(null, trustManagerFactory.getTrustManagers(), new SecureRandom());
            return sslContext.getSocketFactory();
        } catch (Exception e) {
            e.printStackTrace();
        }
        return sslContext.getSocketFactory();
    }

    @Override
    public Call<Response> config(Client client) {

        return null;
    }

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

  //  }
}
