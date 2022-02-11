import HTTP from './axios'

export default {
    /**
     * 登录
     * @param {*} username 
     * @param {*} pwd 
     * @returns 
     */
    login(username, pwd) {
        return HTTP.post('login', {
            username,
            pwd
        });
    },
    /**
     * 注册用户
     * @param {String} name 
     * @param {String} pwd 
     */
    register(data) {
        return HTTP.post('register', data);
    },

    /**
     * 创建新文件
     * @param disk
     * @param path
     * @param name
     * @returns {AxiosPromise<any>}
     */
    createFile(disk, path, name) {
        return HTTP.post('create_file', {
            disk,
            path,
            name
        });
    },

    /**
     * 更新文件
     * @param formData
     * @returns {*}
     */
    updateFile(formData) {
        return HTTP.post('update_file', formData);
    },

    /**
     * 创建新目录
     * @param data
     * @returns {*}
     */
    createDirectory(data) {
        return HTTP.post('create_directory', data);
    },

    /**
     * 上传文件
     * @param data
     * @param config
     * @returns {AxiosPromise<any>}
     */
    upload(data, config) {
        return HTTP.post('upload', data, config);
    },

    /**
     * 删除选中的文件
     * @param data
     * @returns {*}
     */
    delete(data) {
        return HTTP.post('delete', data);
    },

    /**
     * 重命名文件或文件夹
     * @param data
     * @returns {*}
     */
    rename(data) {
        return HTTP.post('rename', data);
    },

    /**
     * 复制 或 剪切文件或文件夹
     * @param data
     * @returns {*}
     */
    paste(data) {
        return HTTP.post('paste', data);
    },

    /**
     * 压缩
     * @param data
     * @returns {*}
     */
    zip(data) {
        return HTTP.post('zip', data);
    },

    /**
     * 解压缩
     * @returns {*}
     * @param data
     */
    unzip(data) {
        return HTTP.post('unzip', data);
    },
}