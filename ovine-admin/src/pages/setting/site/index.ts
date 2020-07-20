export const schema = {
  type: 'page',
  title: "系统配置",
  body: [
    {
      "type": "service",
      "initFetchSchema": true,
      "schemaApi": "GET setting/site"
    }
  ]
}
