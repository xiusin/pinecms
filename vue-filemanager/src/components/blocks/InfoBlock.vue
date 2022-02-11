<template >
  <div class="info-block">
    <el-row type="flex" :gutter="10">
      <el-col :span="6">
        <div class="grid-content bg-purple">
          <span
            @click="showModal('Status')"
            :class="[hasErrors ? 'text-danger' : 'text-success']"
            title="状态"
          >
            <i class="fas fa-info-circle" />
          </span>
          <span
            v-show="clipboardType"
            @click="showModal('Clipboard')"
            :title="`剪切板 - ${clipboardType=='copy' ? '复制':'粘贴'}`"
          >
            <i class="far fa-clipboard" />
          </span>
        </div>
      </el-col>
      <el-col :span="15">
        <div
          class="progress-bar"
          v-show="percentage"
          :aria-valuenow="percentage"
          aria-valuemin="0"
          aria-valuemax="100"
        >
          <el-progress
            :stroke-width="8"
            :text-inside="true"
            :percentage="percentage"
            :color="customColorMethod"
            :format="customFormat"
          ></el-progress>
        </div>
      </el-col>
      <el-col :span="3">
        <span v-show="loadingSpinner">
          <i class="fas fa-spinner fa-pulse" />
        </span>
      </el-col>
    </el-row>
  </div>
</template>
<script>
import helper from "@/mixins/helper";
export default {
  mixins: [helper],
  methods: {
    /**
     * 显示相应的模块,本组件中只显示Clipboard剪切板组件
     * @param modalName
     */
    showModal(modalName) {
      this.$store.commit("fm/modal/setModalState", {
        modalName,
        show: true
      });
    },
    /**
     * 自定义进度条颜色
     */
    customColorMethod(percentage) {
      if (percentage < 30) {
        return "#909399";
      } else if (percentage < 70) {
        return "#e6a23c";
      } else {
        return "#67c23a";
      }
    },
    customFormat(percentage) {
      if (percentage == 100) {
        return "wait...";
      }
    }
  },
  computed: {
    /**
     * 当前磁盘
     * @returns {String}
     */
    activeManager() {
      return this.$store.state.fm.activeManager;
    },

    /**
     * 进度条值 - %
     * @returns {Number}
     */
    percentage() {
      return this.$store.state.fm.messages.actionProgress;
    },

    /**
     * 是否出现错误
     * @returns {Boolean}
     */
    hasErrors() {
      return !!this.$store.state.fm.messages.errors.length;
    },

    /**
     * 选中的文件数量
     * @returns {Number}
     */
    filesCount() {
      return this.$store.getters[`fm/${this.activeManager}/filesCount`];
    },

    /**
     * 选中的文件夹数量
     * @returns {Number}
     */
    directoriesCount() {
      return this.$store.getters[`fm/${this.activeManager}/directoriesCount`];
    },

    /**
     * 所有文件大小总和
     * @returns {Number}
     */
    filesSize() {
      return this.bytesToHuman(
        this.$store.getters[`fm/${this.activeManager}/filesSize`]
      );
    },

    /**
     * 选中的文件和文件夹的数量总和
     * @returns {Number}
     */
    selectedCount() {
      return this.$store.getters[`fm/${this.activeManager}/selectedCount`];
    },

    /**
     * 计算选中的文件大小总和
     * @returns {Number}
     */
    selectedFilesSize() {
      return this.bytesToHuman(
        this.$store.getters[`fm/${this.activeManager}/selectedFilesSize`]
      );
    },

    /**
     * 剪切板类型-copy/cut
     * @returns {String}
     */
    clipboardType() {
      return this.$store.state.fm.clipboard.type;
    },

    /**
     * Spinner
     * @returns {Number}
     */
    loadingSpinner() {
      return this.$store.state.fm.messages.loading;
    }
  }
};
</script>
<style lang="scss" scoped>
.info-block {
  .grid-content {
    display: flex;
    justify-content: space-around;
    align-items: center;
    & > span:not(:first-child) {
      cursor: pointer;
    }
    span.text-danger {
      color: #dc3545;
    }
    span.text-success {
      color: #28a745;
    }
  }
}
</style>