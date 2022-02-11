<template >
  <div class="delete">
    <el-dialog :visible="showModal" width="40%" :before-close="handleClose">
      <span slot="title">
        <strong>删除文件</strong>
      </span>
      <div>
        <div v-if="selectedItems.length">
          <div class="delete-file" v-for="(f, i) in selectedItems" :key="i">
            <div class="text-truncate">
              <span v-if="f.type === 'dir'">
                <i class="far fa-folder" />
                {{ f.basename }}
              </span>
              <span v-else>
                <i class="far" v-bind:class="extensionToIcon(f.extension)" />
                {{ f.basename }}
              </span>
            </div>
            <div class="text-right" v-if="f.type === 'file'">{{ bytesToHuman(f.size) }}</div>
          </div>
        </div>

        <div v-else>
          <span class="text-danger">暂无选中！</span>
        </div>
      </div>
      <span slot="footer" class="dialog-footer">
        <el-button type="danger" size="small" :disabled="!selectedItems.length" @click.native="deleteItems">删除</el-button>
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
     * 获取要删除的文件和文件夹
     * @returns {*}
     */
    selectedItems() {
      return this.$store.getters["fm/selectedItems"];
    }
  },
  methods: {
    /**
     * 删除选中的文件和文件夹
     */
    deleteItems() {
      // 创建删除列表
      const items = this.selectedItems.map(item => ({
        path: item.path,
        type: item.type
      }));

      this.$store.dispatch("fm/delete", items).then(() => {
        // 关闭模块窗口
        this.hideModal();
      });
    }
  }
};
</script>
<style lang="scss" scoped>
.delete {
  .delete-file {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 3px 0;
    .text-truncate {
      overflow: hidden;
      text-overflow: ellipsis;
      white-space: nowrap;
      width: 70%;
    }
  }
  .text-danger {
    color: #dc3545;
  }
}
</style>