export const schema = {
  type: 'page',
  title: "数据表管理",
  body: {
    type: 'crud',
    api: '$preset.apis.list',
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
      {
        type: 'columns-toggler',
        align: 'left',
      },
      '$preset.actions.backup',
      'bulkActions',
    ],
    columns: [
      {
        name: 'id',
        label: '数据表',
        type: 'text',
      },
      {
        name: 'total',
        label: '记录数',
        type: 'text',
      },
      {
        name: 'engine',
        label: '存储引擎',
        type: 'text',
      },
      {
        name: 'comment',
        label: '表注释',
        type: 'text',
      },
      {
        type: 'operation',
        label: '操作',
        width: 160,
        buttons: ['$preset.actions.crud', '$preset.actions.menu'], //'$preset.actions.editField',
      },
    ],
  },
  preset: {
    actions: {
      backup: {
        limits: 'del',
        type: 'action',
        label: '备份数据库',
        size: 'xs',
        align: 'right',
        actionType: 'ajax',
        confirmText: "确定要备份数据库吗? 这可能需要几分钟!",
        api: "$preset.apis.backup"
      },
      crud: {
        type: 'action',
        label: '生成CRUD',
        size: 'xs',
        align: 'right',
        actionType: 'ajax',
        confirmText: "是否要一键生成CRUD内容？",
        api: "$preset.apis.crud"
      },
      menu: {
        type: 'action',
        label: '生成菜单',
        size: 'xs',
        align: 'right',
        actionType: 'ajax',
        confirmText: "确定要备份数据库吗? 这可能需要几分钟!",
        api: "$preset.apis.menu"
      },
    }
  }
}
