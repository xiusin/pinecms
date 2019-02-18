
import Cookie from 'js-cookie'
export default function ({ store, req, redirect, route }) {
  let token = ''
  if (req && req.headers && req.headers.cookie) {
    token = req.headers.cookie.split(';').find(c => c.trim().startsWith('token='))
    if (token) {
      store.commit("SET_TOKEN", token)  // 从nuxt上获取token 并且设置
    } else {
      console.log('server no token')
    }
  } else {
    if (token = Cookie.get('token')) {
      store.commit("SET_TOKEN", token) // 从cookie里设置token
    } else {
      console.log('client no token')
    }
  }

  if (route.path.startsWith('/user/') && !token) {
      return redirect('/login')
  }
}
