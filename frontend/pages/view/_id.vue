<template>
  <div class="me-view-body">
    <el-container class="me-view-container" style="width: 1160px;">
      <el-main>
        <div class="me-view-card">
          <h1 class="me-view-title">{{article.title}}</h1>
          <div class="me-view-content" v-html="article.content"></div>

          <el-card shadow="hover" style="
    padding: 3px;
    background-color: #ecf8ff;
    border-radius: 4px;
    border-left: 5px solid #50bfff;
    font-size: 13px;
" v-if="article">
            <b style="font-size: 14px; margin-bottom: 8px;">资源地址: <a :href="article.source_url" target="_blank">{{this.article.source_url}}</a></b>
            <br/>
            <br/>
            <template v-if="article.pwd_type === 2 && article.money > 0">
              <el-button @click="getPayUrl('alipay')">支付宝去获取资源密码 (积分: {{article.money}})</el-button>
              <el-button @click="getPayUrl('wechat')">微信去获取资源密码 (积分: {{article.money}})</el-button>
            </template>
            <template v-else>
              获取下载密码:<br/>
              &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[打开微信]->[扫描<span style="color: orangered;cursor: pointer"
                                                                         @click="dialogTableVisible=true">二维码</span>]->[关注公众号]
              <br/>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;输入 <i style="color: orangered">{{article.id}}</i> 获取分享码
              如果取消关注本公众号，再次关注即可再次享受服务！
            </template>
          </el-card>

          <el-dialog title="公众号" :visible.sync="dialogTableVisible" width="395px">
            <img src="~/static/qrcode/wechat.jpg" style="width: 350px"/>
          </el-dialog>

          <div class="me-view-tag">
            标签：
            <el-button @click="tagOrCategory('tag', t)" size="mini" type="primary" v-for="t in article.tags" :key="t">
              {{t}}
            </el-button>
          </div>

          <div class="me-view-tag">
            文章分类：
            <el-button @click="tagOrCategory('category', article.category.id)" size="mini" type="primary">
              我是分类标题
            </el-button>
          </div>
        </div>
      </el-main>

    </el-container>
  </div>
</template>

<script>
  import {viewArticle} from '@/api/article'

  export default {
    name: 'BlogView',
    watch: {
      '$route': 'getArticle'
    },
    head() {
      return {
        title: this.article.title
      }
    },
    data() {
      return {
        dialogTableVisible: false,
        article: {
          id: '',
          title: '',
          commentNum: 0,
          viewNum: 0,
          summary: '',
          createTime: '',
          author: {},
          tags: [],
          category: {},
        }
      }
    },
    asyncData({ params, error }) {
      return viewArticle(params.id).then(data => {
        data.data.data.tags = data.data.data.tags.split(',')
        return {
          article: data.data.data
        }
      }).catch(e => {
        error({statusCode: 404, message: '文章加载失败'})
      })
    },
    methods: {
      getPayUrl(payType) {
        this.$router.push({path: `/buy/${this.$route.params.id}/` + payType})
      },
      tagOrCategory(type, id) {
        this.$router.push({path: `/${type}/${id}`})
      }
    },
    // //组件内的守卫 调整body的背景色
    // beforeRouteEnter(to, from, next) {
    //   if (process.client) {
    //     // window.document.body.style.backgroundColor = '#fff';
    //   }
    //   next();
    // },
    // beforeRouteLeave(to, from, next) {
    //   if (process.client) {
    //     // window.document.body.style.backgroundColor = '#f5f5f5';
    //   }
    //   next();
    // }
  }
</script>

<style>
  .me-view-body {
    margin: 100px auto 70px;
    background-color: #fff;
  }

  .me-view-container {
    width: 700px;
  }

  .el-main {
    overflow: hidden;
  }

  .me-view-title {
    font-size: 34px;
    font-weight: 700;
    line-height: 1.3;
  }

  .me-view-author {
    /*margin: 30px 0;*/
    margin-top: 30px;
    vertical-align: middle;
  }

  .me-view-picture {
    width: 40px;
    height: 40px;
    border: 1px solid #ddd;
    border-radius: 50%;
    vertical-align: middle;
    background-color: #010101;
  }

  .me-view-info {
    display: inline-block;
    vertical-align: middle;
    margin-left: 8px;
  }

  .me-view-meta {
    font-size: 12px;
    color: #969696;
  }

  .me-view-end {
    margin-top: 20px;
    padding-right: 25px;
  }

  .me-view-tag {
    margin-top: 20px;
    padding-left: 6px;
    border-left: 4px solid #c5cac3;
  }

  .me-view-tag-item {
    margin: 0 4px;
  }

  .me-view-comment {
    margin-top: 60px;
  }

  .me-view-comment-title {
    font-weight: 600;
    border-bottom: 1px solid #f0f0f0;
    padding-bottom: 20px;
  }

  .me-view-comment-write {
    margin-top: 20px;
  }

  .me-view-comment-text {
    font-size: 16px;
  }

  .v-show-content {
    padding: 8px 25px 15px 0px !important;
  }

  .v-note-wrapper .v-note-panel {
    box-shadow: none !important;
  }

  .v-note-wrapper .v-note-panel .v-note-show .v-show-content, .v-note-wrapper .v-note-panel .v-note-show .v-show-content-html {
    background: #fff !important;
  }


</style>
