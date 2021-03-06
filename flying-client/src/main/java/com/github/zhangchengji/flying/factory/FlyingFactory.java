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

package com.github.zhangchengji.flying.factory;

import com.github.zhangchengji.flying.ConfigService;
import com.github.zhangchengji.flying.exceptions.FlyingConfigException;

import java.util.Properties;

public class FlyingFactory {
    public FlyingFactory(){

    }
    public static ConfigService createConfigService(Properties properties ,Grpc grpc) throws FlyingConfigException {
        return ConfigFactory.createConfigService(properties,grpc);
    }
    public static Grpc createGrpcFactory(String address,String appId) throws FlyingConfigException {
        return ConfigFactory.createGrpcFactory(address, appId);
    }
}
