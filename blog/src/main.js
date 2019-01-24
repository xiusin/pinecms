import Vue from 'vue'
import App from './App'

import router from './router'
import store from './store'
import lodash from 'lodash'
import ElementUI from 'element-ui'

import '@/assets/theme/index.css'
import '@/assets/icon/iconfont.css'
import '@/assets/font/iconfont.css'

import {formatTime} from "./utils/time";

Vue.config.productionTip = false

Vue.use(ElementUI)

// lodash是一个一致性、模块化、高性能的 JavaScript 实用工具库
Object.defineProperty(Vue.prototype, '$_', {value: lodash})

Vue.directive('title', function (el, binding) {
  document.title = el.dataset.title
})

// 格式化时间
Vue.filter('format', formatTime)

new Vue({
  el: '#app',
  router,
  store,
  template: '<App/>',
  components: {App},
  // 添加mounted，不然不会执行预编译
  mounted () {
    document.dispatchEvent(new Event('render-event'))
  }
})
