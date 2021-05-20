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

package com.github.zhangchengji.flying.repository;

import com.github.zhangchengji.flying.spring.FlyingPropertySource;

import java.util.ArrayList;
import java.util.List;
import java.util.concurrent.ConcurrentHashMap;

public final class RemoteConfigRepository {

    private final static ConcurrentHashMap<String, FlyingPropertySource> FLYING_PROPERTY_SOURCE_REPOSITORY = new ConcurrentHashMap<String, FlyingPropertySource>();
    public static List<FlyingPropertySource> getAll() {
        return new ArrayList<>(FLYING_PROPERTY_SOURCE_REPOSITORY.values());
    }

    public static void collectFlyingPropertySource(
            FlyingPropertySource flyingPropertySource) {
        FLYING_PROPERTY_SOURCE_REPOSITORY
                .putIfAbsent(getMapKey(flyingPropertySource.getAppId(),
                        flyingPropertySource.getNamespaceName()), flyingPropertySource);
    }
    public static String getMapKey(String appId, String namespace) {
        return String.join(",", String.valueOf(appId),
                String.valueOf(namespace));
    }
}
