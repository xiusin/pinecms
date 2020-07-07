export default {
  'POST user/login': () => {
    return {
      code: 0,
      data: {
        key: 'X-TOKEN',
        token: 'mockToken',
      },
    }
  },
}
