export const schema = {
  type: 'page',

  body: {
    type: 'lib-crud',
    api: '$preset.apis.list',
    filter: '$preset.forms.filter',
    filterTogglable: true,
    perPageAvailable: [50, 100, 200],
    defaultParams: {
      size: 50,
    },
    perPageField: 'size',
    pageField: 'page',
    headerToolbar: [
      'filter-toggler',
      {
        type: 'columns-toggler',
        align: 'left',
      },
      {
        type: 'pagination',
        align: 'left',
      },
      '$preset.actions.add',
    ],
    footerToolbar: ['statistics', 'switch-per-page', 'pagination'],
    columns: [
      {
        name: 'id',
        label: 'ID',
        type: 'text',
      },
      {
        name: 'listorder',
        label: '排序',
        type: 'text',
      },
      {
        name: 'name',
        label: '名称',
        type: 'text',
      },
      {
        name: 'image',
        label: '图片',
        type: 'image',
      },
      {
        name: 'space_name',
        label: '广告位',
        type: 'text',
      },
      {
        name: 'status',
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
          name: 'name',
          label: '名称',
          type: 'text',
          required: true,
        },
        {
          type: "select",
          name: "space_id",
          label: "广告位",
          // clearable: true,
          required: true,
          source: "GET public/select?type=ad_space",
        },
        {
          name: 'image',
          label: '图片',
          reciever: 'POST public/upload',
          type: 'image',
        },
        {
          name: 'link_url',
          label: '链接',
          type: 'url',
          required: true,
        },
        {
          type: "datetime",
          name: "start_time",
          label: "开始日期",
          maxDate: "$end_time"
        },
        {
          type: "datetime",
          name: "end_time",
          label: "结束日期",
          minDate: "$start_time"
        },
        {
          name: 'listorder',
          value: 30,
          label: '排序',
          type: 'text',
        },
        {
          name: 'status',
          value: true,
          label: '启用',
          type: 'switch',
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
