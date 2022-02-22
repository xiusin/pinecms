<template >
  <transition name="modal">
    <component :is="modalName"></component>
  </transition>
</template>
<script>
import NewFolder from "./views/NewFolder";
import NewFile from "./views/NewFile";
import AudioPlayer from "./views/AudioPlayer";
import VideoPlayer from "./views/VideoPlayer";
import Clipboard from "./views/Clipboard";
import Status from "./views/Status";
import Delete from "./views/Delete";
import Properties from "./views/Properties";
import Rename from "./views/Rename";
import TextEdit from "./views/TextEdit";
import Preview from "./views/Preview";
import Upload from "./views/Upload";
import OfficeViewer from "./views/OfficeViewer";
export default {
  data() {
    return {
      dialogVisible: false
    };
  },
  components: {
    NewFolder,
    NewFile,
    AudioPlayer,
    VideoPlayer,
    Clipboard,
    Status,
    Delete,
    Properties,
    Rename,
    TextEdit,
    Preview,
    Upload,
    OfficeViewer,
  },
  mounted() {
    // 设置模块高度
    // this.$store.commit(
    //   "fm/modal/setModalBlockHeight",
    //   this.$refs.modal.offsetHeight
    // );
  },
  computed: {
    // ...mapState("fm", {
    //   showModal: state => state.modal.showModal
    // }),
    /**
     * 选中的模块名
     * @returns {null|*}
     */
    modalName() {
      return this.$store.state.fm.modal.modalName;
    }
    /**
     * 模块的尺寸大小
     * @returns {{'modal-lg': boolean, 'modal-sm': boolean}}
     */
    // modalSize() {
    //   return {
    //     "modal-xl":
    //       this.modalName === "Preview" || this.modalName === "TextEdit",
    //     "modal-lg": this.modalName === "VideoPlayer",
    //     "modal-sm": false
    //   };
    // }
  },
  methods: {
    /**
     * 隐藏模块窗口
     */
    // hideModal() {
    //   this.$store.commit("fm/modal/clearModal");
    // }
  }
};
</script>
<style lang="scss" >
@import "~plyr/src/sass/plyr.scss";
// .modal {
//   position: absolute;
//   z-index: 9998;
//   top: 0;
//   bottom: 0;
//   left: 0;
//   right: 0;
//   width: 100%;
//   height: 100%;
//   background-color: rgba(0, 0, 0, 0.35);
//   display: block;
//   transition: opacity 0.4s ease;
//   overflow: auto;

//   .modal-xl {
//     max-width: 96%;
//   }
// }

.el-dialog {
  border-radius: 5px;
  .el-dialog__body {
    padding: 0px 20px;
  }
}

.modal-enter-active,
.modal-leave-active {
  transition: opacity 0.5s;
}
.modal-enter,
.modal-leave-to {
  opacity: 0;
}
</style>
