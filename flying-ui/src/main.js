import Vue from 'vue'
import i18n from "./locale";
import App from './App.vue'
import router from './router'
import axios from './router/axios'
import VueAxios from 'vue-axios'
import * as urls from '@/config/env'
import store from './store'
import './permission' // 权限
import Antd from 'ant-design-vue';
import 'ant-design-vue/dist/antd.less';
import './assets/iconfont/iconfont';
import './assets/iconfont/iconfont.css';
import VueInsProgressBar from 'vue-ins-progress-bar'
import GeminiScrollbar from 'vue-gemini-scrollbar'




const options = {
  position: 'fixed',
  show: true,
  height: '3px'
}
Vue.use(GeminiScrollbar)
Vue.use(VueInsProgressBar, options)
Vue.config.productionTip = false
window.axios = axios
Vue.use(VueAxios, axios)
// 加载相关url地址
Object.keys(urls).forEach(key => {
  Vue.prototype[key] = urls[key]
})
Vue.use(Antd)
new Vue({

  router,
  store,
  i18n,
  render: h => h(App)
}).$mount('#app')
