<template >
  <div class="status">
    <el-dialog :visible="showModal" width="30%" :before-close="handleClose">
      <span slot="title">
        <h2>状态</h2>
      </span>
      <div class="status-body">
        <div v-if="errors.length">
          <ul>
            <li v-for="(error, i) in errors" :key="i">{{ error.status }} - {{ error.message }}</li>
          </ul>
        </div>
        <div v-else>暂无错误</div>
      </div>
      <span slot="footer">
        <el-button type="danger" size="small" :disabled="!errors.length" @click.native="clearErrors">清除</el-button>
        <el-button plain size="small" @click.native="hideModal">取消</el-button>
      </span>
    </el-dialog>
  </div>
</template>
<script>
import { mapState } from "vuex";
import modal from "../mixins/modal";
export default {
  mixins: [modal],
  computed: {
    ...mapState("fm", {
      showModal: state => state.modal.showModal
    }),
    /**
     * 程序错误
     * @returns {}
     */
    errors() {
      return this.$store.state.fm.messages.errors;
    }
  },
  methods: {
    /**
     * 清除所有错误
     */
    clearErrors() {
      this.$store.commit("fm/messages/clearErrors");
    }
  }
};
</script>
<style lang="scss" scoped>
.status {
  h2 {
    margin: 0;
  }
  .status-body {
    ul {
      list-style: none;
      padding: 0;
    }
  }
}
</style>