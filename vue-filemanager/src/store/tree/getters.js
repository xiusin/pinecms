export default {
  /**
   * 在目录树中查找目录索引
   * @param state
   * @returns {Number}
   */
  findDirectoryIndex: (state) => (path) => state.directories.findIndex((el) => el.path === path),

  /**
   * 筛选目录列表
   * @param state
   * @param getters
   * @param rootState
   * @returns {*}
   */
  directories(state, getters, rootState) {
    if (rootState.fm.settings.hiddenFiles) {
      return state.directories;
    }
    return state.directories.filter((item) => item.basename.match(new RegExp('^([^.]).*', 'i')));
  },
};