export const schema = {
  type: 'page',
  title: "内容管理",
  aside: {
    type: "wrapper",
    size: "xs",
    body: {
      type: "nav",
      stacked: true,
      source: "GET category/aside-category"
    }
  },
  body: [
    {
      "type": "service",
      "initFetchSchema": true,
      "schemaApi": "GET content/news-list?catid=$catid"
    }
  ]
}
