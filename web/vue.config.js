'use strict'

module.exports = {
    publicPath: './',
    outputDir: 'dist',
    assetsDir: 'static',
    productionSourceMap: false,
    devServer: {
        port: 8080,
        open: true,
        overlay: {
            warnings: false,
            errors: true
        },
        proxy: {
            ['/api']: { //需要代理的路径
                target: `http://127.0.0.1:8321/`, //代理到 目标路径
                changeOrigin: true,
                pathRewrite: { // 修改路径数据
                    ['^/api']: '/api' // '^/api:""' 把路径中的/api字符串删除
                }
            }
        },
    }
}