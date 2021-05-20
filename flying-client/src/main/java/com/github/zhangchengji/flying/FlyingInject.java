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

import com.github.zhangchengji.flying.factory.DefaultPropertiesFactory;
import com.github.zhangchengji.flying.factory.PropertiesFactory;
import com.google.inject.AbstractModule;

import com.google.inject.Guice;
import com.google.inject.Singleton;

public class FlyingInject  {
    private static volatile com.google.inject.Injector f_injector;
    private static final Object lock = new Object();

    private static com.google.inject.Injector getInjector() {
        if (f_injector == null) {
            synchronized (lock) {
                if (f_injector == null) {
                    try {
                        f_injector = Guice.createInjector(new FlyingModule());
                    } catch (Throwable ex) {

                    }
                }
            }
        }
        return f_injector;
    }



    public static  <T> T getInstance(Class<T> aClass) {
        return getInjector().getInstance(aClass);
    }


    public static <T> T getInstance(Class<T> clazz, String name) {
        return null;
    }


    private static class FlyingModule extends AbstractModule {
        @Override
        protected void configure() {
            bind(PropertiesFactory.class).to(DefaultPropertiesFactory.class).in(Singleton.class);
            super.configure();
        }
    }
    }
