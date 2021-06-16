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
import com.github.zhangchengji.flying.constants.PropertySourcesConstants;
import com.github.zhangchengji.flying.exceptions.FlyingConfigException;
import com.github.zhangchengji.flying.factory.Grpc;
import com.github.zhangchengji.flying.factory.GrpcFactory;
import com.github.zhangchengji.flying.repository.RemoteConfigRepository;
import com.github.zhangchengji.flying.client.Client;
import com.github.zhangchengji.flying.client.FlyingConfig;
import com.github.zhangchengji.flying.common.Response;
import com.google.protobuf.InvalidProtocolBufferException;
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
import retrofit2.Call;

import java.util.*;


@Order(0)
public class FlyingPropertySourceLocator implements PropertySourceLocator {
    private final static Logger log = LoggerFactory.getLogger(FlyingPropertySourceLocator.class);
    private FlyingConfigProperties flyingConfigProperties;
    private  boolean isRefreshEnabled;
    private Grpc grpc;

     FlyingPropertySourceLocator(FlyingConfigManager flyingConfigManager){
         this.flyingConfigProperties = flyingConfigManager.getFlyingConfigProperties();
         this.grpc=flyingConfigManager.grpc;
         this.isRefreshEnabled = flyingConfigProperties.isRefreshEnabled();
     }
    @Override
    public PropertySource<?> locate(Environment env) {
         flyingConfigProperties.setEnvironment(env);
        CompositePropertySource composite = new CompositePropertySource(PropertySourcesConstants.FLYING_BOOTSTRAP_PROPERTY_SOURCE_NAME);
        String namespaces = flyingConfigProperties.getNamespace();
        log.info("all namespace："+namespaces);

        String[] namespaceList= namespaces.trim().split(",");
      try {
          log.info("Getting configuration 🚚 ... "); //正在获取配置
        for (String namespace : namespaceList) {
                composite.addFirstPropertySource(getConfig(namespace));
        }
          if (!isRefreshEnabled) {
             // grpc.shutdown();
          }
        } catch (Exception e) {
            log.error("Configuration failed to load");
            e.printStackTrace();
        }
        return composite;
    }


    private Client createClient(String appId, String namespace){
        return Client.newBuilder().setAppId(appId).setNamespace(namespace).build();

    }
    public synchronized FlyingPropertySource getConfig(String namespace)   {
        log.info("⚙️ namespace: "+namespace+"Get configuration");
        try {
           Client c= createClient(this.flyingConfigProperties.getAppId(),namespace);
            Call<Response> resp = grpc.config(c);
            FlyingConfig flyingConfig = null;
            try {
                System.out.println( resp.execute().body().getData());
                flyingConfig = resp.execute().body().getData().unpack(FlyingConfig.class);
            } catch (InvalidProtocolBufferException e) {
                e.printStackTrace();
            }
            if (flyingConfig==null){
               return null;
            }

            FlyingPropertySource source=null;
            log.info("namespace:"+namespace+",Configuration obtained successfully 🎈🎈🎈");
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
                    log.error("Does not match the configuration format"); //未匹配到配置格式
                    throw new FlyingConfigException("Does not match the configuration format");
            }

        RemoteConfigRepository.collectFlyingPropertySource(source);
            return source;

        } catch (Exception e) {
            log.error("Configuration center connection failed. Please check whether the configuration center is running normally?"); //配置中心连接失败，请配置中心检查是否正常运行
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
