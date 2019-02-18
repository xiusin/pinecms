module.exports = {
  head: {
    titleTemplate: '%s - IT资源屋',
    meta: [
      { charset: 'utf-8' },
      { name: 'viewport', content: 'width=device-width, initial-scale=1' },
      {
        hid: 'description',
        name: 'description',
        content: 'Native-like Page Transitions with Vue and Nuxt, A Travel App'
      }
    ],
    script: [
      { src: 'https://cdn.vaptcha.com/v2.js' }  // 不知道怎么自定义到页面加载
    ]
  },
  router: {
    middleware: 'auth'
  },
  /*
  ** Customize the progress bar color
  */
  loading: { color: '#3B8070' },
  /*
  ** Build configuration
  */
  build: {
    vendor: ['external_library','element-ui'],
    extractCSS: { allChunks: true }
  },
  render: {
    bundleRenderer: {
      shouldPreload: (file, type) => {
        return ['script', 'style', 'font'].includes(type)
      }
    }
  },
  plugins: [
    { // 引入elementUI插件
      src: '~/plugins/ElementUI',
      ssr: true,
    },
    { // globle
      src: '~/plugins/global',
      ssr: true,
    }
  ],
}
