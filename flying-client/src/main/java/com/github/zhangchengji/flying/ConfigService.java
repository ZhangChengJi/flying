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

public interface  ConfigService {

   
    String getConfig(String appId, String namespace, long timeoutMs);

  
    String getConfigAndSignListener(String appId, String namespace, long timeoutMs, Listener listener);

 
    void addListener(String appId, String namespace, Listener listener);


    boolean publishConfig(String appId, String namespace, String content);


    boolean removeConfig(String appId, String namespace);


    void removeListener(String appId, String namespace, Listener listener);

    String getServerStatus();


    void shutDown();
}
