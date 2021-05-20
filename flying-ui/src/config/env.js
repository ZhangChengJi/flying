// 配置编译环境和线上环境之间的切换

const env = process.env
const baseUrl = ''
const codeUrl = `${window.location.origin}/code`

export {
    baseUrl,
    codeUrl,
    env
}