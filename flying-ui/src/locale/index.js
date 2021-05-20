import Vue from 'vue';
import VueI18n from "vue-i18n";
import store from '../store';
import enLocale from './en_us'
import zhLocale from './zh_cn'
Vue.use(VueI18n);
if(store.getters.docsite_language===''){
    const language= 'en';
    store.dispatch("SetLanguage", language)
}
const messages = {
    en: {
        ...enLocale

    },
    zh: {
        ...zhLocale

    }
}
const i18n=new VueI18n({
    locale: store.getters.docsite_language||'en',
    messages
})
window.i18n = i18n
export  default i18n