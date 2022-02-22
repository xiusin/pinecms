<template >
  <div class="modal-content modal-folder">
    <el-dialog :visible="showModal" width="90%" style="min-height: 500px" :before-close="handleClose">
      <iframe v-if="fullUrl" :src="'//view.officeapps.live.com/op/view.aspx?src=' + fullUrl" width="100%" height="100%" style="border: none" />
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
<style >
.el-dialog__header {
  padding: 0;
}
</style>
