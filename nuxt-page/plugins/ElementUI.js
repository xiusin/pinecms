import Vue from 'vue'
import lodash from 'lodash'
import ElementUI from 'element-ui'

import '@/assets/theme/index.css'
import '@/assets/icon/iconfont.css'
import '@/assets/font/iconfont.css'

Vue.config.productionTip = false

Vue.use(ElementUI)

// lodash是一个一致性、模块化、高性能的 JavaScript 实用工具库
Object.defineProperty(Vue.prototype, '$_', {value: lodash})

Vue.directive('title', function (el, binding) {
  document.title = el.dataset.title
})