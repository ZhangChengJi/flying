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

package com.github.zhangchengji.flying.constants;

import com.github.zhangchengji.flying.exceptions.FlyingConfigException;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.context.properties.ConfigurationProperties;
import org.springframework.core.env.Environment;
import org.springframework.util.StringUtils;

import javax.annotation.PostConstruct;
import java.util.List;
import java.util.Objects;
import java.util.Properties;

import static com.github.zhangchengji.flying.constants.EnvProperties.*;

@ConfigurationProperties("spring.flying.bootstrap")
public class FlyingConfigProperties {
    private Boolean enabled;
    private String namespace;
    private String appId;
    private List<Address> address;
    private String serverAddr;
    private boolean refreshEnabled = true;
    private int longPollTimeout=50000;

    @Autowired
    private Environment environment;

    @PostConstruct
    public void init() {
        this.overrideFromEnv();
    }

    private void overrideFromEnv() {
        String active = environment.getActiveProfiles()[0];
        if (StringUtils.isEmpty(this.getServerAddr())) {
           Address ad = address.stream().filter(a -> active.equals(a.getName())).findAny().orElse(null);
           if (ad==null){
               throw new FlyingConfigException("未匹配到正确环境,检查配置中心是否存在"+active+"环境");

           }
            this.setServerAddr(Objects.requireNonNull(ad).getUrl());
        }
        if (StringUtils.isEmpty(this.getAppId())) {
            this.setAppId(environment.resolvePlaceholders("${"+PropertySourcesConstants.FLYING_BOOTSTRAP_APPID+":}"));
        }
        if (StringUtils.isEmpty(this.getNamespace())) {
            this.setNamespace(environment.resolvePlaceholders("${"+PropertySourcesConstants.FLYING_BOOTSTRAP_NAMESPACE+":}"));
        }
    }
    public Properties listConfigServiceProperties(){
        Properties properties = new Properties();
        properties.put(SERVER_ADDR, this.serverAddr);
        properties.put(APPID,this.appId);
        properties.put(NAMESPACE,this.namespace);
        properties.put(LONGPOLLTIMEOUT,this.longPollTimeout);
       return properties;
    }
    public Boolean getEnabled() {
        return enabled;
    }

    public void setEnabled(Boolean enabled) {
        this.enabled = enabled;
    }

    public String getNamespace() {
        return namespace;
    }

    public void setNamespace(String namespace) {
        this.namespace = namespace;
    }

    public String getAppId() {
        return appId;
    }

    public void setAppId(String appId) {
        this.appId = appId;
    }

    public List<Address> getAddress() {
        return address;
    }

    public void setAddress(List<Address> address) {
        this.address = address;
    }

    public String getServerAddr() {
        return serverAddr;
    }

    public void setServerAddr(String serverAddr) {
        this.serverAddr = serverAddr;
    }

    public Environment getEnvironment() {
        return environment;
    }

    public void setEnvironment(Environment environment) {
        this.environment = environment;
    }

    public boolean isRefreshEnabled() {
        return refreshEnabled;
    }

    public void setRefreshEnabled(boolean refreshEnabled) {
        this.refreshEnabled = refreshEnabled;
    }

    public int getLongPollTimeout() {
        return longPollTimeout;
    }

    public void setLongPollTimeout(int longPollTimeout) {
        this.longPollTimeout = longPollTimeout;
    }
}
