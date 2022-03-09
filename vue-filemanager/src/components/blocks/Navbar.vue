<template >
  <el-row type="flex" class="row-bg" justify="space-between" :gutter="20">
    <el-col :span="8">
      <el-button type="primary" size="mini" title="选择文件上传" @click.native="showModal('Upload')">
        <i class="fa fa-lg fa-cloud-upload-alt"></i>
        <span class="action-text">&nbsp;上传</span>
      </el-button>
      <el-button
        type="primary"
        size="mini"
        title="新建文件夹"
        plain
        @click.native="showModal('NewFolder')"
      >
        <i class="fa fa-lg fa-folder-plus"></i>
        <span class="action-text">&nbsp;新建文件夹</span>
      </el-button>
      <el-button
        type="primary"
        size="mini"
        title="新建文件"
        plain
        @click.native="showModal('NewFile')"
      >
        <i class="fa fa-lg fa-file-alt"></i>
        <span class="action-text">&nbsp;新建文件</span>
      </el-button>
    </el-col>
    <el-col :span="14">
      <div v-show="selectedItems.length || clipboard.type">
        <el-button
          type="primary"
          size="mini"
          title="复制"
          plain
          :disabled="!isAnyItemSelected"
          @click.native="toClipboard('copy')"
        >
          <i class="fa fa-lg fa-copy"></i>
          <span class="action-text">&nbsp;复制</span>
        </el-button>
        <el-button
          type="primary"
          size="small"
          title="剪切"
          plain
          :disabled="!isAnyItemSelected || (!itemAuthor&&acl&&selectedItemAcl)"
          @click.native="toClipboard('cut')"
        >
          <i class="fa fa-lg fa-cut"></i>
          <span class="action-text">&nbsp;剪切</span>
        </el-button>
        <el-button
          type="primary"
          size="mini"
          title="粘贴"
          plain
          :disabled="!clipboardType || (!authorIsUser&&this.selectedDisk == 'public')"
          @click.native="paste"
        >
          <i class="fa fa-lg fa-paste"></i>
          <span class="action-text">&nbsp;粘贴</span>
        </el-button>
        <el-button
          size="mini"
          title="删除"
          plain
          :disabled="!isAnyItemSelected || (!itemAuthor&&acl&&selectedItemAcl)"
          @click.native="showModal('Delete')"
          style="color:red;border-color:red;"
        >
          <i class="fa fa-lg fa-trash-alt"></i>
          <span class="action-text">&nbsp;删除</span>
        </el-button>
      </div>
    </el-col>
    <el-col :span="2" style="text-align: end;line-height: 2;">
      <i
        class="fa fa-lg"
        :class="viewType === 'table'?'fa-th-large':'fa-align-justify'"
        @click="selectView"
      ></i>
    </el-col>
  </el-row>
</template>
<script>
export default {
  props: {
    manager: {
      type: String,
      required: true
    }
  },
  data() {
    return {};
  },
  computed: {
    /**
     * 视图类型 - Grid 或 table
     * @returns {String}
     */
    viewType() {
      return this.$store.state.fm[this.manager].viewType;
    },
    /**
     * 当前目录的拥有者
     */
    ownerOfPath() {
      let username = this.$store.state.fm.username;
      return this.selectedItems.every(function(item) {
        return username === item.author;
      });
    },
    authorIsUser() {
      let author = this.$store.getters[
        `fm/${this.activeManager}/selectedDirectoryOwner`
      ];
      let username = this.$store.state.fm.username;
      return username === author;
    },
    itemAuthor() {
      let username = this.$store.state.fm.username;
      return this.selectedItems.every(function(item) {
        return item.author === username;
      });
    },
    /**
     * 当前磁盘
     * @returns {String}
     */
    activeManager() {
      return this.$store.state.fm.activeManager;
    },
    /**
     * ACL on/off
     */
    acl() {
      return this.$store.state.fm.settings.acl;
    },
    /**
     * 选择硬盘
     * @returns {*}
     */
    selectedDisk() {
      return this.$store.getters["fm/selectedDisk"];
    },
    /**
     * 当前路径
     */
    selectedDir() {
      return this.$store.getters["fm/selectedDirectory"];
    },
    /**
     * 获取要删除的文件和文件夹
     * @returns {*}
     */
    selectedItems() {
      return this.$store.getters["fm/selectedItems"];
    },
    /**
     * 判断选中的文件的acl
     * @returns
     */
    selectedItemAcl() {
      return this.selectedItems.every(function(item) {
        return item.acl === 1;
      });
    },
    /**
     * 文件或目录是否被选中
     * @returns {boolean}
     */
    isAnyItemSelected() {
      return (
        this.$store.state.fm[this.activeManager].selected.files.length > 0 ||
        this.$store.state.fm[this.activeManager].selected.directories.length > 0
      );
    },
    /**
     * 剪贴板-操作类型 -copy | cut
     * @returns {String}
     */
    clipboardType() {
      return this.$store.state.fm.clipboard.type;
    },
    /**
     * 剪切板状态
     * @returns {*}
     */
    clipboard() {
      return this.$store.state.fm.clipboard;
    }
  },
  methods: {
    /**
     * 选择视图类型-table/grid
     */
    selectView() {
      if (this.viewType == "table") {
        this.$store.commit(`fm/${this.activeManager}/setView`, "grid");
        return;
      }
      this.$store.commit(`fm/${this.activeManager}/setView`, "table");
    },
    /**
     * 显示 相应的模块
     * @param modalName
     */
    showModal(modalName) {
      if (
        (modalName == "NewFolder" || modalName == "NewFile") &&
        this.acl &&
        this.selectedDisk == "public" &&
        this.selectedDir !== null &&
        !this.authorIsUser
      ) {
        this.$notify.warning({
          title: "警告",
          message: "该目录下你没有操作权限"
        });
        return;
      }
      this.$store.commit("fm/modal/setModalState", {
        modalName,
        show: true
      });
    },
    /**
     * 粘贴
     */
    paste() {
      if (
        this.acl &&
        this.selectedDisk == "public" &&
        this.selectedDir !== null &&
        !this.authorIsUser
      ) {
        this.$notify.warning({
          title: "警告",
          message: "该目录下你没有操作权限"
        });
        return;
      }
      this.$store.dispatch("fm/paste");
      this.resetClipboard();
    },
    /**
     * 复制
     * @param type
     */
    toClipboard(type) {
      this.$store.dispatch("fm/toClipboard", type);
      // 显示通知
      if (type === "cut") {
        this.$message({
          message: "剪切到粘贴板!",
          type: "success"
        });
        // EventBus.$emit("addNotification", {
        //   status: "success",
        //   message: this.lang.notifications.cutToClipboard
        // });
      } else if (type === "copy") {
        this.$message({
          message: "复制到粘贴板!",
          type: "success"
        });
        // EventBus.$emit("addNotification", {
        //   status: "success",
        //   message: this.lang.notifications.copyToClipboard
        // });
      }
    },
    getOwner() {
      let author = this.$store.dispatch(
        `fm/${this.activeManager}/getOwnerOfDir`
      );
      let username = this.$store.state.fm.username;
      return author == username;
    },
    /**
     * 重置剪贴板
     */
    resetClipboard() {
      this.$store.commit("fm/resetClipboard");
    }
  }
};
</script>
<style lang="scss" scoped>
@media screen and (max-width: 1140px) {
  .action-text {
    display: none;
  }
}
</style>
