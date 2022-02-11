const path = require('path');
const resolve = dir => path.join(__dirname, dir);
const CompressionWebpackPlugin = require('compression-webpack-plugin');
const UglifyJsPlugin = require('uglifyjs-webpack-plugin');
const productionGzipExtensions = /\.(js|css|json|txt|html|ico|svg)(\?.*)?$/i;
const IS_PROD = ["production", "prod"].includes(process.env.NODE_ENV);
/* eslint-disable no-irregular-whitespace */
module.exports = {
    publicPath: './',
    outputDir: '../src/application/controllers/backend/filemanager/dist',
    productionSourceMap: false,
    indexPath: 'index.html', //指定生成的 index.html 的输出路径(相对于 outputDir)也可以是一个绝对路径。
    css: {
        extract: true, //是否使用css分离插件ExtractTextPlugin
        sourceMap: !IS_PROD, //开启CSS source maps
        loaderOptions: {}, // css预设器配置项
        requireModuleExtension: true // 启用 CSS modules for all css / pre-processor files.
    },
    chainWebpack: config => {
        config.plugins.delete('preload')
        config.plugins.delete('prefetch')
        const cdn = {
            css: [
                // 'https://cdn.bootcdn.net/ajax/libs/twitter-bootstrap/4.5.0/css/bootstrap.min.css',
                'https://cdn.bootcdn.net/ajax/libs/font-awesome/5.14.0/css/all.css',
                'https://cdn.bootcdn.net/ajax/libs/cropperjs/1.5.9/cropper.min.css',
                // 'https://cdn.bootcdn.net/ajax/libs/plyr/3.6.2/plyr.css'
            ],
            js: [
                'https://cdn.bootcdn.net/ajax/libs/cropperjs/1.5.9/cropper.min.js',
                // 'https://cdn.bootcdn.net/ajax/libs/plyr/3.6.2/plyr.min.js',
            ]
        };
        if (IS_PROD) {
            config.module.rule("images")
                .use("image-webpack-loader")
                .loader("image-webpack-loader")
                .options({
                    bypassOnDebug: true
                    // mozjpeg: {
                    //     progressive: true,
                    //     quality: 65
                    // },
                    // optipng: {
                    //     enabled: false
                    // },
                    // pngquant: {
                    //     quality: [0.65, 0.9],
                    //     speed: 4
                    // },
                    // gifsicle: {
                    //     interlaced: false
                    // }
                    // webp: { quality: 75 } //ios不支持
                }).end()
        }
        // 添加别名
        config.resolve.alias
            .set("@", resolve("src"))
            .set("components", resolve("src/components"));
        config
            .plugin('html')
            .tap(args => {
                args[0].title = '文件管理'
                args[0].cdn = cdn
                return args
            })

    },
    configureWebpack: config => {
        const plugins = [];
        config.externals = {
                // 'bootstrap': 'bootstrap',
                'cropper': 'Cropper',
                // 'plyr': 'Plyr'
            },
            config.optimization = {
                splitChunks: {
                    cacheGroups: {
                        common: {
                            name: "chunk-common",
                            chunks: "initial",
                            minChunks: 2,
                            maxInitialRequests: 5,
                            minSize: 0,
                            priority: 1,
                            reuseExistingChunk: true,
                            enforce: true
                        },
                        vendors: {
                            name: "chunk-vendors",
                            test: /[\\/]node_modules[\\/]/,
                            chunks: "initial",
                            priority: 2,
                            reuseExistingChunk: true,
                            enforce: true
                        },
                        codemirror: {
                            name: "chunk-codemirror",
                            test: /[\\/]node_modules[\\/](vue-)?codemirror[\\/]/,
                            chunks: "all",
                            priority: 3,
                            reuseExistingChunk: true,
                            enforce: true
                        }
                    }
                }
            },
            plugins.push(
                new CompressionWebpackPlugin({
                    filename: '[path][base].gz',
                    algorithm: 'gzip',
                    test: productionGzipExtensions,
                    threshold: 10240,
                    minRatio: 0.6
                }),
                new UglifyJsPlugin({
                    uglifyOptions: {
                        compress: {
                            drop_console: true,
                            drop_debugger: false,
                            pure_funcs: ['console.log'] //移除console
                        }
                    },
                    sourceMap: false,
                    parallel: true
                })
            );
        config.plugins = [
            ...config.plugins,
            ...plugins,
        ]
    },
    // devServer: {
    //     proxy: {
    //         '/': {
    //             target: "http://localhost:3000",
    //             secure: false,
    //             changeOrigin: true,

    //         }
    //     }
    // }
}