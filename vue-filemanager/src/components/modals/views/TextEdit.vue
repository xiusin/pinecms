<template >
  <div class="text-edit">
    <el-dialog :visible="showModal" width="80%" :before-close="handleCloseTip">
      <span slot="title">
        <strong>编辑</strong>
      </span>
      <codemirror ref="codeEditor" v-model="code" :options="options"></codemirror>
      <span slot="footer">
        <el-button
          type="success"
          size="small"
          @click.native="updateFile"
        >提交</el-button>
        <el-button
         size="small"
         plain @click.native="hideModal">取消</el-button>
      </span>
    </el-dialog>
  </div>
</template>
<script>
import { mapState } from "vuex";
import { codemirror } from "vue-codemirror";
import "codemirror/mode/shell/shell";
import "codemirror/mode/css/css";
import "codemirror/mode/sass/sass";
import "codemirror/mode/htmlmixed/htmlmixed";
import "codemirror/mode/javascript/javascript";
import "codemirror/mode/vue/vue";
import "codemirror/mode/markdown/markdown";
import "codemirror/mode/xml/xml";
import "codemirror/mode/clike/clike";
import "codemirror/mode/php/php";
import "codemirror/mode/sql/sql";
import "codemirror/mode/lua/lua";
import "codemirror/mode/python/python";
import "codemirror/mode/swift/swift";
import "codemirror/mode/go/go";
import "codemirror/mode/yaml/yaml";
import "codemirror/mode/properties/properties";
import modal from "../mixins/modal";
export default {
  mixins: [modal],
  components: {
    codemirror
  },
  data() {
    return {
      code: ""
    };
  },
  mounted() {
    this.$store
      .dispatch("fm/getFile", {
        disk: this.selectedDisk,
        path: this.selectedItem.path
      })
      .then(resp => {
        // 添加代码
        if (this.selectedItem.extension === "json") {
          this.code = JSON.stringify(resp.data, null, 4);
        } else {
          this.code = String(resp.data);
        }

        // 设置大小
        this.$refs.codeEditor.codemirror.setSize(null, this.editorHeight);
      });
  },
  computed: {
    ...mapState("fm", {
      showModal: state => state.modal.showModal
    }),
    /**
     * ACL on/off
     */
    acl() {
      return this.$store.state.fm.settings.acl;
    },
    itemAuthor() {
      return this.selectedItem.author === this.$store.state.fm.username;
    },
    /**
     * 选择磁盘
     * @returns {*}
     */
    selectedDisk() {
      return this.$store.getters["fm/selectedDisk"];
    },
    /**
     * 选择文件
     * @returns {*}
     */
    selectedItem() {
      return this.$store.getters["fm/selectedItems"][0];
    },
    /**
     * 代码编辑器配置项
     * @returns {*}
     */
    options() {
      return {
        mode: this.$store.state.fm.settings.textExtensions[
          this.selectedItem.extension
        ],
        theme: "oceanic-next",
        lineNumbers: true,
        line: true
      };
    },
    /**
     * 计算代码编辑器的高度
     * @returns {number}
     */
    editorHeight() {
      if (this.$store.state.fm.modal.modalBlockHeight) {
        return this.$store.state.fm.modal.modalBlockHeight - 200;
      }
      return 450;
    }
  },
  methods: {
    // 更新文件
    updateFile() {
      const formData = new FormData();
      // 添加磁盘名
      formData.append("disk", this.selectedDisk);
      // 添加路径
      formData.append("path", this.selectedItem.dirname);
      // 添加更新的文件
      formData.append(
        "file",
        new Blob([this.code]),
        this.selectedItem.basename
      );

      this.$store.dispatch("fm/updateFile", formData).then(response => {
        // 如果文件更新成功
        if (response.data.result.status === "success") {
          // 关闭模块窗口
          this.hideModal();
        }
      });
    }
  }
};
</script>
<style lang="scss" scoped>
@import "~codemirror/lib/codemirror.css";
@import "~codemirror/theme/oceanic-next.css";
::v-deep .CodeMirror-vscrollbar {
  &::-webkit-scrollbar {
    width: 5px;
  }
  &::-webkit-scrollbar-track {
    background-color: #f5f5f5;
    -webkit-box-shadow: inset 0 0 3px rgba(0, 0, 0, 0.1);
    border-radius: 5px;
  }
  &::-webkit-scrollbar-thumb {
    background-color: rgba(140, 199, 181, 0.8);
    border-radius: 5px;
  }
  &::-webkit-scrollbar-button {
    display: none;
    background-color: #eee;
  }
  &::-webkit-scrollbar-corner {
    background-color: black;
  }
}
</style>