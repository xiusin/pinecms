<template>
  <div :class="['header', device]">
    <div class="logo-box">
      <div
        v-if="device == 'mobile'"
        class="header-menu-button"
        @click="onShowMenuClick"
      >
        <Icon type="menu" />
      </div>
      <div class="logo">
        <img :src="logoPath" style="width: 30px;" alt="" />
      </div>
      <div
        v-if="
          !(device == 'mobile' && config && (config.cache || config.with_cache))
        "
        class="logo-text"
      >
        {{ config && config.title ? config.title : "Api Doc" }}
      </div>
    </div>
    <div class="header-menu"></div>
    <div class="user-wrapper">
      <div class="select-host" v-if="hostList && hostList.length > 1">
        <span v-if="device != 'mobile'">HOST: </span>
        <Select
          v-model="currentHost"
          :style="{
            width: device == 'mobile' ? '100px' : '170px'
          }"
          @change="onHostChange"
        >
          <SelectOption
            v-for="(item, index) in hostList"
            :key="index"
            :value="item.host"
          >
            {{ item.title }}
          </SelectOption>
        </Select>
      </div>

      <div class="select-app" v-if="apps.length">
        <span v-if="device != 'mobile'"
          >{{
            config.apps_title
              ? config.apps_title
              : this.config.versions
              ? "Select Version"
              : "应用"
          }}:
        </span>
        <app-select
          :style="{
            width: device == 'mobile' ? '100px' : '170px'
          }"
          :value="currentApp"
          :options="apps"
          @change="onAppChange"
        />
      </div>

      <div
        class="select-log"
        v-if="
          config &&
            (config.cache || config.with_cache) &&
            apiData &&
            apiData.cacheFiles &&
            apiData.cacheFiles.length > 1
        "
      >
        <Select
          v-model="currentCache"
          :style="{
            width: device == 'mobile' ? '100px' : '170px'
          }"
          @change="onCacheChange"
        >
          <SelectOption
            v-for="(item, index) in apiData.cacheFiles"
            :key="index"
            :value="item"
          >
            {{ item }}
          </SelectOption>
        </Select>
      </div>
      <div class="actions">
        <Tooltip placement="bottom">
          <template slot="title">
            设置全局请求参数
          </template>
          <Button
            v-if="isGlobalParams"
            icon="global"
            type="primary"
            @click="showGlobalParams"
            ><span v-if="device != 'mobile'">全局参数</span></Button
          >
          <Button v-else icon="global" @click="showGlobalParams"
            ><span v-if="device != 'mobile'">全局参数</span></Button
          >
        </Tooltip>

        <Tooltip
          v-if="
            (config && config.with_cache) ||
              (config && config.cache && config.cache.enable && config.debug)
          "
          placement="bottom"
        >
          <template slot="title">
            重新生成接口数据
          </template>

          <Button icon="reload" @click="reloadData" style="margin-left:8px;"
            ><span v-if="device != 'mobile'">Reload</span></Button
          >
        </Tooltip>
      </div>
    </div>
  </div>
</template>

<script>
import { Select, Tooltip, Button, Icon } from "ant-design-vue";
import GlobalParamsModal from "./globalParamsModal";
import AppSelect from "./AppSelect";
import { ls } from "@/utils/cache";
import { getUrlQuery } from "@/utils/utils";

export default {
  components: {
    Select,
    SelectOption: Select.Option,
    Tooltip,
    Button,
    Icon,
    AppSelect
  },
  props: {
    config: {
      type: Object,
      default: () => {}
    },
    apiData: {
      type: Object,
      default: () => {}
    },
    device: {
      type: String,
      default: "xl"
    }
  },
  data() {
    return {
      isGlobalParams: false,
      currentCache: "",
      logoPath: "./logo.png",
      currentApp: "",
      currentHost: "",
      hostList: [],
      urlQuery: {}
    };
  },
  computed: {
    apps() {
      if (this.config.apps && this.config.apps.length) {
        return this.config.apps;
      } else if (this.config.versions && this.config.versions.length) {
        return this.config.versions;
      }
      return [];
    }
  },
  watch: {
    apiData(val) {
      this.currentCache = val.cacheName;
      this.currentApp = val.appKey;
    }
  },
  created() {
    this.urlQuery = getUrlQuery();
    const globalParams = ls.get("globalParams");
    if (globalParams) {
      this.isGlobalParams = true;
    }
    // this.hostList = config.HOSTS;
    if (this.urlQuery && this.urlQuery.host) {
      this.currentHost = this.urlQuery.host;
    }
  },
  methods: {
    showGlobalParams() {
      GlobalParamsModal({
        config: this.config,
        success: this.setGlobalParamsSuccess
      });
    },
    setGlobalParamsSuccess(val) {
      if (val) {
        this.isGlobalParams = true;
      } else {
        this.isGlobalParams = false;
      }
    },
    onCacheChange(val) {
      this.$emit("cacheChange", this.currentApp, val);
    },
    reloadData() {
      this.$emit("reload", this.currentApp, "", true);
    },
    onShowMenuClick() {
      this.$emit("showSideMenu");
    },
    onAppChange(key) {
      if (this.config.versions && this.config.versions.length) {
        const find = this.config.versions.find(p => p.folder === key);
        this.$emit("appChange", find.title);
      } else {
        this.$emit("appChange", key);
      }
    },
    onHostChange(v) {
      console.log(v);
      this.currentHost = v;
      // config.HOST = v;

      const pathname = window.location.pathname ? window.location.pathname : "";
      const url = pathname + "?host=" + v;
      window.location.href = url;
    }
  }
};
</script>
<style lang="less" scoped>
.header {
  width: 100%;
  height: 40px;
  border-bottom: 1px solid #ddd;
  display: flex;
  padding: 0 5px;
  &.mobile {
    padding: 0 10px;
  }
  .header-menu-button {
    float: left;
    text-align: center;
    font-size: 20px;
    padding-right: 15px;
    padding-left: 5px;
  }
  .logo-box {
    padding: 4px 0;
    overflow: hidden;
    .logo {
      float: left;
      & > img {
      }
    }

    .logo-text {
      float: left;
      line-height: 32px;
      margin-left: 10px;
      font-weight: bold;
    }
  }
  .header-menu {
    flex: 1;
  }
  .user-wrapper {
    display: flex;
    .select-app {
      padding: 4px;
      flex: 1;
    }
    .select-log,
    .select-host {
      padding: 4px;
    }
    .actions {
      // margin-left: 16px;
      padding: 4px;
      .action {
        padding: 10px;
        display: inline-block;
        color: #555;
      }
    }
  }
}
</style>
