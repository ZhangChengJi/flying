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

import com.github.zhangchengji.flying.FlyingUtils;
import org.springframework.core.env.PropertiesPropertySource;

import java.util.Map;
import java.util.Properties;

public class FlyingPropertySource extends PropertiesPropertySource {

    private String appId;
    private String namespaceName;
    private String releaseKey;

    public FlyingPropertySource(String appId,String name, String source) {
        super(name, FlyingUtils.toProperties(source));
        this.appId=appId;
        this.namespaceName=name;
    }
    public FlyingPropertySource(String appId,String name, Properties source) {
        super(name, source);
        this.appId=appId;
        this.namespaceName=name;
    }

    public FlyingPropertySource(String appId, String name, Map<String, Object> source) {
        super(name, source);
        this.appId=appId;
        this.namespaceName=name;


    }

    public String getAppId() {
        return appId;
    }

    public void setAppId(String appId) {
        this.appId = appId;
    }

    public String getNamespaceName() {
        return namespaceName;
    }

    public void setNamespaceName(String namespaceName) {
        this.namespaceName = namespaceName;
    }

    public String getReleaseKey() {
        return releaseKey;
    }

    public void setReleaseKey(String releaseKey) {
        this.releaseKey = releaseKey;
    }
}
