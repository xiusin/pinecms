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
          name: 'group',
          label: '配置分组',
          type: 'select',
          creatable: true,
          source: 'GET setting/group',
          required: true,
        },
        {
          name: 'form',
          label: '配置名称',
          type: 'text',
          required: true,
        },
        {
          name: 'key',
          label: '配置Key',
          type: 'text',
          required: true,
        },
        {
          name: 'value',
          label: '配置值',
          type: 'text',
          required: true,
        },
        {
          name: 'editor',
          label: '配置类型',
          type: 'select',
          value: "text",
          options: [
            {
              "label": "文本类型",
              "value": "text"
            },
            {
              "label": "数字类型",
              "value": "number"
            },
            {
              "label": "开关类型",
              "value": "witch"
            },
            {
              "label": "下拉类型",
              "value": "select" //todo 下拉设置启动另一个字段合并
            },
            {
              "label": "图片类型",
              "value": "image"
            },
            {
              "label": "代码编辑器",
              "value": "editor"
            },
          ],
        },
        {
          name: 'extra',
          label: '额外配置',
          type: 'textarea',
          value: "text",
        },
      ],
    },
  },
}
