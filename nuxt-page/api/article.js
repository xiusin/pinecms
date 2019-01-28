import request from '@/request'

export function getArticles(query, page) {
  return request({
    url: '/RngNb/list',
    method: 'get',
    params: {
      pageNo: page.pageNo,
      pageSize: page.pageSize,
      six: page.name,
      sort: page.sort,
      catid: query.id,
      tagId: query.tagId
    }
  })
}

export function goPay(data) {
  return request({
    url: '/content/pay',
    method: 'post',
    data: data
  })
}

export function getNewArtices() {
  return request({
    url: '/article/new',
    method: 'get'
  })
}

export function viewArticle(id) {
  return request({
    url: `/content/info/${id}`,
    method: 'get'
  })
}

export function getArticleById(id) {
  return request({
    url: `/article/${id}`,
    method: 'get'
  })
}

export function publishArticle(article) {
  return request({
    url: '/article/publish',
    method: 'post',
    data: article
  })
}

export function listArchives() {
  return request({
    url: '/article/archives',
    method: 'get'
  })
}


/*
 * 以下俩接口暂时未用到
 * 可通过/article/list接口实现
 */
export function getArticlesByCategory(id) {
  return request({
    url: `/article/category/${id}`,
    method: 'get'
  })
}

export function getArticlesByTag(id) {
  return request({
    url: `/article/tag/${id}`,
    method: 'get'
  })
}

