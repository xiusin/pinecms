<template>
  <div id="verifyUser" v-title :data-title="getTitle()">

    <div class="me-login-box me-login-box-radius">
      <h1>{{title}} 验证邮箱</h1>

      <div style="text-align: center">
        <i class="icon iconfont icon-duihao" style="font-size: 220px; color: #0e931c" v-if="verify_result"></i>
        <i class="icon iconfont icon-cuohao" style="font-size: 220px; color: #ff0000" v-else></i>
      </div>

    </div>
  </div>
</template>

<script>
  export default {
    name: 'VerifyUser',
    data() {
      return {
        verify_result: false,
        title: '',
      }
    },
    mounted:function () {
    },
    methods: {
      getTitle(){
        this.title = window.title
        return '验证邮箱 - ' + window.title + ' - ' + window.keywords + ' - ' + window.description
      }
    }
  }
</script>

<style scoped>
  .me-login-box {
    position: absolute;
    width: 300px;
    height: 320px;
    background-color: white;
    margin-top: 150px;
    margin-left: -180px;
    left: 50%;
    padding: 30px;
  }

  .me-login-box-radius {
    border-radius: 10px;
    box-shadow: 0px 0px 1px 1px rgba(161, 159, 159, 0.1);
  }

  .me-login-box h1 {
    text-align: center;
    font-size: 24px;
    margin-bottom: 20px;
    vertical-align: middle;
  }

  .me-login-design {
    text-align: center;
    font-family: 'Open Sans', sans-serif;
    font-size: 18px;
  }

  .me-login-design-color {
    color: #5FB878 !important;
  }

  .me-login-button {
    text-align: center;
  }

  .me-login-button button {
    width: 100%;
  }

</style>
