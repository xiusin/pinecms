<template>
  <div class="doc-content">
    <div v-if="headerData && headerData.length">
      <h2>请求头Headers</h2>
      <div class="api-param-table">
        <Table
          :columns="headersColumns"
          size="small"
          rowKey="name"
          :bordered="true"
          :pagination="false"
          :data-source="headerData"
          :scroll="tableScroll"
        >
          <template slot="headerValue" slot-scope="text, record">
            <TableInput
              :style="{ width: device === 'mobile' ? '200px' : '350px' }"
              :data="text"
              @change="onHeaderCellChange(record.name, 'default', $event)"
            />
          </template>
        </Table>
      </div>
    </div>
    <div v-if="apiData.param && apiData.param.length">
      <h2>请求参数Parameters</h2>
      <div class="api-param-textarea">
        <div
          v-if="
            apiData.paramType === 'formdata' || apiData.paramType === 'route'
          "
          class="param-box"
        >
          <Form :label-col="{ span: 5 }" :wrapper-col="{ span: 12 }">
            <FormItem
              v-for="item in apiData.param"
              :key="item.name"
              :label="item.name"
            >
              <div v-if="item.type === 'file'">
                <Upload
                  :file-list="fileList[item.name]"
                  :remove="
                    file => {
                      fileHandleRemove(file, item.name);
                    }
                  "
                  :before-upload="
                    file => {
                      fileBeforeUpload(file, item.name);
                      return false;
                    }
                  "
                  :name="item.name"
                >
                  <Button> Select File </Button>
                </Upload>
              </div>
              <Input v-else v-model="formdata[item.name]" />
            </FormItem>
          </Form>
        </div>
        <codemirror class="code" v-model="parameters" :options="cmOptions"></codemirror>
      </div>
    </div>

    <div class="api-debug-action">
      <Button type="primary" :loading="loading" block @click="excute"
        >执行</Button
      >
    </div>

    <div>
      <h2>响应结果Responses</h2>
      <div v-if="returnData && returnData.status" class="api-param-table">
        <Alert
          v-if="returnData.status === 200"
          :message="returnData.status"
          type="success"
          show-icon
        />
        <Alert v-else :message="returnData.status" type="error" show-icon />
        <div class="api-param-code">
          <div class="code">
            <div
              v-if="returnString"
              class="string-code"
              v-html="returnString"
            ></div>
            <json-viewer
              :value="typeof returnData.data !== 'undefined' ? returnData.data : {}"
              :expand-depth="3"
              copyable
              boxed
            ></json-viewer>
          </div>
        </div>
      </div>
      <div v-else class="api-param-empty">
        <Empty :description="false" />
      </div>
    </div>
  </div>
</template>

<script>
import {
  Input,
  Button,
  Alert,
  Empty,
  Table,
  Form,
  Upload,
  message
} from "ant-design-vue";
import { sendRequest } from "@/utils/request";
import TableInput from "@/utils/Input";
import cloneDeep from "lodash/cloneDeep";
import { ls } from "@/utils/cache";
import JsonViewer from "vue-json-viewer";
import {codemirror} from 'vue-codemirror';

import 'codemirror/lib/codemirror.css'
import 'codemirror/keymap/sublime'
import "codemirror/theme/dracula.css"
import "codemirror/mode/vue/vue.js"
import 'codemirror/addon/selection/active-line'


