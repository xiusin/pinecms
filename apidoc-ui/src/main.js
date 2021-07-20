import Vue from "vue";
import App from "./App.vue";
import Storage from "vue-ls";
import VueHighlightJS from "vue-highlight.js";
import "highlight.js/styles/atom-one-dark.css";
import "vue-highlight.js/lib/allLanguages";

Vue.config.productionTip = false;

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
