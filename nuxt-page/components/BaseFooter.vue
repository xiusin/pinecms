<template>
  <el-footer class="me-area" height="115px">
    <div class="me-footer">
        <div class="container">
          <div class="nav-footer">
            <a href="https://www.waitsun.com/aboutus">联系我们</a>
            <a href="https://www.waitsun.com/aboutus/copyright">版权声明</a>
            <a href="http://help.waitsun.com">网站帮助</a>
            <a href="https://www.waitsun.com/tags-page">热门标签</a>
            <a href="https://www.waitsun.com/xpay-html">赞助我们</a>
          </div>
          <div class="copyright-footer">
            <p>Copyright © 2010 - 2019 <a class="site-link" href="https://www.waitsun.com" title="爱情守望者" rel="home">爱情守望者</a> · WU Theme Source <a href="http://www.suxing.me" target="_blank">Suxing</a> </p> 
          </div>
          <div class="links-footer">
            <span>友情链接：</span>
            <a href="https://www.waitsun.com" rel="me" target="_blank">Mac软件分享</a>
            <a href="https://www.paurl.com" target="_blank">互联网运营培训</a>
            <a href="http://www.sqlsec.com" title="白帽子国光的信息安全学习记录博客" target="_blank">国光的博客</a>
          </div>
      </div>
    </div>
  </el-footer>
</template>

<script>
  export default {
    name: 'BaseFooter',
    data() {
      return {
        title: ''
      }
    }
  }
</script>

<style>

  .me-footer {
    width: 1160px;
    height: 95px;
    overflow: auto;
    margin: 10px auto;
    font-size: 16px;
    line-height: 30px;
  }

  .copyright-footer {
    font-size:14px;
    color: #696969;
  }

  .copyright-footer a {
    color: #929292;
  }

  .nav-footer {
      font-size:16px;
  }
  .nav-footer a{ 
     color:#34495e;
  }

  .links-footer {
      font-size: 12px;
      color: #696969;
      border-top: 1px solid rgba(255,255,255,.05);
  }

   .links-footer a {
     color: #696969;
   }

  .me-login-design-color {
    color: #010101 !important;
  }

</style>
