import request from '@/request'

export function getAllCategorys() {
  return request({
    url: '/category/list',
    method: 'get',
  })
}

export function getAllCategorysDetail() {
  return request({
    url: '/category/detail',
    method: 'get',
  })
}

export function getCategoryDetail(id) {
  return request({
    url: `/category/detail/${id}`,
    method: 'get',
  })
}

 /*
  * 暂时未用到
  */
export function getCategory(id) {
  return request({
    url: `/category/${id}`,
    method: 'get',
  })
}
