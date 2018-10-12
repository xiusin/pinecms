<template>
  <div id="verifyUser" v-title data-title="验证邮箱 - For Fun">

    <div class="me-login-box me-login-box-radius">
      <h1>ForFun 验证邮箱</h1>

      <div class="me-login-design">
        <p>Designed by
          <strong>
            <router-link to="/" class="me-login-design-color">ForFun</router-link>
          </strong>
        </p>
      </div>

    </div>
  </div>
</template>

<script>
  export default {
    name: 'VerifyUser',
    data() {
      return {
      }
    },
    mounted:function () {
    },
    methods: {
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
