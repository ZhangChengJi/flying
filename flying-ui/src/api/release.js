import request from "@/router/axios";

export function createRelease(obj) {
    return request({
        url: '/release/createRelease',
        method: 'post',
        data: obj
    })
}