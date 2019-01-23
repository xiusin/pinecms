<template>
  <div >
    <el-container>
      <el-aside style="width: 150px; float: left; height: 690px;">
        <el-menu :router="true" style="height: 100%;border:none;">
          <el-menu-item index="/user/center">
            <i class="icon iconfont icon-geren"></i>
            <span slot="title">个人信息</span>
          </el-menu-item>
          <el-menu-item index="/user/money">
            <i class="icon iconfont icon-qian"></i>
            <span slot="title">我的资产</span>
          </el-menu-item>
          <el-menu-item index="/user/cart">
            <i class="icon iconfont icon-renwu"></i>
            <span slot="title">消费清单</span>
          </el-menu-item>
          <el-menu-item index="/user/recharge">
            <i class="icon iconfont icon-xiaofeimingxidan"></i>
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
          <el-form-item label="原密码" prop="account" size="mini" style="width: 450px;">
            <el-input v-model="ruleForm.name" disabled></el-input>
          </el-form-item>

          <el-form-item label="新密码" prop="email" size="mini" style="width: 450px;">
            <el-input v-model="ruleForm.name"></el-input>
          </el-form-item>

          <el-form-item label="确认新密码" prop="telphone" size="mini" style="width: 450px;">
            <el-input v-model="ruleForm.name"></el-input>
          </el-form-item>

          <el-form-item>
            <el-button type="primary" @click="submitForm('ruleForm')">立即创建</el-button>
            <el-button @click="resetForm('ruleForm')">重置</el-button>
          </el-form-item>
        </el-form>
      </el-main>

    </el-container>
  </div>
</template>

<script>
  export default {
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
        console.log(file, fileList);
      },
      handlePictureCardPreview(file) {
        this.dialogImageUrl = file.url;
        this.dialogVisible = true;
      }
    }
  }
</script>

