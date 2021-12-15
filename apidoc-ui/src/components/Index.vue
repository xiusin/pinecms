<template>
  <div :class="[`layout-size_${currentSize}`]">
    <Header
      ref="header"
      :config="config"
      :apiData="apiData"
      :device="device"
      @cacheChange="getApiList"
      @reload="getApiList"
      @appChange="getApiList"
      @showSideMenu="onShowSideMenu"
    />
    <div class="spin-box" v-if="loading">
      <Spin tip="加载中..." :spinning="loading"> </Spin>
    </div>
    <div v-else-if="error.status">
      <error-box :error="error" />
      <div style="text-align: center;">
        <Button icon="reload" size="large" @click="reloadPage">刷新</Button>
      </div>
    </div>
    <div v-else>
      <splitpanes style="height: calc(100vh - 50px)">
        <pane v-if="device != 'mobile'" size="15" min-size="10" max-size="40">
          <Card
            :bordered="false"
            style="height:100%"
            :bodyStyle="{ padding: 0 }"
          >
            <DocMenu
              ref="sideMenu"
              :apiData="apiData.list"
              :groups="apiData.groups"
              :tags="apiData.tags"
              :docs="apiData.docs"
              :config="config"
              @change="menuChange"
              @showCrud="onShowCrud"
            />
          </Card>
        </pane>
        <pane>
          <Card
            :bordered="false"
            style="height:100%;overflow: auto;"
            id="pageContainer"
            :bodyStyle="{ padding: device == 'mobile' ? '10px' : '10px' }"
          >
            <DocApiContent
              v-if="currentApiData && currentApiData.url"
              :apiData="currentApiData"
              :config="config"
            />
            <DocMdContent
              v-else-if="currentDocData && currentDocData.type === 'md'"
              :docData="currentDocData"
              :appKey="currentAppKey"
              :device="device"
            />
            <DocHome v-else :apiData="apiData" :config="config" />
          </Card>
        </pane>
      </splitpanes>
      <Drawer
        v-if="device == 'mobile'"
        :title="config && config.title ? config.title : 'Api Doc'"
        placement="left"
        :visible="visible.sideMenu"
        width="80%"
        :bodyStyle="{ padding: 0 }"
        @close="onSideMenuClose"
      >
        <DocMenu
          ref="sideMenu"
          :apiData="apiData.list"
          :groups="apiData.groups"
          :tags="apiData.tags"
          :docs="apiData.docs"
          :config="config"
          :device="device"
          @change="menuChange"
        />
      </Drawer>
    </div>
  </div>
</template>

