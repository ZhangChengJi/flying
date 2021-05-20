import Vue from 'vue'
import Vuex from 'vuex'
import {getStore, removeStore, setStore} from '@/utils/store'
import {loginIn} from "@/api/user";
import { notification } from 'ant-design-vue';
import {generateTitle} from "@/utils/i18n";

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    appNodeId: getStore({
      name: 'appNodeId'
    }) || [],
    asId: getStore({
      name: 'asId'
    }) ||'',
    nodeStatus: getStore({
      name: 'nodeStatus'
    }) || true,
    loading: getStore({
      name: 'loading'
    })|| false,
    expires_in: getStore({
      name: 'expires_in'
    }) || '',
    username:  getStore({
      name: 'username'
    }) || '',
    headerImg: getStore({
               name: 'headerImg'
             }) || '',
    x_token: getStore({
      name: 'x_token'
    }) || '',
    isLock: getStore({ name: 'isLock' }) || false,
    docsite_language: getStore({name: 'docsite_language'})|| ''
  },
  mutations: {
    SET_X_TOKEN: (state, x_token) => {
      state.x_token = x_token
      setStore({
        name: 'x_token',
        content: state.x_token,
        type: 'session'
      })
    },
    SET_EXPIRES_IN: (state, expires_in) => {
      state.expires_in = expires_in
      setStore({
        name: 'expires_in',
        content: state.expires_in,
        type: 'session'
      })
    },
    SET_USERNAME:  (state, username) => {
      state.username = username
      setStore({
        name: 'username',
        content: state.username,
        type: 'session'
      })
    },
    SET_HEADER_IMG:  (state, headerImg) => {
      state.headerImg = headerImg
      setStore({
        name: 'headerImg',
        content: state.headerImg,
        type: 'session'
      })
    },
    changeAppNodeId(state,value){
      state.appNodeId[0]=value;
      setStore({name: 'appNodeId',content: state.appNodeId})

    },
    clearAppNodeId(state){
      state.appNodeId=[];
      setStore({name: 'appNodeId',content: state.appNodeId})

    },
    changeAsId: (state,value)=> {
      state.asId=value
      setStore({name: 'asId',content: state.asId})
    },
    changeNodeStatus:  (state,value)=> {
      state.nodeStatus=value
      setStore({name: 'nodeStatus',content: state.nodeStatus})
    },
    changLoading: (state,value)=> {
        state.loading=value
         setStore({name: 'loading',content: state.loading})
    },
    CLEAR_LOCK: (state) => {
      state.isLock = false
      state.lockPasswd = ''
      removeStore({
        name: 'lockPasswd'
      })
      removeStore({
        name: 'isLock'
      })
    },
    changeLanguage:(state,value) => {
      state.docsite_language=value
      setStore({name: 'docsite_language',content: state.docsite_language})
    },
  },
  actions: {
     LoginIn({commit},loginInfo){
       return new Promise((resolve, reject) => {
         loginIn(loginInfo).then(response => {
           if(response.data.code===200){
             notification.success({
               message: window.i18n.t('Login.loginSuccessfully')
             });
            commit('SET_X_TOKEN', response.data.data.token)
            commit('SET_EXPIRES_IN', response.data.data.expiresAt)
             commit('SET_USERNAME', response.data.data.user.userName)
            commit('SET_HEADER_IMG',response.data.data.user.headerImg)
            commit('CLEAR_LOCK')
            resolve()
           }
         }).catch(error => {
           reject(error)
         })
       })
    },
    FedLogOut({commit}) {
      return new Promise(resolve => {
        commit('SET_X_TOKEN', '')
        commit('SET_EXPIRES_IN','')
        commit('CLEAR_LOCK')
        resolve()
      })
    },
    SetLanguage({commit},language) {
      return new Promise(resolve => {
        commit('changeLanguage', language)
        resolve()
      })
    }
  },
  modules: {
  },
  getters:{
    appNodeId: state => state.appNodeId,
    asId: state => state.asId,
    nodeStatus: state => state.nodeStatus,
    loadding: state => state.loadding,
    expires_in: state => state.user.expires_in,
    x_token: state => state.x_token,
    username: state => state.username,
    headerImg: state => state.headerImg,
    docsite_language: state => state.docsite_language,
  }
})
