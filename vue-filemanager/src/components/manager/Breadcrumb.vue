<template >
  <div class="breadcrumb">
    <el-breadcrumb separator-class="el-icon-arrow-right">
      <el-breadcrumb-item>
        <span v-show="!isRootPath" @click="levelUp" class="up">返回上一级 | &nbsp;</span>
        <a href="#" @click.prevent="selectMainDirectory">全部文件</a>
      </el-breadcrumb-item>
      <el-breadcrumb-item v-for="(item, index) in breadcrumb" :key="index">
        <a href="#" @click.prevent="selectDirectory(index)">{{ item }}</a>
      </el-breadcrumb-item>
    </el-breadcrumb>
    <div style="font-size: 14px;">{{ `已全部加载,${directoriesCount}目录/${filesCount}文件:${filesSize}` }}</div>
  </div>
</template>
<script>
// import translate from "../../mixins/translate";
import helper from "../../mixins/helper";
import managerHelper from "./mixins/manager";
export default {
  mixins: [helper, managerHelper],
  props: {
    manager: { type: String, required: true }
  },
  data() {
    return {};
  },
  computed: {
    /**
     * 该磁盘下当前选择的所有目录
     * @returns {String}
     */
    selectedDirectory() {
      return this.$store.state.fm[this.manager].selectedDirectory;
    },
    /**
     * 面包屑
     * @returns {*}
     */
    breadcrumb() {
      return this.$store.getters[`fm/${this.manager}/breadcrumb`];
    },
    /**
     * 当前选中的磁盘管理器
     * @returns {String}
     */
    activeManager() {
      return this.$store.state.fm.activeManager;
    },
    /**
     * 当前路径下文件夹的数量
     * @returns {*}
     */
    directoriesCount() {
      return this.$store.getters[`fm/${this.activeManager}/directoriesCount`];
    },
    /**
     * 当前路径中文件的数量
     * @returns {*}
     */
    filesCount() {
      return this.$store.getters[`fm/${this.activeManager}/filesCount`];
    },
    /**
     * 当前路径下所有的文件大小
     * @returns {*|string}
     */
    filesSize() {
      return this.bytesToHuman(
        this.$store.getters[`fm/${this.activeManager}/filesSize`]
      );
    }
  },
  methods: {
    /**
     * 加载选中的目录
     * @param index
     */
    selectDirectory(index) {
      const path = this.breadcrumb.slice(0, index + 1).join("/");
      // 仅当未选择此路径时
      if (path !== this.selectedDirectory) {
        // 加载目录
        this.$store.dispatch(`fm/${this.manager}/selectDirectory`, {
          path,
          history: true
        });
      }
    },

    /**
     * 选择主目录
     */
    selectMainDirectory() {
      if (this.selectedDirectory) {
        this.$store.dispatch(`fm/${this.manager}/selectDirectory`, {
          path: null,
          history: true
        });
      }
    }
  }
};
</script>
<style lang="scss" scoped>
.up {
  color: #09aaff;
  cursor: pointer;
}
.breadcrumb {
  display: flex;
  justify-content: space-between;
}
</style>