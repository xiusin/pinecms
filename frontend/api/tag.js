import request from '@/request'

export function getAllTags() {
  return request({
    url: '/tags/list',
    method: 'get',
  })
}

export function getHotTags() {
  return request({
    url: '/tags/hot',
    method: 'get',
  })
}

export function getAllTagsDetail() {
  return request({
    url: '/tags/detail',
    method: 'get',
  })
}

export function getTagDetail(id) {
  return request({
    url: `/tags/detail/${id}`,
    method: 'get',
  })
}


/*
 * 暂时未用到
 */
export function getTag(id) {
  return request({
    url: `/tags/${id}`,
    method: 'get',
  })
}

