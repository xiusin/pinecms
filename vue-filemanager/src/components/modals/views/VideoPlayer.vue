<template >
  <div class="video-player">
    <el-dialog :visible="showModal" :before-close="handleCloseTip">
      <span slot="title" class="title">
        <strong>播放音频</strong>
        <small class="text-truncate">{{ videoFile.basename }}</small>
      </span>
      <div class="video-body">
        <video ref="video" controls></video>
      </div>
      <span slot="footer"></span>
    </el-dialog>
  </div>
</template>
<script>
import { mapState } from "vuex";
import Plyr from "plyr";
import modal from "../mixins/modal";
import HTTP from "@/http/axios";
export default {
  mixins: [modal],
  data() {
    return {
      player: {}
    };
  },
  mounted() {
    this.$nextTick(function() {
      // 初始化视频播放器
      this.player = new Plyr(this.$refs.video);
      // load source
      // this.player.source = {
      //   type: 'video',
      //   title: this.videoFile.filename,
      //   sources: [{
      //     src: `${this.$store.getters['fm/settings/baseUrl']}stream-file?disk=${this.selectedDisk}&path=${encodeURIComponent(this.videoFile.path)}`,
      //     type: `audio/${this.videoFile.extension}`,
      //   }],
      // };
      HTTP.get(this.$store.getters["fm/settings/baseUrl"] + "stream-file", {
        params: {
          disk: this.selectedDisk,
          path: this.videoFile.path
        }
      }).then(resp => {
        this.player.source = {
          type: "video",
          title: this.videoFile.filename,
          sources: [
            {
              // src: `${this.$store.getters['fm/settings/baseUrl']}stream-file?disk=${this.selectedDisk}&path=${encodeURIComponent(this.videoFile.path)}`,
              src: resp.data,
              type: `audio/${this.videoFile.extension}`
            }
          ]
        };
      });
    });
  },
  beforeDestroy() {
    this.player.destroy();
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
     * 视频文件
     * @returns {*}
     */
    videoFile() {
      return this.$store.getters["fm/selectedItems"][0];
    }
  }
};
</script>
<style lang="scss" scoped>
.video-player {
  .title {
    display: inline-block;
    overflow: hidden;
    text-overflow: ellipsis;
    width: -25%;
    width: 90%;
    white-space: nowrap;
    .text-truncate {
      padding-left: 1rem;
      color: #6c757d;
    }
  }
}
</style>