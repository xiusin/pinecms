<template>
  <div id="register" v-title :data-title="getTitle()">
    <!--<video preload="auto" class="me-video-player" autoplay="autoplay" loop="loop">
          <source src="../../static/vedio/sea.mp4" type="video/mp4">
      </video>-->

    <div class="me-login-box me-login-box-radius">
      <h1>{{title}} 注册</h1>

      <el-form ref="userForm" :model="userForm" :rules="rules">
        <el-form-item prop="account">
          <el-input placeholder="用户名" v-model="userForm.account"></el-input>
        </el-form-item>

        <el-form-item prop="email">
          <el-input placeholder=" 邮箱" v-model="userForm.email"></el-input>
        </el-form-item>


        <el-form-item>
          <div class="vaptcha-container" id="vaptcha-reg-container" data-config='{"vid": "5bbc46c6fc650e3be06e5869","type": "click"}' style="width:300px;height:36px;">
            <div class="vaptcha-init-main">
              <div class="vaptcha-init-loading">
                <a href="https://vaptcha.com" target="_blank">
                  <img src="https://cdn.vaptcha.com/vaptcha-loading.gif" />
                </a>
                <span class="vaptcha-text"> VAPTCHA启动中...</span>
              </div>
            </div>
          </div>
        </el-form-item>

        <el-form-item prop="password">
          <el-input placeholder="密码" v-model="userForm.password"></el-input>
        </el-form-item>

        <el-form-item size="small" class="me-login-button">
          <el-button type="primary" @click.native.prevent="register('userForm')">注册</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
  

  import {register} from '@/api/login'
 
  export default {
    head: {
	    script: [
	      { src: 'https://cdn.vaptcha.com/v2.js' }
	    ]
	  },
    name: 'Register',
    data() {
      return {
        title: '',
        captcha: null,
        userForm: {
          account: '',
          email: '',
          password: '',
          token: ''
        },
        rules: {
          account: [
            {required: true, message: '请输入用户名', trigger: 'blur'},
            {max: 16, message: '不能大于16个字符', trigger: 'blur'}
          ],
          email: [
            {required: true, message: '请输入邮箱', trigger: 'blur'},
            {max: 20, message: '不能大于20个字符', trigger: 'blur'}
          ],
          password: [
            {required: true, message: '请输入密码', trigger: 'blur'},
            {max: 16, message: '不能大于16个字符', trigger: 'blur'}
          ]
        }
      }
    },
    mounted:function () {
      this.createVcaptcha()
    },
    methods: {
      getTitle(){
          return '注册账户 ' 
      },
      createVcaptcha() {
        let that = this
        if (process.client) {
          window.vaptcha({
            //配置参数
            vid: '5bbc46c6fc650e3be06e5869', // 验证单元id
            type: 'click', // 展现类型 点击式
            container: window.document.getElementById("vaptcha-reg-container") // 按钮容器，可为Element 或者 selector
          }).then((vaptchaObj) => {
            that.captcha = vaptchaObj
            vaptchaObj.listen('pass', function() {
              that.userForm.token = vaptchaObj.getToken()
            })
            vaptchaObj.render()// 调用验证实例 vaptchaObj 的 render 方法加载验证按钮
          })
        }
      },
      register(formName) {
        let that = this
        this.$refs[formName].validate((valid) => {
          if (valid) {
            that.$store.dispatch('register', that.userForm).then((data) => {
              that.$message({message: '注册成功,请尽快通过邮箱验证账号', type: 'success', showClose: true,onClose: function () {
                  that.$router.push({path: '/'})
                }});
            }).catch((error) => {
              if (error !== 'error') {
                that.$message({message: error, type: 'error', showClose: true});
              }
            })

          } else {
            return false;
          }
        });

      }

    }
  }
</script>

<style scoped>
  .me-container {
    margin: 90px auto 25px;
    width: 1160px;
    height: 680px;
  }
  #login {
    min-width: 100%;
    min-height: 100%;
  }
  .el-footer {
      padding: 0 20px;
      box-sizing: border-box;
      position: fixed;
      bottom: 0px;
  }
  .me-video-player {
    background-color: transparent;
    width: 100%;
    height: 100%;
    object-fit: fill;
    display: block;
    position: absolute;
    left: 0;
    z-index: 0;
    top: 0;
  }

  .me-login-box {
    position: absolute;
    width: 300px;
    height: 300px;
    background-color: white;
    margin-top: 150px;
    margin-left: -180px;
    left: 50%;
    padding: 30px;
  }

  .me-login-box-radius {
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
    color: #010101 !important;
  }

  .me-login-button {
    text-align: center;
  }

  .me-login-button button {
    width: 100%;
  }

</style>
