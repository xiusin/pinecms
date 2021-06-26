<template>
  <Modal
    title="快速创建Crud"
    :width="modalWidth"
    :visible="visible"
    :destroyOnClose="true"
    :maskClosable="false"
    okText="确认创建"
    cancelText="取消"
    :bodyStyle="{ padding: ' 16px 16px' }"
    @cancel="handleCancel"
    @ok="handleOk"
  >
    <div class="crud-form">
      <Base
        ref="base"
        :config="config"
        :currentAppKey="appKey"
        @titleEnter="onTitleEnter"
        @appkeyChange="onAppkeyChange"
      />
      <ClassConfig
        ref="classConfig"
        :currentAppKey="appKey"
        :config="config"
        @className="onClassNameChange"
      />

      <Table ref="table" :config="config" />
    </div>
  </Modal>
</template>

<script>
import { Modal, message } from "ant-design-vue";
import ClassConfig from "./ClassConfig";
import Table from "./Table";
import Base from "./Base";
// import { treeTransArray } from "../../utils/utils";
import { createCrud } from "@/api/app";

export default {
  components: {
    Modal,
    ClassConfig,
    Table,
    Base
  },
  props: {
    config: {
      type: Object,
      default: () => {}
    },
    currentAppKey: String,
    success: Function,
    cancel: Function,
    clientWidth: Number
  },
  data() {
    return {
      visible: false,
      loading: false,
      appKey: "",
      modalWidth: 1200
    };
  },
  computed: {},
  created() {
    this.appKey = this.currentAppKey;
    if (this.clientWidth < 1200) {
      this.modalWidth = "95%";
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
      const baseData = this.$refs.base.getData();
      const classConfig = this.$refs.classConfig.getData();
      const tableData = this.$refs.table.getData();

      if (!classConfig || !tableData) {
        return false;
      }

      classConfig.model.table = tableData;
      const json = {
        ...baseData,
        ...classConfig,
        appKey: this.appKey
      };
      createCrud(json)
        .then(() => {
          message.success("创建成功");
          this.$emit("success");
          this.handleCancel();
        })
        .catch(err => {
          const status =
            err.response && err.response.status ? err.response.status : 404;
          const error = {
            status: status,
            message:
              err.response && err.response.data && err.response.data.message
                ? err.response.data.message
                : err.message
          };
          Modal.warning({
            title: "创建失败",
            content: error.message,
            okText: "确认"
          });
        });
    },
    onClassNameChange(name) {
      this.$refs.service.setClassName(name);
      this.$refs.model.setClassName(name);
    },
    onTitleEnter() {
      this.$refs.classConfig.focus();
    },
    onAppkeyChange(val) {
      this.appKey = val;
    }
  }
};
</script>

<style lang="less" scoped></style>
