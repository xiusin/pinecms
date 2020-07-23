export const schema = {
  type: 'page',
  title: "内容管理",
  // initApi : "https://houtai.baidu.com/api/mock2/form/initData?id=${id}",
  aside: {
    type: "wrapper",
    size: "xs",
    body: {
      type: "nav",
      stacked: true,
      source: "GET category/aside-category"
      // "links": [
      //   {
      //     "label": "页面1",
      //     "to": "?id=1"
      //   },
      //   {
      //     "label": "页面2",
      //     "children": [
      //       {
      //         "label": "页面2-1",
      //         "to": "?id=2-1"
      //       },
      //       {
      //         "label": "页面2-2",
      //         "to": "?id=2-2"
      //       },
      //       {
      //         "label": "页面2-3（disabled）",
      //         "disabled": true,
      //         "to": "?id=2-3"
      //       }
      //     ]
      //   },
      //   {
      //     "label": "页面3",
      //     "to": "?id=3"
      //   }
      // ]
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
