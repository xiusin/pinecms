import request from '@/request'

export function getAllFreeVideoList() {
  return request({
    url: '/free/video/list',
    method: 'get',
  })
}

export function getAllFreeBookList() {
  return request({
    url: '/free/book/list',
    method: 'get',
  })
}

export function getAllPaidVideoList() {
  return request({
    url: '/paid/video/list',
    method: 'get',
  })
}

export function getAllPaidBookList() {
  return request({
    url: '/paid/book/list',
    method: 'get',
  })
}
