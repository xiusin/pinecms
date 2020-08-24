export default {
  limits: {
    $page: {
      label: '查看列表',
    },
    del: {
      label: '清理',
    },
  },
  apis: {
    list: {
      url: 'GET setting/cache'
    },
    dels: {
      url: 'POST setting/del-cache?keys=${keys|raw}',
      limits: 'del',
    },
    del: {
      url: 'POST setting/del-cache?key=$key',
      limits: 'del',
    },
  },
}
