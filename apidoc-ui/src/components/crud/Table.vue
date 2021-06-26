<template>
  <div class="crud-model form-wraper">
    <div class="form-table">
      <table>
        <thead>
          <tr>
            <th>字段名</th>
            <th>注释</th>
            <th>类型</th>
            <th>
              <a-tooltip>
                <template slot="title">
                  小数位用,（逗号）分隔
                </template>
                长度
              </a-tooltip>
            </th>
            <th>默认值</th>
            <th v-if="validateRules && validateRules.length">
              <a-tooltip>
                <template slot="title">
                  数据验证规则
                </template>
                验证
              </a-tooltip>
            </th>
            <th>非null</th>
            <th>主键</th>
            <th>自增</th>

            <th>
              <a-tooltip>
                <template slot="title">
                  分页列表查询条件
                </template>
                查询
              </a-tooltip>
            </th>
            <th>
              <a-tooltip>
                <template slot="title">
                  分页列表返回的字段
                </template>
                列表
              </a-tooltip>
            </th>
            <th>
              <a-tooltip>
                <template slot="title">
                  明细查询返回的字段
                </template>
                明细
              </a-tooltip>
            </th>
            <th>
              <a-tooltip>
                <template slot="title">
                  新增提交的字段
                </template>
                新增
              </a-tooltip>
            </th>
            <th>
              <a-tooltip>
                <template slot="title">
                  编辑提交的字段
                </template>
                编辑
              </a-tooltip>
            </th>
            <th></th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="(item, index) in data" :key="index">
            <td>
              <a-input v-model="item.field" @blur="onFieldBlur($event, item)" />
            </td>
            <td style="width:150px;"><a-input v-model="item.desc" /></td>
            <td style="width:130px;">
              <a-select
                v-model="item.type"
                style="width:120px;"
                :options="fieldTypeOptions"
                show-search
                @change="onTypeChange($event, item)"
              />
            </td>
            <td style="width:70px;">
              <a-input v-model="item.length"></a-input>
            </td>

            <td style="width:100px;">
              <a-input v-model="item.default"></a-input>
            </td>
            <td
              v-if="validateRules && validateRules.length"
              style="width:110px;"
            >
              <a-select
                v-model="item.validate"
                style="width:100px;"
                :options="validateRules"
                show-search
                @change="onValidateRulesChange($event, item)"
              />
            </td>
            <td style="width:60px;">
              <a-checkbox v-model="item.not_null"></a-checkbox>
            </td>
            <td style="width:50px;">
              <a-checkbox v-model="item.main_key"></a-checkbox>
            </td>
            <td style="width:50px;">
              <a-checkbox v-model="item.incremental"></a-checkbox>
            </td>

            <td style="width:50px;">
              <a-checkbox v-model="item.query"></a-checkbox>
            </td>
            <td style="width:50px;">
              <a-checkbox v-model="item.list"></a-checkbox>
            </td>
            <td style="width:50px;">
              <a-checkbox v-model="item.detail"></a-checkbox>
            </td>
            <td style="width:50px;">
              <a-checkbox v-model="item.add"></a-checkbox>
            </td>
            <td style="width:50px;">
              <a-checkbox v-model="item.edit"></a-checkbox>
            </td>
            <td style="width:40px;">
              <a-button
                icon="close"
                type="link"
                style="color:#999"
                size="small"
                @click="removeRow(index)"
              ></a-button>
            </td>
          </tr>
        </tbody>
      </table>
      <a-button style="margin-top:5px;" @click="addField">+ 添加字段</a-button>
    </div>
  </div>
</template>

<script>
import {
  Input,
  Tag,
  InputNumber,
  Select,
  Checkbox,
  Button,
  Tooltip,
  message
} from "ant-design-vue";
import cloneDeep from "lodash/cloneDeep";

const fieldDefaultData = {
  field: "",
  desc: "",
  type: "",
  length: null,
  default: "",
  not_null: false,
  main_key: false,
  incremental: false,
  validate: "",
  query: false,
  list: true,
  detail: true,
  add: true,
  edit: true
};
export default {
  components: {
    [Input.name]: Input,
    [InputNumber.name]: InputNumber,
    [Select.name]: Select,
    [Checkbox.name]: Checkbox,
    [Button.name]: Button,
    [Tag.name]: Tag,
    [Tooltip.name]: Tooltip
  },
  props: {
    config: {
      type: Object,
      default: () => {}
    }
  },
  data() {
    return {
      fieldTypeOptions: [],
      data: []
    };
  },
  computed: {
    validateRules() {
      if (
        this.config &&
        this.config.crud &&
        this.config.crud.validate &&
        this.config.crud.validate.rules &&
        this.config.crud.validate.rules.length
      ) {
        return this.config.crud.validate.rules.map(item => {
          return {
            value: item.rule,
            label: item.name
          };
        });
      }
      return [];
    }
  },
  created() {
    if (
      this.config.crud &&
      this.config.crud.model &&
      this.config.crud.model.default_fields
    ) {
      this.data = cloneDeep(this.config.crud.model.default_fields);
    }
    if (
      this.config.crud &&
      this.config.crud.model &&
      this.config.crud.model.fields_types
    ) {
      this.fieldTypeOptions = cloneDeep(
        this.config.crud.model.fields_types
      ).map(p => {
        return {
          value: p,
          label: p
        };
      });
    }
    this.addField();
  },
  methods: {
    addField() {
      this.data.push(cloneDeep(fieldDefaultData));
    },
    getData() {
      const { data } = this;
      //首位可以是字母以及下划线。首位之后可以是字母，数字以及下划线。下划线后不能接下划线
      const fieldReg = /(^_([a-zA-Z0-9]_?)*$)|(^[a-zA-Z](_?[a-zA-Z0-9])*_?$)/;
      let error = false;
      let isMainKey = false;
      const list = data
        .filter(p => p.field && p.type)
        .map(p => {
          if (!fieldReg.test(p.field)) {
            error = `字段【${p.field}】字段名错误，请重新输入`;
          }
          if (p.main_key && p.incremental) {
            isMainKey = true;
          }
          return p;
        });
      if (error) {
        message.error(error);
        return false;
      }
      if (!isMainKey) {
        message.error("至少存在一个自增主键字段");
        return false;
      }
      return list;
    },
    onFieldBlur(e, item) {
      const { value } = e.target;
      if (value && value.indexOf("time") > -1 && !item.type) {
        item.type = "int";
        item.length = 10;
      } else if (value && !item.type) {
        item.type = "varchar";
        item.length = 255;
      }
    },
    onTypeChange(val, item) {
      if (val === "int") {
        item.length = 11;
      } else if (val === "varchar" && !item.length) {
        item.length = 255;
      } else if (val === "text" && !item.length) {
        item.length = "";
      }
    },
    removeRow(index) {
      this.data.splice(index, 1);
    },
    onValidateRulesChange(v, item) {
      if (v.indexOf("require") > -1) {
        item.not_null = true;
      }
    }
  }
};
</script>

<style lang="less" scoped>
.form-wraper {
  .form-item {
    margin-bottom: 10px;
    .form-item_text {
      margin-right: 10px;
    }
    .form-item_input {
      width: 150px;
    }
  }
}
.form-table {
  table {
    border: 1px solid #ddd;
    width: 100%;
    th,
    td {
      padding: 5px;
      text-align: center;
      border-bottom: 1px solid #ddd;
    }
    th {
      background: #fafafa;
    }
    td {
    }
  }
}
</style>
