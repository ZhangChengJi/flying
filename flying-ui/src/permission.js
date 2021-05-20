/**
 * 全站权限配置
 *
 */
import router from './router'
import store from '@/store'
import {validatenull} from '@/utils/validate'
import NProgress from 'nprogress' // progress bar
import 'nprogress/nprogress.css' // progress bar style
NProgress.configure({showSpinner: false})

/**
 * 导航守卫，相关内容可以参考:
 * https://router.vuejs.org/zh/guide/advanced/navigation-guards.html
 */
router.beforeEach((to, from, next) => {
   // NProgress.start()

  if (store.getters.x_token) {

   if (to.path === '/login') {
     return  next({
       name: "首页"
     })
    }else{
     return   next()
   }
  } else {
    if (to.path === '/login') {

      return  next()
  }else{

      return  next({path:'/login'})

    }
  }
})

router.afterEach((to,from)=>{ //这里不接收next
  console.log(to);
  console.log(from);
//  NProgress.done()
 // const title = store.getters.tag.label

})
