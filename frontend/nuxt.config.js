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
    ]
  },
  router: {
    middleware: 'test'
  },
  /*
  ** Customize the progress bar color
  */
  loading: { color: '#3B8070' },
  /*
  ** Build configuration
  */
  build: {
    vendor: ['external_library','element-ui']
  },
  plugins: [
    { // 引入elementUI插件
      src: '~/plugins/ElementUI',
      ssr: true,
    },
    { // globle
      src: '~/plugins/globle',
      ssr: true,
    }
  ],
}
