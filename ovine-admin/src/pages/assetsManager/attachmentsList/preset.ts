export default {
  limits: {
    $page: {
      label: '查看列表',
    },
    del: {
      label: '删除',
    }
  },
  apis: {
    list: {
      url: 'GET assets-manager/attachments-list',
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
    del: {
      url: 'POST assets-manager/attachments-delete?id=$linkid',
      limits: 'del',
    },
  },
}
