<template >
  <div class="register">
    <div class="body">
      <div class="info">
        <h3>用科技</h3>
        <p>让复杂的世界更简单</p>
      </div>
      <div class="content">
        <div class="top">
          <h3>欢迎注册</h3>
          <p>
            已有账号?
            <router-link to="/login">登录</router-link>
          </p>
        </div>
        <el-form
          status-icon
          label-position="left"
          label-width="90px"
          :rules="rules"
          :model="formLabelAlign"
          ref="ruleForm"
          @submit.native.prevent
        >
          <el-form-item label="昵称" prop="nickName">
            <el-input placeholder="用于登录显示(可不填)" clearable v-model="formLabelAlign.nickName"></el-input>
          </el-form-item>
          <el-form-item label="账号" prop="account">
            <el-input placeholder="用于登录和找回密码" clearable v-model="formLabelAlign.account"></el-input>
          </el-form-item>
          <el-form-item label="密码" prop="pwd">
            <el-input
              placeholder="请输入登录密码"
              autocomplete="off"
              show-password
              clearable
              v-model="formLabelAlign.pwd"
            ></el-input>
          </el-form-item>
          <el-form-item label="确认密码" prop="againPwd">
            <el-input
              placeholder="请再次输入登录密码"
              autocomplete="off"
              show-password
              clearable
              v-model="formLabelAlign.againPwd"
            ></el-input>
          </el-form-item>
          <el-form-item></el-form-item>
        </el-form>
        <el-button type="primary" @click.native="submitForm()">注册</el-button>
      </div>
    </div>
  </div>
</template>
<script>
import GET from "@/http/get";
import POST from "@/http/post";
export default {
  name: "Register",
  data() {
    const reg = new RegExp('[\\\\/:*?"<>|]');
    const legitimacy = (rule, value, callback) => {
      if (reg.test(value)) {
        callback(new Error('含有非法字符\\/:*?"<>|'));
      } else {
        callback();
      }
    };
    const validator = (rule, value, callback) => {
      if (!value) {
        callback(new Error("请输入内容"));
      } else if (reg.test(value)) {
        callback(new Error('含有非法字符\\/:*?"<>|'));
      } else {
        callback();
      }
    };
    const validateName = (rule, value, callback) => {
      const reg = new RegExp(/^[a-z][a-z0-9]{4,15}$/, "g");
      if (!reg.test(value)) {
        callback(new Error("账号不合法:只允许小写字母、数字,长度5-16"));
      }
      GET.findUsername(value).then(resp => {
        if (resp.status == 200 && resp.data.isExist) {
          callback(new Error("账号已存在!"));
          return;
        } else {
          callback();
        }
      });
    };
    const validatePass = (rule, value, callback) => {
      const reg = new RegExp(/^(?=.*\d)(?=.*[a-z])(?=.*[A-Z]).{8,15}$/, "g");
      if (!reg.test(value)) {
        callback(new Error("必须包含大小写字母和数字的组合，不能使用特殊字符"));
      } else {
        if (this.formLabelAlign.againPwd !== "") {
          this.$refs.ruleForm.validateField("againPwd");
        }
        callback();
      }
    };
    const validatePass2 = (rule, value, callback) => {
      if (value !== this.formLabelAlign.pwd) {
        callback(new Error("两次输入密码不一致!"));
      } else {
        callback();
      }
    };
    return {
      formLabelAlign: {
        nickName: "",
        account: "",
        pwd: "",
        againPwd: ""
      },
      rules: {
        nickName: [{ validator: legitimacy, trigger: "blur" }],
        account: [
          { validator: validator, trigger: "blur" },
          {
            validator: validateName,
            trigger: "blur"
          }
        ],
        pwd: [
          { validator: validator, trigger: "blur" },
          {
            min: 8,
            max: 10,
            message: "长度在 8 到 10 个字符",
            trigger: "blur"
          },
          { validator: validatePass, trigger: "blur" }
        ],
        againPwd: [
          { validator: validator, trigger: "blur" },
          { validator: validatePass2, trigger: "blur" }
        ]
      }
    };
  },
  methods: {
    submitForm() {
      this.$refs.ruleForm.validate(valid => {
        if (valid) {
          // 去注册
          POST.register(this.formLabelAlign).then(resp => {
            if (resp.data.result.status == "success") {
              this.$router.push({ path: "login" });
              return;
            }
            this.$refs.ruleForm.resetFields();
          });
        } else {
          this.$notify.error({
            title: "错误",
            message: "请填写正确信息!"
          });
        }
      });
    }
  }
};
</script>
<style lang="scss" scoped>
.register {
  font-family: Tahoma, Helvetica, "Microsoft Yahei", "微软雅黑", Arial, STHeiti;
  position: relative;
  height: 100%;
  width: 100%;
  // background-image: url("../assets/register.jpg");
  background-image: url("https://525-save.oss-cn-zhangjiakou.aliyuncs.com/%E6%96%87%E6%A1%A3/register.jpg");
  background-repeat: no-repeat;
  background-size: cover;
  display: flex;
  align-items: center;
  .body {
    width: 100%;
    height: 540px;
    display: flex;
    align-items: center;
    justify-content: space-around;
    .info {
      color: #fff;
      h3 {
        font-size: 3.375rem;
        letter-spacing: 0;
        font-weight: 700;
      }
      p {
        font-size: 32px;
        letter-spacing: 3.81px;
        font-weight: 300;
      }
    }
    .content {
      width: 480px;
      height: 540px;
      border-radius: 10px;
      background: rgba(255, 255, 255, 0.9);
      box-shadow: rgba(0, 0, 0, 0.3) 0 0 50px;
      padding: 30px 39px 39px 39px;
      box-sizing: border-box;
      .top {
        h3 {
          font-size: 2.25rem;
          color: #000;
          padding-bottom: 4px;
        }
        p {
          font-size: 14px;
          color: #9b9b9b;
          a {
            color: #2e58ff;
            text-decoration: none;
          }
        }
      }
      ::v-deep .el-form {
        margin-top: 40px;
        .el-form-item {
          .el-form-item__label {
            padding: 0;
          }
          .el-form-item__content {
            margin-left: 50px;
          }
        }
      }
    }
  }
}
h3,
p {
  margin: 0;
  padding: 0;
}
::v-deep .el-button {
  width: 100%;
  height: 50px;

  border-radius: 50px;
}
</style>