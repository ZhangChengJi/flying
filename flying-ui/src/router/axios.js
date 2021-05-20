import axios from 'axios';
import NProgress from 'nprogress';
import 'nprogress/nprogress.css'
import { notification } from 'ant-design-vue'
import router from "@/router"
import store from "@/store"; // progress bar style
import {text} from "@/utils/codeText"
import {loginMessage} from "@/utils/alert";
// message的基础配置
// message.config({
//     duration: 2, // 提示时常单位为s
//     top: '40px', // 距离顶部的位置
//     maxCount: 3 // 最多显示提示信息条数(后面的会替换前面的提示)
// })

axios.defaults.timeout = 10000;
// 返回其他状态吗
axios.defaults.validateStatus = function (status) {
    return status >= 200 && status <= 500 // 默认的
}
// 跨域请求，允许保存cookie
axios.defaults.withCredentials = true
// NProgress Configuration
NProgress.configure({
    showSpinner: false
});
// HTTPrequest拦截
axios.interceptors.request.use(config => {
    NProgress.start();
    let token =  store.getters.x_token
    if (token){
        config.headers['x-token'] = token // token
    }
    return config
}, error => {
    return Promise.reject(error)
});

// HTTPresponse拦截
axios.interceptors.response.use(res => {
    NProgress.done();
    const status = Number(res.status) || 200
    if (status === 401) {
        loginMessage(res.data)
        store.dispatch('FedLogOut').then(() => {
            router.push({path: '/login'})
        })
        return
    }

    // if(res.status!=200 || res.data.code === 1){
    //     notification.error({
    //         message: text(500)
    //
    //     });
    //
    // }

    return res
}, error => {
    NProgress.done()
    return Promise.reject(new Error(error))
})

export default axios