<template >
  <div class="rename">
    <el-dialog :visible="showModal" width="30%" :before-close="handleClose">
      <span slot="title">
        <h2>重命名</h2>
      </span>
      <div class="rename-content">
        <el-form label-position="top" :model="ruleForm" :rules="rules" @submit.native.prevent>
          <el-form-item label="请输入名称" label-width="80px" prop="name">
            <el-input v-model="ruleForm.name" autofocus clearable></el-input>
          </el-form-item>
        </el-form>
      </div>
      <span slot="footer" class="dialog-footer">
        <el-button type="success" :disabled="submitDisable" @click.native="rename">提交</el-button>
        <el-button plain @click.native="hideModal">取消</el-button>
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
    const validateName = (rule, value, callback) => {
      if (value && value !== this.selectedItem.basename) {
        // 如果是文件夹
        if (this.selectedItem.type === "dir") {
          // 检查文件夹名称是否存在
          this.directoryExist = this.$store.getters[
            `fm/${this.activeManager}/directoryExist`
          ](value);
        } else {
          // 检查文件名是否存在
          this.fileExist = this.$store.getters[
            `fm/${this.activeManager}/fileExist`
          ](value);
        }
        if (this.directoryExist) {
          callback(new Error("文件夹已存在!"));
        } else if (this.fileExist) {
          callback(new Error("文件已存在!"));
        } else if (reg.test(value)) {
          callback(new Error('含有非法字符有\\/:*?"<>|'));
        } else if (
          this.selectedItem.type == "file" &&
          !fileExtension.test(value)
        ) {
          callback(new Error("文件缺少扩展名!"));
        } else {
          callback();
        }
      }
    };
    return {
      ruleForm: {
        name: "" // 新建的文件夹的名字
      },
      rules: {
        name: [
          { required: true, message: "内容不能为空", trigger: "blur" },
          { validator: validateName, vatrigger: "blur" }
        ]
      },
      directoryExist: false,
      fileExist: false
    };
  },
  computed: {
    ...mapState("fm", {
      showModal: state => state.modal.showModal
    }),
    /**
     * 禁止点击提交按钮
     * @returns {Boolean}
     */
    submitDisable() {
      return (
        this.checkName || this.ruleForm.name === this.selectedItem.basename
      );
    },
    /**
     * 检查文件名
     * @returns {Boolean}
     */
    checkName() {
      return this.directoryExist || this.fileExist || !this.ruleForm.name;
    },
    /**
     * 当前选择的文件
     * @returns {*}
     */
    selectedItem() {
      return this.$store.getters[`fm/${this.activeManager}/selectedList`][0];
    }
  },
  mounted() {
    // 文件名
    this.ruleForm.name = this.selectedItem.basename;
  },
  methods: {
    /**
     * 重命名
     */
    rename() {
      // 用路径创建新名称
      const newName = this.selectedItem.dirname
        ? `${this.selectedItem.dirname}/${this.ruleForm.name}`
        : this.ruleForm.name;
        console.log(this.selectedItem.dirname, newName)
      this.$store
        .dispatch("fm/rename", {
          type: this.selectedItem.type,
          newName,
          oldName: this.selectedItem.path
        })
        .then(() => {
          this.hideModal();
        });
    }
  }
};
</script>
<style lang="scss" scoped>
.rename {
  h2 {
    margin: 0;
  }
  ::v-deep .el-form-item {
    .el-form-item__label::before {
      content: "";
    }
  }
}
</style>