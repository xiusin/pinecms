<template >
  <div class="modal-content modal-folder">
    <el-dialog ref="elDialog" :visible="showModal" height="100vh" :fullscreen="true" :before-close="handleClose" custom-class="officeViewer" :destroy-on-close="true">
      <iframe v-if="fullUrl" :src="'//view.officeapps.live.com/op/view.aspx?src=' + fullUrl" width="100%" style="border: none; height: calc(100% - 1px)" />
    </el-dialog>
  </div>
</template>
<script>
import { mapState } from "vuex";
import modal from "../mixins/modal";
import GET from "@/http/get";
export default {
  mixins: [modal],
  data() {
    return {
      fullUrl: ""
    };
  },
  computed: {
    ...mapState("fm", {
      showModal: state => state.modal.showModal
    }),
    selectedItem() {
      return this.$store.getters[`fm/${this.activeManager}/selectedList`][0];
    }
  },
  mounted() {
    this.getFullUrl()
    this.$nextTick(() => function () {
      this.$refs.elDialog[0].$el.firstChild.style.height = '90%'
    });
  },
  methods: {
    getFullUrl: function () {
      GET.url(this.selectedDisk, this.selectedItem.path).then(
          response => {
            this.fullUrl = response.data.url;
          }
      )
    }
  }
};
</script>
<style lang="scss">
  .officeViewer >>> .el-dialog__header {
    padding: 0;
  }
  .officeViewer >>> .el-dialog__body {
    padding: 0;
    overflow: hidden;
    overflow-y: auto;
  }
  .el-dialog >>> .el-dialog__body {
    height: 100%;
  }
</style>
