export default {
  /**
   * 公共请求URL
   * @param state
   * @returns {String}
   */
  baseUrl(state) {
    return state.baseUrl;
  },

  /**
   * 请求头
   * @param state
   * @return {*}
   */
  headers(state) {
    return state.headers;
  },

  /**
   * 请求头是否包含Authorization
   * @param state
   * @return {Boolean}
   */
  authHeader(state) {
    return Object.prototype.hasOwnProperty.call(state.headers, 'Authorization');
  },
};