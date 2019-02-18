<template>
  <div >
    <el-container>
      <el-aside style="width: 150px; float: left; height: 690px;">
        <el-menu :router="true" style="height: 100%;border:none;">

          <el-menu-item index="/user/center">
            <i class="icon iconfont icon-geren"></i>
            <span slot="title">个人信息</span>
          </el-menu-item>

          <el-menu-item index="/user/task">
            <i class="icon iconfont icon-renwu"></i>
            <span slot="title">每日任务</span>
          </el-menu-item>

          <el-menu-item index="/user/money">
            <i class="icon iconfont icon-qian"></i>
            <span slot="title">我的资产</span>
          </el-menu-item>

          <el-menu-item index="/user/cart">
            <i class="icon iconfont icon-dixiaofei"></i>
            <span slot="title">消费清单</span>
          </el-menu-item>

          <el-menu-item index="/user/recharge">
            <i class="icon iconfont icon-lishijilu"></i>
            <span slot="title">充值记录</span>
          </el-menu-item>

          <el-menu-item index="/user/ref">
            <i class="icon iconfont icon-daochu"></i>
            <span slot="title">线上推广</span>
          </el-menu-item>

          <el-menu-item index="/user/outmo">
            <i class="icon iconfont icon-qian"></i>
            <span slot="title">站内提现</span>
          </el-menu-item>
          <el-menu-item index="/user/modipwd">
            <i class="icon iconfont icon-jiesuo"></i>
            <span slot="title">修改密码</span>
          </el-menu-item>
          <!--<el-menu-item index="4">-->
            <!--<i class="icon iconfont icon-weixin"></i>-->
            <!--<span slot="title">快捷绑定</span>-->
          <!--</el-menu-item>-->

          <el-menu-item index="/user/publish">
            <i class="icon iconfont icon-dianping"></i>
            <span slot="title">发布资源</span>
          </el-menu-item>

        </el-menu>
      </el-aside>

      <el-main style="width: 800px; float: right;background-color: #fff; margin-left: 20px; padding: 30px;">
        <el-form :model="ruleForm" :rules="rules" ref="ruleForm" label-width="100px" class="demo-ruleForm">
          <el-form-item label="登录名" prop="account" size="mini" style="width: 450px;">
            <el-input v-model="ruleForm.name" disabled></el-input>
          </el-form-item>

          <el-form-item label="邮箱" prop="email" size="mini" style="width: 450px;">
            <el-input v-model="ruleForm.name"></el-input>
          </el-form-item>

          <el-form-item label="手机" prop="telphone" size="mini" style="width: 450px;">
            <el-input v-model="ruleForm.name"></el-input>
          </el-form-item>

          <el-form-item label="头像" prop="account">
            <el-upload
              action="https://jsonplaceholder.typicode.com/posts/"
              list-type="picture-card"
              :on-preview="handlePictureCardPreview"
              :on-remove="handleRemove">
              <i class="el-icon-plus"></i>
            </el-upload>
            <el-dialog :visible.sync="dialogVisible">
              <img width="100%" :src="dialogImageUrl" alt="">
            </el-dialog>
          </el-form-item>

          <el-form-item label="注册时间" prop="reg_time" size="mini" style="width: 450px;">
            <el-input v-model="ruleForm.name"></el-input>
          </el-form-item>

          <el-form-item label="网址" prop="url" size="mini" style="width: 450px;">
            <el-input v-model="ruleForm.name" placeholder="github或个人网站"></el-input>
          </el-form-item>

          <el-form-item label="QQ" prop="qq" size="mini" style="width: 450px;">
            <el-input v-model="ruleForm.name"></el-input>
          </el-form-item>

          <el-form-item label="微信" prop="wechat" size="mini" style="width: 450px;">
            <el-input v-model="ruleForm.name"></el-input>
          </el-form-item>

          <el-form-item>
            <el-button type="primary" size="mini" @click="submitForm('ruleForm')">修改</el-button>
          </el-form-item>
        </el-form>
      </el-main>

    </el-container>
  </div>
</template>

<script>
  export default {
    ayncData () {

    },
    data() {
      return {
        dialogImageUrl: '',
        dialogVisible: false,
        ruleForm: {
          name: '',
          region: '',
          date1: '',
          date2: '',
          delivery: false,
          type: [],
          resource: '',
          desc: ''
        },
        rules: {
          name: [
            { required: true, message: '请输入活动名称', trigger: 'blur' },
            { min: 3, max: 5, message: '长度在 3 到 5 个字符', trigger: 'blur' }
          ],
          region: [
            { required: true, message: '请选择活动区域', trigger: 'change' }
          ],
          date1: [
            { type: 'date', required: true, message: '请选择日期', trigger: 'change' }
          ],
          date2: [
            { type: 'date', required: true, message: '请选择时间', trigger: 'change' }
          ],
          type: [
            { type: 'array', required: true, message: '请至少选择一个活动性质', trigger: 'change' }
          ],
          resource: [
            { required: true, message: '请选择活动资源', trigger: 'change' }
          ],
          desc: [
            { required: true, message: '请填写活动形式', trigger: 'blur' }
          ]
        }
      };
    },
    methods: {
      submitForm(formName) {
        this.$refs[formName].validate((valid) => {
          if (valid) {
            alert('submit!');
          } else {
            console.log('error submit!!');
            return false;
          }
        });
      },
      resetForm(formName) {
        this.$refs[formName].resetFields();
      },
      handleRemove(file, fileList) {
      },
      handlePictureCardPreview(file) {
        this.dialogImageUrl = file.url;
        this.dialogVisible = true;
      }
    }
  }
</script>

