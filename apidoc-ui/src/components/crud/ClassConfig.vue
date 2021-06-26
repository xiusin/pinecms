<template>
  <div class="form-wraper">
    <!-- <a-row :gutter="10" style="margin-bottom:5px;">
      <a-col :span="colspan" v-for="item in crudConfigList" :key="item.name"> -->
    <div class="crud-file-list">
      <div
        v-for="item in crudConfigList"
        :key="item.name"
        class="crud-file-item"
      >
        <div>
          <label style="color:#999">{{ item.name }}</label>
          <a-input
            :ref="`input_${item.name}`"
            class="input-link"
            v-model="item.class_name"
            :placeholder="`请输入${item.name}文件名`"
            @blur="onClassNameChange($event, item)"
            @pressEnter="onClassNameChange($event, item)"
          />
        </div>
        <p><label>Path：</label>{{ item.path }}</p>
      </div>
      <!-- </a-col>
    </a-row> -->
    </div>
    <a-alert
      style="margin:5px 0;"
      message="文件命名：可输入首字母大写的字母+数字组合。并确保以上文件不存在！"
      banner
    />
  </div>
</template>

<script>
import { Row, Col, Input, Alert, message } from "ant-design-vue";
import { treeTransArray, getTreeMaxlevel } from "../../utils/utils";
import cloneDeep from "lodash/cloneDeep";

export default {
  components: {
    [Row.name]: Row,
    [Col.name]: Col,
    [Input.name]: Input,
    [Alert.name]: Alert
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
      crudConfigList: []
    };
  },
  computed: {
    colspan() {
      return 24 / this.crudConfigList.length;
    }
  },
  watch: {
    currentAppKey() {
      this.renderNamespace();
    }
  },

  created() {
    this.init();
    this.renderNamespace();
  },
  methods: {
    init() {
      const { config } = this;
      let crudConfigList = [];
      for (const key in config.crud) {
        crudConfigList.push({
          ...config.crud[key],
          name: key,
          class_name: ""
        });
      }
      this.crudConfigList = crudConfigList;
    },
    onClassNameChange(e, item) {
      const { value } = e.target;
      const { crudConfigList } = this;
      if (item.name === "controller") {
        this.crudConfigList = crudConfigList.map(p => {
          if (p.name != "controller") {
            p.class_name = value;
          }
          return p;
        });
      } else {
        this.crudConfigList = crudConfigList.map(p => {
          if (p.name != item.name && !p.class_name) {
            p.class_name = value;
          }
          return p;
        });
      }
    },

    getData() {
      const { crudConfigList } = this;
      const configJson = {};
      let error = false;
      for (let i = 0; i < crudConfigList.length; i++) {
        const item = crudConfigList[i];
        configJson[item.name] = item;
        // 验证文件名必填
        if (item.class_name) {
          const reg = /^[A-Z]{1}[A-Za-z0-9]{1,32}$/;
          if (!reg.test(item.class_name)) {
            message.error(`${item.class_name}文件名不合法`);
            error = true;
          }
        } else {
          message.error(`请填写${item.name}文件名`);
          error = true;
        }
      }
      if (error) {
        return false;
      }
      return configJson;
    },
    focus() {
      const controllerInput = this.$refs.input_controller;
      if (controllerInput) {
        controllerInput[0].focus();
      }
    },
    renderNamespace() {
      const { config } = this;

      let currentApps = [];
      if (this.config && this.config.apps) {
        const list = treeTransArray(this.config.apps, "items");
        if (this.currentAppKey.indexOf("_") > -1) {
          const keyArr = this.currentAppKey.split("_");
          for (let i = 0; i < keyArr.length; i++) {
            const key = keyArr[i];
            const find = list.find(p => p.folder === key);
            if (find) {
              currentApps.push(find);
            }
          }
        } else if (this.currentAppKey) {
          const find = list.find(p => p.folder === this.currentAppKey);
          if (find) {
            currentApps.push(find);
          }
        }
      }
      const maxTreelevel = getTreeMaxlevel(this.config.apps, "items");

      this.crudConfigList = this.crudConfigList.map(p => {
        let currentPath =
          config.crud[p.name] && config.crud[p.name].path
            ? cloneDeep(config.crud[p.name].path)
            : "";
        if (maxTreelevel) {
          for (let i = 0; i < maxTreelevel; i++) {
            if (i < currentApps.length) {
              const item = currentApps[i];
              for (const key in item) {
                const keyStr = `\${app[${i}].${key}}`;
                if (currentPath.indexOf(keyStr) > -1) {
                  currentPath = currentPath.replace(keyStr, item[key]);
                }
              }
            } else if (currentApps && currentApps.length) {
              const item = currentApps[0];
              for (const key in item) {
                const keyStr = `\${app[${i}].${key}}`;
                if (currentPath.indexOf(keyStr) > -1) {
                  currentPath = currentPath.replace(keyStr, "");
                }
              }
            }
          }
        }
        // 去除多余反斜杠
        if (currentPath.indexOf("\\\\") > -1) {
          currentPath = currentPath.replace("\\\\", "\\");
        }
        // 去除最后一个反斜杠
        if (currentPath.charAt(currentPath.length - 1) == "\\") {
          currentPath = currentPath.substr(0, currentPath.length - 1);
        }
        p.path = currentPath;
        return p;
      });
    }
  }
};
</script>

<style lang="less" scoped>
.crud-file-list {
  display: flex;
  .crud-file-item {
    // border: 1px solid #fafafa;
    background: #fafafa;
    padding: 10px;
    flex: 1;
    margin: 0 3px;
    p {
      margin-bottom: 5px;
    }
  }
}
.input-link {
  border: none;
  border-bottom: 1px solid #ddd;
  background: none;
  font-size: 16px;
  font-weight: 500;
  &:focus {
    box-shadow: none;
  }
}
</style>
