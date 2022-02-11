<template >
  <div class="properties">
    <el-dialog :visible="showModal" width="32%" :before-close="handleClose">
      <span slot="title">
        <h2>属性</h2>
      </span>
      <div class="properties-content">
        <el-row>
          <el-col :span="4">磁盘：</el-col>
          <el-col :span="19">{{ selectedDisk }}</el-col>
          <el-col :span="1">
            <i @click="copyToClipboard(selectedDisk)" title="复制" class="far fa-copy"></i>
          </el-col>
        </el-row>
        <el-row>
          <el-col :span="4">名称：</el-col>
          <el-col :span="19">{{ selectedItem.basename }}</el-col>
          <el-col :span="1">
            <i @click="copyToClipboard(selectedItem.basename)" title="复制" class="far fa-copy"></i>
          </el-col>
        </el-row>
        <el-row>
          <el-col :span="4">路径:</el-col>
          <el-col :span="19">{{ selectedItem.path }}</el-col>
          <el-col :span="1">
            <i @click="copyToClipboard(selectedItem.path)" title="复制" class="far fa-copy"></i>
          </el-col>
        </el-row>
        <template v-if="selectedItem.type === 'file'">
          <el-row>
            <el-col :span="4">大小：</el-col>
            <el-col :span="19">{{ bytesToHuman(selectedItem.size) }}</el-col>
            <el-col :span="1">
              <i
                @click="copyToClipboard(bytesToHuman(selectedItem.size))"
                title="复制"
                class="far fa-copy"
              ></i>
            </el-col>
          </el-row>
          <el-row>
            <el-col :span="4">链接：</el-col>
            <el-col :span="19">
              <span v-if="url">{{ url }}</span>
              <span v-else>
                <button class="url" type="button" @click="getUrl">
                  <i class="fas fa-sm fa-link"></i>Get URL
                </button>
              </span>
            </el-col>
            <el-col :span="1" v-if="url">
              <i @click="copyToClipboard(url)" title="复制" class="far fa-copy"></i>
            </el-col>
          </el-row>
        </template>
        <template v-if="selectedItem.hasOwnProperty('timestamp')">
          <el-row>
            <el-col :span="4">时间：</el-col>
            <el-col :span="19">{{ timestampToDate(selectedItem.timestamp) }}</el-col>
            <el-col :span="1">
              <i
                @click="copyToClipboard(timestampToDate(selectedItem.timestamp))"
                title="复制"
                class="far fa-copy"
              ></i>
            </el-col>
          </el-row>
        </template>
        <template v-if="selectedItem.hasOwnProperty('acl')">
          <el-row>
            <el-col :span="4">创建者</el-col>
            <el-col :span="19">{{ author }}</el-col>
            <el-col :span="1">
              <i @click="copyToClipboard(author)" title="复制" class="far fa-copy"></i>
            </el-col>
          </el-row>
          <el-row>
            <el-col :span="4">授权</el-col>
            <el-col :span="19">{{ lang.modal.properties['access_' + selectedItem.acl] }}</el-col>
            <el-col :span="1"></el-col>
          </el-row>
        </template>
      </div>
      <span slot="footer"></span>
    </el-dialog>
  </div>
</template>
<script>
import { mapState } from "vuex";
import modal from "../mixins/modal";
import helper from "@/mixins/helper";
import translate from "@/mixins/translate";
export default {
  mixins: [modal, helper, translate],
  data() {
    return {
      url: null
    };
  },
  computed: {
    ...mapState("fm", {
      showModal: state => state.modal.showModal
    }),
    /**
     * 选择的磁盘
     * @returns {*}
     */
    selectedDisk() {
      return this.$store.getters["fm/selectedDisk"];
    },

    /**
     * 选择的文件
     * @returns {*}
     */
    selectedItem() {
      return this.$store.getters["fm/selectedItems"][0];
    },
    /**
     * 选择的文件的创建者
     */
    author() {
      return this.selectedItem.author;
    }
  },
  methods: {
    /**
     * 获得该文件的 URL
     */
    getUrl() {
      this.$store
        .dispatch("fm/url", {
          disk: this.selectedDisk,
          path: this.selectedItem.path
        })
        .then(resp => {
          if (resp.data.result.status === "success") {
            this.url = resp.data.url;
          }
        });
    },
    /**
     * 复制文本到剪贴板
     * @param text
     */
    copyToClipboard(text) {
      // 创建 input
      const copyInputHelper = document.createElement("input");
      copyInputHelper.className = "copyInputHelper";
      document.body.appendChild(copyInputHelper);
      // 添加文本
      copyInputHelper.value = text;
      copyInputHelper.select();
      // 复制文本到剪切板
      document.execCommand("copy");
      // 移除该元素
      document.body.removeChild(copyInputHelper);
      // 提醒
      this.$message({
        message: "已复制!",
        type: "success",
        duration: 2000
      });
    }
  }
};
</script>
<style lang="scss" scoped>
.properties {
  .properties-content {
    ::v-deep .el-row {
      display: flex;
      padding: 0.3rem 0;
      //   align-items: center;
      &:hover {
        background-color: #f8f9fa;
        .el-col {
          .fa-copy {
            display: block;
          }
        }
      }
      .el-col.el-col-4 {
        font-weight: 600;
        font-size: 1rem;
      }
      .el-col {
        .fa-copy {
          display: none;
          cursor: pointer;
        }
        .url {
          color: #212529;
          background-color: #f8f9fa;
          border-color: #f8f9fa;
          cursor: pointer;
          border: none;
          padding: 0.25rem 0.5rem;
          font-size: 0.875rem;
          line-height: 1.5;
          border-radius: 0.2rem;
          &:hover {
            background-color: #e2e6ea;
            border-color: #dae0e5;
          }
        }
      }
    }
  }
  h2 {
    margin: 0;
  }
}
</style>