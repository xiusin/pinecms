import mutations from './mutations';

export default {
  namespaced: true,
  state() {
    return {
      // 模块窗口
      showModal: false,

      // 模块名称
      modalName: null,

      // 主模块高度
      modalBlockHeight: 0,
    };
  },
  mutations,
};
