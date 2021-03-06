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

package com.github.zhangchengji.flying.listener;

import java.util.concurrent.Executor;

@SuppressWarnings("PMD.AbstractClassShouldStartWithAbstractNamingRule")
public abstract class AbstractSharedListener implements Listener {
    
    private volatile String appId;
    
    private volatile String namespace;
    
    public final void fillContext(String appId, String namespace) {
        this.appId = appId;
        this.namespace = namespace;
    }
    
    @Override
    public final void receiveConfigInfo(String configInfo) {
        innerReceive(appId, namespace, configInfo);
    }
    
    @Override
    public Executor getExecutor() {
        return null;
    }
    
    /**
     * receive.
     *
     * @param appId     data ID
     * @param namespace      namespace
     * @param configInfo content
     */
    public abstract void innerReceive(String appId, String namespace, String configInfo);
}
