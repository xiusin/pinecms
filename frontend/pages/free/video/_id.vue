<template>
  <div style="width: 1160px; ">
    <el-container v-loading="loading" style="min-height: 700px;">
      <el-aside class="me-area">
        <ul class="me-month-list">
          <li v-for="category in categories" :key="category.Catid" class="me-month-item">
              <el-button @click="changeArchive(category.Catid,category.Catname)" size="small" style="width: 160px;">{{category.Catname}}
              </el-button>
          </li>
        </ul>
      </el-aside>

      <el-main class="me-articles" style="width: 720px;">
        <el-card v-for="article in articles" :key="article.id" class="me-area" :body-style="{ padding: '16px' }">
          <div class="me-article-header">
            <a @click="view(article.id)" class="me-article-title">{{article.title}}</a>
            <el-button v-if="article.weight > 0" class="me-article-icon" type="text">置顶</el-button>
            <span class="me-pull-right me-article-count">
              <i class="me-icon-comment"></i>&nbsp;{{article.commentNum}}
            </span>
                  <span class="me-pull-right me-article-count">
              <i class="el-icon-view"></i>&nbsp;{{article.viewNum}}
            </span>
                </div>
                <div class="me-artile-description">
                  {{article.summary}}
                </div>

                <div class="me-article-footer">
                  <span class="me-article-author">
                    <i class="me-icon-author"></i>&nbsp;{{article.nickname}}
                  </span>
                  <el-tag  v-for="tag in article.tags" :key="t" size="mini" type="success">{{tag}}</el-tag>
                  <span class="me-pull-right me-article-count">
              <i class="el-icon-time"></i>&nbsp;{{article.createTime}}
            </span>
          </div>
        </el-card>
      </el-main>
    </el-container>
  </div>
</template>

<script>
  import {getAllCategoryList} from "~/api/category"
  import {getArticles} from "~/api/article"
  export default {
    name: "FreeCategory",
     head () {
      return {
        title: this.title
      }
    },
    asyncData (context) {
      let params = context.route.params
      return getAllCategoryList(context.route.path.replace( params.id ? '/' + params.id : '' , '') + '/list').then((data => {
          const categories = data.data.data
          let catName = ''
          let curID = params.id || 27
          let innerPage = {
              pageSize: 10,
              pageNo: 1,
              name: 'id',
              sort: 'desc'
          }
          let query = {
            id: curID
          }
          for(let i = 0; i < categories.length; i++) {
            if (categories[i].Catid == params.id) {
              catName = categories[i].Catname
            }
          }
          let c = getArticles(query, innerPage).then(data => {
            return data.data.data
          }).catch(error => {
            return []
          })

          console.log(JSON.stringify(c)) //todo 如何直接拿到结果呢

          return {
            loading: false,
            categories: categories,
            title: params.id ? catName + ' - 免费视频' : '全部免费视频',
            innerPage: innerPage,
            query: query
          }
        })).catch(error => {
          console.log(error)
          return {
            loading: false,
            categories: [],
            title: '全部免费视频'
          }
        })
    },
    data() {
      return {
        loading: true,
        categories: [],
        articles: [],

      }
    },
    created() {
    },
    watch: {
      '$route'(old, val) {
        console.log(old,val)
        if (this.$route.params.id) {
          this.article.query.id = this.$route.params.id
        }
      }
    },
    methods: {
      changeArchive(catid,catname) {
        let route = this.$route.path.replace('/' + this.$route.params.id , '')
        this.$router.push({path: route + '/' + catid})
        this.$setSeo(catname + ' - 免费视频')
      }
    }
  }
</script>

<style scoped>

  .el-aside {
    margin-right: 50px;
    width: 200px !important;
  }

  .el-main {
    padding: 0px;
    line-height: 16px;
    flex: none;
  }

  .me-month-list {
    margin-top: 10px;
    margin-bottom: 10px;
    text-align: center;
    list-style-type: none;
  }

  .me-month-item {
    margin-top: 18px;
    padding: 4px;
    font-size: 18px;
    color: #010101;
  }

  .me-order-list {
    float: right;
  }

  .me-month-title {
    margin-left: 4px;
    margin-bottom: 12px;
  }
</style>
