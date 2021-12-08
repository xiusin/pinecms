const vueConfig = {
  publicPath: "",
  outputDir: "apidoc",
  assetsDir: "static",
  productionSourceMap: false,
  devServer: {
    port: 9998
  },
  css: {
    loaderOptions: {
      less: {
        lessOptions: {
          modifyVars: {
            "primary-color": "#1DA57A",
            "link-color": "#1DA57A",
            "border-radius-base": "2px"
          },
          javascriptEnabled: true
        }
      }
    }
  }
};

module.exports = vueConfig;
