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


import com.github.zhangchengji.flying.constants.FlyingConfigProperties;
import com.github.zhangchengji.flying.constants.PropertySourcesConstants;
import com.github.zhangchengji.flying.exceptions.FlyingConfigException;
import com.github.zhangchengji.flying.repository.RemoteConfigRepository;
import com.github.zhangchengji.flying.util.GrpcClient;
import com.github.zhangchengji.proto.client.FlyingConfig;
import com.google.common.base.Joiner;
import com.google.common.escape.Escaper;
import com.google.common.net.UrlEscapers;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.config.YamlPropertiesFactoryBean;
import org.springframework.cloud.bootstrap.config.PropertySourceLocator;
import org.springframework.core.annotation.Order;
import org.springframework.core.env.CompositePropertySource;
import org.springframework.core.env.Environment;
import org.springframework.core.env.PropertySource;
import org.springframework.core.io.ByteArrayResource;
import org.springframework.core.io.Resource;

import java.util.*;


@Order(0)
public class FlyingPropertySourceLocator implements PropertySourceLocator {
    private final static Logger log = LoggerFactory.getLogger(FlyingPropertySourceLocator.class);
    private static final Joiner.MapJoiner MAP_JOINER = Joiner.on("&").withKeyValueSeparator("=");
    private static final Escaper queryParamEscaper = UrlEscapers.urlFormParameterEscaper();
    private FlyingConfigProperties flyingConfigProperties;
    private GrpcClient grpcClient;
    private final boolean isRefreshEnabled;
     FlyingPropertySourceLocator(FlyingConfigProperties flyingConfigProperties,GrpcClient grpcClient){
         this.flyingConfigProperties = flyingConfigProperties;
         this.isRefreshEnabled = flyingConfigProperties.isRefreshEnabled();
         this.grpcClient=grpcClient;
     }
    @Override
    public PropertySource<?> locate(Environment environment) {
         flyingConfigProperties.setEnvironment(environment);
        CompositePropertySource composite = new CompositePropertySource(PropertySourcesConstants.FLYING_BOOTSTRAP_PROPERTY_SOURCE_NAME);
        String namespaces = flyingConfigProperties.getNamespace();
        log.info("所有namespace："+namespaces);

        String[] namespaceList= namespaces.trim().split(",");
      try {
          log.info("正在请求获取配置... 🚗 🚗 🚗 ");
        for (String namespace : namespaceList) {
                composite.addFirstPropertySource(getConfig(namespace));
        }
          if (!isRefreshEnabled) {
           grpcClient.shutdown();
          }
        } catch (Exception e) {
            log.error("配置加载失败");
            e.printStackTrace();
        }
        return composite;
    }



    public synchronized FlyingPropertySource getConfig(String namespace)   {
        log.info("⚙️ namespace: "+namespace+"进行获取配置");
        try {
            FlyingConfig flyingConfig = grpcClient.config(namespace);
            if (flyingConfig==null){
               return null;
            }
            FlyingPropertySource source=null;
            log.info("namespace:"+namespace+",配置获取成功 🎈🎈🎈");
            switch(flyingConfig.getFormat()){
                case "yaml":
                    YamlPropertiesFactoryBean yamlFactory = new YamlPropertiesFactoryBean();
                    yamlFactory.setResources(new Resource[]{new ByteArrayResource(flyingConfig.getValue().getBytes())});
                     source = new FlyingPropertySource(this.flyingConfigProperties.getAppId(),namespace, propertiesToMap(yamlFactory.getObject()));
                     break;
                case "properties":
                    source = new FlyingPropertySource(this.flyingConfigProperties.getAppId(),namespace, flyingConfig.getValue());
                    break;
                default:
                    log.error("未匹配到配置格式");
                    throw new FlyingConfigException("未匹配到配置格式");
            }

        RemoteConfigRepository.collectFlyingPropertySource(source);
            return source;

        } catch (Exception e) {
            log.error("配置中心无法连接错误，请检查配置");
           e.printStackTrace();
        }
        return null;
    }

    private Map<String, Object> propertiesToMap(Properties properties) {
        Map<String, Object> result = new HashMap(16);
        Enumeration keys = properties.propertyNames();
        while(keys.hasMoreElements()) {
            String key = (String)keys.nextElement();
            Object value = properties.getProperty(key);
            if (value != null) {
                result.put(key, ((String)value).trim());
            } else {
                result.put(key, (Object)null);
            }
        }

        return result;
    }

}
