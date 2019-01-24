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
        <template>
          <el-tabs type="card" style="height: 150px;">
            <el-tab-pane label="我的资产" style="font-size: 14px;">
              消费友币：10
              <br/>
              剩余友币：0
            </el-tab-pane>
          </el-tabs>

          <el-tabs type="card" style="height: 300px;">
            <el-tab-pane label="在线充值">
              <el-form :model="payPointForm" ref="payPointForm" label-width="100px">
                <el-form-item label="充值友币" prop="money" >
                    <el-input placeholder="1人民币=10友币">
                      <template slot="prepend">￥</template>
                    </el-input>
                </el-form-item>

                <el-form-item label="支付方式" prop="money" >
                   <el-radio-group v-model="payType">
                     <el-radio-button label="weixin"><i class="icon iconfont icon-weixinzhifu2"></i> &nbsp;&nbsp;微信支付</el-radio-button>
                     <el-radio-button label="zhifubao"><i class="icon iconfont icon-zhifubao"></i> &nbsp;支付宝支付</el-radio-button>
                   </el-radio-group>
                </el-form-item>

                <el-form-item>
                  <el-button type="primary">立即充值</el-button>
                </el-form-item>
              </el-form>
            </el-tab-pane>
          </el-tabs>

          <el-tabs type="card" style="height: 200px;">
            <el-tab-pane label="开通VIP">
              <el-form :model="payPointForm" ref="payPointForm" label-width="100px">
                <el-form-item label="充值友币" prop="money" >
                  <el-input readonly="readonly" value="一年内免费下载任意资源" disabled="disabled">
                    <template slot="prepend"><i class="icon iconfont icon-zuanshi" style="color: #5FB878;"></i></template>
                  </el-input>
                </el-form-item>

                <el-form-item label="支付方式" prop="money" >
                  <el-radio-group v-model="payType">
                    <el-radio-button label="weixin"><i class="icon iconfont icon-weixinzhifu2"></i> &nbsp;&nbsp;微信支付</el-radio-button>
                    <el-radio-button label="zhifubao"><i class="icon iconfont icon-zhifubao"></i> &nbsp;支付宝支付</el-radio-button>
                  </el-radio-group>
                </el-form-item>

                <el-form-item>
                  <el-button type="primary">立即充值</el-button>
                </el-form-item>
              </el-form>
            </el-tab-pane>
          </el-tabs>

        </template>
      </el-main>

    </el-container>
  </div>
</template>

<script>
  export default {
    data() {
      return {
        payType:"zhifubao",
        payPointForm: {
          "money":0,
          "payType": "zhifubao"
        },
      };
    },
    methods: {
    }
  }
</script>

