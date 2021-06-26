<script>
import { Select, Icon } from "ant-design-vue";

export default {
  components: {
    [Select.name]: Select,
    [Select.OptGroup.name]: Select.OptGroup,
    [Select.Option.name]: Select.Option,
    [Icon.name]: Icon
  },
  props: {
    value: {
      type: String,
      default: ""
    },
    options: {
      type: Array,
      default: () => []
    }
  },

  data() {
    return {
      optionData: []
    };
  },
  created() {
    this.optionData = this.handleData(this.options);
  },
  methods: {
    handleData(data, pItem) {
      return data
        ? data.map(item => {
            if (pItem && pItem.folder) {
              item._key = `${pItem.folder}_${item.folder}`;
              item.checkedTitle = `${pItem.title}-${item.title}`;
            } else {
              item._key = `${item.folder}`;
              item.checkedTitle = item.title;
            }
            if (item.items) {
              item.items = this.handleData(item.items, item);
            }
            return item;
          })
        : [];
    },
    renderOptions(item) {
      if (item.items) {
        // 分组
        return this.renderGroup(item);
      }
      return this.renderOption(item);
    },
    renderGroup(item) {
      const itemArr = [];
      if (item.items && item.items.length) {
        item.items.forEach(p => itemArr.push(this.renderOptions(p)));
        return (
          <a-select-opt-group>
            <span slot="label">{item.title}</span>
            {itemArr}
          </a-select-opt-group>
        );
      }
      return null;
    },
    renderOption(item) {
      const lockIcon = item.hasPassword ? <a-icon type="lock" /> : "";
      return (
        <a-select-option value={item._key} label={item.checkedTitle}>
          <div class="app-select-option">
            <div class="app-select-option_title">{item.title}</div>
            <div class="app-select-option_icon">{lockIcon}</div>
          </div>
        </a-select-option>
      );
    },
    onChange(val) {
      this.$emit("change", val);
    }
  },
  render() {
    const { onChange, value } = this;
    const selectOptions = this.optionData.map(item => {
      return this.renderOptions(item);
    });
    return (
      <a-select
        style="width: 200px"
        {...{
          props: { "option-label-prop": "label", value: value },
          on: { change: onChange }
        }}
      >
        {selectOptions}
      </a-select>
    );
  }
};
</script>
<style lang="less" scoped>
.app-select-option {
  display: flex;
  &_title {
    flex: 1;
    white-space: nowrap;
    text-overflow: ellipsis;
    overflow: hidden;
    word-break: break-all;
  }
  &_icon {
    color: #999;
  }
}
</style>
