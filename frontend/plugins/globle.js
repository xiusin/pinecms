import Vue from 'vue'
// 注入一些全局方法 , 混合模式 .可以在全局使用this.method
Vue.mixin({
  methods: {
    $setSeo (title, content) {
        return {
            title: title, 
            meta: [{ 
                hid: 'description', 
                name: 'description', 
                content: content 
            }] 
        }
    },
  }
})