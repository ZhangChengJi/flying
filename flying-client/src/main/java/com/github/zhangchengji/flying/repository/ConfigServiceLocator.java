//package com.flying.client.repository;
//
//import cn.hutool.core.util.StrUtil;
//import cn.hutool.json.JSON;
//import cn.hutool.json.JSONObject;
//import cn.hutool.json.JSONUtil;
//import com.flying.client.FlyingInject;
//import com.flying.client.constants.PropertySourcesConstants;
//import com.flying.client.factory.PropertiesFactory;
//import com.flying.client.model.dto.FlyingConfig;
//import com.flying.client.model.dto.R;
//import com.flying.client.spring.FlyingPropertySource;
//import com.google.common.base.Joiner;
//import com.google.common.collect.Maps;
//import com.google.common.escape.Escaper;
//import com.google.common.net.UrlEscapers;
//import kong.unirest.HttpResponse;
//import kong.unirest.Unirest;
//import kong.unirest.UnirestInstance;
//import lombok.extern.slf4j.Slf4j;
//import org.springframework.beans.factory.config.YamlPropertiesFactoryBean;
//import org.springframework.core.env.CompositePropertySource;
//import org.springframework.core.env.ConfigurableEnvironment;
//import org.springframework.core.io.ByteArrayResource;
//import org.springframework.core.io.Resource;
//
//import java.util.Enumeration;
//import java.util.HashMap;
//import java.util.Map;
//import java.util.Properties;
//import java.util.concurrent.atomic.AtomicReference;
//
//@Slf4j
//public class ConfigServiceLocator {
//    private static final Joiner.MapJoiner MAP_JOINER = Joiner.on("&").withKeyValueSeparator("=");
//    private static final Escaper queryParamEscaper = UrlEscapers.urlFormParameterEscaper();
//    private volatile AtomicReference<FlyingConfig> m_configCache;
//    public static String PROFILES_ACTIVE;
//    public static String URI;
//    public static String APPID;
//    protected PropertiesFactory propertiesFactory = FlyingInject.getInstance(PropertiesFactory.class);
//    public void initActive(ConfigurableEnvironment environment){
//       PROFILES_ACTIVE=environment.getProperty(PropertySourcesConstants.SPRING_PROFILES_ACTIVE).trim();
//      if( PROFILES_ACTIVE==null){
//          log.error(PropertySourcesConstants.SPRING_PROFILES_ACTIVE+"未设置");
//          return;
//      }
//
//    }
//    public void initUri(ConfigurableEnvironment environment){
//        URI=environment.getProperty(PropertySourcesConstants.FLYING_BOOTSTRAP_URI+"."+PROFILES_ACTIVE+".meta");
//        if( URI==null){
//            log.error(PropertySourcesConstants.FLYING_BOOTSTRAP_URI+"未设置");
//        }else{
//            log.info("🧨 正在使用"+PROFILES_ACTIVE+"环境，配置中心地址>>> :"+URI);
//
//        }
//    }
//    public void initAppId(ConfigurableEnvironment environment){
//       APPID=environment.getProperty(PropertySourcesConstants.FLYING_BOOTSTRAP_APPID).trim();
//        if( APPID==null){
//            log.error(PropertySourcesConstants.FLYING_BOOTSTRAP_APPID+"未设置");
//            return;
//        }
//        if(StrUtil.endWith(APPID,"/")){
//            APPID=StrUtil.removeSuffix(APPID,"/");
//        }
//    }
//    private String metaServiceUrl(String namespace){
//        Map<String, String> queryParams = Maps.newHashMap();
//        queryParams.put("appId", queryParamEscaper.escape(APPID));
//        queryParams.put("namespace",namespace);
//        String template= "{}/service/config?{}";
//        return StrUtil.format(template,URI,MAP_JOINER.join(queryParams));
//    }
//    public synchronized FlyingPropertySource getConfig(String namespace){
//        log.info("⚙️ 开始进行根据namespace进行获取配置");
//        String url=metaServiceUrl(namespace);
//        log.info("namespace:"+namespace+",获取配置路径为>>>"+url);
//        log.info("正在请求获取配置... 🚗 🚗 🚗 ");
//        HttpResponse<FlyingConfig> f =Unirest.get(url).asObject(FlyingConfig.class);
//       if (f.mapError(Error.class)!=null){
//           log.error("当前namespace：%s,获取配置失败",namespace);
//           return null;
//        }
//        log.info("namespace:"+namespace+",配置获取成功 🎈🎈🎈");
//        YamlPropertiesFactoryBean yamlFactory = new YamlPropertiesFactoryBean();
//        yamlFactory.setResources(new Resource[]{new ByteArrayResource(f.getBody().getValue().getBytes())});
//        return new FlyingPropertySource(namespace,this.propertiesToMap(yamlFactory.getObject()));
//    }
//    private Map<String, Object> propertiesToMap(Properties properties) {
//        Map<String, Object> result = new HashMap(16);
//        Enumeration keys = properties.propertyNames();
//
//        while(keys.hasMoreElements()) {
//            String key = (String)keys.nextElement();
//            Object value = properties.getProperty(key);
//            if (value != null) {
//                result.put(key, ((String)value).trim());
//            } else {
//                result.put(key, (Object)null);
//            }
//        }
//
//        return result;
//    }
//
//}
