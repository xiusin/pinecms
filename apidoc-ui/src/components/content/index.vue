<template>
  <div class="doc-content">
    <div class="doc-header">
      <h1 class="title">{{ apiData.title }}</h1>
      <div class="text-list" style="margin-bottom:10px">
        <div v-if="apiData.author" class="text-list-item">
          <span class="text-label">作者：</span>
          <span class="text-value">{{ apiData.author }}</span>
        </div>
        <div v-if="tags && tags.length" class="text-list-item">
          <span class="text-label">Tags：</span>
          <span class="text-value">
            <Tag v-for="(item, index) in tags" :key="index">{{ item }}</Tag>
          </span>
        </div>
      </div>
      <div class="title-sub" v-html="desc"></div>
    </div>
    <div class="api-url-box">
      <div
        v-if="methodTags"
        :class="['api-method-select']"
        :title="apiData.method"
      >
        <a-select v-model="currentMethod" style="width:100%;">
          <a-select-option
            v-for="item in methodTags"
            :key="item"
            :value="item"
            >{{ item }}</a-select-option
          >
        </a-select>
      </div>
      <div v-else class="api-url-tag" :style="{ background: methodColor }">
        {{ apiData.method }}
      </div>
      <div :class="['api-url-input', { 'method-multiple': methodTags }]">
        <input v-model="url" readonly />
      </div>
      <div class="api-url-copy">
        <Icon type="copy" @click="copyUrl" />
      </div>
    </div>

    <Tabs default-active-key="1">
      <TabPane key="1" tab="文档">
        <DocTable :apiData="apiData" :config="config" />
      </TabPane>
      <TabPane key="2" tab="Json" force-render>
        <DocJson :apiData="apiData" :config="config" />
      </TabPane>
      <TabPane key="3" tab="调试">
        <DocDebug
          :url="url"
          :apiData="apiData"
          :currentMethod="currentMethod"
        />
      </TabPane>
    </Tabs>
  </div>
</template>

<script>
import { Icon, Tabs, message, Tag, Select } from "ant-design-vue";
import DocTable from "./DocTable";
import DocJson from "./DocJson";
import DocDebug from "./DocDebug";
import { textToHtml } from "../../utils/utils";

export default {
  components: {
    Icon,
    Tabs: Tabs,
    TabPane: Tabs.TabPane,
    DocTable,
    DocJson,
    DocDebug,
    Tag,
    [Select.name]: Select,
    [Select.Option.name]: Select.Option
  },
  props: {
    apiData: {
      type: Object,
      default: () => {}
    },
    config: {
      type: Object,
      default: () => {}
    }
  },
  computed: {
    methodColor() {
      let color = "";
      switch (this.apiData.method) {
        case "GET":
          color = "#87d068";
          break;
        case "POST":
          color = "#2db7f5";
          break;
        case "PUT":
          color = "#ff9800";
          break;
        case "DELETE":
          color = "#ff4d4f";
          break;
        default:
          color = "#ddd";
          break;
      }
      return color;
    },
    methodTags() {
      if (this.apiData.method.indexOf(",") > -1) {
        const tags = this.apiData.method.split(",");
        return tags;
      }
      return "";
    },
    tags() {
      let tags = [];
      if (this.apiData.tag && this.apiData.tag.indexOf(" ") > -1) {
        tags = this.apiData.tag.split(" ");
      } else if (this.apiData.tag) {
        tags = [this.apiData.tag];
      }
      return tags;
    },
    desc() {
      if (this.apiData.desc) {
        return textToHtml(this.apiData.desc);
      }
      return "";
    }
  },
  data() {
    return {
      url: "",
      currentMethod: ""
    };
  },
  watch: {
    "apiData.url"(url) {
      this.url = url;
    },
    "apiData.method"(method) {
      if (method.indexOf(",") > -1) {
        const methodTags = method.split(",");
        this.currentMethod = methodTags[0];
      } else {
        this.currentMethod = method;
      }
    }
  },
  created() {
    this.url = this.apiData.url;
    if (this.methodTags && this.methodTags.length) {
      this.currentMethod = this.methodTags[0];
    } else {
      this.currentMethod = this.apiData.method;
    }
  },
  methods: {
    copyUrl() {
      const text = this.url;
      this.$copyText(text)
        .then(() => {
          message.success("复制成功");
        })
        .catch(() => {
          message.error("复制失败");
        });
    }
  }
};
</script>

<style lang="less" scoped>
.doc-content {
  margin: 0 auto;
  .title-label {
    color: #333;
  }
  .title {
    font-size: 26px;
  }
  .title-sub {
    font-size: 16px;
    color: #999;
    margin-bottom: 16px;
  }

  .api-url-box {
    width: 100%;
    position: relative;
    height: 38px;
    border-radius: 4px;
    margin-bottom: 24px;

    .api-url-tag,
    .api-method-select {
      width: 70px;
      text-align: center;
      line-height: 38px;
      color: #fff;
      position: absolute;
      top: 0;
      left: 0;
      border-top-left-radius: 4px;
      border-bottom-left-radius: 4px;
    }
    .api-method-select {
      width: 105px;
      background: #fafafa;
      height: 36px;
      top: 1px;
      left: 1px;
      /deep/.ant-select-selection {
        border: none;
        background: none;
        box-shadow: none;
        .ant-select-selection-selected-value {
          width: 100%;
        }
      }
    }
    .api-url-input {
      line-height: 33px;
      input {
        box-sizing: border-box;
        margin: 0;
        font-variant: tabular-nums;
        list-style: none;
        font-feature-settings: "tnum";
        width: 100%;
        height: 38px;
        color: rgba(0, 0, 0, 0.65);
        font-size: 16px;
        line-height: 1.5;
        background-color: #fff;
        background-image: none;
        border: 1px solid #d9d9d9;
        border-radius: 4px;
        transition: all 0.3s;
        font-family: monospace;
        padding: 0 50px 0 80px;
        &:hover {
          border-color: #40a9ff;
          border-right-width: 1px !important;
        }
        &:focus {
          outline: 0;
          box-shadow: 0 0 0 2px rgba(24, 144, 255, 0.2);
        }
      }
      &.method-multiple {
        input {
          padding-left: 110px;
        }
      }
      // &.method-num-2 {
      //   input {
      //     padding-left: 90px;
      //   }
      // }
      // &.method-num-3 {
      //   input {
      //     padding-left: 120px;
      //   }
      // }
      // &.method-num-4 {
      //   input {
      //     padding-left: 175px;
      //   }
      // }
    }
    .api-url-copy {
      padding: 8px 15px;
      position: absolute;
      top: 0;
      right: 0;
    }
  }

  .api-param-table {
    margin-bottom: 16px;
  }
  .text-list-item {
    margin-bottom: 16px;
  }
}
</style>
