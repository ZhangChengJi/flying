import Vue from 'vue'
import VueRouter from 'vue-router'
import Layout from '@/layout'
import SonView from "@/SonView";
Vue.use(VueRouter)
// 解决报错
const originalPush = VueRouter.prototype.push
const originalReplace = VueRouter.prototype.replace
// push
VueRouter.prototype.push = function push (location, onResolve, onReject) {
  if (onResolve || onReject) return originalPush.call(this, location, onResolve, onReject)
  return originalPush.call(this, location).catch(err => err)
}
// replace
VueRouter.prototype.replace = function push (location, onResolve, onReject) {
  if (onResolve || onReject) return originalReplace.call(this, location, onResolve, onReject)
  return originalReplace.call(this, location).catch(err => err)
}

const routes = [
  {
    path: '/login',
    name: 'login',
    component: () => import('@/views/Login'),
  },
  {
    path: '/',
    name: '主页',
    redirect: '/wel'
  },
  {
    path: '/wel',
    component: Layout,
    redirect: '/wel/index',
    children: [{
      path: 'index',
      name: '首页',
      component: () =>
          import ( /* webpackChunkName: "views" */ '@/views/Home')
    }]
  },
  {
    path: '/app',
   // redirect: '/app/detail',
    component: SonView,
    name:'app',
    children: [
      {
        path: 'detail',
        component: () => import('@/views/Detail'),
        name: 'Detail'
      },
      {
        path: 'namespace',
        component: () => import('@/components/Namespace'),
        name: 'Namespace'
      }
    ]
  },
  {
    path: '/namespace',
    children: [{
      path: 'namespace',
      name: '首页',
      component: () =>
          import ( '@/views/Namespace')
    }]
  },
  {
    path: '/list',
    component: Layout,
    name:'list',
    children: [
      {
        path: 'node',
        component: () => import('@/views/NodeList'),
        name: 'Node'
      },
      {
        path: 'app',
        component: () => import('@/views/AppList'),
        name: 'App'
      }
    ]

  }
]

const router = new VueRouter({
  mode: 'hash',
 // base: process.env.BASE_URL,
  routes
})

export default router
