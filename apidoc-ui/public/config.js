// eslint-disable-next-line no-unused-vars
var config = {
  // 请求host
  HOST: "",
  // 菜单配置
  MENU: {
    // 是否显示控制器类名
    SHOW_CONTROLLER_CLASS: true,
    // 是否显示接口url
    SHOW_API_URL: true,
    // 是否显示接口请求类型
    SHOW_API_METHOD: true
  },
  // 当字段无默认值时，使用字段类型为默认值
  USE_TYPE_DEFAULT_VALUE: true,
  HOSTS: [
    {
      title: "本地测试",
      host: "https://apidoc.demo.hg-code.com/"
    },
    {
      title: "正式环境",
      host: "http://www.apidoc.net.cn"
    },
  ]
};
