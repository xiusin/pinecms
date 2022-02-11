<template >
  <div class="upload">
    <el-dialog :visible="showModal" width="45%" :before-close="handleClose">
      <span slot="title">
        <strong>上传文件</strong>
      </span>
      <div class="upload-body">
        <div class="upload-area">
          <el-upload
            ref="upload"
            class="upload-demo"
            action="#"
            :auto-upload="false"
            :file-list="fileList"
            :http-request="httpRequest"
            :on-remove="handleRemove"
            :before-upload="beforeAvatarUpload"
            :on-change="fileChange"
            drag
            multiple
          >
            <i class="el-icon-upload"></i>
            <div class="el-upload__text">
              将文件拖到此处，或
              <em>点击上传</em>
            </div>
            <div class="el-upload__tip" slot="tip">
              <div>
                <span>注意：单个文件大小不得超过1GB</span>
                <span v-show="countFiles">文件数量:{{ countFiles }}/总大小{{ allFilesSize }}</span>
              </div>
              <div>
                若文件存在,是否覆盖原文件:
                <el-radio v-model="overwrite" label="0">忽略</el-radio>
                <el-radio v-model="overwrite" label="1">覆盖</el-radio>
              </div>
            </div>
          </el-upload>
        </div>
      </div>
      <span slot="footer" class="dialog-footer">
        <el-button
          type="info"
          size="small"
          :class="[countFiles ?  'btn-light':'btn-info']"
          :disabled="!countFiles || isprogressing"
          @click.native="uploadFiles"
        >提交</el-button>
        <el-button 
        size="small"
        plain @click.native="hideModal">取消</el-button>
      </span>
    </el-dialog>
  </div>
</template>
<script>
import modal from "../mixins/modal";
import helper from "@/mixins/helper";
import { mapState } from "vuex";
export default {
  mixins: [modal, helper],
  data() {
    return {
      // 选择的文件
      fileList: [],
      // 上传携带的参数
      params: {},
      // 是否覆盖已存在的文件
      overwrite: "0"
    };
  },
  computed: {
    ...mapState("fm", {
      showModal: state => state.modal.showModal
    }),
    authorIsUser() {
      let author = this.$store.getters[
        `fm/${this.activeManager}/selectedDirectoryOwner`
      ];
      let username = this.$store.state.fm.username;
      return username === author;
    },
    /**
     * 进度条值 - %
     * @returns {number}
     */
    progressBar() {
      return this.$store.state.fm.messages.actionProgress;
    },
    isprogressing() {
      return Boolean(this.progressBar);
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
     * 上传文件的数量
     * @returns {number}
     */
    countFiles() {
      return this.fileList.length;
    },

    /**
     * 计算上传文件的大小
     * @returns {*|string}
     */
    allFilesSize() {
      let size = 0;
      for (let i = 0; i < this.fileList.length; i++) {
        size += this.fileList[i].size;
      }

      return this.bytesToHuman(size);
    }
  },
  methods: {
    /**
     * 选择一个或多个文件
     * @param event
     */
    selectFiles(event) {
      if (event.target.files.length === 0) {
        // 没有选择文件
        this.newFiles = [];
      } else {
        // 有文件
        this.newFiles = event.target.files;
      }
    },
    /**
     * 文件列表移除文件时的钩子
     */
    handleRemove(file, fileList) {
      this.fileList = fileList;
    },
    /**
     * 文件状态改变时的钩子，添加文件、上传成功和上传失败时都会被调用
     */
    fileChange(file, fileList) {
      this.fileList = fileList;
    },
    /**
     * 上传文件之前的钩子,把文件限制在1G以下
     */
    beforeAvatarUpload(file) {
      const isLt1G = file.size / 1024 / 1024 / 1024 < 1;
      if (!isLt1G) {
        this.$message.error("上传的单个文件不能超过1GB");
      }
      return isLt1G;
    },

    /**
     * 上传文件
     */
    uploadFiles() {
      if (
        this.acl &&
        this.selectedDisk == "public" &&
        this.selectedDir !== null &&
        !this.authorIsUser
      ) {
        this.$notify.warning({
          title: "警告",
          message: "该目录下无法上传文件!你没有权限"
        });
        return;
      }
      this.$refs.upload.submit();
    },
    httpRequest(param) {
      //若有文件要上传
      if (this.countFiles) {
        this.$store
          .dispatch("fm/upload", {
            files: [param.file],
            overwrite: Number(this.overwrite),
            fileParam: param
          })
          .then(response => {
            // 上传成功
            if (response.data.result.status === "success") {
              // 关闭模块窗口
              // this.hideModal();
              param.onSuccess(response);
            }
          });
      }
    }
  }
};
</script>
<style lang="scss" scoped>
.upload {
  .upload-body {
    ::v-deep .upload-demo {
      width: 100%;
      .el-upload {
        width: 100%;
        .el-upload-dragger {
          width: 100% !important;
        }
      }
      .el-upload__tip {
        & > div:first-child {
          display: flex;
          justify-content: space-between;
          align-items: center;
        }
        & > div:last-child {
          margin-top: 10px;
        }
      }
    }
  }
  .dialog-footer {
    .btn-light {
      color: #fff;
      background-color: #17a2b8;
      border-color: #17a2b8;
      &:hover {
        background-color: #138496;
        border-color: #117a8b;
      }
    }
  }
}
</style>