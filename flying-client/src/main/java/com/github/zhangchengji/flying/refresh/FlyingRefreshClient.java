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

package com.github.zhangchengji.flying.refresh;

import com.github.zhangchengji.flying.constants.CacheData;
import com.github.zhangchengji.flying.constants.GroupKey;
import com.github.zhangchengji.flying.listener.Listener;
import com.github.zhangchengji.flying.util.GrpcClient;
import com.google.common.base.Joiner;
import com.google.common.escape.Escaper;
import com.google.common.net.UrlEscapers;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import static com.github.zhangchengji.flying.constants.EnvProperties.*;
import java.util.*;
import java.util.concurrent.Executors;
import java.util.concurrent.ScheduledExecutorService;
import java.util.concurrent.ThreadFactory;
import java.util.concurrent.TimeUnit;
import java.util.concurrent.atomic.AtomicReference;

public class FlyingRefreshClient  {
    private final static Logger log = LoggerFactory.getLogger(FlyingRefreshClient.class);
    private static final Escaper queryParamEscaper = UrlEscapers.urlFormParameterEscaper();
    private static final Joiner.MapJoiner MAP_JOINER = Joiner.on("&").withKeyValueSeparator("=");
     ScheduledExecutorService executor;
     ScheduledExecutorService executorService;
     private Properties properties;
    private GrpcClient grpcClient;
    private final AtomicReference<Map<String, CacheData>> cacheMap = new AtomicReference<Map<String, CacheData>>(
            new HashMap<String, CacheData>());
    public void addTenantListeners(String appId, String namespace, List<? extends Listener> listeners) {
        CacheData cache = addCacheDataIfAbsent(appId, namespace);
        for (Listener listener : listeners) {
            cache.addListener(listener);
        }
        checkConfigInfo();
    }
    public CacheData addCacheDataIfAbsent(String appId, String namespace)   {

        synchronized (cacheMap) {
            String key = GroupKey.getKeyTenant(appId, namespace, "");
            CacheData cache = new CacheData(appId, namespace,properties.getProperty(SERVER_ADDR));
                    cache.setValue(grpcClient.config(namespace).getValue());
            Map<String, CacheData> copy = new HashMap<String, CacheData>(this.cacheMap.get());
            copy.put(key, cache);
            cacheMap.set(copy);
            return cache;
            }


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

    public FlyingRefreshClient(Properties properties, GrpcClient grpcClient){
        this.properties = properties;
        this.grpcClient=grpcClient;
        this.executor = Executors.newScheduledThreadPool(1);
        this.executorService = Executors
                .newScheduledThreadPool(Runtime.getRuntime().availableProcessors(), new ThreadFactory() {
                    @Override
                    public Thread newThread(Runnable r) {
                        Thread t = new Thread(r);
                        t.setName("com.flying.client.Worker.longPolling.");
                        t.setDaemon(true);
                        return t;
                    }
                });

    }



    private double currentLongingTaskCount = 0;
    public void checkConfigInfo() {

        int listenerSize = cacheMap.get().size();
        // Round up the longingTaskCount.


        //         if(f.getBody().ok()){
        //
        //         }
        double perTaskConfigSize = 3000;
        int longingTaskCount = (int) Math.ceil(listenerSize / perTaskConfigSize);
        if (longingTaskCount > currentLongingTaskCount) {
            for (int i = (int) currentLongingTaskCount; i < longingTaskCount; i++) {
                // The task list is no order.So it maybe has issues when changing.
                executorService.execute(new LongPollingRunnable(i));
            }
            currentLongingTaskCount = longingTaskCount;
        }
    }

    class LongPollingRunnable implements Runnable {

        private final int taskId;
        private boolean connectionStatus = true;
        public LongPollingRunnable(int taskId) {
            this.taskId = taskId;
        }
        @Override
        public void run() {
            List<CacheData> cacheDatas = new ArrayList<CacheData>();
            try {
                // check failover config
                for (CacheData cacheData : cacheMap.get().values()) {
                    if (cacheData.getTaskId() == taskId) {
                        cacheDatas.add(cacheData);

                            if(grpcClient.listener()){
                                cacheData.checkListenerMd5();
                            }
                            if(!connectionStatus){
                                log.info("The configuration center resumes normal connection successfully"); //配置中心已经恢复正常连接
                                connectionStatus=true;
                            }
                        }

                    }
                executorService.execute(this);
            } catch (Throwable e) {
                connectionStatus=false;
                log.error("longPolling error : Configuration center connection failed ", e);
                executorService.schedule(this, 30000, TimeUnit.MILLISECONDS);
            }
        }

    }



     //   adapter.receiveConfigInfo(f.getBody().getValue());
       // CacheData cache = cacheMap.get().get(GroupKey.getKeyTenant(flyingConfigProperties.getAppId(), flyingConfigProperties.getNamespace()));

      //  cache.checkListenerMd5();


}
