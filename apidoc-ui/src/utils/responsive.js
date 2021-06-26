import enquireJs from "enquire.js";

// 响应式
const responsiveMap = {
  xs: "(max-width: 575px)",
  sm: "(min-width: 576px)",
  md: "(min-width: 768px)",
  lg: "(min-width: 992px)",
  xl: "(min-width: 1200px)",
  xxl: "(min-width: 1600px)"
};

export const DEVICE_TYPE = {
  DESKTOP: "desktop",
  TABLET: "tablet",
  MOBILE: "mobile"
};

const DEVICE_SIZE = {
  DESKTOP: ["xl", "xxl"],
  TABLET: ["lg"],
  MOBILE: ["xs", "sm", "md"]
};

export default {
  data() {
    return {
      currentSize: ""
    };
  },
  computed: {
    device() {
      const { currentSize } = this;
      let type = "";
      for (const key in DEVICE_SIZE) {
        if (DEVICE_SIZE[key].includes(currentSize)) {
          type = DEVICE_TYPE[key];
          break;
        }
      }
      return type;
    }
  },
  mounted() {
    const that = this;
    this.$nextTick(() => {
      const keys = Object.keys(responsiveMap);
      keys.map(screen =>
        enquireJs.register(responsiveMap[screen], {
          match: () => {
            that.currentSize = screen;
          },
          unmatch: () => {
            const keyIndex = keys.findIndex(p => p === screen);
            if (keyIndex > 0) {
              const newKeyIndex = keyIndex - 1;
              that.currentSize = keys[newKeyIndex];
            }
          },
          destroy() {}
        })
      );
    });
  },
  beforeDestroy() {
    Object.keys(responsiveMap).map(screen =>
      enquireJs.unregister(responsiveMap[screen])
    );
  }
};
