package com.github.zhangchengji.flying;

import com.github.zhangchengji.flying.constants.FlyingConfigProperties;
import com.github.zhangchengji.flying.exceptions.FlyingConfigException;
import com.github.zhangchengji.flying.factory.FlyingFactory;
import com.github.zhangchengji.flying.util.GrpcClient;

import java.util.Objects;
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

public class FlyingConfigManager {

    private static ConfigService service = null;
    private static GrpcClient grpcClient;
    private FlyingConfigProperties flyingConfigProperties;
    public FlyingConfigManager(FlyingConfigProperties flyingConfigProperties, GrpcClient grpcClient){
        this.flyingConfigProperties=flyingConfigProperties;
        this.grpcClient=grpcClient;
        createConfigService(flyingConfigProperties);
    }
    static ConfigService createConfigService(FlyingConfigProperties flyingConfigProperties){
//        if (Objects.isNull(service)) {
//            synchronized (FlyingConfigManager.class) {
                try {
                    if (Objects.isNull(service)) {
                        service = FlyingFactory.createConfigService(flyingConfigProperties.listConfigServiceProperties(),grpcClient);
                    }
                }
                catch (Exception e) {
                    throw new FlyingConfigException("出问题lalalala",e);

                }
           // }
   //     }
        return service;
    }
    public ConfigService getConfigService() {
        if (Objects.isNull(service)) {
            createConfigService(this.flyingConfigProperties);
        }
        return service;
    }
    public FlyingConfigProperties getFlyingConfigProperties() {
        return flyingConfigProperties;
    }


}
