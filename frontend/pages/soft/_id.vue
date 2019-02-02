<template>
  <div style="width: 1160px; ">
    <el-container v-loading="loading" style="min-height: 700px;">
      <el-aside class="me-area">
        <ul class="me-month-list">
          <li v-for="a in archives" :key="a.Catid" class="me-month-item">
              <el-button @click="changeArchive(a.Catid,a.Catname)" size="small" style="width: 160px;">{{a.Catname}}
              </el-button>
          </li>
        </ul>
      </el-aside>

      <el-main class="me-articles" style="width: 720px;">
        <article-scroll-page v-bind="article"></article-scroll-page>
      </el-main>
    </el-container>
  </div>
</template>

<script>
  import ArticleScrollPage from '~/pages/common/ArticleScrollPage'
  import {getAllCategoryList} from "~/api/category"

  export default {
    name: "FreeCategory",
    components: {
      ArticleScrollPage
    },
     head () {
      return {
        title: this.title
      }
    },
    asyncData (context) {
      let params = context.route.params
      return getAllCategoryList(context.route.path.replace( params.id ? '/' + params.id : '' , '') + '/list').then((data => {
          const archives = data.data.data
          let catName = ''
          for(let i = 0; i < archives.length; i++) {
            if (archives[i].Catid == params.id) {
              catName = archives[i].Catname
            }
          }
          return {
            loading: false,
            archives: archives,
            title: params.id ? catName + ' - 软件' : '全部软件'
          }
        })).catch(error => {
          return {
            loading: false,
            archives: [],
            title: '全部软件'
          }
        })
    },
    data() {
      return {
        loading: true,
        article: {
          query: {
            id: this.$route.params.id,
          }
        },
        archives: []
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
        this.$setSeo(catname + ' - 软件')
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
