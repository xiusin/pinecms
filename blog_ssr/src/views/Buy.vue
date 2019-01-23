<template>
  <div id="pay" v-title :data-title="getTitle()">

    <div class="me-login-box me-login-box-radius">
      <h1>{{data.title}} : 订单号{{data.order_id}}</h1>

      <div class="me-login-design">
      <iframe :src="data.url" style="border: none; width: 100%; height: 600px"/>
      </div>

    </div>
  </div>
</template>

<script>
  import {goPay} from "../api/article";

  export default {
    name: 'Buy',
    data() {
      return {
        data: {},
      }
    },
    created: function () {
      this.gopay()
    },
    methods: {
      getTitle() {
        return '验证邮箱 - ' + window.title + ' - ' + window.keywords + ' - ' + window.description
      },
      gopay() {
        goPay({
          "id": this.$route.params.id,
          "paytype": this.$route.params.paytype,
        }).then((data) => {
          this.data = data.data;
        })
      }
    }
  }
</script>

<style scoped>
  .me-login-box {
    width: 800px;
    height: 650px;
    background-color: white;
    margin: 0 auto;
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
