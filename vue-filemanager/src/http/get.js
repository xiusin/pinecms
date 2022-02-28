import HTTP from './axios'
export default {

    /**
     * 获取登录状态
     * @returns {*}
     */
    loginStatus() {
        return HTTP.get('me');
    },

    /**
     * 退出
     * @returns
     */
    logout() {
        return HTTP.get('logout');
    },
    /**
     * 校验用户名
     * @param {String} name
     * @returns
     */
    findUsername(name) {
        return HTTP.get('validate-name', {
            params: {
                name
            }
        })
    },

    /**
     * 查找文件创建者
     * @param {Stirng} filePath
     * @returns
     */
    findAuthor(filePath) {
        return HTTP.get('file-author', {
            params: {
                filePath
            }
        })
    },

    /**
     * 从服务器获取配置数据
     * @returns {*}
     */
    initialize() {
        return HTTP.get('initialize');
    },
    /**
     * 获取指定磁盘下的文件夹和文件
     * @param disk
     * @param path
     * @returns {*}
     */
    content(disk, path) {
        return HTTP.get('content', {
            params: {
                disk,
                path
            }
        });
    },
    /**
     * 获取目录
     * @param {*} disk
     * @param {*} path
     */
    tree(disk, path) {
        return HTTP.get('tree', {
            params: {
                disk,
                path
            }
        })
    },
    /**
     * 选择磁盘
     * @param disk
     * @returns {*}
     */
    selectDisk(disk) {
        return HTTP.get('select-disk', {
            params: {
                disk
            }
        });
    },

    /**
     * 属性
     */
    /* properties(disk, path) {
      return HTTP.get('properties', { params: { disk, path } });
    }, */

    /**
     * URL
     * @param disk
     * @param path
     * @returns {*}
     */
    url(disk, path) {
        return HTTP.get('url', {
            params: {
                disk,
                path
            }
        });
    },

    /**
     * 选择一个文件去编辑或显示
     * @param disk
     * @param path
     * @returns {*}
     */
    getFile(disk, path) {
        return HTTP.get('download', {
            params: {
                disk,
                path
            }
        });
    },

    /**
     * 获取文件- ArrayBuffer
     * @param disk
     * @param path
     * @returns {*}
     */
    getFileArrayBuffer(disk, path) {
        return HTTP.get('download', {
            responseType: 'arraybuffer',
            params: {
                disk,
                path
            },
        });
    },

    /**
     * 图片缩略图
     * @param disk
     * @param path
     * @returns {*}
     */
    // thumbnail(disk, path) {
    //   return HTTP.get('thumbnails', {
    //     responseType: 'arraybuffer',
    //     params: { disk, path },
    //   });
    // },

    /**
     * 图片缩略图链接
     * @param {*} disk
     * @param {*} path
     */
    thumbnailLink(disk, path) {
        return HTTP.get('thumbnails_link', {
            params: {
                disk,
                path
            },
        });
    },

    /**
     * 图片预览
     * @param disk
     * @param path
     * @return {*}
     */
    preview(disk, path) {
        return HTTP.get('preview', {
            responseType: 'arraybuffer',
            params: {
                disk,
                path
            },
        });
    },

    /**
     * 下载文件
     * @param disk
     * @param path
     * @return {*}
     */
    download(disk, path) {
        return HTTP.get('download', {
            responseType: 'arraybuffer',
            params: {
                disk,
                path
            },
        });
    },

    /**
     * 下载文件
     * @param disk
     * @param path
     * @return {*}
     */
    downloadFile(disk, path) {
        return HTTP.get('download_file', {
            // responseType: 'arraybuffer',
            params: {
                disk,
                path
            },
        });
    },
}
