export default {
  /**
   * 获取磁盘列表 (1)
   * @param state
   * @returns {String[]}
   */
  diskList(state) {
    return Object.keys(state.disks);
  },

  /**
   * 选择磁盘
   * @param state
   * @returns {*}
   */
  selectedDisk(state) {
    return state[state.activeManager].selectedDisk;
  },

  /**
   * 为活动管理器选择目录
   * @param state
   * @returns {*}
   */
  selectedDirectory(state) {
    return state[state.activeManager].selectedDirectory;
  },

  /**
   * 获取用户选择的文件、文件夹
   * @param state
   * @param getters
   * @returns {*}
   */
  selectedItems(state, getters) {
    return getters[`${state.activeManager}/selectedList`];
  },

  /**
   * 非活动状态的管理器
   * @param state
   * @returns {String}
   */
  inactiveManager(state) {
    return state.activeManager === 'left' ? 'right' : 'left';
  },

  /**
   * 获取用户名
   * @param {*} state 
   * @returns 
   */
  getUsername(state) {
    return {
      username: state.username,
      nickname: state.nickname
    };
  }
};