<script>
import Vue from "vue";
import { Splitpanes, Pane } from "splitpanes";
import "splitpanes/dist/splitpanes.css";
import { Card, Spin, Drawer, Button } from "ant-design-vue";
import DocMenu from "./Menu";
import DocApiContent from "./content";
import DocHome from "./DocHome";
import VueClipboard from "vue-clipboard2";
import { ls } from "@/utils/cache";
import {
  setCurrentUrl,
  getUrlQuery,
  getTreeFirstNode,
  changeUrlArg,
  treeTransArray,
  deleteUrlArg,
  getCurrentAppConfig
} from "@/utils/utils";
import PasswordModal from "./auth/passwordModal";
import responsiveMixin from "@/utils/responsive";
import { getConfig, getApiData } from "@/api/app";
import DocMdContent from "./DocMdContent";
import CrudModal from "./crud";
import Header from "./Header";
import "./index.less";
import ErrorBox from "./ErrorBox.vue";
Vue.use(VueClipboard);
export default {
  components: {
    Card,
    Spin,
    Splitpanes,
    Pane,
    DocMenu,
    DocApiContent,
    DocHome,
    Header,
    Drawer,
    DocMdContent,
    Button,
    ErrorBox
  },
  mixins: [responsiveMixin],
  data() {
    return {
      loading: true,
      apiData: {},
      currentApiData: {},
      currentDocData: {},
      currentAppKey: "",
      config: {},
      visible: {
        sideMenu: false
      },
      error: {
        status: "",
        message: ""
      },
      clientWidth: 1920,
      urlQuery: {}
    };
  },
  created() {
    this.urlQuery = getUrlQuery();
    this.getConfig([this.urlQuery.appKey]);
    this.currentAppKey = this.urlQuery.appKey;

    // this.getApiList();
  },
  mounted() {
    this.clientWidth = document.body.clientWidth;
  },
  methods: {
    getApiList(appKey = "", cacheFileName = "", reload = false) {
      let version = null;
      if (appKey) {
        this.currentAppKey = appKey;
        if (this.config.versions && this.config.versions.length) {
          version = this.currentAppKey;
        }
      } else {
        // 默认版本/应用
        if (!this.currentAppKey && this.config.apps) {
          const firstNodes = getTreeFirstNode(this.config.apps, "items");
          if (firstNodes && firstNodes.length) {
            const keys = firstNodes.map(p => p.folder);
            this.currentAppKey = keys.join("_");
          }
        } else if (
          !this.currentAppKey &&
          this.config.versions &&
          this.config.versions.length
        ) {
          // 兼容低版本versions
          const firstNode = this.config.versions[0];
          if (firstNode) {
            this.currentAppKey = firstNode["title"];
            version = this.currentAppKey;
          }
        }
      }
      this.loading = true;
      // 更新url
      const url = changeUrlArg(
        window.location.href,
        "appKey",
        this.currentAppKey
      );
      setCurrentUrl(url);
      const json = {
        appKey: this.currentAppKey,
        version: version,
        cacheFileName: cacheFileName,
        reload: reload
      };
      getApiData(json)
        .then(res => {
          this.loading = false;
          const json = {
            ...res.data.data,
            appKey: this.currentAppKey
          };
          this.apiData = json;
          this.currentApiData = {};
          this.currentDocData = {};

          setTimeout(() => {
            if (this.urlQuery.md && this.$refs.sideMenu) {
              // 跳转url指定md
              const getMenuData = this.$refs.sideMenu.getMenuData();
              const mdList = treeTransArray(getMenuData[0].items, "items");
              const mdFind = mdList.find(p => p.path === this.urlQuery.md);
              if (mdFind) {
                this.$refs.sideMenu.onMenuClick(mdFind);
              }
            } else if (this.urlQuery.api) {
              const apiList = treeTransArray(this.apiData.list, "children");
              const apiFind = apiList.find(p => p.url === this.urlQuery.api);
              if (apiFind) {
                this.$refs.sideMenu.onMenuClick(apiFind);
              }
            }
          }, 100);
        })
        .catch(err => {
          const status =
            err.response && err.response.status ? err.response.status : 500;
          if (status === 401) {
            ls.remove("token");
            PasswordModal({
              appKey: this.currentAppKey,
              success: () => {
                window.location.reload();
              }
            });
          } else {
            this.error = {
              status: status,
              message:
                err.response && err.response.data && err.response.data.message
                  ? err.response.data.message
                  : err.message
            };
            this.loading = false;
          }
        });
    },
    menuChange(currentApiData) {
      if (currentApiData.type === "md") {
        // docs文档
        this.currentDocData = currentApiData;
        this.currentApiData = {};
      } else {
        this.currentApiData = currentApiData;
        this.currentDocData = {};
        let url = changeUrlArg(
          window.location.href,
          "api",
          this.currentApiData.url
        );
        url = deleteUrlArg(url, "md");
        setCurrentUrl(url);
      }
      this.visible.sideMenu = false;
    },

    getConfig(option) {
      getConfig().then(res => {
        if (res.data && res.data.title) {
          this.config = res.data;
        } else if (res.data && res.data.data) {
          this.config = res.data.data;
        }
        ls.set("config", this.config);
        document.title = this.config.title;
        this.verifyAuth(option);
      });
      // .catch(err => {
      //   const status =
      //     err.response && err.response.status ? err.response.status : 404;
      //   this.error = {
      //     status: status,
      //     message:
      //       err.response && err.response.data && err.response.data.message
      //         ? err.response.data.message
      //         : err.message
      //   };
      //   this.loading = false;
      // });
    },
    verifyAuth(option) {
      // 默认版本/应用
      // let currentAppKey = "";
      // if (this.currentAppKey && this.config.apps) {
      //   const appList = treeTransArray(this.config.apps, "items");
      //   const currentApp = appList.find(p => p.folder === this.currentAppKey);
      //   currentAppKey = currentApp;
      // }
      const that = this;
      const currentApp = getCurrentAppConfig(
        this.currentAppKey,
        this.config.apps
      );
      const tokenKey =
        currentApp && currentApp.hasPassword ? this.currentAppKey : "global";
      const token = ls.get("token_" + tokenKey);
      console.log(token, currentApp, this.currentAppKey, option);
      ls.set("current_app", this.currentAppKey);
      if (
        !token &&
        ((this.config && this.config.auth && this.config.auth.enable) ||
          (currentApp && currentApp.hasPassword))
      ) {
        // 不存在token并需要登录
        // 密码验证方法
        PasswordModal({
          appKey: this.currentAppKey,
          success: () => {
            window.location.reload();
          }
        });
      } else {
        that.getApiList(...option);
      }
    },
    // 显示移动端侧边栏
    onShowSideMenu() {
      this.visible.sideMenu = true;
    },
    onSideMenuClose() {
      this.visible.sideMenu = false;
    },
    onShowCrud() {
      const that = this;
      CrudModal({
        config: this.config,
        currentAppKey: this.currentAppKey,
        clientWidth: this.clientWidth,
        success: () => {
          that.getApiList();
        }
      });
    },
    reloadPage() {
      window.location.href = "/apidoc/";
    }
  }
};
</script>

<style lang="less" scoped>
.spin-box {
  width: 100%;
  height: 100vh;
  text-align: center;
  padding-top: 100px;
}
</style>
