export default {
  /**
   * 设置消息存在时的action结果
   * @param state
   * @param status
   * @param message
   */
  setActionResult(state, {
    status,
    message
  }) {
    state.actionResult.status = status;
    state.actionResult.message = message;
  },

  /**
   * 清除action结果
   * @param state
   */
  clearActionResult(state) {
    state.actionResult.status = null;
    state.actionResult.message = null;
  },

  /**
   * 上传进度条 (%)
   * @param state
   * @param progress
   */
  setProgress(state, progress) {
    state.actionProgress = progress;
  },

  /**
   * 清除进度
   * @param state
   */
  clearProgress(state) {
    state.actionProgress = 0;
  },

  /**
   * 记录 新操作
   * @param state
   */
  addLoading(state) {
    state.loading += 1;
  },

  /**
   * 操作结束
   * @param state
   */
  subtractLoading(state) {
    state.loading -= 1;
  },

  /**
   * 清除
   * @param state
   */
  clearLoading(state) {
    state.loading = 0;
  },

  /**
   * 设置错误信息
   * @param state
   * @param error
   */
  setError(state, error) {
    state.errors.push(error);
  },

  /**
   * 清除错误
   * @param state
   */
  clearErrors(state) {
    state.errors = [];
  },
};