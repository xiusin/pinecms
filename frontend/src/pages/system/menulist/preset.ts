export default {
  limits: {
    $page: {
      label: '查看列表',
    },
    add: {
      label: '添加',
    },
    edit: {
      label: '编辑',
    },
    del: {
      label: '删除',
    },
  },
  apis: {
    list: {
      url: 'GET system/menulist',
      limits: '$page',
      onPreRequest: (source) => {
        const { dateRange } = source.data
        if (dateRange) {
          const arr = dateRange.split('%2C')
          source.data = {
            ...source.data,
            startDate: arr[0],
            endDate: arr[1],
          }
        }
        return source
      },
    },
    add: {
      url: 'POST system/menu-add',
      limits: 'add',
    },
    edit: {
      url: 'POST system/menu-edit?id=$id',
      limits: 'edit',
    },
    del: {
      url: 'POST system/menu-delete?id=$id',
      limits: 'del',
    },
  },
}
