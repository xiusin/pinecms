<template>
  <div class="crud-base form-wraper">
    <div class="form-item">
      <span v-if="config.apps && config.apps.length">
        <span class="form-item_label"
          >{{ config.apps_title ? config.apps_title : "App/Version" }}：</span
        >
        <app-select
          :value="appKey"
          :options="config.apps"
          @change="onAppChange"
        />
      </span>

      <span v-if="config.groups && config.groups.length">
        <span class="form-item_label">分组：</span>
        <a-select style="width:130px;" v-model="group">
          <a-select-option
            v-for="(item, index) in config.groups"
            :key="index"
            :value="item.name"
            >{{ item.title }}</a-select-option
          >
        </a-select>
      </span>

      <span class="form-item_label">标题：</span>
      <a-input
        class="form-item_input"
        placeholder="控制器标题"
        v-model="title"
        @pressEnter="onTitleEnter"
      />
    </div>
  </div>
</template>

<script>
import { Form, Input, Tag, Select } from "ant-design-vue";
import AppSelect from "../AppSelect";
export default {
  components: {
    [Form.name]: Form,
    [Form.Item.name]: Form.Item,
    [Input.name]: Input,
    [Tag.name]: Tag,
    [Select.name]: Select,
    [Select.Option.name]: Select.Option,
    AppSelect
  },
  props: {
    currentAppKey: String,
    config: {
      type: Object,
      default: () => {}
    }
  },
  data() {
    return {
      appKey: "",
      group: "",
      title: ""
    };
  },

  created() {
    this.appKey = this.currentAppKey;
  },
  methods: {
    getData() {
      const { version, group, title } = this;
      return {
        version,
        group,
        title
      };
    },
    onTitleEnter() {
      this.$emit("titleEnter");
    },
    onVersionChange(val) {
      // this.version = val;
      this.$emit("versionChange", val);
    },
    onAppChange(val) {
      this.appKey = val;
      this.$emit("appkeyChange", val);
    }
  }
};
</script>

<style lang="less" scoped>
.form-wraper {
  .form-item {
    margin-bottom: 10px;
    .form-item_label {
      margin-left: 16px;
    }
    .form-item_input {
      width: 150px;
    }
  }
}
.action-title-tag {
  width: 50px;
  text-align: center;
  padding: 0 3px;
}
</style>
