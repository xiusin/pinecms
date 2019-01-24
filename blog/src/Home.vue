<template>
  <div id="home">
    <el-container direction="vertical">

      <base-header :activeIndex="activeIndex"></base-header>

      <router-view class="me-container"/>

      <base-footer v-show="footerShow"></base-footer>

    </el-container>

  </div>

</template>

<script>
  import BaseFooter from '@/components/BaseFooter'
  import BaseHeader from '@/views/BaseHeader'

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
      'base-footer': BaseFooter
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
  .me-container {
    margin: 90px auto 25px;
    width: 1160px;
  }
</style>
