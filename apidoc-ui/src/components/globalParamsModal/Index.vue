<template>
  <Modal
    title="设置全局请求参数"
    :width="800"
    :visible="visible"
    :destroyOnClose="true"
    :maskClosable="false"
    :bodyStyle="{ padding: '0 10px 10px' }"
    @cancel="handleCancel"
  >
    <a-tabs default-active-key="1">
      <a-tab-pane key="1" tab="全局Header">
        <a-alert
          type="info"
          show-icon
          style="margin-bottom:10px"
          message="发送请求时，所有接口将自动携带以下Header参数"
        ></a-alert>
        <params-table ref="headerTable" :data="headers" />
      </a-tab-pane>
      <a-tab-pane key="2" tab="全局Params">
        <a-alert
          type="info"
          show-icon
          style="margin-bottom:10px"
          message="发送请求时，所有接口将自动携带以下Params参数；如请求参数中存在，则替换全局参数"
        ></a-alert>
        <params-table ref="paramsTable" :data="params" />
      </a-tab-pane>
    </a-tabs>
    <template slot="footer">
      <a-popconfirm
        title="确认清空所有全局参数吗?"
        ok-text="确认"
        cancel-text="取消"
        @confirm="handleDelete"
      >
        <Button type="danger" ghost>清空</Button>
      </a-popconfirm>
      <Button type="primary" @click="handleOk">确认</Button>
    </template>
  </Modal>
</template>

<script>
import {
  Modal,
  Button,
  message,
  Tabs,
  Alert,
  Popconfirm
} from "ant-design-vue";
import { ls } from "@/utils/cache";
import ParamsTable from "./ParamsTable";

export default {
  components: {
    Modal,
    Button,
    ParamsTable,
    [Tabs.name]: Tabs,
    [Tabs.TabPane.name]: Tabs.TabPane,
    [Alert.name]: Alert,
    [Popconfirm.name]: Popconfirm
  },
  props: {
    config: {
      type: Object,
      default: () => {}
    },
    success: Function,
    cancel: Function
  },
  data() {
    return {
      visible: false,
      loading: false,
      headers: [],
      params: []
    };
  },
  created() {
    const globalParams = ls.get("globalParams");
    if (globalParams && globalParams.headers && globalParams.headers.length) {
      this.headers = globalParams.headers;
    } else if (
      this.config &&
      this.config.headers &&
      this.config.headers.length
    ) {
      this.headers = this.config.headers.map((p, i) => {
        return {
          id: i + 1,
          key: p.name,
          desc: p.desc,
          isconfig: true //配置中的参数
        };
      });
    } else {
      this.headers = [{ id: 1, key: "", value: "", desc: "" }];
    }
    if (globalParams && globalParams.params && globalParams.params.length) {
      this.params = globalParams.params;
    } else if (
      this.config &&
      this.config.parameters &&
      this.config.parameters.length
    ) {
      this.params = this.config.parameters.map((p, i) => {
        return {
          id: i + 1,
          key: p.name,
          desc: p.desc,
          isconfig: true
        };
      });
    } else {
      this.params = [{ id: 1, key: "", value: "", desc: "" }];
    }
  },
  mounted() {
    this.visible = true;
  },
  methods: {
    handleCancel() {
      this.$emit("cancel");
      this.visible = false;
    },
    handleOk() {
      const headerTable = this.$refs.headerTable;
      const paramsTable = this.$refs.paramsTable;
      const headers = headerTable
        .getData()
        .filter(p => (p.key && !p.isconfig) || (p.isconfig && p.value));
      let params = [];
      if (paramsTable) {
        params = paramsTable.getData();
      } else {
        params = this.params;
      }
      const json = {
        headers,
        params: params.filter(
          p => (p.key && !p.isconfig) || (p.isconfig && p.value)
        )
      };
      if (!json.headers.length && !json.params.length) {
        ls.remove("globalParams");
        this.$emit("success", false);
      } else {
        ls.set("globalParams", json);
        this.$emit("success", json);
      }
      message.success("设置成功");
      this.visible = false;
    },
    handleDelete() {
      ls.remove("globalParams");
      message.success("操作成功");
      this.visible = false;
      this.$emit("success", false);
    }
  }
};
</script>

<style lang="less" scoped></style>
