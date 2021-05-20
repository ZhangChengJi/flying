package com.github.zhangchengji.flying.refresh;

import com.github.zhangchengji.flying.ConfigService;
import com.github.zhangchengji.flying.FlyingConfigManager;
import com.github.zhangchengji.flying.constants.FlyingConfigProperties;
import com.github.zhangchengji.flying.exceptions.FlyingConfigException;
import com.github.zhangchengji.flying.listener.AbstractSharedListener;
import com.github.zhangchengji.flying.listener.Listener;
import com.github.zhangchengji.flying.spring.FlyingPropertySource;
import com.github.zhangchengji.flying.repository.RemoteConfigRepository;
import com.github.zhangchengji.flying.spring.FlyingConfigBootstrapConfiguration;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.BeansException;
import org.springframework.boot.context.event.ApplicationReadyEvent;
import org.springframework.cloud.endpoint.event.RefreshEvent;
import org.springframework.context.ApplicationContext;
import org.springframework.context.ApplicationContextAware;
import org.springframework.context.ApplicationListener;
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
import java.util.Map;
import java.util.concurrent.ConcurrentHashMap;

public class FlyingContextRefresher implements ApplicationListener<ApplicationReadyEvent>, ApplicationContextAware {
    private final static Logger log = LoggerFactory.getLogger(FlyingConfigBootstrapConfiguration.class);

    private FlyingConfigProperties flyingConfigProperties;
        private final ConfigService configService;
        private final boolean isRefreshEnabled;
        private ApplicationContext applicationContext;
        private Map<String, Listener> listenerMap = new ConcurrentHashMap<String, Listener>(16);
   // private final ConfigService configService = FlyingInject.getInstance(FlyingConfigService.class);
     public FlyingContextRefresher(FlyingConfigManager flyingConfigManager ){
         this.flyingConfigProperties=flyingConfigManager.getFlyingConfigProperties();
         this.configService=flyingConfigManager.getConfigService();
         this.isRefreshEnabled=this.flyingConfigProperties.isRefreshEnabled();


     }
    /***
     * 加载Spring配置文件时，如果Spring配置文件中所定义的Bean类实现了ApplicationContextAware 接口，那么在加载Spring配置文件时，
         * 会自动调用ApplicationContextAware 接口中的setApplicationContext
     * @param applicationContext
     * @throws BeansException
     */
    @Override
    public void setApplicationContext(ApplicationContext applicationContext) throws BeansException {
    this.applicationContext=applicationContext;
    }

    protected void registerFlyingListenerForApplication(){
        for(FlyingPropertySource propertySource : RemoteConfigRepository.getAll()){
            registerFlyingListener(propertySource.getAppId(),propertySource.getNamespaceName());
        }
    }
    protected  void registerFlyingListener(String appId, String namespace){
      String key= RemoteConfigRepository.getMapKey(appId,namespace);
        //computeIfAbsent，若key对应的value为空，会将第二个参数的返回值存入并返回
        Listener  listener=listenerMap.computeIfAbsent(key,
                lst->  new AbstractSharedListener(){
                    @Override
                    public void innerReceive(String appId, String namespace, String configInfo) {
                        applicationContext.publishEvent( new RefreshEvent(this, null, "Refresh Flying config"));
                    }
                });
        try {
            configService.addListener(appId,namespace,listener);
        }
        catch (FlyingConfigException e) {
            log.warn("添加监听失败", e);
        }

    }
    @Override
    public void onApplicationEvent(ApplicationReadyEvent event) {
        if(isRefreshEnabled()){
            this.registerFlyingListenerForApplication();
        }


    }

        public boolean isRefreshEnabled() {
            if (null == flyingConfigProperties) {
                return isRefreshEnabled;
            }
            // Compatible with older configurations
            if (flyingConfigProperties.isRefreshEnabled() && !isRefreshEnabled) {
                return false;
            }
            return isRefreshEnabled;
        }

    }
