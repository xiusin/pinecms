import GET from '../../http/get';

export default {
  /**
   * 初始化目录树
   * @param state
   * @param commit
   * @param disk
   * @returns {Promise}
   */
  initTree({
    state,
    commit
  }, disk) {
    return GET.tree(disk, null).then((response) => {
      // 成功
      if (response.data.result.status === 'success') {
        // 清除旧目录树
        if (state.directories) commit('cleanTree');

        // 初始化目录树
        commit('addDirectories', {
          parentId: 0,
          directories: response.data.directories,
        });
      }
    });
  },

  /**
   * 给目录树添加新文件夹
   * @param state
   * @param commit
   * @param getters
   * @param parentPath
   * @param newDirectory
   */
  addToTree({
    state,
    commit,
    getters
  }, {
    parentPath,
    newDirectory
  }) {
    //如果不是根目录
    if (parentPath) {
      // 找到父目录的下标
      const parentDirectoryIndex = getters.findDirectoryIndex(parentPath);

      if (parentDirectoryIndex !== -1) {
        console.log("newDirectory:", newDirectory);
        // 添加新目录
        commit('addDirectories', {
          directories: newDirectory,
          parentId: state.directories[parentDirectoryIndex].id,
        });

        // 更新父目录属性
        commit('updateDirectoryProps', {
          index: parentDirectoryIndex,
          props: {
            hasSubdirectories: true,
            showSubdirectories: true,
            subdirectoriesLoaded: true,
          },
        });
      } else {
        commit('fm/messages/setError', {
          message: '文件夹未找到!'
        }, {
          root: true
        });
      }
    } else {
      // 添加磁盘根目录
      commit('addDirectories', {
        directories: newDirectory,
        parentId: 0,
      });
    }
  },

  /**
   * 删除目录和子目录
   * @param state
   * @param commit
   * @param getters
   * @param dispatch
   * @param directories
   */
  deleteFromTree({
    state,
    commit,
    getters,
    dispatch,
  }, directories) {
    directories.forEach((item) => {
      // 查找目录
      const directoryIndex = getters.findDirectoryIndex(item.path);

      if (directoryIndex !== -1) {
        // 根据下标删除对应目录
        commit('addToTempArray', directoryIndex);

        // 如果有子目录
        if (state.directories[directoryIndex].props.hasSubdirectories) {
          // 查找子目录
          dispatch('subDirsFinder', state.directories[directoryIndex].id);
        }
      }
    });

    // 过滤目录
    const temp = state.directories.filter((item, index) => {
      if (state.tempIndexArray.indexOf(index) === -1) {
        return item;
      }
      return false;
    });

    // 替换
    commit('replaceDirectories', temp);

    // 清除临时数组
    commit('clearTempArray');
  },

  /**
   * 查找子目录
   * @param state
   * @param commit
   * @param dispatch
   * @param parentId
   */
  subDirsFinder({
    state,
    commit,
    dispatch
  }, parentId) {
    state.directories.forEach((item, index) => {
      if (item.parentId === parentId) {
        // 添加目录下标
        commit('addToTempArray', index);

        // 如果有子目录
        if (item.props.hasSubdirectories) {
          // 查找子目录
          dispatch('subDirsFinder', item.id);
        }
      }
    });
  },


  /**
   * 显示子目录
   * @param state
   * @param commit
   * @param getters
   * @param dispatch
   * @param path
   * @returns {*}
   */
  showSubdirectories({
    state,
    commit,
    getters,
    dispatch,
  }, path) {
    const promise = Promise.resolve();
    // 查找父目录下标
    const parentDirectoryIndex = getters.findDirectoryIndex(path);

    if (parentDirectoryIndex !== -1) {
      // 子目录是否加载
      if (state.directories[parentDirectoryIndex].props.subdirectoriesLoaded) {
        // 更新目录属性
        commit('updateDirectoryProps', {
          index: parentDirectoryIndex,
          props: {
            showSubdirectories: true,
          },
        });
      } else {
        // 加载子目录
        return dispatch('getSubdirectories', {
          path: state.directories[parentDirectoryIndex].path,
          parentId: state.directories[parentDirectoryIndex].id,
          parentIndex: parentDirectoryIndex,
        }).then(() => {
          // 更新父目录属性
          commit('updateDirectoryProps', {
            index: parentDirectoryIndex,
            props: {
              showSubdirectories: true,
            },
          });
        });
      }
    } else {
      commit('fm/messages/setError', {
        message: '目录未找到!'
      }, {
        root: true
      });
    }

    return promise;
  },

  /**
   * 从服务器得到所选目录的子目录
   * @param commit
   * @param rootGetters
   * @param path
   * @param parentId
   * @param parentIndex
   * @returns {Promise}
   */
  getSubdirectories({
    commit,
    rootGetters
  }, {
    path,
    parentId,
    parentIndex
  }) {
    return GET.tree(rootGetters['fm/selectedDisk'], path).then((response) => {
      // 成功
      if (response.data.result.status === 'success') {
        // 添加目录
        commit('addDirectories', {
          parentId,
          directories: response.data.directories,
        });

        // 更新父目录属性
        commit('updateDirectoryProps', {
          index: parentIndex,
          props: {
            subdirectoriesLoaded: true,
          },
        });
      }
    });
  },


  /**
   * 隐藏子目录
   * @param commit
   * @param getters
   * @param path
   */
  hideSubdirectories({
    commit,
    getters
  }, path) {
    // 查找父目录下标
    const parentDirectoryIndex = getters.findDirectoryIndex(path);

    if (parentDirectoryIndex !== -1) {
      // 隐藏子目录
      commit('updateDirectoryProps', {
        index: parentDirectoryIndex,
        props: {
          showSubdirectories: false,
        },
      });
    } else {
      commit('fm/messages/setError', {
        message: '目录未找到!'
      }, {
        root: true
      });
    }
  },

  /**
   * 重新打开选择的路径
   * @param dispatch
   * @param path
   * @returns {Promise<void>}
   */
  reopenPath({
    dispatch
  }, path) {
    let promises = Promise.resolve();

    if (path) {
      const splitPath = path.split('/');

      for (let i = 0; splitPath.length > i; i += 1) {
        promises = promises.then(() => dispatch(
          'showSubdirectories',
          splitPath.slice(0, i + 1).join('/'),
        ));
      }

      return promises;
    }

    return promises;
  },
};