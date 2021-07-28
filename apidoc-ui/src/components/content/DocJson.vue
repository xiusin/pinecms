<template>
  <div class="doc-content">
    <div v-if="apiData.param && apiData.param.length">
      <h2>请求参数Parameters</h2>
      <div class="api-param-code">
        <div class="code">
          {{ paramCode }}
        </div>
      </div>
    </div>

    <Tabs defaultActiveKey="1" type="card">
      <TabPane tab="响应结果Responses" key="1">
        <json-viewer :value="getJsonViewData()" :expand-depth="3" copyable boxed></json-viewer>
      </TabPane>
    </Tabs>
  </div>
</template>

<script>
import {renderParamsCode} from "@/utils/utils";
import {Icon, Popover, Tabs} from "ant-design-vue";
import JsonViewer from "vue-json-viewer";

export default {
  components: {
    Popover,
    Icon,
    Tabs,
    TabPane: Tabs.TabPane,
    JsonViewer
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
    paramCode() {
      return renderParamsCode(this.apiData.param, 0, true);
    },
    returnCode() {
      return renderParamsCode(this.apiData.return, 0, true);
    }
  },
  data() {
    return {};
  },

  created() {},
  methods: {
    getJsonViewData() {
      try{
        return JSON.parse(this.apiData.raw_return)
      } catch (e) {
        return {}
      }
    }
  }
};
</script>

<style lang="less" scoped>
.code-textarea,
.code-textarea:focus {
  border: none;
  resize: none;
}
.note-text {
  color: #999;
  font-size: 14px;
  border-top: 1px solid #ddd;
  padding-top: 5px;
}
.api-param-code {
  margin-top: 16px;
  max-height: 500px;
  overflow: auto;
  .code {
    margin-bottom: 1em;
    pre {
      margin-bottom: 0;
      border-radius: 4px;
    }
  }
}
</style>
