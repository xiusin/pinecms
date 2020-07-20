export const schema = {
  type: 'page',
  body: {
    type: 'lib-crud',
    bulkActions: [
      {
        label: '优化',
        size: 'xs',
        actionType: "ajax",
        confirmText: "确定要优化选中数据表吗? 这可能需要几分钟!",
        api: '$preset.apis.backup'
      },
      {
        label: '修复',
        size: 'xs',
        actionType: "ajax",
        confirmText: "确定要修复选中数据表吗? 这可能需要几分钟!",
        api: '$preset.apis.backup'
      },
    ],
    headerToolbar: [
      'bulkActions',
    ],
    api: '$preset.apis.list',
    columns: [
      {
        name: 'name',
        label: '缓存',
        type: 'text',
      },
      {
        name: 'description',
        label: '描述',
        type: 'text',
      },
      {
        type: 'operation',
        label: '操作',
        width: 60,
        limits: ['edit', 'del'],
        limitsLogic: 'or',
        buttons: ['$preset.actions.del'],
      },
    ],
  },
  preset: {
    actions: {
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
  },
}
