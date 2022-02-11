import mutations from './mutations';
import getters from './getters';
import actions from './actions';

export default {
  namespaced: true,
  state() {
    return {
      /**
       * directories.id (int), 目录id
       * directories.basename                       (string), 文件名
       * directories.dirname                        (string) 目录名
       * directories.path                           (string), 目录路径
       * directories.props                          (object), 目录属性
       * directories.props.hasSubdirectories        (boolean), 是否有子目录,
       * directories.props.subdirectoriesLoaded     (boolean), 子目录加载
       * directories.props.showSubdirectories       (boolean), 显示隐藏子目录
       * directories.parentId                       (int), 父目录id
       */
      directories: [],

      // 目录计数器
      counter: 1,

      // 临时目录索引数组
      tempIndexArray: [],
    };
  },
  mutations,
  getters,
  actions,
};