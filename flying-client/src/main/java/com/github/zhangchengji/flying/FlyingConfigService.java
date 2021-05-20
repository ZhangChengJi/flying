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

package com.github.zhangchengji.flying;

import com.github.zhangchengji.flying.listener.Listener;
import com.github.zhangchengji.flying.refresh.FlyingRefreshClient;
import com.github.zhangchengji.flying.util.GrpcClient;

import java.util.*;

@SuppressWarnings("PMD.ServiceOrDaoClassShouldEndWithImplRule")
public class FlyingConfigService implements ConfigService {
    private Properties properties;
    private FlyingRefreshClient flyingRefreshClient;
    public  FlyingConfigService(Properties properties, GrpcClient grpcClient){
        this.properties=properties;
       this.flyingRefreshClient=new FlyingRefreshClient(properties,grpcClient);


    }

    @Override
    public String getConfig(String appId, String namespace, long timeoutMs) {
        return null;
    }

    @Override
    public String getConfigAndSignListener(String appId, String namespace, long timeoutMs, Listener listener) {
        return null;
    }

    @Override
    public void addListener(String appId, String namespace, Listener listener) {
        flyingRefreshClient.addTenantListeners(appId,namespace,Arrays.asList(listener));
    }

    @Override
    public boolean publishConfig(String appId, String namespace, String content) {
        return false;
    }

    @Override
    public boolean removeConfig(String appId, String namespace) {
        return false;
    }

    @Override
    public void removeListener(String appId, String namespace, Listener listener) {

    }

    @Override
    public String getServerStatus() {
        return null;
    }

    @Override
    public void shutDown() {

    }
}
