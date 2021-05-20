import request from '@/router/axios';

export function loginIn(obj) {
    return request({
        url: '/base/login',
        method: 'post',
        data: obj
    })
}
export function upPass(obj) {
    return request({
        url: '/base/upPass',
        method: 'post',
        data: obj
    })
}