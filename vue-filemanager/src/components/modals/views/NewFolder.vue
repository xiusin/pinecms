<template >
  <div class="modal-content modal-folder">
    <el-dialog :visible="showModal" width="30%" :before-close="handleClose">
      <span slot="title">
        <strong>创建文件夹</strong>
      </span>
      <el-form label-position="top" :model="ruleForm" :rules="rules" @submit.native.prevent>
        <el-form-item size="middle" prop="directoryName">
          <el-input v-model="ruleForm.directoryName" autofocus clearable></el-input>
        </el-form-item>
      </el-form>
      <span slot="footer" class="dialog-footer">
        <el-button type="success" size="small" :disabled="!submitActive" @click.native="addFolder">提交</el-button>
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
    const validateDirectoryName = (rule, value, callback) => {
      if (value) {
        //检查文件夹是否存在
        this.directoryExist = this.$store.getters[
          `fm/${this.activeManager}/directoryExist`
        ](value);
        if (this.directoryExist) {
          callback(new Error("目录已存在!"));
        } else if (reg.test(value)) {
          callback(new Error('含有非法字符\\/:*?"<>|'));
        } else {
          callback();
        }
      } else {
        this.directoryExist = false;
      }
    };
    return {
      ruleForm: {
        directoryName: "" // 新建的文件夹名
      },
      directoryExist: false, // 文件夹是否存在
      rules: {
        directoryName: [{ validator: validateDirectoryName, vatrigger: "blur" }]
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
      return this.ruleForm.directoryName && !this.directoryExist;
    }
  },
  methods: {
    /**
     * 检查文件夹是否存在
     */
    // validateDirName() {
    //   if (this.ruleForm.directoryName) {
    //     this.directoryExist = this.$store.getters[
    //       `fm/${this.activeManager}/directoryExist`
    //     ](this.ruleForm.directoryName);
    //   } else {
    //     this.directoryExist = false;
    //   }
    // },

    /**
     * 创建文件夹
     */
    addFolder() {
      this.$store
        .dispatch("fm/createDirectory", this.ruleForm.directoryName)
        .then(response => {
          if (response.data.result.status === "success") {
            this.hideModal(); // 如果新目录创建成功,关闭模块窗口
          }
        });
    }
  }
};
</script>
<style lang="scss">
</style>