export default {
  /**
   * 文件列表
   * @param state
   * @param getters
   * @param rootState
   */
  files(state, getters, rootState) {
    if (rootState.fm.settings.hiddenFiles) {
      return state.files;
    }

    return state.files.filter((item) => item.basename.match(new RegExp('^([^.]).*', 'i')));
  },
  /**
   * 排序字段
   * @param {*} state 
   */
  sort(state) {
    return state.sort;
  },
  /**
   * 目录列表
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

  /**
   * 图片文件
   * @param {*} state 
   * @param {*} getters 
   * @param {*} rootState 
   * @returns 
   */
  imageFiles(state, getters, rootState) {
    return state.files.filter((item) => rootState.fm.settings.imageExtensions.includes(item.extension.toLowerCase()))
  },
  /**
   * 图片文件链接
   * @param {*} state 
   * @returns 
   */
  imageUrl(state) {
    return state.imageUrl;
  },
  /**
   * 面包屑(获取当前选择的目录)
   * @param state
   * @returns {*}
   */
  breadcrumb(state) {
    if (state.selectedDirectory) {
      return state.selectedDirectory.split('/');
    }
    return null;
  },
  /**
   * 当前路径所有者
   * @param {*} state 
   * @returns 
   */
  selectedDirectoryOwner(state) {
    return state.selectedDirectoryOwner;
  },
  /**
   * 文件数量
   * @param state
   * @param getters
   * @returns {*}
   */
  filesCount(state, getters) {
    return getters.files.length;
  },

  /**
   * 目录计算器
   * @param state
   * @param getters
   * @returns {*}
   */
  directoriesCount(state, getters) {
    return getters.directories.length;
  },

  /**
   * 文件大小 - bytes
   * @param state
   * @param getters
   * @returns {*}
   */
  filesSize(state, getters) {
    if (getters.files.length) {
      return getters.files.reduce((previous, current) => previous + Number(current.size), 0);
    }
    return 0;
  },

  /**
   * 选择的文件和文件夹数量
   * @param state
   * @param getters
   * @returns {Number}
   */
  selectedCount(state, getters) {
    return getters.selectedList.length;
  },

  /**
   * 选中的文件大小
   * @param state
   * @returns {Number}
   */
  selectedFilesSize(state) {
    const selectedFiles = state.files.filter((file) => state.selected.files.includes(file.path));
    if (selectedFiles.length) {
      return selectedFiles.reduce((previous, current) => previous + Number(current.size), 0);
    }
    return 0;
  },

  /**
   * 选择的文件和文件夹 (1)
   * @param state
   */
  selectedList(state) {
    const selectedDirectories = state.directories.filter((directory) => state.selected.directories.includes(directory.path));
    const selectedFiles = state.files.filter((file) => state.selected.files.includes(file.path));
    return selectedDirectories.concat(selectedFiles);
  },
  /**
   * @returns  是否全选
   */
  isCheckedAll(state) {
    return state.isCheckedAll;
  },
  /** 
   * @returns 全选效果属性
   */
  isIndeterminate(state) {
    return state.isIndeterminate;
  },
  /**
   * @returns 返回选择的文件和文件夹
   */
  checkedFilesList(state, getters) {
    return getters.selectedList;
  },

  /**
   * 比较文件夹名
   * @param state
   */
  directoryExist: (state) => (basename) => state.directories.some((el) => el.basename === basename),

  /**
   * 比较文件名
   * @param state
   */
  fileExist: (state) => (basename) => state.files.some((el) => el.basename === basename),
};