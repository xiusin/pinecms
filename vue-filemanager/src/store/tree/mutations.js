/* eslint-disable no-param-reassign,no-restricted-syntax */
export default {
  /**
   * 清理左侧的目录树
   * @param state
   */
  cleanTree(state) {
    state.directories = [];
    state.counter = 1;
  },

  /**
   * 向左侧的目录树中添加目录
   * @param state
   * @param directories
   * @param parentId
   */
  addDirectories(state, {
    parentId,
    directories
  }) {
    directories.forEach((directory) => {
      // 向目录添加属性
      directory.id = state.counter;
      directory.parentId = parentId;
      directory.props.subdirectoriesLoaded = false;
      directory.props.showSubdirectories = false;

      state.counter += 1;

      state.directories.push(directory);
    });
  },

  /**
   * 替换目录
   * @param state
   * @param directories
   */
  replaceDirectories(state, directories) {
    state.directories = directories;
  },

  /**
   * 更新目录属性
   * @param state
   * @param index
   * @param props
   */
  updateDirectoryProps(state, {
    index,
    props
  }) {
    for (const property in props) {
      if (Object.prototype.hasOwnProperty.call(props, property)) {
        state.directories[index].props[property] = props[property];
      }
    }
  },

  /**
   * 添加到临时索引数组
   * @param state
   * @param index
   */
  addToTempArray(state, index) {
    state.tempIndexArray.push(index);
  },

  /**
   * 清除临时索引数组
   * @param state
   */
  clearTempArray(state) {
    state.tempIndexArray = [];
  },
};