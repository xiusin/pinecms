export default {
  /**
   * 设置将显示的模块
   * @param state
   * @param show
   * @param modalName
   */
  setModalState(state, {
    show,
    modalName
  }) {
    state.showModal = show;
    state.modalName = modalName;
  },

  /**
   * 关闭模块
   * @param state
   */
  clearModal(state) {
    state.showModal = false;
    state.modalName = null;
  },

  /**
   * 设置模块的高度
   * @param state
   * @param height
   */
  setModalBlockHeight(state, height) {
    state.modalBlockHeight = height;
  },
};