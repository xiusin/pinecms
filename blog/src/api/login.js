import request from '@/request'

export function register(account, nickname, password) {
  const data = {
    account,
    nickname,
    password
  }
  return request({
    url: '/user/register',
    method: 'post',
    data
  })
}

export function login(account, password) {
  // const data = {
  //   account,
  //   password
  // }
  return request({
    url: '/user/login',
    method: 'post',
    account,
    password
  })
}

export function logout() {
  return request({
    url: '/user/logout',
    method: 'get'
  })
}

export function getUserInfo() {
  return request({
    url: '/user/current',
    method: 'get'
  })
}

