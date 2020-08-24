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
  },
  apis: {
    list: {
      url: 'GET assets-manager/list',
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
      url: 'POST assets-manager/add',
      limits: 'add',
    },
    edit: {
      url: 'POST assets-manager/edit?name=$name',
      limits: 'edit',
    }
  },
}
