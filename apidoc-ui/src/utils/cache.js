import Vue from "vue";

/**
 * @description: 缓存
 * @param {String} name key
 * @param {String, Object} value value
 * @return: cache value
 * @use 存： ls.set("name", value);
 *      取： const value = ls.get("name");
 */
export const ls = {
  set(name, value, expire) {
    if (expire) {
      return Vue.ls.set(name, value, expire);
    } else {
      return Vue.ls.set(name, value);
    }
  },
  get(name) {
    return Vue.ls.get(name);
  },
  remove(name) {
    return Vue.ls.remove(name);
  }
};
