import {text} from './codeText'
import { notification } from 'ant-design-vue'
export  function alertMessage(t,data){
    const m=text(data)
    t.$notification[m.type]({
        message: m.title,
        description: m.msg,
    });
     return  m.type === 'success'
}
export  function loginMessage(data){
    const m=text(data)
    notification[m.type]({
        message: m.title,
        description: m.msg,
    });

}