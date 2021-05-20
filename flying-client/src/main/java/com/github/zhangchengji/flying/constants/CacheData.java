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

import com.github.zhangchengji.flying.listener.AbstractConfigChangeListener;
import com.github.zhangchengji.flying.listener.AbstractSharedListener;
import com.github.zhangchengji.flying.listener.Listener;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import java.util.ArrayList;
import java.util.List;
import java.util.concurrent.CopyOnWriteArrayList;


public class CacheData {
    private final static Logger log = LoggerFactory.getLogger(CacheData.class);
    private String namespace;
    private String appId;
    private String serverAddr;
    private String releaseKey;
    private volatile String value;
    private int taskId;

    private final CopyOnWriteArrayList<CacheData.ManagerListenerWrap> listeners;
    private volatile String md5;


    public void addListener(Listener listener) {
        if (null == listener) {
            throw new IllegalArgumentException("listener is null");
        }
        ManagerListenerWrap wrap =
                (listener instanceof AbstractConfigChangeListener) ? new ManagerListenerWrap(listener, md5, value)
                        : new ManagerListenerWrap(listener, md5);
        if (listeners.addIfAbsent(wrap)) {
            log.info("[{}] [add-listener] ok, tenant={}, dataId={}, group={}, cnt={}", namespace, null, appId, namespace,
                    listeners.size());
        }

    }

    public CacheData(String namespace, String appId, String serverAddr) {
        this.namespace = namespace;
        this.appId = appId;
        this.serverAddr = serverAddr;
         this.listeners = new CopyOnWriteArrayList<ManagerListenerWrap>();

    }
    public  void checkListenerMd5() {
        for (ManagerListenerWrap wrap : listeners) {
//            if (!md5.equals(wrap.lastCallMd5)) {
                safeNotifyListener(appId, namespace, value, md5, wrap);
       //     }
        }
    }
    private void safeNotifyListener(final String appId, final String namespace, final String value,
                                    final String md5, final ManagerListenerWrap listenerWrap) {
        final Listener listener = listenerWrap.listener;
        Runnable job = new Runnable() {
            @Override
            public void run() {
                ClassLoader myClassLoader = Thread.currentThread().getContextClassLoader();
                ClassLoader appClassLoader = listener.getClass().getClassLoader();
                try {
                    if (listener instanceof AbstractSharedListener) {
                        AbstractSharedListener adapter = (AbstractSharedListener) listener;
                        adapter.fillContext(appId, namespace);
                        log.info("[{}] [notify-context] dataId={}, group={}, md5={}", namespace, appId, namespace, md5);
                    }
                    // 执行回调之前先将线程classloader设置为具体webapp的classloader，以免回调方法中调用spi接口是出现异常或错用（多应用部署才会有该问题）。
                    Thread.currentThread().setContextClassLoader(appClassLoader);
                    listener.receiveConfigInfo(value);

                    listenerWrap.lastCallMd5 = md5;
                } catch (Exception ex) {
                    log.error("[{}] [notify-error] dataId={}, group={}, md5={}, listener={} errCode={} errMsg={}",
                            namespace, appId, namespace, md5, listener);
                } catch (Throwable t) {
                    log.error("[{}] [notify-error] dataId={}, group={}, md5={}, listener={} tx={}", namespace, appId,
                            namespace, md5, listener, t.getCause());
                } finally {
                    Thread.currentThread().setContextClassLoader(myClassLoader);
                }
                }
        };

        final long startNotify = System.currentTimeMillis();
        try {
            if (null != listener.getExecutor()) {
                listener.getExecutor().execute(job);
            } else {
                job.run();
            }
        } catch (Throwable t) {
            log.error("[{}] [notify-error] dataId={}, group={}, md5={}, listener={} throwable={}", namespace, appId,
                    namespace, md5, listener, t.getCause());
        }
    }
    public List<Listener> getListeners() {
        List<Listener> result = new ArrayList<Listener>();
        for (ManagerListenerWrap wrap : listeners) {
            result.add(wrap.listener);
        }
        return result;
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

    public String getServerAddr() {
        return serverAddr;
    }

    public void setServerAddr(String serverAddr) {
        this.serverAddr = serverAddr;
    }

    public String getReleaseKey() {
        return releaseKey;
    }

    public void setReleaseKey(String releaseKey) {
        this.releaseKey = releaseKey;
    }

    public String getValue() {
        return value;
    }

    public void setValue(String value) {
        this.value = value;
    }

    public int getTaskId() {
        return taskId;
    }

    public void setTaskId(int taskId) {
        this.taskId = taskId;
    }

    public static String getMd5String(String config) {
        return null == config ? "" : MD5Utils.md5Hex(config, "UTF-8");
    }

    private static class ManagerListenerWrap {
        final Listener listener;

        String lastCallMd5 = CacheData.getMd5String(null);

        String lastContent = null;

        ManagerListenerWrap(Listener listener) {
            this.listener = listener;
        }

        ManagerListenerWrap(Listener listener, String md5) {
            this.listener = listener;
            this.lastCallMd5 = md5;
        }

        ManagerListenerWrap(Listener listener, String md5, String lastContent) {
            this.listener = listener;
            this.lastCallMd5 = md5;
            this.lastContent = lastContent;
        }

    }
}
