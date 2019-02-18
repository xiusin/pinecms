<template>
  <el-header class="me-area">
    <el-row class="me-header" :gutter="20" style="width: 1160px;margin: 0 auto;">

      <el-col :span="2" class="me-header-left">
        <router-link to="/" class="me-title">
          <img src="../assets/img/logo.png"/>
        </router-link>
      </el-col>

      <el-col v-if="!simple" :span="14" :offset="0">
        <el-menu :router=true menu-trigger="click" active-text-color="#010101" :default-active="activeIndex"
                 mode="horizontal">
          <el-menu-item index="/">首页</el-menu-item>
          <el-menu-item index="/free/video">免费视频</el-menu-item>
          <el-menu-item index="/free/book">免费图书</el-menu-item>
          <el-menu-item index="/paid/video">付费视频</el-menu-item>
          <el-menu-item index="/paid/book">付费图书</el-menu-item>
          <el-menu-item index="/soft">软件下载</el-menu-item>
          <el-menu-item index="/about">重要说明</el-menu-item>

        </el-menu>
      </el-col>

      <template v-else>
        <slot></slot>
      </template>

      <el-col :span="4" :offset="1" style="float: right">
        <el-menu :router=true menu-trigger="click" mode="horizontal" active-text-color="#010101" style="border: none">
          <template v-if="!user.login">
             <el-menu-item index="/login" style="float:right;">
              <el-button type="text">登录</el-button>
            </el-menu-item>

            <el-menu-item index="/register" style="float:right;">
              <el-button type="text">注册</el-button>
            </el-menu-item>
          </template>

          <template v-else>
            <el-submenu index>
              <template slot="title">
                <img class="me-header-picture" @click="goto('/user/center')" :src="user.avatar" onerror="this.src='../../static/default_avatar.png'"/>
              </template>
              <el-menu-item index @click="logout"><i class="el-icon-back"></i>退出</el-menu-item>
            </el-submenu>
          </template>
        </el-menu>
      </el-col>

    </el-row>
  </el-header>
</template>

<script>
  export default {
    name: 'BaseHeader',
    props: {
      activeIndex: String,
      simple: {
        type: Boolean,
        default: false
      }
    },
    data() {
      return {}
    },
    computed: {
      user() {
        let login = this.$store.state.token
        let avatar = this.$store.state.avatar
        return {
          login, avatar
        }
      }
    },
    methods: {
      logout() {
        let that = this
        this.$store.dispatch('logout').then(() => {
          this.$router.push({path: '/'})
        }).catch((error) => {
          if (error !== 'error') {
            that.$message({message: error, type: 'error', showClose: true});
          }
        })
      },
      goto(str) {
        this.$router.push({path: str})
      }
    }
  }
</script>

<style>

  .el-header {
    position: fixed;
    z-index: 1024;
    min-width: 100%;
    box-shadow: 0 2px 3px hsla(0, 0%, 7%, .1), 0 0 0 1px hsla(0, 0%, 7%, .1);
  }

  .me-title {
    margin-top: 10px;
    font-size: 24px;
  }

  .me-header-left {
    margin-top: 10px;
  }

  .me-title img {
    max-height: 2.4rem;
    max-width: 100%;
  }

  .me-header-picture {
    width: 36px;
    height: 36px;
    border: 1px solid #ddd;
    border-radius: 50%;
    vertical-align: middle;
    background-color: #010101;
  }
</style>
