/* eslint-disable max-len,prefer-destructuring,object-curly-newline */
import GET from '../http/get';
import POST from '../http/post';

export default {
  /**
   * 从服务器获取初始化数据
   * @param state
   * @param commit
   * @param getters
   * @param dispatch
   */
  initializeApp({
    // state,
    commit,
    getters,
    dispatch
  }) {
    GET.initialize().then((response) => {
      if (response.data.result.status === 'success') {
        commit('settings/initSettings', response.data.config);
        commit('setDisks', response.data.config.disks);

        let leftDisk = response.data.config.leftDisk ?
          response.data.config.leftDisk :
          getters.diskList[0];

        // 路径
        let leftPath = response.data.config.leftPath;

        // 在URL中查找磁盘和路径设置
        if (window.location.search) {
          const params = new URLSearchParams(window.location.search);
          if (params.get('leftDisk')) {
            leftDisk = params.get('leftDisk');
          }

          if (params.get('leftPath')) {
            leftPath = params.get('leftPath');
          }

        }

        commit('left/setDisk', leftDisk);

        // 如果leftPath不为空,添加浏览目录路径 和 访问历史记录
        if (leftPath) {
          commit('left/setSelectedDirectory', leftPath);
          commit('left/addToHistory', leftPath);
        }

        dispatch('getLoadContent', {
          manager: 'left',
          disk: leftDisk,
          path: leftPath,
        });
        // if (state.settings.windowsConfig === 2) {
        console.log("leftDisk:", leftDisk, "leftPath:", leftPath);
        dispatch('tree/initTree', leftDisk).then(() => {
          if (leftPath) {
            // 路径不为空则重新打开
            dispatch('tree/reopenPath', leftPath);
          }
        });
        // }
      }
    });
  },

  /**
   * 获取指定磁盘下所有的文件及文件夹 
   * @param context
   * @param manager
   * @param disk
   * @param path
   */
  getLoadContent(context, {
    manager,
    disk,
    path
  }) {
    GET.content(disk, path).then((response) => {
      if (response.data.result.status === 'success') {
        context.commit(`${manager}/setDirectoryContent`, response.data);
      }
    });
  },

  /**
   * 选择磁盘
   * @param state
   * @param commit
   * @param dispatch
   * @param disk
   * @param manager
   */
  selectDisk({
    commit,
    dispatch
  }, {
    disk,
    manager
  }) {
    GET.selectDisk(disk).then((response) => {
      // 改变磁盘
      if (response.data.result.status === 'success') {
        // 设置磁盘名称
        commit(`${manager}/setDisk`, disk);

        // 重置 history
        commit(`${manager}/resetHistory`);

        // 如果目录树显示了，重新初始化
        // if (state.settings.windowsConfig === 2) {
        dispatch('tree/initTree', disk);
        // }

        // 下载根路径的内容
        dispatch(`${manager}/selectDirectory`, {
          path: null,
          history: false
        });
      }
    });
  },

  /**
   * 创建新文件
   * @param getters
   * @param dispatch
   * @param fileName
   * @returns {Promise}
   */
  createFile({
    getters,
    dispatch
  }, fileName) {
    //新文件的目录
    const selectedDirectory = getters.selectedDirectory;

    // 服务器端创建新文件
    return POST.createFile(getters.selectedDisk, selectedDirectory, fileName)
      .then((response) => {
        // 更新文件列表
        dispatch('updateContent', {
          response,
          oldDir: selectedDirectory,
          commitName: 'addNewFile',
          type: 'file',
        });

        return response;
      });
  },

  /**
   * 获取文件内容
   * @param context
   * @param disk
   * @param path
   * @returns {*}
   */
  getFile(context, {
    disk,
    path
  }) {
    return GET.getFile(disk, path);
  },

  /**
   * 更新文件
   * @param getters
   * @param dispatch
   * @param formData
   * @returns {PromiseLike | Promise}
   */
  updateFile({
    getters,
    dispatch
  }, formData) {
    return POST.updateFile(formData).then((response) => {
      // 更新文件列表
      dispatch('updateContent', {
        response,
        oldDir: getters.selectedDirectory,
        commitName: 'updateFile',
        type: 'file',
      });

      return response;
    });
  },

  /**
   * 创建新目录
   * @param getters
   * @param dispatch
   * @param name
   * @returns {*}
   */
  createDirectory({
    getters,
    dispatch
  }, name) {
    // 新文件夹的目录
    const selectedDirectory = getters.selectedDirectory;

    // 创建新目录，服务器端
    return POST.createDirectory({
      disk: getters.selectedDisk,
      path: selectedDirectory,
      name,
    }).then((response) => {
      // 更新文件列表
      dispatch('updateContent', {
        response,
        oldDir: selectedDirectory,
        commitName: 'addNewDirectory',
        type: 'directory',
      });

      return response;
    });
  },

  /**
   * 上传文件
   * @param getters
   * @param commit
   * @param dispatch
   * @param files
   * @param overwrite
   * @returns {Promise}
   */
  upload({
    getters,
    commit,
    dispatch
  }, {
    files,
    overwrite,
    fileParam
  }) {
    // 将要上传文件的目录
    const selectedDirectory = getters.selectedDirectory;
    const data = new FormData();
    data.append('disk', getters.selectedDisk);
    data.append('path', selectedDirectory || '');
    data.append('overwrite', overwrite);
    // 添加文件
    for (let i = 0; i < files.length; i += 1) {
      data.append('files[]', files[i]);
    }

    // axios 配置- 进度条
    const config = {
      onUploadProgress(progressEvent) {
        const progress = Math.round((progressEvent.loaded * 100) / progressEvent.total);
        fileParam.file.percent = progress;
        fileParam.onProgress(fileParam.file)
        commit('messages/setProgress', progress);
      },
    };

    // 上传文件
    return POST.upload(data, config).then((response) => {
      // 清除进度条
      commit('messages/clearProgress');

      // 上传成功
      if (
        response.data.result.status === 'success' &&
        selectedDirectory === getters.selectedDirectory
      ) {
        // 刷新内容
        dispatch('refreshManagers');
      }

      return response;
    }).catch(() => {
      // 清除进度条
      commit('messages/clearProgress');
    });
  },

  /**
   * 删除选择的文件和文件夹
   * @param state
   * @param getters
   * @param dispatch
   * @param items
   * @returns {*}
   */
  delete({
    getters,
    dispatch
  }, items) {
    return POST.delete({
      disk: getters.selectedDisk,
      items,
    }).then((response) => {
      // 如果成功删除所有项目
      if (response.data.result.status === 'success') {
        // 刷新内容
        dispatch('refreshManagers');

        // 从目录树中删除文件夹
        // if (state.settings.windowsConfig === 2) {
        const onlyDir = items.filter((item) => item.type === 'dir');
        dispatch('tree/deleteFromTree', onlyDir);
        // }
      }

      return response;
    });
  },

  /**
   * 粘贴文件和文件夹
   * @param state
   * @param commit
   * @param getters
   * @param dispatch
   */
  paste({
    state,
    commit,
    getters,
    dispatch
  }) {
    POST.paste({
      disk: getters.selectedDisk,
      path: getters.selectedDirectory,
      clipboard: state.clipboard,
    }).then((response) => {
      if (response.data.result.status === 'success') {
        // 刷新内容
        // dispatch('refreshAll');

        // 剪切、清除剪贴板
        if (state.clipboard.type === 'cut') {
          commit('resetClipboard');
        }
      }
      dispatch('refreshAll');
    });
  },

  /**
   * 重命名文件或文件夹
   * @param getters
   * @param dispatch
   * @param type
   * @param newName
   * @param oldName
   * @returns {Promise}
   */
  rename({
    getters,
    dispatch
  }, {
    type,
    newName,
    oldName
  }) {
    return POST.rename({
      disk: getters.selectedDisk,
      newName,
      oldName,
      type,
    }).then((response) => {
      // 刷新内容
      if (type === 'dir') {
        dispatch('refreshAll');
      } else {
        dispatch('refreshManagers');
      }

      return response;
    });
  },

  /**
   * 获取文件url
   * @param store
   * @param disk
   * @param path
   * @returns {Promise}
   */
  url(store, {
    disk,
    path
  }) {
    return GET.url(disk, path);
  },

  /**
   * 压缩文件和文件夹
   * @param state
   * @param getters
   * @param dispatch
   * @param name
   * @returns {*}
   */
  zip({
    state,
    getters,
    dispatch
  }, name) {
    const selectedDirectory = getters.selectedDirectory;

    return POST.zip({
      disk: getters.selectedDisk,
      path: selectedDirectory,
      name,
      elements: state[state.activeManager].selected,
    }).then((response) => {
      // 如果压缩成功
      if (response.data.result.status === 'success' &&
        selectedDirectory === getters.selectedDirectory
      ) {
        // 刷新内容
        dispatch('refreshManagers');
      }

      return response;
    });
  },

  /**
   * 解压缩
   * @param getters
   * @param dispatch
   * @param folder
   * @returns {*}
   */
  unzip({
    getters,
    dispatch
  }, folder) {
    const selectedDirectory = getters.selectedDirectory;

    return POST.unzip({
      disk: getters.selectedDisk,
      path: getters.selectedItems[0].path,
      folder,
    }).then((response) => {
      // 解压成功
      if (response.data.result.status === 'success' &&
        selectedDirectory === getters.selectedDirectory
      ) {
        // 刷新
        dispatch('refreshAll');
      }
      return response;
    });
  },

  /**
   * 将所选项目添加到剪贴板
   * @param state
   * @param commit
   * @param getters
   * @param type
   */
  toClipboard({
    state,
    commit,
    getters
  }, type) {
    // 如果选择了文件
    if (getters[`${state.activeManager}/selectedCount`]) {
      commit('setClipboard', {
        type,
        disk: state[state.activeManager].selectedDisk,
        directories: state[state.activeManager].selected.directories.slice(0),
        files: state[state.activeManager].selected.files.slice(0),
      });
    }
  },

  /**
   * 刷新管理器
   * @param dispatch
   * @param state
   * @returns {*}
   */
  refreshManagers({
    dispatch
  }) {
    // 选择需要更新的内容
    // if (state.settings.windowsConfig === 3) {
    //   return Promise.all([
    //     dispatch('left/refreshDirectory'),
    //     dispatch('right/refreshDirectory'),
    //   ]);
    // }

    return dispatch('left/refreshDirectory');
  },

  /**
   * 刷新所有
   * @param state
   * @param getters
   * @param dispatch
   * @returns {*}
   */
  refreshAll({
    state,
    getters,
    dispatch
  }) {
    // if (state.settings.windowsConfig === 2) {
    // 刷新目录树
    return dispatch('tree/initTree', state.left.selectedDisk).then(() => Promise.all([
      // 重新打开目录
      dispatch('tree/reopenPath', getters.selectedDirectory),
      // 刷新管理器
      dispatch('refreshManagers'),
    ]));
    // }
    // 刷新管理器
    // return dispatch('refreshManagers');
  },

  /**
   * 重复排序
   * @param state
   * @param dispatch
   * @param manager
   */
  repeatSort({
    state,
    dispatch
  }, manager) {
    dispatch(`${manager}/sortBy`, {
      field: state[manager].sort.field,
      direction: state[manager].sort.direction,
    });
  },

  /**
   * 更新内容-创建或更新后的文件、文件夹
   * @param state
   * @param commit
   * @param getters
   * @param dispatch
   * @param response
   * @param oldDir
   * @param commitName
   * @param type
   */
  updateContent({
    state,
    commit,
    getters,
    dispatch
  }, {
    response,
    oldDir,
    commitName,
    type
  }) {
    // 如果操作成功
    if (
      response.data.result.status === 'success' &&
      oldDir === getters.selectedDirectory
    ) {
      // 在文件列表中更新替换之前的旧文件
      commit(`${state.activeManager}/${commitName}`, response.data[type]);
      // 重复排序
      // dispatch('repeatSort', state.activeManager);
      if (type === 'directory') {
        // 更新树模块
        dispatch('tree/addToTree', {
          parentPath: oldDir,
          newDirectory: response.data.tree,
        });
      } else {
        // 在非当前磁盘的“文件/文件夹”列表中添加/更新文件/文件夹
        commit(`${getters.inactiveManager}/${commitName}`, response.data[type]);
        // 重复排序
        dispatch('repeatSort', getters.inactiveManager);
      }

      // 树模块显示
      // if (type === 'directory' && state.settings.windowsConfig === 2) {
      // // 更新树模块
      // dispatch('tree/addToTree', {
      //   parentPath: oldDir,
      //   newDirectory: response.data.tree,
      // });
      // 如果两个管理器显示同一个文件夹
      // } else if (
      //   state.settings.windowsConfig === 3 &&
      //   state.left.selectedDirectory === state.right.selectedDirectory &&
      //   state.left.selectedDisk === state.right.selectedDisk
      // ) {
      //   // 在“文件/文件夹”列表中添加/更新文件/文件夹（非活动管理器）
      //   commit(`${getters.inactiveManager}/${commitName}`, response.data[type]);
      //   // 重复排序
      //   dispatch('repeatSort', getters.inactiveManager);
      // }
    }
  },

  /**
   * 重置程序状态
   * @param state
   * @param commit
   */
  resetState({
    commit
  }) {
    // 左管理器
    commit('left/setDisk', null);
    commit('left/setSelectedDirectory', null);
    commit('left/setDirectoryContent', {
      directories: [],
      files: []
    });
    commit('left/resetSelected');
    commit('left/resetSortSettings');
    commit('left/resetHistory');
    commit('left/setView', 'table');
    // 清除模块
    commit('modal/clearModal');
    // 消息
    commit('messages/clearActionResult');
    commit('messages/clearProgress');
    commit('messages/clearLoading');
    commit('messages/clearErrors');

    // if (state.settings.windowsConfig === 3) {
    //   // 右管理器
    //   commit('right/setDisk', null);
    //   commit('right/setSelectedDirectory', null);
    //   commit('right/setDirectoryContent', {
    //     directories: [],
    //     files: []
    //   });
    //   commit('right/resetSelected');
    //   commit('right/resetSortSettings');
    //   commit('right/resetHistory');
    //   commit('right/setView', 'table');
    // } else if (state.settings.windowsConfig === 2) {
    commit('tree/cleanTree');
    commit('tree/clearTempArray');
    // }

    commit('resetState');
  },

  /**
   * 打开 PDF
   * @param context
   * @param disk
   * @param path
   */
  openPDF(context, {
    disk,
    path
  }) {
    const win = window.open();

    GET.getFileArrayBuffer(disk, path).then((response) => {
      const blob = new Blob([response.data], {
        type: 'application/pdf'
      });

      win.document.write(`<iframe src="${URL.createObjectURL(blob)}" allowfullscreen height="100%" width="100%"></iframe>`);
    });
  },

  /**
   * 请求获取登录状态
   * @param {*} param
   */
  checkLoginStatus({
    commit
  }) {
    return GET.loginStatus().then(response => {
      commit('setLoginStatus', response.data.isLogin)
      return response;
    })
  }

};