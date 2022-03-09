<template >
  <div style="height: 100%">
    <el-row style="height: 100%">
      <el-col :span="12" style="height: 100%">
        <div class="login">
          <div class="content">
            <div class="login-body">
              <div class="top">
                <p class="title">账号密码登录</p>
                <router-link to="/register">
                  立即注册
                  <i class="el-icon-arrow-right"></i>
                </router-link>
              </div>
              <el-form
                  status-icon
                  label-position="top"
                  label-width="80px"
                  :rules="rules"
                  :model="formLabelAlign"
                  ref="ruleForm"
                  @submit.native.prevent
              >
                <el-form-item label="账号" prop="account">
                  <el-input placeholder="请输入账号" clearable v-model="formLabelAlign.account"></el-input>
                </el-form-item>
                <el-form-item label="密码" prop="pwd">
                  <el-input placeholder="请输入密码" show-password clearable v-model="formLabelAlign.pwd"></el-input>
                </el-form-item>
                <el-form-item>
                  <div class="check">
                    <el-checkbox v-model="remember">记住密码</el-checkbox>
                    <el-checkbox v-model="autoLogin">自动登录</el-checkbox>
                  </div>
                </el-form-item>
                <el-form-item>
                  <el-button type="primary" @click="submitForm()">登录</el-button>
                  <router-link to>忘记密码?</router-link>
                </el-form-item>
              </el-form>
              <div class="footer"></div>
            </div>
          </div>
        </div>
      </el-col>
      <el-col :span="12" style="height: 100%">
        <div class="login-bg">
<!--          <monaco-with-tree-->
<!--              :files="files"-->
<!--              :default-open-files="defaultOpenFiles"-->
<!--              :get-file-content="getFileContent"-->
<!--          />-->
        </div>
      </el-col>
    </el-row>
  </div>
</template>
<script>
import POST from "@/http/post";
export default {
  name: "Login",
  components: {
    // AmEditor
    // MonacoWithTree
  },
  data() {
    const reg = new RegExp('[\\\\/:*?"<>|]');
    const validator = (rule, value, callback) => {
      if (!value) {
        callback(new Error("请输入内容"));
      } else if (reg.test(value)) {
        callback(new Error('含有非法字符\\/:*?"<>|'));
      } else {
        callback();
      }
    };
    return {
      files: ['package.json',  'README.md', 'index.js', 'src/test.js', 'src/index.js', 'public/index.html'],
      defaultOpenFiles: ['package.json'],
      formLabelAlign: {
        account: "",
        pwd: ""
      },
      rules: {
        account: [{ validator: validator, trigger: "blur" }],
        pwd: [{ validator: validator, trigger: "blur" }]
      },
      remember: false,
      autoLogin: false
    };
  },
  mounted() {
    this.getAccount();
  },
  methods: {
    getFileContent(filePath) {
      return [`${filePath}-left`, `${filePath}-right`];
    },
    submitForm() {
      this.$refs.ruleForm.validate(valid => {
        if (valid) {
          POST.login(this.formLabelAlign.account, this.formLabelAlign.pwd).then(
            resp => {
              if (resp.data.result.status == "success") {
                this.$store.commit(`fm/setAutoLogin`, this.autoLogin);
                this.$store.commit("fm/setLoginStatus", resp.data.isLogin);
                this.$store.commit("fm/setUserName", {
                  name: this.formLabelAlign.account,
                  nickname: resp.data.nickname
                });
                if (this.remember && window.localStorage) {
                  window.localStorage.setItem(
                    "auto",
                    JSON.stringify({
                      a: this.$store.state.fm.autoLogin
                    })
                  );
                  this.saveAccount(
                    this.formLabelAlign.account,
                    this.formLabelAlign.pwd
                  );
                }
                this.$router.push({ path: "/file" });
              }
            }
          );
        } else {
          this.$message.error("请输入合理内容!");
          return false;
        }
      });
    },
    saveAccount(t = "", b = "") {
      window.localStorage.setItem("hallelujah", JSON.stringify({ t, b }));
    },
    getAccount() {
      console.log(
        "this.$store.state.fm.autoLogin",
        this.$store.state.fm.autoLogin
      );
      this.autoLogin = this.$store.state.fm.autoLogin;
      if (window.localStorage && window.localStorage.getItem("hallelujah")) {
        this.remember = true;
        let user = JSON.parse(window.localStorage.getItem("hallelujah"));
        this.formLabelAlign.account = user.t;
        this.formLabelAlign.pwd = user.b;
      }
    }
  },
  watch: {
    remember(newV) {
      if (!newV && window.localStorage) {
        window.localStorage.removeItem("hallelujah");
      }
    }
  }
};
</script>
<style lang="scss" scoped>
.login-bg {
  height: 100%;
  width: 100%;
  background-image: url("../assets/bg.jpg");
  background-repeat: no-repeat;
  background-size: cover;
}
.login {
  .header {
    padding-top: 5px;
    padding-left: 10px;
    img {
      width: 60px;
      height: 60px;
    }
    span {
      display: inline-block;
      font-size: 1.8rem;
      font-family: cursive;
      color: papayawhip;
      margin-left: 7px;
      line-height: 2;
      vertical-align: text-bottom;
    }
  }
  .content {
    position: relative;
    width: 100%;
    height: 100%;
    .login-body {
      background: rgba(255, 255, 255, 0.2);
      width: 352px;
      height: 405px;
      margin-top: calc(50% - 190px);
      margin-left: calc(50% - 176px);
      box-shadow: rgba(0, 0, 0, 0.3) 0 0 50px;
      padding: 0 20px 20px 20px;
      .top {
        display: flex;
        align-items: center;
        justify-content: space-between;
        .title {
          font-weight: 600;
          font-size: 1.2rem;
        }
        a {
          text-decoration: none;
          margin: 0 4px;
          color: #fff;
          &:hover {
            color: #09aaff;
          }
        }
      }

      ::v-deep .el-form {
        .el-form-item {
          margin-bottom: 10px;
          .el-form-item__label {
            padding: 0;
          }
          .el-form-item__content {
            button {
              width: 100%;
            }
            .check {
              display: flex;
              justify-content: space-between;
              align-content: center;
              .el-checkbox {
                color: #cdc6c6;
              }
            }
          }
        }
        .el-form-item:last-child {
          margin-top: 20px;
          a {
            color: #cdc6c6;
            text-decoration: none;
            font-size: 14px;
            position: relative;
            top: 14px;
            margin: 0;
            &:hover {
              color: #09aaff;
            }
          }
        }
      }
    }
  }
}
</style>
