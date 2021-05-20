import request from '@/router/axios';
export function getNodeAll() {
    return request({
        url: '/node/getNodeList',
        method: 'get',
    })
}
export function getAppNode(appId) {
    return request({
        url: '/node/getAppNode/'+appId,
        method: 'get',

    })
}
export function getSyncNode(query) {
    return request({
        url: '/node/getSyncNode',
        method: 'get',
        params: query

    })
}
export function getNodeList(query) {
    return request({
        url: '/node/getNodeList',
        method: 'get',
        params: query
    })
}
export function deleteNode(obj) {
    return request({
        url: '/node/deleteNode',
        method: 'delete',
        data: obj
    })
}
export function createNode(obj) {
    return request({
        url: '/node/createNode',
        method: 'post',
        data: obj
    })
}
export function findNode(id) {
    return request({
        url: '/node/findNode/'+id,
        method: 'get',

    })
}
export function updateNode(obj) {
    return request({
        url: '/node/updateNode',
        method: 'put',
        data: obj

    })
}
export function ping(id) {
    return request({
        url: '/node/ping/'+id,
        method: 'get',

    })
}