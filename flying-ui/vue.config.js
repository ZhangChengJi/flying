const url = 'http://192.168.1.125:8888'
let publicPath = '/'
module.exports = {
    publicPath: publicPath,
    lintOnSave: false,
    css: {
        loaderOptions: {
            less: {
                lessOptions: {
                    modifyVars: {
                        'primary-color': '#42b983',
                        'link-color': '#42b983',
                        'border-radius-base': '3px'
                    },
                    javascriptEnabled: true,
                },
            },
        },
    },
    productionSourceMap: false,

    // 它支持webPack-dev-server的所有选项

    devServer: {

        proxy: {
            '/api': {
                target: url,
                ws: true,
                pathRewrite: {
                    '^/': '/'
                }
            }
        }
    }
    };