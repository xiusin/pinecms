<template>
  <div style="width: 1160px; ">
    <el-container v-loading="loading" style="min-height: 700px;">
      <el-aside class="me-area">
        <div class="me-category-title">资源分类</div>
        <ul class="me-month-list">
          <li v-for="category in categories" :key="category.Catid" class="me-month-item">
            <el-button @click="changeArchive(category.Catid,category.Catname)" size="small" style="width: 160px;">
              {{category.Catname}}
            </el-button>
          </li>
        </ul>
      </el-aside>

      <el-main class="me-articles" style="width: 885px;">
        <el-card v-if="articles.length" v-for="article in articles" :key="article.id" class="me-area" :body-style="{ padding: '16px' }">
          <div class="me-article-header">
            <nuxt-link :to="view(article.id)" class="me-article-title">{{article.title}}</nuxt-link>
            <el-button v-if="article.weight > 0" class="me-article-icon" type="text">置顶</el-button>
            <span class="me-pull-right me-article-count">
              <i class="el-icon-view"></i>&nbsp;150
            </span>
          </div>
          <div class="me-artile-description">
            {{article.summary}}
          </div>

          <div class="me-article-footer">
            <el-tag v-for="tag in article.tags" :key="tag" size="mini" type="success">{{tag}}</el-tag>
            <span class="me-pull-right me-article-count">
              <i class="el-icon-time"></i>&nbsp;2018-12-12 12:12:12
            </span>
          </div>
        </el-card>
        <el-card class="me-area" :body-style="{padding: '16px'}" v-else>
          没有任何内容
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
    head() {
      return {
        title: this.title
      }
    },
    asyncData(context) {
      let params = context.route.params
      return getAllCategoryList(context.route.path.replace(params.id ? '/' + params.id : '', '') + '/list').then((data => {
        const categories = data.data.data
        let catName = ''
        let curID = params.id || 30
        let innerPage = {
          pageSize: 10,
          pageNo: 1,
          name: 'id',
          sort: 'desc'
        }
        let query = {
          id: curID
        }
        for (let i = 0; i < categories.length; i++) {
          if (categories[i].Catid === Number(params.id)) {
            catName = categories[i].Catname
          }
        }
        return getArticles(query, innerPage).then(data => {
          data = data.data.data
          for (let i = 0; i < data.length; i++) {
            data[i].tags = data[i].tags.split(',')
          }
          return {
            loading: false,
            categories: categories,
            title: params.id ? catName + ' - 免费书籍' : '全部付费书籍',
            innerPage: innerPage,
            query: query,
            articles: data
          }
        }).catch(error => {
          error({statusCode: 404, message: 'Post not found'})
        })

      })).catch(error => {
        error({statusCode: 404, message: 'Post not found'})
      })
    },
    data() {
      return {
        loading: true,
        categories: [],
        articles: []
      }
    },
    created() {
    },
    watch: {
      '$route'(old, val) {
        console.log(old, val)
        if (this.$route.params.id) {
          this.article.query.id = this.$route.params.id
        }
      }
    },
    methods: {
      changeArchive(catid, catname) {
        let route = this.$route.path.replace('/' + this.$route.params.id, '')
        this.$router.push({path: route + '/' + catid})
      },
      view(artid) {
        return '/view/' + artid
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

  .me-article-header {
    /*padding: 10px 18px;*/
    padding-bottom: 10px;
  }

  .me-article-title {
    font-weight: 600;
  }

  .me-article-icon {
    padding: 3px 6px;
  }

  .me-article-count {
    color: #a6a6a6;
    padding-left: 8px;
    font-size: 13px;
  }

  .me-pull-right {
    float: right;
  }

  .me-article-description {
    font-size: 13px;
    line-height: 24px;
    margin-bottom: 10px;
  }
  a:visited {
    color: #303133;
  }
  .me-article-author {
    color: #a6a6a6;
    padding-right: 18px;
    font-size: 13px;
  }
  .me-category-title {
    width: 100%;
    font-size: 14px;
    height: 40px;
    line-height: 40px;
    text-align: center;
    border-bottom: 1px solid #cccccc;
  }
  .el-tag {
    margin-right: 6px;
    color: #010101;
    border-radius: 0;
    background-color: rgba(235,235,235, .12);
    border-color: rgba(150,150,150,1);
  }
  .el-card {
    margin-bottom: 10px;
    box-shadow: none;
  }
</style>
