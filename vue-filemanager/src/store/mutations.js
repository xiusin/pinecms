/* eslint-disable object-curly-newline */
export default {
  /**
   * 设置磁盘从服务器返回的磁盘
   * @param state
   * @param disks
   */
  setDisks(state, disks) {
    state.disks = disks;
  },

  /**
   * 设置剪贴板
   * @param state
   * @param type
   * @param disk
   * @param directories
   * @param files
   */
  setClipboard(state, {
    type,
    disk,
    directories,
    files
  }) {
    state.clipboard.type = type;
    state.clipboard.disk = disk;
    state.clipboard.directories = directories;
    state.clipboard.files = files;
  },

  /**
   * 删除剪贴板中的相应内容
   * @param state
   * @param type 文件类型
   * @param path 文件路径
   */
  truncateClipboard(state, {
    type,
    path
  }) {
    const itemIndex = state.clipboard[type].indexOf(path);
    if (itemIndex !== -1) state.clipboard[type].splice(itemIndex, 1);
    if (!state.clipboard.directories.length && !state.clipboard.files.length) {
      state.clipboard.type = null;
    }
  },

  /**
   * 重置剪贴板
   * @param state
   */
  resetClipboard(state) {
    state.clipboard.type = null;
    state.clipboard.disk = null;
    state.clipboard.directories = [];
    state.clipboard.files = [];
  },

  /**
   * 选择管理器（显示2个文件管理器窗口时）
   * @param state
   * @param managerName
   */
  setActiveManager(state, managerName) {
    state.activeManager = managerName;
  },

  /**
   * 设置 文件回调
   * @param state
   * @param callback
   */
  setFileCallBack(state, callback) {
    state.fileCallback = callback;
  },

  /**
   * 屏幕模式切换
   * @param state
   */
  // screenToggle(state) {
  //   state.fullScreen = !state.fullScreen;
  // },

  /**
   * 重置状态
   * @param state
   */
  resetState(state) {
    state.activeManager = 'left';
    state.clipboard = {
      type: null,
      disk: null,
      directories: [],
      files: [],
    };
    state.disks = [];
    state.fileCallback = null;
    // state.fullScreen = false;
  },

  /**
   * 设置登录状态
   * @param {*} state 
   * @param {Boolean} status 
   */
  setLoginStatus(state, status) {
    state.isLogin = status;
  },
  /**
   * 设置自动登录
   * @param {*} state 
   * @param {Boolean} isAutoLogin 
   */
  setAutoLogin(state, isAutoLogin) {
    state.autoLogin = isAutoLogin;
  },
  /**
   * 设置用户名
   * @param {*} state 
   * @param {String} name 
   */
  setUserName(state, {
    name,
    nickname
  }) {
    state.username = name;
    state.nickname = nickname;
  }
};