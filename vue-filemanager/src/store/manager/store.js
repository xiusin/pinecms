import mutations from './mutations';
import getters from './getters';
import actions from './actions';

export default {
  namespaced: true,
  state() {
    return {
      // 当前选择的磁盘
      selectedDisk: null,

      // 当前选择的目录路径
      selectedDirectory: null,

      // 当前选择的目录所有者
      selectedDirectoryOwner: null,
      // 历史指针
      historyPointer: 0,

      // 浏览文件历史记录
      history: [null],

      // 当前路径下的文件夹
      directories: [],

      // 当前路径下的文件
      files: [],

      // 当前路径下的图片链接
      imageUrl: [],

      // 当前选中的文件和文件夹
      selected: {
        directories: [],
        files: [],
      },

      // 当前选中的复选框文件列表
      // checkedFiles: [],

      //是否全选
      isCheckedAll: false,

      // 以表示 checkbox 的不确定状态，用于实现全选的效果
      isIndeterminate: false,

      // 排序设置
      sort: {
        field: 'name',
        direction: 'up',
      },

      // 视图类型-表格或网格-（默认值-表格）
      viewType: 'table',
    };
  },
  mutations,
  getters,
  actions,
};