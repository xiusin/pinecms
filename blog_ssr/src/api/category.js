import request from '@/request'

export function getAllCategoryList(path) {
  return request({
    url: path,
    method: 'get',
  })
}