export default {
  components: {
    Input,
    Button,
    Alert,
    Empty,
    Table,
    TableInput,
    Form,
    codemirror,
    FormItem: Form.Item,
    Upload,
    JsonViewer
  },
  props: {
    apiData: {
      type: Object,
      default: () => {}
    },
    device: {
      type: String,
      default: "xl"
    },
    url: {
      type: String,
      default: ""
    },
    currentMethod: {
      type: String,
      default: "GET"
    }
  },

  data() {
    return {
      code: "",
      cmOptions: {
        tabSize: 4,// tab的空格个数
        theme: 'dracula',//主题样式
        lineNumbers: true,//是否显示行数
        lineWrapping: true, //是否自动换行
        styleActiveLine: true,//line选择是是否加亮
        matchBrackets: true,//括号匹配
        mode: "javascript", //实现javascript代码高亮
        readOnly: false//只读
      },
      returnString: "",
      returnData: {},
      parameters: "",
      method: "",
      headersColumns: [
        {
          title: "Key",
          dataIndex: "name",
          width: 240
        },
        {
          title: "Value",
          dataIndex: "default",
          width: this.device === "mobile" ? 150 : 350,
          scopedSlots: { customRender: "headerValue" }
        },
        {
          title: "说明",
          dataIndex: "desc"
        }
      ],
      headerData: [],
      fileList: {},
      formdata: {},
      loading: false,
      tableScroll: {
        x: "600px",
        y: "100%"
      },
      config: {}
    };
  },
  watch: {
    apiData() {
      this.returnData = {};
      this.initApiData();
    },
    currentMethod(v) {
      this.method = v;
    }
  },

  created() {
    this.config = ls.get("config");
    this.initApiData();
    this.method = this.currentMethod;
  },
  methods: {
    initApiData() {
      this.handleParameters();
      if (this.apiData.paramType === "formdata") {
        let fileList = {};
        let formdata = {};
        this.apiData.param.forEach(item => {
          if (item.type === "file") {
            fileList[item.name] = [];
          } else {
            if (
              this.globalParams &&
              this.globalParams.params &&
              this.globalParams.params.length
            ) {
              const paramsItem = this.globalParams.params.find(
                p => p.key === item.name
              );
              if (paramsItem && paramsItem.value) {
                formdata[item.name] = paramsItem.value;
              } else {
                formdata[item.name] = item.default ? item.default : "";
              }
            } else {
              formdata[item.name] = item.default ? item.default : "";
            }
          }
        });
        this.fileList = fileList;
        this.formdata = formdata;
      }
      this.headerData = this.renderHeaderData(this.apiData.header);
    },
    handleParameters() {
      let newParams = JSON.parse(this.apiData.raw_param);
      // 处理全局请求参数
      this.globalParams = ls.get("globalParams");
      if (
        this.globalParams &&
        this.globalParams.params &&
        this.globalParams.params.length
      ) {
        for (let i = 0; i < this.globalParams.params.length; i++) {
          const globalParamItem = this.globalParams.params[i];
          const paramsItem = newParams.find(
            p => p.name === globalParamItem.key
          );
          if (paramsItem && !paramsItem.default) {
            paramsItem.default = globalParamItem.value;
          }
        }
      }

      this.parameters = JSON.stringify(newParams, null, 4); // renderParamsCode(params);
    },

    excute() {
      const that = this;
      let url = cloneDeep(this.url);
      this.loading = true;
      let json = {};
      if (this.apiData.paramType == "formdata") {
        const formData = new FormData();
        this.apiData.param.forEach(item => {
          if (item.type === "file") {
            const fileList = this.fileList[item.name];
            if (fileList && fileList.length) {
              formData.append(item.name, fileList[0]);
            }
          } else {
            formData.append(item.name, this.formdata[item.name]);
          }
        });
        json = formData;
      } else if (this.apiData.paramType == "route") {
        // 路由参数，将参数拼接到url中
        this.apiData.param.forEach(item => {
          const placeholderKeys = [
            `:${item.name}`,
            `<${item.name}>`,
            `<${item.name}?>`,
            `[:${item.name}]`
          ];
          for (let i = 0; i < placeholderKeys.length; i++) {
            const key = placeholderKeys[i];
            if (url.indexOf(key) > -1) {
              const reg = new RegExp(key, "g");
              url = url.replace(reg, this.formdata[item.name]);
            }
          }
        });
      } else {
        const string = this.parameters;
        try {
          json = eval("(" + string + ")");
        } catch (error) {
          message.error("json 格式错误，请检查");
          this.loading = false;
          return false;
        }
      }
      const method = this.method.toLowerCase();
      const headers = {};
      if (this.headerData && this.headerData.length) {
        this.headerData.forEach(item => {
          headers[item.name] = item.default;
        });
      }
      // 添加全局请求头参数
      const globalParams = ls.get("globalParams");
      if (globalParams && globalParams.headers && globalParams.headers.length) {
        for (let i = 0; i < globalParams.headers.length; i++) {
          const globalHeaderParam = globalParams.headers[i];
          if (!headers[globalHeaderParam.key]) {
            headers[globalHeaderParam.key] = globalHeaderParam.value;
          }
        }
      }
      // 添加全局请求参数
      if (globalParams && globalParams.params && globalParams.params.length) {
        for (let i = 0; i < globalParams.params.length; i++) {
          const globalParamItem = globalParams.params[i];
          if (!json[globalParamItem.key]) {
            json[globalParamItem.key] = globalParamItem.value;
          }
        }
      }
      if (this.apiData.paramType === "formdata") {
        headers["Content-Type"] = "application/x-www-form-urlencoded";
      }

      let host = "http://localhost:2019";
      for (let i in this.config.apps) {
        if (this.config.apps[i].folder == ls.get("current_app")) {
          host = this.config.apps[i].host;
        }
      }
      console.log(host + url, json, method, headers);
      sendRequest(host + url, json, method, headers)
        .then(res => {
          this.loading = false;
          if (res.data && typeof res.data === "string") {
            that.returnString = res.data;
            that.returnData = res;
          } else {
            that.returnString = "";
            that.returnData = res;
          }
        })
        .catch(err => {
          this.loading = false;
          if (err.response) {
            this.returnData = err.response;
          } else {
            this.returnData = {
              status: 500
            };
          }
        });
    },
    renderHeaderData(headerData) {
      const data = cloneDeep(headerData);
      if (data && data.length) {
        const globalParams = ls.get("globalParams");
        if (
          globalParams &&
          globalParams.headers &&
          globalParams.headers.length
        ) {
          return data.map(item => {
            const globalParamFind = globalParams.headers.find(
              p => p.key === item.name
            );
            if (globalParamFind && globalParamFind.value) {
              item.default = globalParamFind.value;
            }
            return item;
          });
        } else {
          return data;
        }
      }

      return [];
    },
    onHeaderCellChange(key, dataIndex, value) {
      const dataSource = [...this.headerData];
      const target = dataSource.find(item => item.name === key);
      if (target) {
        target[dataIndex] = value;
        this.headerData = dataSource;
      }
    },
    fileBeforeUpload(file, name) {
      this.fileList[name] = [file];

      return false;
    },
    fileHandleRemove(file, name) {
      let fileList = this.fileList[name];
      const index = fileList.indexOf(file);
      const newFileList = fileList.slice();
      newFileList.splice(index, 1);
      this.fileList[name] = newFileList;
    }
  }
};
</script>

<style lang="less" scoped>
.api-param-textarea textarea {
  font-family: "SFMono-Regular", Consolas, "Liberation Mono", Menlo, Courier,
    monospace;
}
.api-debug-action {
  padding: 16px 0;
}
.api-param-code {
  margin-top: 16px;
  /deep/.hljs {
    max-height: 500px;
  }
  .code {
    margin-bottom: 1em;
    pre {
      margin-bottom: 0;
      border-radius: 4px;
    }
  }
}
.api-param-empty {
  border: 1px solid #ddd;
  border-radius: 4px;
  padding: 16px;
}
.api-param-table {
  margin-bottom: 16px;
}
/deep/
  .ant-table-small
  > .ant-table-content
  > .ant-table-body
  > table
  > .ant-table-thead
  > tr
  > th {
  background: #fafafa;
  font-weight: 600;
}
/deep/ .ant-table-small > .ant-table-content > .ant-table-body {
  margin: 0;
}
.string-code {
  display: block;
  overflow-x: auto;
  padding: 0.5em;
  color: #abb2bf;
  background: #282c34;
  border-radius: 4px;
}
</style>
