/* eslint-disable object-curly-newline */
import GET from '../../http/get';

export default {
  /**
   * 加载所选目录下的文件和文件夹
   * @param state
   * @param commit
   * @param dispatch
   * @param rootState
   * @param path
   * @param history
   * @returns {Promise}
   */
  selectDirectory({
    state,
    commit,
    dispatch
  }, {
    path,
    history
  }) {
    // 重置 内容
    commit('setDirectoryContent', {
      directories: [],
      files: [],

    });

    // 获取所选目录的内容
    return GET.content(state.selectedDisk, path).then((response) => {
      if (response.data.result.status === 'success') {
        commit('resetSelected');
        commit('resetSortSettings');
        commit('setDirectoryContent', response.data);
        commit('setSelectedDirectory', path);

        if (history) commit('addToHistory', path);

        // 显示目录
        if (path && response.data.directories.length) {
          dispatch('fm/tree/showSubdirectories', path, {
            root: true
          });
        }
      }
    });
  },

  /**
   * 返回当前目录的所有者
   * @param {*} param0 
   * @returns 
   */
  getOwnerOfDir({
    commit
  }, {
    path
  }) {
    return GET.findAuthor(path).then(res => {
      if (res.data.result.status === "success") {
        commit('setSelectedDirectoryOwner', res.data.author);
      }
    });
  },

  /**
   * 刷新所选目录中的内容
   * @param state
   * @param commit
   * @param dispatch
   */
  refreshDirectory({
    state,
    commit,
    dispatch
  }) {
    GET.content(state.selectedDisk, state.selectedDirectory).then((response) => {
      commit('resetSelected');
      commit('resetSortSettings');
      commit('resetHistory');

      // 添加到历史选择目录
      if (state.selectedDirectory) commit('addToHistory', state.selectedDirectory);

      if (response.data.result.status === 'success') {
        commit('setDirectoryContent', response.data);
      } else if (response.data.result.status === 'danger') {
        // 如果找不到目录，则加载主目录
        commit('setSelectedDirectory', null);
        dispatch('refreshDirectory');
      }
    });
  },

  /**
   * 回退历史
   * @param state
   * @param commit
   * @param dispatch
   */
  historyBack({
    state,
    commit,
    dispatch
  }) {
    dispatch('selectDirectory', {
      path: state.history[state.historyPointer - 1],
      history: false,
    });
    commit('pointerBack');
  },

  /**
   * 前进历史
   * @param state
   * @param commit
   * @param dispatch
   */
  historyForward({
    state,
    commit,
    dispatch
  }) {
    dispatch('selectDirectory', {
      path: state.history[state.historyPointer + 1],
      history: false,
    });
    commit('pointerForward');
  },

  /**
   * 按字段排序数据
   * @param context
   * @param field
   * @param direction
   */
  sortBy({
    state,
    commit
  }, {
    field,
    direction
  }) {
    if (state.sort.field === field && !direction) {
      commit('setSortDirection', state.sort.direction === 'up' ? 'down' : 'up');
    } else if (direction) {
      commit('setSortDirection', direction);
      commit('setSortField', field);
    } else {
      commit('setSortDirection', 'up');
      commit('setSortField', field);
    }
    // 按字段类型排序
    switch (field) {
      case 'name':
        commit('sortByName');
        break;
      case 'size':
        commit('sortBySize');
        break;
      case 'type':
        commit('sortByType');
        break;
      case 'date':
        commit('sortByDate');
        break;
      default:
        break;
    }
  },
};