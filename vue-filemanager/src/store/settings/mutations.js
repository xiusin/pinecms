import Vue from 'vue';

export default {
  /**
   * 设置配置
   * @param state
   * @param data
   */
  manualSettings(state, data) {
    // 重写 Axios 的 headers
    if (Object.prototype.hasOwnProperty.call(data, 'headers')) {
      state.headers = data.headers;
    }
    // axios请求的公共url
    if (Object.prototype.hasOwnProperty.call(data, 'baseUrl')) {
      state.baseUrl = data.baseUrl;
    }
    // windows配置
    if (Object.prototype.hasOwnProperty.call(data, 'windowsConfig')) {
      state.windowsConfig = data.windowsConfig;
    }
    // 语言
    if (Object.prototype.hasOwnProperty.call(data, 'lang')) {
      state.lang = data.lang;
    }
    // 译文
    if (Object.prototype.hasOwnProperty.call(data, 'translation')) {
      //在对象添加属性
      Vue.set(state.translations, data.translation.name, Object.freeze(data.translation.content));
    }
  },

  /**
   * 初始化 Axios 的 baseUrl 和 headers
   * @param state
   */
  initAxiosSettings(state) {
    // 如果未手动设置，则初始化基本url
    if (!state.baseUrl) {
      if (process.env.VUE_APP_AXIOS_BASE_URL) {
        // vue .env变量
        state.baseUrl = process.env.VUE_APP_AXIOS_BASE_URL;
      } else {
        let baseUrl = `${window.location.protocol}//${window.location.hostname}`;
        if (window.location.port.length) {
          baseUrl += `:${window.location.port}/file/`;
        } else {
          baseUrl += '/file/';
        }
        state.baseUrl = baseUrl;
      }
    }

    // 如果未手动设置，则初始化基本headers
    if (Object.keys(state.headers).length === 0) {
      // 如有需要，关闭laravel csrf-token
      if (process.env.VUE_APP_CSRF_TOKEN === 'OFF') {
        state.headers = {
          'X-Requested-With': 'XMLHttpRequest'
        };
      } else {
        // CSRF token
        const token = document.head.querySelector('meta[name="csrf-token"]');
        if (!token) {
          state.headers = {
            'X-Requested-With': 'XMLHttpRequest',
          };
          // eslint-disable-next-line
          console.error('CSRF token not found: https://laravel.com/docs/csrf#csrf-x-csrf-token');
        } else {
          state.headers = {
            'X-Requested-With': 'XMLHttpRequest',
            'X-CSRF-TOKEN': token.content,
          };
        }
      }
    }
  },

  /**
   * 从服务器初始化应用设置
   * @param state
   * @param data
   */
  initSettings(state, data) {
    // if (!state.lang) state.lang = data.lang;
    // if (!state.windowsConfig) state.windowsConfig = data.windowsConfig;
    state.acl = data.acl;
    // state.hiddenFiles = data.hiddenFiles;
  },

  /**
   * 显示/隐藏文件
   * @param state
   */
  toggleHiddenFiles(state) {
    state.hiddenFiles = !state.hiddenFiles;
  },
};