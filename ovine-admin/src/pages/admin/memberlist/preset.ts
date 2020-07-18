export default {
  limits: {
    $page: {
      label: '查看列表',
    },
    addItem: {
      label: '添加',
    },
    editItem: {
      label: '编辑',
    },
    delItem: {
      label: '删除',
      needs: ['addItem', 'editItem'],
    },
  },
  apis: {
    list: {
      url: 'GET admin/memberlist',
      limits: '$page',
    },
    add: {
      url: 'POST admin/member-add',
      limits: 'addItem',
    },
    edit: {
      url: 'PUT admin/member-edit?id=$id',
      limits: 'editItem',
    },
    del: {
      url: 'POST admin/member-delete?id=$userid',
      limits: 'delItem',
    },
  },
}
