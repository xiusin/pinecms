<template>
  <div id="home">
    <el-container direction="vertical">

      <base-header></base-header>
      <!-- <router-view class="me-container"/> -->
      <nuxt class="me-container"/>
      <base-footer></base-footer>

    </el-container>

  </div>

</template>

<script>
  import BaseFooter from '~/components/BaseFooter'
  import BaseHeader from '~/pages/BaseHeader'
  import GoTop from '~/components/gotop/GoTop'

  export default {
    name: 'Home',
    data() {
      return {
        activeIndex: '/',
        footerShow: true
      }
    },
    components: {
      'base-header': BaseHeader,
      'base-footer': BaseFooter,
      'go-top': GoTop
    },
    beforeRouteEnter(to, from, next) {
      next(vm => {
        vm.activeIndex = to.path
      })
    },
    beforeRouteUpdate(to, from, next) {
      this.footerShow = true
      this.activeIndex = to.path
      next()
    }
  }
</script>

<style>
 
  * {
    margin: 0;
    padding: 0;
  }

  body {
    background-color: #f5f5f5;
    font-weight: 400;
    font-family: wu,"Open Sans",Arial,"Hiragino Sans GB","Microsoft YaHei","STHeiti","WenQuanYi Micro Hei",SimSun,sans-serif;
    line-height: 1.5;
  }

  a {
    cursor: pointer;
    text-decoration: none;
    transition: none 86ms ease-out;
  }

  a:hover {
    color: #010101;
  }

  .me-area {
    background-color: #fff;
  }

  html {
    height: 100%;
  }

  body {
    min-height: 100%;
  }

  .me-pull-right {
    float: right;
  }

   .me-container {
    margin: 90px auto 25px;
    width: 1160px;
  }
</style>
