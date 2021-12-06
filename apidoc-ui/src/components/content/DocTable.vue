<template>
  <div class="doc-content">
    <div v-if="apiData.header && apiData.header.length">
      <h2>请求头Headers</h2>
      <div class="api-param-table">
        <Table
          :columns="paramsColumns"
          size="small"
          rowKey="name"
          :scroll="tableScroll"
          :bordered="true"
          :pagination="false"
          :data-source="apiData.header"
          childrenColumnName="params"
        >
        </Table>
      </div>
    </div>
    <div>
      <h2>
        请求参数Parameters &nbsp;
        <Tooltip title="编辑请求参数信息" v-if="!recordEditable">
          <Button type="primary" size="small" @click="recordEditable = true"
            >编辑</Button
          >
        </Tooltip>
        <template v-else>
          <Button
            type="primary"
            size="small"
            @click="saveData(apiData.param, 'request')"
            >保存</Button
          >
          &nbsp;&nbsp;
          <Button type="info" size="small" @click="addParam(apiData.param)"
            >新增</Button
          >
          &nbsp;&nbsp;
          <Tooltip title="清除后， 系统将根据请求响应生成新的文档">
            <Button type="danger" size="small">清除保存数据</Button>
          </Tooltip>
        </template>
      </h2>
      <div class="api-param-table">
        <Table
          :columns="paramsColumns"
          size="small"
          :rowKey="renterRowKey"
          :bordered="true"
          :pagination="false"
          :data-source="apiData.param"
          :scroll="tableScroll"
          defaultExpandAllRows
          childrenColumnName="params"
        >
          <template
            v-for="col in ['name', 'default', 'address', 'desc']"
            :slot="col"
            slot-scope="text, record"
          >
            <div :key="col">
              <template v-if="recordEditable">
                <Input
                  style="margin: -5px 0"
                  :value="text"
                  @blur="e => (record[col] = e.target.value)"
                >
                  <Icon
                    v-if="col === 'name'"
                    type="delete"
                    slot="addonAfter"
                    @click="
                      () => {
                        removeParam(record);
                      }
                    "
                  />
                </Input>
              </template>

              <div v-else-if="col === 'desc'" v-html="textToHtml(text)"></div>
              <template v-else>
                {{ text }}
              </template>
            </div>
          </template>

          <template slot="require" slot-scope="text, record">
            <template v-if="recordEditable">
              <Checkbox
                key="require"
                @change="e => (record.require = e.target.checked)"
                :checked="text"
              />
            </template>
            <template v-else>
              <Icon
                key="require"
                type="check"
                style="color:#1890ff"
                v-if="text"
              />
            </template>
          </template>

          <template slot="type" slot-scope="text, record">
            <template v-if="recordEditable">
              <Select
                :defaultValue="text"
                style="width: 90px"
                @change="v => (record.type = v)"
              >
                <Option v-for="item in types" :key="item">
                  {{ item }}
                </Option>
              </Select>
            </template>
            <template v-else>
              {{ text }}
            </template>
          </template>
        </Table>
      </div>
    </div>

    <h2>
      响应结果Responses
      <Popover title="统一响应体">
        <template slot="content">
          <textarea
            class="code-textarea"
            cols="30"
            rows="8"
            readonly
            v-model="apiData.raw_return"
          ></textarea>
          <div class="note-text">
            <span style="color:#f00;">*</span>响应案例
          </div>
        </template>
        <Icon
          style="float:right;color:#999;font-size:18px;"
          type="question-circle"
        />
      </Popover>
    </h2>
    <div class="api-param-table">
      <Table
        :columns="returnColumns"
        size="small"
        rowKey="_key"
        :bordered="true"
        :pagination="false"
        :data-source="apiData.return"
        :scroll="tableScroll"
        defaultExpandAllRows
        :expandedRowKeys="expandedRowKeys"
        childrenColumnName="params"
        @expandedRowsChange="onExpandedRowsChange"
      >
      </Table>
    </div>

    <h2>
      响应结果Responses示例
    </h2>

    <div class="api-param-table">
      <json-viewer
          :value="getJsonViewData()"
          :expand-depth="3"
          copyable
          boxed
      ></json-viewer>
    </div>
  </div>
