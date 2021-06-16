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
import org.springframework.context.ApplicationEvent;
import org.springframework.context.ApplicationEventPublisher;

import java.lang.reflect.Constructor;
import java.util.Properties;

public class ConfigFactory {
    /**
     * Create Config.
     *
     * @param properties init param
     * @return ConfigService
     * @throws  Exception
     */
    public static ConfigService createConfigService(Properties properties,Grpc grpc ) throws FlyingConfigException {
        try {
            Class<?> driverImplClass = Class.forName("com.github.zhangchengji.flying.FlyingConfigService");
            Constructor constructor = driverImplClass.getConstructor(Properties.class,Grpc.class);
            ConfigService vendorImpl = (ConfigService) constructor.newInstance(properties,grpc);
            return vendorImpl;
        } catch (Throwable e) {
            throw new FlyingConfigException("出现问题了",e);
        }
    }
    public static Grpc createGrpcFactory(String address,String appId) throws FlyingConfigException {
        try {
            Class<?> driverImplClass = Class.forName("com.github.zhangchengji.flying.factory.GrpcFactory");
            Constructor constructor = driverImplClass.getConstructor(String.class,String.class);
            Grpc vendorImpl = (Grpc) constructor.newInstance(address,appId);
            return vendorImpl;
        } catch (Throwable e) {
            throw new FlyingConfigException("出现问题了",e);
        }
    }

}
