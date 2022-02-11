<template >
  <div class="clipboard">
    <el-dialog :visible="showModal" width="36%" :before-close="handleClose">
      <span slot="title">
        <strong>剪切板</strong>
      </span>
      <div class="clipboard-body">
        <template v-if="clipboard.type">
          <div class="body-h">
            <div>
              <i class="far fa-hdd" />
              {{ clipboard.disk }}
            </div>
            <div>
              <span :title="`类型 - ${clipboard.type=='copy' ? '复制':'粘贴'}`">
                <i v-if="clipboard.type === 'copy'" class="fas fa-copy" />
                <i v-else class="fas fa-cut" />
              </span>
            </div>
          </div>
          <hr />
          <div class="body-content" v-for="(d,i) in directories" :key="`d-${i}`">
            <div class="text-truncate">
              <i class="far fa-folder" />
              {{ d.name }}
            </div>
            <div class="text-right">
              <button
                type="button"
                class="close"
                title="删除"
                @click="deleteItem('directories', d.path)"
              >
                <span aria-hidden="true">&times;</span>
              </button>
            </div>
          </div>
          <div class="body-content" v-for="(f, i) in files" :key="`f-${i}`">
            <div class="text-truncate">
              <i class="far" :class="f.icon" />
              {{ f.name }}
            </div>
            <div class="text-right">
              <button type="button" class="close" title="删除" @click="deleteItem('files', f.path)">
                <span aria-hidden="true">&times;</span>
              </button>
            </div>
          </div>
        </template>
        <template v-else>
          <span>未选中文件!</span>
        </template>
      </div>
      <span slot="footer">
        <el-button type="danger" size="small" :disabled="!clipboard.type" @click.native="resetClipboard">清除</el-button>
        <el-button plain size="small" @click.native="hideModal">取消</el-button>
      </span>
    </el-dialog>
  </div>
</template>
<script>
import { mapState } from "vuex";
import modal from "../mixins/modal";
import helper from "@/mixins/helper";
export default {
  mixins: [modal, helper],
  computed: {
    ...mapState("fm", {
      showModal: state => state.modal.showModal
    }),
    /**
     * 剪切板状态
     * @returns {Object}
     */
    clipboard() {
      return this.$store.state.fm.clipboard;
    },
    /**
     * 文件夹的路径和名字为对象组成的数组
     * @returns {{path: *, name: *}[]}
     */
    directories() {
      return this.$store.state.fm.clipboard.directories.map(item => ({
        path: item,
        name: item.split("/").slice(-1)[0]
      }));
    },
    /**
     * 文件名，路径和图标为对象组成的数组
     * @returns {{path: *, name: *, icon: *}[]}
     */
    files() {
      return this.$store.state.fm.clipboard.files.map(item => {
        const name = item.split("/").slice(-1)[0];
        return {
          path: item,
          name,
          icon: this.extensionToIcon(name.split(".").slice(-1)[0])
        };
      });
    }
  },
  methods: {
    /**
     * 删除剪贴板中的内容
     * @param type
     * @param path
     */
    deleteItem(type, path) {
      this.$store.commit("fm/truncateClipboard", { type, path });
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
.clipboard {
  .clipboard-body {
    .body-h {
      display: flex;
      justify-content: space-between;
    }
    hr {
      border: 0;
      border-top: 1px solid rgba(0, 0, 0, 0.1);
    }
    .body-content {
      display: flex;
      justify-content: space-between;
      align-items: center;
      .text-truncate {
        overflow: hidden;
        text-overflow: ellipsis;
        white-space: nowrap;
        width: 75%;
      }
      .close {
        border: 0;
        padding: 0;
        background: transparent;
        font-size: 1.5rem;
        opacity: 0.5;
      }
    }
  }
}
</style>