</template>

<script>
import {
  Table,
  Icon,
  Popover,
  Input,
  Checkbox,
  Select,
  Button,
  Tooltip
} from "ant-design-vue";
import { textToHtml } from "../../utils/utils";
import request from "../../utils/request";
import { url } from "@/api/app";
import JsonViewer from "vue-json-viewer";

let paramsRowKey = 0;
export default {
  components: {
    Table,
    Icon,
    Popover,
    JsonViewer,
    Input,
    Checkbox,
    Select,
    Option: Select.Option,
    Button,
    Tooltip
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
  computed: {},
  data() {
    return {
      recordEditable: false,
      types: ["string", "number", "bool", "any"],
      paramsColumns: [
        {
          title: "名称",
          dataIndex: "name",
          width: 240,
          scopedSlots: { customRender: "name" }
        },
        {
          title: "类型",
          dataIndex: "type",
          align: "center",
          width: 130,
          scopedSlots: { customRender: "type" }
        },
        {
          title: "必填",
          dataIndex: "require",
          width: 60,
          align: "center",
          scopedSlots: { customRender: "require" }
        },
        {
          title: "默认值",
          dataIndex: "default",
          align: "center",
          scopedSlots: { customRender: "default" },
          width: 80
        },
        {
          title: "说明",
          dataIndex: "desc",
          scopedSlots: { customRender: "desc" }
        }
        // ,
        // {
        //   title: "启用",
        //   dataIndex: "enable",
        //   align: "center",
        //   width: 50,
        //   scopedSlots: { customRender: "enable" }
        // }
      ],
      returnColumns: [
        {
          title: "名称",
          dataIndex: "name",
          width: 240
        },
        {
          title: "类型",
          dataIndex: "type",
          align: "center",
          width: 130,
          customRender: (text, record) => {
            if (text === "array" && record.childrenType) {
              return `${text}<${record.childrenType}>`;
            } else {
              return text;
            }
          }
        },
        {
          title: "默认值",
          dataIndex: "default",
          align: "center",
          width: 80
        },
        {
          title: "说明",
          dataIndex: "desc",
          scopedSlots: { customRender: "rowDesc" }
        }
      ],
      tableScroll: {
        x: "700px",
        y: "100%"
      },
      expandedRowKeys: [],
      returnData: []
    };
  },
  watch: {
    apiData(val) {
      this.returnData = this.handleReturnData(val.return);
    }
  },
  created() {
    this.returnData = this.handleReturnData(this.apiData.return);
  },
  methods: {
    getJsonViewData() {
      try {
        return JSON.parse(this.apiData.raw_return);
      } catch (e) {
        return {};
      }
    },
    saveData(record, type) {
      request.post(
        url.edit + "?type=" + type + "&menu_key=" + this.apiData.menu_key,
        record
      );
      this.recordEditable = false;
    },
    removeParam(record) {
      this.apiData.param.splice(this.apiData.param.indexOf(record), 1);
    },
    addParam(record) {
      for (const idx in record) {
        if (record[idx].name === "") {
          this.$message.error("已存在未填写的参数行");
          return;
        }
      }
      record.push({
        name: "",
        default: "",
        rowDesc: "",
        type: "string",
        require: false,
        enable: true
      });
    },
    textToHtml,
    handleReturnData(data) {
      return data
        ? data.map(item => {
            paramsRowKey++;
            item._key = `${item.name}_${paramsRowKey}`;
            if (item.params) {
              this.expandedRowKeys.push(item._key);
              item.params = this.handleReturnData(item.params);
            }
            return item;
          })
        : [];
    },
    onExpandedRowsChange(expandedRows) {
      this.expandedRowKeys = expandedRows;
    },
    // 处理table行rowKey防止key重复
    renterRowKey(record) {
      paramsRowKey++;
      return `${record.name}_${paramsRowKey}`;
    }
  }
};
</script>

<style lang="less" scoped>
* {
  font-size: 13px;
}
.api-param-table {
  margin-bottom: 16px;
}
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
</style>
