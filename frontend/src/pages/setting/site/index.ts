export const schema = {
  type: 'page',
  title: "系统配置",
  body: [
    {
      "type": "service",
      "initFetchSchema": true,
      "schemaApi": "GET setting/site"
    }
  ],
  definitions: {
    updateControls: {
      controls: [
        {
          name: 'name',
          label: '所属分组',
          type: 'tree-select',
          source: 'GET category/category-select',
          required: true,
        },
        {
          name: 'catname',
          label: '栏目名称',
          type: 'text',
          required: true,
        },
        {
          name: 'type',
          label: '栏目类型',
          type: 'select',
          required: true,
          options: [
            {
              "label": "栏目",
              "value": "0"
            },

            {
              "label": "单页",
              "value": "1"
            },

            {
              "label": "链接",
              "value": "2"
            },
          ]
        },
        {
          name: 'model_id',
          label: '文档模型',
          type: 'select',
          source: "GET public/select?type=models",
          required: true,
          hiddenOn: "data.type != 0"
        },
        {
          name: 'thumb',
          label: '缩略图',
          reciever: 'POST public/upload',
          type: 'image',
        },
        {
          name: 'url',
          label: '链接',
          type: 'text',
          visibleOn: "data.type == 2"
        },
        {
          name: 'dir',
          label: '栏目目录',
          type: 'text',
          hiddenOn: "data.type == 2"
        },
        {
          name: 'list_tpl',
          label: '列表页面',
          type: 'select',
          source: "GET public/select?type=tpl_list",
          clearable: true,
          hiddenOn: "data.type != 0"
        },
        {
          name: 'detail_tpl',
          label: '详情页面',
          type: 'select',
          source: "GET public/select?type=tpl_list",
          clearable: true,
          hiddenOn: "data.type == 2",
        },
        {
          name: 'keywords',
          label: '关键词',
          type: 'text',
        },
        {
          name: 'description',
          label: '描述',
          type: 'textarea',
        },
        {
          name: 'ismenu',
          value: true,
          label: '启用',
          type: 'switch',
        },
        {
          name: 'listorder',
          value: 30,
          label: '排序',
          type: 'number',
        }
      ],
    },
  },
}
