export const schema = {
  type: 'page',
  title: "内容管理",
  aside: {
    type: "form",
    wrapWithPanel: false,
    target: "window",
    submitOnChange: true,
    controls: [
      {
        type: "tree",
        name: "catid",
        inputClassName: "no-border",
        source: "GET content/aside-category"
      }
    ]
  },
  body: [
    {
      type: "service",
      name: "window",
      initFetchSchema: false,
      schemaApi: "GET content/news-crud?catid=$catid"
    },
  ],
  definitions: {
    updateControls: {
      controls: [
        {
          name: 'name',
          label: '名称',
          type: 'text',
          required: true,
        },
        {
          name: 'logo',
          label: 'Logo',
          reciever: 'POST public/upload',
          type: 'image',
        },
        {
          name: 'url',
          label: '链接',
          type: 'url',
          required: true,
        },
        {
          name: 'introduce',
          label: '描述',
          type: 'textarea',
        },
        {
          name: 'passed',
          value: true,
          label: '启用',
          type: 'switch',
        },
        {
          name: 'listorder',
          value: 30,
          label: '排序',
          type: 'text',
        }
      ],
    },
  },
  preset: {
    actions: {
      add: {
        limits: 'add',
        type: 'button',
        align: 'right',
        actionType: 'dialog',
        label: '添加',
        icon: 'fa fa-plus pull-left',
        size: 'sm',
        primary: true,
        dialog: {
          title: '新增文档',
          size: 'lg',
          body: {
            type: 'form',
            api: '$preset.apis.add',
            mode: 'normal',
            $ref: 'updateControls',
          },
        },
      },
      edit: {
        limits: 'edit',
        type: 'button',
        icon: 'fa fa-pencil',
        tooltip: '编辑',
        actionType: 'dialog',
        dialog: {
          title: '编辑文档',
          size: 'lg',
          body: {
            type: 'form',
            mode: 'normal',
            api: '$preset.apis.edit',
            $ref: 'updateControls',
          },
        },
      },
      del: {
        limits: 'del',
        type: 'action',
        icon: 'fa fa-times text-danger',
        actionType: 'ajax',
        tooltip: '删除',
        confirmText: '您确认要删除?',
        api: {
          $preset: 'apis.del',
        },
        messages: {
          success: '删除成功',
          failed: '删除失败',
        },
      },
    },
    forms: {
      filter: {
        controls: [
          {
            type: 'date-range',
            name: 'dateRange',
            label: '创建时间范围',
            format: 'YYYY-MM-DD',
          },
          {
            type: 'submit',
            className: 'm-l',
            label: '搜索',
          },
        ],
      },  // 搜索
    },
  },
}
