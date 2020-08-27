export const schema = {
  type: 'page',
  body: {
    type: 'crud',
    api: '$preset.apis.list',
    quickSaveItemApi: "POST system/menu-quick-save",
    expandConfig: {
      accordion: true
    },
    headerToolbar: [
      '$preset.actions.add',
    ],
    columns: [
      {
        name: 'id',
        label: 'ID',
        type: 'text',
        width: 60,
      },
      {
        name: 'name',
        label: '菜单名称',
        quickEdit: true
      },
      {
        name: 'listorder',
        label: '排序',
        quickEdit: {
          saveImmediately: true,
        },
      },
      {
        type: 'operation',
        label: '操作',
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
          source: 'GET system/menu-tree',
          required: true,
        },
        {
          name: 'name',
          label: '菜单名称',
          type: 'text',
          required: true,
        },
        {
          name: 'c',
          label: '控制器名称',
          type: 'text',
          required: true,
        },
        {
          name: 'a',
          label: '路由名称',
          type: 'text',
          required: true,
        },
        {
          name: 'data',
          label: '附加数据',
          type: 'textarea',
        },
        {
          name: 'listorder',
          value: 30,
          label: '排序',
          type: 'number',
        },
        {
          name: 'display',
          value: true,
          label: '是否显示',
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
          title: '新增菜单',
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
  }
}
