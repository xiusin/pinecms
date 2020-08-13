export const schema = {
  type: 'page',
  body: {
    type: 'lib-crud',
    api: '$preset.apis.list',
    draggable: true,
    saveOrderApi: 'POST /category/category-order',
    expandConfig: {
      accordion: true
    },
    headerToolbar: [
      'filter-toggler',
      {
        type: 'columns-toggler',
        align: 'left',
      },
      '$preset.actions.add',
    ],
    columns: [
      {
        name: 'catid',
        label: 'ID',
        type: 'text',
      },
      {
        name: 'listorder',
        label: '排序',
        type: 'text',
      },
      {
        name: 'catname',
        label: '栏目名称',
        type: 'text',
      },
      {
        name: 'dir',
        label: '静态目录',
        type: 'text',
      },
      {
        name: 'type',
        type: 'tpl',
        label: '类型',
        tpl: '<div class="<%= ["text-success", "text-info", "text-primary"][data.type] %>"><%=["列表栏目", "频道单页", "外部链接"][data.type] %> </div>',
      },
      {
        name: 'model_id',
        label: '模型',
        type: 'tpl',
        tpl: '<div class="text-info">${model_id}</div>',
      },
      {
        type: 'operation',
        label: '操作',
        width: 60,
        limits: ['edit', 'del'],
        limitsLogic: 'or',
        buttons: ['$preset.actions.edit', '$preset.actions.del'],
      },
    ],
  },
  definitions: {
    updateControls: {
      controls: [
        {
          name: 'parentid',
          label: '上级栏目',
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
          visibleOn: "data.type != 2"
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
            mode: 'horizontal',
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
            mode: 'horizontal',
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
