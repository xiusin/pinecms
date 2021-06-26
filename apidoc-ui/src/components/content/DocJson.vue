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

    <h2>
      响应结果Responses
      <Popover
        v-if="config && config.responses && config.responses.jsonStr"
        title="统一响应体"
      >
        <template slot="content">
          <textarea
            class="code-textarea"
            cols="30"
            rows="8"
            readonly
            v-model="config.responses.jsonStr"
          ></textarea>
          <div class="note-text">
            <span style="color:#f00;">*</span>以下只展示{{
              config.responses.main && config.responses.main.desc
                ? config.responses.main.desc
                : "业务数据"
            }}内容
          </div>
        </template>
        <Icon
          style="float:right;color:#999;font-size:18px;"
          type="question-circle"
        />
      </Popover>
    </h2>
    <div class="api-param-table">
      <div class="api-param-code">
        <div class="code">
          {{ returnCode }}
        </div>
      </div>
    </div>
  </div>
</template>

<script>
// import VueHighlightJS from "vue-highlight.js";
import "highlight.js/styles/atom-one-dark.css";
// import javascript from "highlight.js/lib/languages/javascript";
// import json from "highlight.js/lib/languages/json";
import { renderParamsCode } from "@/utils/utils";
import { Popover, Icon } from "ant-design-vue";

// Vue.use(VueHighlightJS, {
//   languages: {
//     javascript,
//     json
//   }
// });

export default {
  components: {
    Popover,
    Icon
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
      const code = renderParamsCode(this.apiData.param, 0, true);
      return code;
    },
    returnCode() {
      const code = renderParamsCode(this.apiData.return, 0, true);
      return code;
    }
  },
  data() {
    return {};
  },

  created() {},
  methods: {}
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
