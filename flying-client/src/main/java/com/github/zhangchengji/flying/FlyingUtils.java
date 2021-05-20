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


import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.util.StringUtils;

import java.io.IOException;
import java.io.StringReader;
import java.util.Properties;


public class FlyingUtils {
    private final static Logger log = LoggerFactory.getLogger(FlyingUtils.class);

    public static Properties toProperties(String text) {
        Properties properties = new Properties();

        try {
            if (StringUtils.hasText(text)) {
                properties.load(new StringReader(text));
            }
        } catch (IOException var3) {
            log.debug("配置转化失败",var3);
        }

        return properties;
    }
}
