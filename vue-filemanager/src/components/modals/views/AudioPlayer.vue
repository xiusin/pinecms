<template >
  <div class="audio-player">
    <el-dialog :visible="showModal" :before-close="handleCloseTip">
      <span slot="title">
        <strong>播放音频</strong>
      </span>
      <div class="audio-body">
        <audio ref="audio" controls></audio>
        <hr />
        <div
          class="track"
          :class="playingIndex === index ? 'bg-light':''"
          v-for="(audio, index) in audioFiles"
          :key="index"
        >
          <div class="text-truncate">
            <span>{{ index+1 }}.</span>
            {{ audio.basename }}
          </div>
          <template v-if="playingIndex === index">
            <div v-if="status === 'playing'">
              <i @click="togglePlay()" class="fas fa-play active"></i>
            </div>
            <div v-else>
              <i @click="togglePlay()" class="fas fa-pause"></i>
            </div>
          </template>
          <template v-else>
            <div>
              <i @click="selectTrack(index)" class="fas fa-play"></i>
            </div>
          </template>
        </div>
      </div>
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
      player: {},
      playingIndex: 0,
      status: "paused"
    };
  },
  mounted() {
    this.$nextTick(function() {
      // 初始化播放器
      this.player = new Plyr(this.$refs.audio, {
        speed: {
          selected: 1,
          options: [0.5, 1, 1.5]
        }
      });
      // 选择列表中的第一项
      this.setSource(this.playingIndex);
      // 添加事件监听器
      this.player.on("play", () => {
        this.status = "playing";
      });

      this.player.on("pause", () => {
        this.status = "paused";
      });
      this.player.on("ended", () => {
        if (this.audioFiles.length >= this.playingIndex + 1) {
          // 播放下一首曲目
          this.selectTrack(this.playingIndex + 1);
        }
      });
    });
  },
  beforeDestroy() {
    // 销毁播放器
    this.player.destroy();
  },
  computed: {
    ...mapState("fm", {
      showModal: state => state.modal.showModal
    }),
    /**
     * 选择硬盘
     * @returns {*}
     */
    selectedDisk() {
      return this.$store.getters["fm/selectedDisk"];
    },

    /**
     * 音频文件列表
     * @returns {*}
     */
    audioFiles() {
      console.log(
        "音频【获取用户选择的文件】",
        this.$store.getters["fm/selectedItems"]
      );
      return this.$store.getters["fm/selectedItems"];
    }
  },
  methods: {
    /**
     * 播放或暂停
     */
    togglePlay() {
      this.player.togglePlay();
    },
    /**
     * 选择并播放曲目
     */
    selectTrack(i) {
      if (this.player.playing) this.player.stop();
      this.setSource(i); // 加载新资源
      this.playingIndex = i;
      this.player.play();
      this.player.autoplay = true;
      console.log("播放");
    },
    /**
     * 添加音频资源
     */
    setSource(i) {
      HTTP.get("stream_file", {
        params: {
          disk: this.selectedDisk,
          path: this.audioFiles[i].path
        }
      }).then(resp => {
        this.player.source = {
          type: "audio",
          title: this.audioFiles[i].filename,
          sources: [
            {
              src: resp.data,
              type: `audio/${this.audioFiles[i].extension}`
            }
          ]
        };
      });
      // this.player.source = {
      //   type: 'audio',
      //   title: this.audioFiles[index].filename,
      //   sources: [{
      //     src: `${this.$store.getters['fm/settings/baseUrl']}stream-file?disk=${this.selectedDisk}&path=${encodeURIComponent(this.audioFiles[index].path)}`,
      //     type: `audio/${this.audioFiles[index].extension}`,
      //   }],
      // };
    }
  }
};
</script>
<style lang="scss" scoped>
.audio-player {
  .audio-body {
    padding-bottom: 1rem;
    hr {
      border: 0;
      border-top: 1px solid rgba(0, 0, 0, 0.1);
    }
    .track {
      display: flex;
      justify-content: space-between;
      align-items: center;
      padding: 0.5rem;
      max-height: 30%;
      overflow-y: auto;
      .text-truncate {
        width: 75%;
        overflow: hidden;
        text-overflow: ellipsis;
        white-space: nowrap;
      }
      .fas.fa-play {
        color: gray;
        opacity: 0.1;
        cursor: pointer;

        &:hover {
          opacity: 0.5;
        }

        &.active {
          opacity: 1;
          color: deepskyblue;
        }
      }
      .fas.fa-pause {
        color: gray;
        opacity: 0.5;
        cursor: pointer;
      }
    }
    .bg-light {
      background: #f8f9fa;
    }
  }
}
</style>
