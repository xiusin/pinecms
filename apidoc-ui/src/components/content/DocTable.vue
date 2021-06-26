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
    <div v-if="apiData.param && apiData.param.length">
      <h2>请求参数Parameters</h2>
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
          <template slot="rowDesc" slot-scope="text">
            <div v-html="textToHtml(text)"></div>
          </template>
        </Table>
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
        <template slot="rowDesc" slot-scope="text">
          <div v-html="textToHtml(text)"></div>
        </template>
      </Table>
    </div>
  </div>
</template>

<script>
import { Table, Icon, Popover } from "ant-design-vue";
import { textToHtml } from "../../utils/utils";

let paramsRowKey = 0;
export default {
  components: {
    Table,
    Icon,
    Popover
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
      paramsColumns: [
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
            if (text == "array" && record.childrenType) {
              return `${text}<${record.childrenType}>`;
            } else {
              return text;
            }
          }
        },
        {
          title: "必填",
          dataIndex: "require",
          width: 60,
          align: "center",
          customRender: text => {
            if (text == 1) {
              return <Icon type="check" style="color:#1890ff" />;
            } else {
              return "";
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
            if (text == "array" && record.childrenType) {
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
