import request from '@/router/axios';
export function getAppList(query) {
    return request({
        url: '/app/getAppList',
        method: 'get',
        params: query
    })
}
export function createApp(obj) {
    return request({
        url: '/app/createApp',
        method: 'post',
        data: obj
    })
}
export function deleteApp(obj) {
    return request({
        url: '/app/deleteApp',
        method: 'delete',
        data: obj
    })
}
export function findApp(id) {
    return request({
        url: '/app/findApp/'+id,
        method: 'get',

    })
}
export function updateApp(obj) {
    return request({
        url: '/app/updateApp',
        method: 'put',
        data: obj

    })
}