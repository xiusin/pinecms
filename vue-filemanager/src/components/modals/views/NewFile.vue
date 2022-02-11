<template >
  <div class="modal-content modal-folder">
    <el-dialog :visible="showModal" width="30%" :before-close="handleClose">
      <span slot="title">创建文件</span>
      <el-form label-position="top" :model="ruleForm" :rules="rules" @submit.native.prevent>
        <el-form-item size="middle" prop="fileName">
          <el-input v-model="ruleForm.fileName" :autofocus="true" clearable></el-input>
        </el-form-item>
      </el-form>
      <span slot="footer" class="dialog-footer">
        <el-button size="small" type="success" :disabled="!submitActive" @click.native="addFile">提交</el-button>
        <el-button size="small" plain @click.native="hideModal">取消</el-button>
      </span>
    </el-dialog>
  </div>
</template>
<script>
import { mapState } from "vuex";
import modal from "../mixins/modal";
export default {
  mixins: [modal],
  data() {
    const reg = new RegExp('[\\\\/:*?"<>|]');
    const fileExtension = new RegExp(/.+\.+\w+/, "ig");
    const validateFileName = (rule, value, callback) => {
      if (value) {
        //检查文件是否存在
        this.fileExist = this.$store.getters[
          `fm/${this.activeManager}/fileExist`
        ](value);
        if (this.fileExist) {
          callback(new Error("文件已存在!"));
        } else if (reg.test(value)) {
          callback(new Error('含有非法字符\\/:*?"<>|'));
        } else if (!fileExtension.test(value)) {
          callback(new Error("文件缺少扩展名!"));
        } else {
          callback();
        }
      } else {
        this.fileExist = false;
      }
    };
    return {
      ruleForm: {
        fileName: "" // 新创建的文件名
      },
      fileExist: false, // 文件名是否存在
      rules: {
        fileName: [{ validator: validateFileName, vatrigger: "blur" }]
      }
    };
  },
  computed: {
    ...mapState("fm", {
      showModal: state => state.modal.showModal
    }),
    /**
     * 提交状态
     * @returns {Boolean}
     */
    submitActive() {
      return this.ruleForm.fileName && !this.fileExist;
    }
  },
  methods: {
    /**
     * 检查文件是否存在
     */
    // validateFileName() {
    //   if (this.fileName) {
    //     this.fileExist = this.$store.getters[
    //       `fm/${this.activeManager}/fileExist`
    //     ](this.fileName);
    //   } else {
    //     this.fileExist = false;
    //   }
    // },

    /**
     * 创建新文件
     */
    addFile() {
      this.$store
        .dispatch("fm/createFile", this.ruleForm.fileName)
        .then(response => {
          if (response.data.result.status === "success") {
            // 如果新文件创建成功,关闭模块窗口
            this.hideModal();
          }
        });
    }
  }
};
</script>
