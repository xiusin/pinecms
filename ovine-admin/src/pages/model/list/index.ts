export const schema = {
  type: 'page',
  body: {
    type: 'lib-crud',
    api: '$preset.apis.list',
    filter: '$preset.forms.filter',
    filterTogglable: true,
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
        name: 'id',
        label: 'ID',
        type: 'text',
      },
      {
        name: 'name',
        label: '栏目名称',
        type: 'text',
      },
      {
        name: 'enabled',
        label: '启用',
        type: 'switch',
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
          name: "name",
          label: "模型名称",
          type: "text",
          require: true,
          size: "full"
        },
        {
          name: "table",
          label: "模型Table",
          type: "text",
          require: true,
          size: "full"
        },
        {
          name: "enabled",
          label: "是否启用",
          type: "switch",
          require: true,
          size: "full"
        },
        {
          "type": "combo",
          "name": "combo101",
          "multiple": true,
          "multiLine": true,
          "label": "模型定义",
          addButtonText: ' 字段',
          "value": [],
          "tabsMode": true,
          "tabsStyle": "radio",
          mode: 'horizontal',
          "maxLength": 40,
          "tabsLabelTpl": "$name",
          "controls": [
            {
              name: "name",
              label: "表单名称",
              type: "text",
              require: true,
            },
            {
              name: "field",
              label: "字段名称",
              type: "text",
              require: true,
            },
            {
              name: "listorder",
              label: "字段排序",
              type: "text",
              require: true,
            },
            {
              name: "required",
              label: "是否必填",
              type: "switch",
            },
            {
              name: "required_tips",
              label: "必填提醒内容",
              type: "text",
            },
            {
              name: "type",
              require: true,
              label: "字段类型",
              source: "GET public/select?type=fields",
              loadDataOnce: true,
              type: "select",
            },
            {
              name: "datasource",
              label: "数据源",
              type: "textarea",
              placeholder: "组件数据源, 写入Json或者接口地址"
            },
            {
              name: "default",
              label: "数据源",
              type: "text",
              placeholder: "数据源选项, 仅对单选,多选,开关按钮,多个用|分割"
            },
            {
              name: "validator",
              label: "字段验证器",
              type: "text",
              placeholder: "amis验证规则或自定义"
            },
          ]
        },
        {
          name: 'tpl_list',
          label: '列表页面',
          type: 'select',
          source: "GET public/select?type=tpl_list",
          clearable: true,
        },
        {
          name: 'tpl_detail',
          label: '详情页面',
          type: 'select',
          source: "GET public/select?type=tpl_list",
          clearable: true,
        },
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
          title: '新增模型',
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
