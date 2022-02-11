<template >
  <div class="preview">
    <button title="裁剪" class="cropper" @click="showCrop">
      <i class="fas fa-crop-alt" />
    </button>
    <!-- <template v-if="showCropper">
      <Cropper
        :imgSrc="imgSrc"
        :showCropper="showCropper"
        :maxHeight="maxHeight"
        :closeCropper="closeCropper"
      />
    </template>-->
  </div>
</template>
<script>
import imagePreview from "image-preview-vue";
import GET from "@/http/get";
import modal from "../mixins/modal";
import EventBus from "@/eventBus.js";
export default {
  mixins: [modal],
  data() {
    return {
      // showCropper: false,
      isHidden: false,
      config: {},
      imgSrc: "",
      imagePreview: null
    };
  },
  mounted() {
    this.preview();
  },
  methods: {
    showCrop() {
      EventBus.$emit("showCropper", true, this.imgSrc);
      this.imagePreview.close();
    },
    preview() {
      let leftImgs = [],
        rightImgs = [];
      this.config = {
        initIndex: 0,
        images: [],
        isEnableBlurBackground: true,
        isEnableLoopToggle: true,
        initViewMode: "contain",
        containScale: 1,
        shirnkAndEnlargeDeltaRatio: 0.2,
        wheelScrollDeltaRatio: 1,
        isEnableImagePageIndicator: true,
        maskBackgroundColor: "rgba(0,0,0,0.6)",
        zIndex: 4000,
        isEnableKeyboardShortcuts: true
      };
      this.config.onClose = () => {
        this.hideModal();
      };

      this.config.initIndex = this.imageItems.findIndex(
        img => img.basename == this.selectedItem.basename
      );
      console.log("this.config.images2", this.config.images);
      console.log("this.config.initIndex", this.config.initIndex);
      // 图片链接不存在则请求
      if (!this.imageUrl.length) {
        let sort = this.$store.getters[`fm/${this.activeManager}/sort`];
        this.$store.commit(`fm/${this.activeManager}/setImagesSortField`, {
          field: sort.field,
          direction: sort.direction
        });
        for (let key = 0; key < this.config.initIndex; key++) {
          leftImgs.push(
            GET.thumbnailLink(this.selectedDisk, this.imageItems[key].path)
          );
          if (key == this.config.initIndex) return;
        }
        for (
          let key = this.config.initIndex + 1;
          key < this.imageItems.length;
          key++
        ) {
          rightImgs.push(
            GET.thumbnailLink(this.selectedDisk, this.imageItems[key].path)
          );
        }
        this.getImgUrl([
          ...leftImgs,
          GET.thumbnailLink(this.selectedDisk, this.selectedItem.path),
          ...rightImgs
        ]).then(data => {
          for (let url of data) {
            this.config.images.push(url.data);
          }
          this.imagePreview = imagePreview(this.config);
        });
      } else {
        this.config.images = this.imageUrl;
        console.log(
          "this.config.images3",
          this.config.images,
          "this.imageItems",
          this.imageItems
        );
        // this.config.images = this.imageItems;
        this.imgSrc = this.config.images[this.config.initIndex];
        this.imagePreview = imagePreview(this.config);
      }
    },
    async getImgUrl(arr) {
      let res = [];
      for (let fn of arr) {
        let data = await fn;
        res.push(data);
      }
      return await res;
    },
    /**
     * 关闭裁剪
     */
    closeCropper() {
      this.showCropper = false;
      // this.loadImage();
    },
    /**
     * 加载图片
     */
    loadImage() {
      // if authorization required
      // if (this.auth) {
      //   GET.preview(
      //     this.selectedDisk,
      //     this.selectedItem.path,
      //   ).then((response) => {
      //     const mimeType = response.headers['content-type'].toLowerCase();
      //     const imgBase64 = Buffer.from(response.data, 'binary').toString('base64');

      //     this.imgSrc = `data:${mimeType};base64,${imgBase64}`;
      //   });
      // } else {
      //   this.imgSrc = `${this.$store.getters['fm/settings/baseUrl']}preview?disk=${this.selectedDisk}&path=${encodeURIComponent(this.selectedItem.path)}&v=${this.selectedItem.timestamp}`;
      // }
      GET.thumbnailLink(this.selectedDisk, this.selectedItem.path).then(
        response => {
          this.imgSrc = response.data;
        }
      );
    }
  },
  beforeDestroy() {
    this.imagePreview = null;
  },
  computed: {
    /**
     * 当前磁盘
     * @returns {String}
     */
    activeManager() {
      return this.$store.state.fm.activeManager;
    },
    /**
     * 选择的磁盘
     * @returns {String}
     */
    selectedDisk() {
      return this.$store.getters["fm/selectedDisk"];
    },

    /**
     * 选择的文件
     * @returns {Object}
     */
    selectedItem() {
      return this.$store.getters["fm/selectedItems"][0];
    },
    /**
     * 返回该目录下所有图片文件
     * @returns
     */
    imageItems() {
      return this.$store.getters[`fm/${this.activeManager}/imageFiles`];
    },
    /**
     * 返回改目录下的图片链接
     */
    imageUrl() {
      return this.$store.getters[`fm/${this.activeManager}/imageUrl`];
    }
  },
  watch: {
    "config.images.length": {
      handler: function(newV) {
        if (newV == this.imageItems.length && this.imageUrl.length == 0) {
          this.$store.commit(
            `fm/${this.activeManager}/setImageUrl`,
            this.config.images
          );
        }
      },
      deep: true
    },
    "imagePreview.$instance.currentIndex": {
      handler(newV) {
        this.imgSrc = this.config?.images[newV];
        let path = this.imageItems[newV]?.path;
        if (path) {
          this.$store.commit(`fm/${this.activeManager}/removeAllSelected`);
          this.$store.commit(`fm/${this.activeManager}/setSelected`, {
            type: "files",
            path: path
          });
        }
      },
      deep: true
      // immediate: true
    }
  }
};
</script>
<style lang="scss" scoped>
.preview {
  .cropper {
    position: absolute;
    border: 0;
    padding: 0;
    right: 80px;
    bottom: 16px;
    z-index: 9999;
    color: #fff;
    background-color: #17a2b8;
    border-color: #17a2b8;
    padding: 0.375rem 0.75rem;
    font-size: 1rem;
    line-height: 1.5;
    border-radius: 0.25rem;
    &:hover {
      color: #fff;
      background-color: hsla(0, 0%, 100%, 0.7);
      border-collapse: hsla(0, 0%, 100%, 0.7);
      text-decoration: none;
      color: #333333;
      cursor: pointer;
    }
  }
}
</style>