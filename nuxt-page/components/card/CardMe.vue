<template>
  <div>
    <el-card>
      <h1 class="me-author-name">站长微信</h1>
      <div class="me-author-description">
        <img src="~/static/qrcode/ideamecn.jpg" style="width: 100%" />
      </div>
      <div class="me-author-tool">
        <i @click="showTool(qq)" :title="qq.title" class="me-icon-QQ"></i>
        <i @click="showTool(gitee)" :title="gitee.title" class="me-icon-github"></i>
      </div>
    </el-card>
    <div style="height: 10px;"></div>
    <el-card>
      <h1 class="me-author-name">相关公众号</h1>
      <div class="me-author-description">
        <img src="~/static/qrcode/wechat.jpg" style="width: 100%" />
      </div>
    </el-card>
  </div>
</template>

<script>
  export default {
    name: 'CardMe',
    data() {
      return {
        qq: {title: 'QQ', message: '826466266'},
        gitee: {
          title: 'gitee',
          message: '<a target="_blank" href="https://gitee.com/xiusin">https://gitee.com/xiusin</a>'
        }
      }
    },
    methods: {
      showTool(tool) {
        this.$message({
          duration: 0,
          showClose: true,
          dangerouslyUseHTMLString: true,
          message: '<strong>' + tool.message + '</strong>'
        });
      }
    }
  }
</script>

<style scoped>
  .me-author-name {
    text-align: center;
    font-size: 30px;
    border-bottom: 1px solid #010101;
  }

  .me-author-description {
    padding: 8px 0;
  }

  .me-icon-job {
    padding-left: 16px;
  }

  .me-author-tool {
    text-align: center;
    padding-top: 10px;
  }

  .me-author-tool i {
    cursor: pointer;
    padding: 4px 10px;
    font-size: 30px;
  }
</style>
