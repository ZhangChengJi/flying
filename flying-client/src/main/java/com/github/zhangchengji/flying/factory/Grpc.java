package com.github.zhangchengji.flying.factory;
import com.github.zhangchengji.proto.client.FlyingConfig;

public interface Grpc {

    FlyingConfig config(String namespace);
    boolean  listener();
    void shutdown();
}
