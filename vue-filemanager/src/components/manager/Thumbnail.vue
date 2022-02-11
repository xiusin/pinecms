<template >
  <figure class="thumbnail">
    <transition name="fade" mode="out-in">
      <i v-if="!src" class="fa fa-5x fa-file-image" style="color: #ff7743;"></i>
      <img v-else :src="src" :alt="file.filename" class="img-thumbnail" />
    </transition>
  </figure>
</template>
<script>
import GET from "@/http/get";
export default {
  props: {
    disk: {
      type: String,
      required: true
    },
    file: {
      type: Object,
      required: true
    }
  },
  data() {
    return {
      src: ""
    };
  },
  watch: {
    "file.timestamp": "loadImage"
  },
  computed: {
    /**
     * 获取授权
     */
    auth() {
      return this.$store.getters["fm/settings/authHeader"];
    }
  },
  mounted() {
    if (window.IntersectionObserver) {
      const observer = new IntersectionObserver(
        (entries, obs) => {
          entries.forEach(entry => {
            if (entry.isIntersecting) {
              // 是否出现在可视区
              this.loadImage();
              obs.unobserve(this.$el);
            }
          });
        },
        {
          root: null, // 指定目标元素所在的容器节点（即根元素）
          threshold: "0.5" //交叉比为0时触发回调函数
        }
      );
      observer.observe(this.$el); // 添加观察元素
    } else {
      this.loadImage();
    }
  },
  methods: {
    /**
     * 加载图片
     */
    loadImage() {
      GET.thumbnailLink(this.disk, this.file.path).then(resp => {
        this.src = resp.data;
      });

      // if (this.auth) {
      //   GET.thumbnail(
      //     this.disk,
      //     this.file.path,
      //   ).then((response) => {
      //     const mimeType = response.headers['content-type'].toLowerCase();
      //     const imgBase64 = Buffer.from(response.data, 'binary').toString('base64');
      //     this.src = `data:${mimeType};base64,${imgBase64}`;
      //   });
      // } else {
      //   this.src = `${this.$store.getters['fm/settings/baseUrl']}thumbnails?disk=${this.disk}&path=${encodeURIComponent(this.file.path)}&v=${this.file.timestamp}`;
      // }
    }
  }
};
</script>
<style lang="scss" scoped>
.thumbnail {
  display: flex;
  justify-content: center;
  align-items: center;
  margin: 0;
  .img-thumbnail {
    object-fit: cover;
    width: 88px;
    height: 88px;
  }

  .fade-enter-active,
  .fade-leave-active {
    transition: opacity 0.3s;
  }
  .fade-enter,
  .fade-leave-to {
    opacity: 0;
  }
}
</style>