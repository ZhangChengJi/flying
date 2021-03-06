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
import com.github.zhangchengji.flying.refresh.FlyingContextRefresher;
import org.springframework.beans.factory.BeanFactoryUtils;
import org.springframework.boot.autoconfigure.condition.ConditionalOnProperty;
import org.springframework.context.ApplicationContext;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

@Configuration(proxyBeanMethods = false) //proxyBeanMethods为flase 不会被代理，如果为true会被CGLIB代理，如果只是普通类的话建议设置为 flase ,这样能提升性能
@ConditionalOnProperty(name = {"spring.flying.bootstrap.refresh-enabled"}, matchIfMissing = false)
public class FlyingConfigAutoConfiguration {
    @Bean
    public FlyingConfigProperties flyingConfigProperties(ApplicationContext context) {
        if (context.getParent() != null
                && BeanFactoryUtils.beanNamesForTypeIncludingAncestors(
                context.getParent(), FlyingConfigProperties.class).length > 0) {
            return BeanFactoryUtils.beanOfTypeIncludingAncestors(context.getParent(),
                    FlyingConfigProperties.class);
        }
        return new FlyingConfigProperties();
    }

    @Bean
    public FlyingConfigManager flyingConfigManager(
            FlyingConfigProperties flyingConfigProperties ) {
        return new FlyingConfigManager(flyingConfigProperties);
    }

    @Bean
    public FlyingContextRefresher flyingContextRefresher(FlyingConfigManager flyingConfigManager ){
        return new FlyingContextRefresher(flyingConfigManager);
    }
}
