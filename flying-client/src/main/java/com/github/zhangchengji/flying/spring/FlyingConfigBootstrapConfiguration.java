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

package com.github.zhangchengji.flying.spring;

import com.github.zhangchengji.flying.FlyingConfigManager;
import com.github.zhangchengji.flying.constants.FlyingConfigProperties;
import com.github.zhangchengji.flying.util.GrpcClient;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.boot.autoconfigure.condition.ConditionalOnMissingBean;
import org.springframework.boot.autoconfigure.condition.ConditionalOnProperty;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

@Configuration(proxyBeanMethods = false) //proxyBeanMethods为flase 不会被代理，如果为true会被CGLIB代理，如果只是普通类的话建议设置为 flase ,这样能提升性能
@ConditionalOnProperty(name = {"spring.flying.bootstrap.enabled"}, matchIfMissing = false)
public class FlyingConfigBootstrapConfiguration {
    private final static Logger log = LoggerFactory.getLogger(FlyingConfigBootstrapConfiguration.class);

    @Bean
    @ConditionalOnMissingBean
    public FlyingConfigProperties flyingConfigProperties(){return new FlyingConfigProperties();}
    @Bean
    public GrpcClient grpcClient(FlyingConfigProperties flyingConfigProperties){
        return new GrpcClient(flyingConfigProperties);
    }
    @Bean
    public FlyingConfigManager flyingConfigManager(
            FlyingConfigProperties flyingConfigProperties,GrpcClient grpcClient) {
        return new FlyingConfigManager(flyingConfigProperties,grpcClient);
    }

    @Bean
    public FlyingPropertySourceLocator flyingPropertySourceLocator(FlyingConfigProperties flyingConfigProperties,GrpcClient grpcClient){
        return new FlyingPropertySourceLocator(flyingConfigProperties,grpcClient);
    }
}
