import Vue from "vue";
import App from "./App.vue";
import Storage from "vue-ls";
import VueHighlightJS from "vue-highlight.js";
import "highlight.js/styles/atom-one-dark.css";
import "vue-highlight.js/lib/allLanguages";

import { message } from "ant-design-vue";

Vue.config.productionTip = false;

Vue.prototype.$message = message;

message.config({
  duration: 2, // 持续时间
  top: `100px`, // 到页面顶部距离
  maxCount: 3 // 最大显示数, 超过限制时，最早的消息会被自动关闭
});

const storageOptions = {
  namespace: "apidoc_",
  name: "ls",
  storage: "local"
};

Vue.use(Storage, storageOptions);

Vue.use(VueHighlightJS);

new Vue({
  render: h => h(App)
}).$mount("#app");
