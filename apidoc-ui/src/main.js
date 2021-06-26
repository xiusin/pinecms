import Vue from "vue";
import App from "./App.vue";
import Storage from "vue-ls";
Vue.config.productionTip = false;

const storageOptions = {
  namespace: "apidoc_",
  name: "ls",
  storage: "local"
};

Vue.use(Storage, storageOptions);

new Vue({
  render: h => h(App)
}).$mount("#app");
