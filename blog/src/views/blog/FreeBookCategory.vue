<template>
  <div v-title :data-title="title" style="margin: 0px auto; width: 980px; padding-top: 80px;">
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
        <article-scroll-page v-bind="article" ></article-scroll-page>
      </el-main>
    </el-container>
  </div>
</template>

<script>
  import ArticleScrollPage from '@/views/common/ArticleScrollPage'
  import {getAllFreeVideoList,getAllCategoryList} from "@/api/category"

  export default {
    name: "FreeCategory",
    components: {
      ArticleScrollPage
    },
    data() {
      return {
        currentArchive: '全部免费书籍',
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
      this.listArchives()
    },
    watch: {
      '$route'(old, val) {
        console.log(old,val)
        if (this.$route.params.id) {
          this.article.query.id = this.$route.params.id
        }
      }
    },
    computed: {
      title (){
        return this.currentArchive + ' - ' + window.title + ' - ' + window.keywords + ' - ' + window.description
      }
    },
    methods: {
      changeArchive(catid,catName) {
        //剔除其他数据
        let route = this.$route.path.replace('/' + this.$route.params.id , '')
        this.$router.push({path: route + '/' + catid})
        this.currentArchive = catName + ' - 免费书籍'
      },
      listArchives() {
        let that = this
        getAllCategoryList(this.$route.path.replace('/' + this.$route.params.id , '') + '/list').then((data => {
          this.loading = false
          that.archives = data.data
        })).catch(error => {
          that.$message({type: 'error', message: '分类加载失败!', showClose: true})
        })
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
    color: #5FB878;
  }

  .me-order-list {
    float: right;
  }

  .me-month-title {
    margin-left: 4px;
    margin-bottom: 12px;
  }
</style>
