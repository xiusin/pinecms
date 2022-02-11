import mutations from './mutations';

export default {
  namespaced: true,
  state() {
    return {
      actionResult: {
        status: null,
        message: null,
      },

      // 进度条完成状态
      actionProgress: 0,

      // 加载spinner
      loading: 0,

      // 应用错误信息
      errors: [],
    };
  },
  mutations,
};