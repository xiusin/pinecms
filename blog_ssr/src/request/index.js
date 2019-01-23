import axios from 'axios'
import {Message} from 'element-ui'
import store from '@/store'
import {getToken} from '@/request/token'

const service = axios.create({
  baseURL: process.env.BASE_API,
  timeout: 10000
})

//request拦截器
service.interceptors.request.use(config => {
  if (store.state.token) {
    const token = getToken()
    if (token) config.headers['Authorization'] = "Bearer " + token
  }
  return config
}, error => {
  Promise.reject(error)
})

// respone拦截器
service.interceptors.response.use(
  response => {
    console.log('headers',response.headers)
    //全局统一处理 Session超时
    if (response.headers['Session_time_out'] == 'timeout') {
      store.dispatch('fedLogOut')
    }
    const res = response.data;
    //0 为成功状态
    if (res.status === false) {
      Message({
        type: 'warning',
        showClose: true,
        message: res.msg
      })
      return Promise.reject(res.msg);
    } else {
      return response.data;
    }
  },
  error => {
    Message({
      type: 'warning',
      showClose: true,
      message: '连接超时'
    })
    return Promise.reject('error')
  })

export default service
