import request from '@/router/axios';
export function getAppNamespace(query) {
    return request({
        url: '/namespace/getAppNamespace',
        method: 'get',
        params: query
    })
}
export function findNamespace(query) {
    return request({
        url: '/namespace/findNamespace',
        method: 'get',
        params: query
    })
}
export function createNamespace(appId,name,nodes,format,value) {
    return request({
        url: '/namespace/createNamespace',
        method: 'post',
        data: {
            appId: appId,
            name: name,
            nodes: nodes,
            format: format,
            value: value,

        }
    })
}
export function createRelease(obj) {
    return request({
        url: '/release/createRelease',
        method: 'post',
        data: obj
    })
}
export function updateNamespace(obj) {
    return request({
        url: '/namespace/updateNamespace',
        method: 'put',
        data: obj
    })
}
export function deleteNamespace(id,nodeId) {
    return request({
        url: '/namespace/deleteNamespace',
        method: 'delete',
        data:{
            id:id,
            nodeId:nodeId
        }

    })
}
export function findRemoteNamespace(obj) {
    return request({
        url: '/namespace/findRemoteNamespace',
        method: 'post',
        data: obj

    })
}
export function syncNamespace(obj) {
    return request({
        url: '/namespace/syncNamespace',
        method: 'post',
        data: obj

    })
}