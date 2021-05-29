/*
 * Copyright 2013-2018 the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package com.github.zhangchengji.flying.util;

import com.github.zhangchengji.flying.constants.Constants;
import com.github.zhangchengji.flying.constants.FlyingConfigProperties;
import com.github.zhangchengji.proto.client.Client;
import com.github.zhangchengji.proto.client.FlyingConfig;
import com.google.protobuf.InvalidProtocolBufferException;
import io.grpc.*;
import io.grpc.netty.shaded.io.grpc.netty.GrpcSslContexts;
import io.grpc.netty.shaded.io.grpc.netty.NettyChannelBuilder;
import io.grpc.netty.shaded.io.netty.handler.ssl.SslContext;
import io.grpc.netty.shaded.io.netty.handler.ssl.SslContextBuilder;

import javax.net.ssl.SSLException;
import java.io.IOException;
import java.io.InputStream;
import java.util.Properties;
import java.util.concurrent.TimeUnit;


import om.github.zhangchengji.proto.client.ClientServiceGrpc;
import om.github.zhangchengji.proto.common.Common;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.core.io.ClassPathResource;

/**
 * grpc 连接
 */

public class GrpcClient {
    private final static Logger log = LoggerFactory.getLogger(GrpcClient.class);

    private FlyingConfigProperties flyingConfigProperties;
    private ManagedChannel channel;

    private final ClientServiceGrpc.ClientServiceBlockingStub blockingStub;
    private Properties properties;
    private Client client;

    public GrpcClient(FlyingConfigProperties flyingConfigProperties) {
        this.flyingConfigProperties = flyingConfigProperties;
        blockingStub = ClientServiceGrpc.newBlockingStub(initGrpcConnection());
    }

    private ManagedChannel initGrpcConnection() {
        ClassPathResource classPathResource = new ClassPathResource("ca.pem");
        InputStream stream = null;
        try {
            stream = classPathResource.getInputStream();
        } catch (IOException e) {
            e.printStackTrace();
        }
        SslContextBuilder builder = GrpcSslContexts.forClient();
        SslContext sslContext = null;
        try {
            sslContext = builder.trustManager(stream).build();
        } catch (SSLException e) {
            e.printStackTrace();
        }
        // 构建 Channel
        this.channel = NettyChannelBuilder
                .forTarget(this.flyingConfigProperties.getServerAddr())
                .overrideAuthority(Constants.SERVER_NAME)
                .sslContext(sslContext)
                .build();
        return channel;
    }

    private Client createClient(String namespace ) {
        client= Client.newBuilder().setAppId(this.flyingConfigProperties.getAppId()).setNamespace(namespace).build();
        return client;
     }

    public FlyingConfig config(String namespace) {
        FlyingConfig flyingConfig = null;

        Common.Response resp = blockingStub.config( createClient(namespace));
        if (resp.getCode() == 200) {
            try {
                flyingConfig = resp.getData().unpack(FlyingConfig.class);
                if (flyingConfig==null){
                    throw new RuntimeException("获取配置信息为null，请检查配置中心是否存在此项目");
                }
            } catch (InvalidProtocolBufferException e) {
                e.printStackTrace();
            }
        }else if (resp.getCode() == 500) {
            log.error(resp.getMessage());
        }

        return flyingConfig;
    }

    public boolean listener() {
        Common.Response resp = blockingStub.listener(client);
       if (resp.getCode() ==304){
           log.debug(resp.getMessage());
           return false;
       }else return resp.getCode() == 200;
    }
    public  void shutdown(){
        try {
            channel.shutdown().awaitTermination(1, TimeUnit.SECONDS);
        } catch (InterruptedException e) {
            e.printStackTrace();
        }
    }


}
