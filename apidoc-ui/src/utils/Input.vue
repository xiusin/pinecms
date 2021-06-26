<template>
  <div class="data-grid-cell-edit" @click="check">
    <div v-if="editable" class="input-wrapper">
      <Input
        ref="input"
        :value="value"
        @change="handleChange"
        @blur="onBlur"
        @pressEnter="onEnter"
        :placeholder="placeholder"
      />
    </div>
    <div v-else class="text-wrapper">
      {{ value || " " }}
    </div>
  </div>
</template>
<script>
import { Input } from "ant-design-vue";

export default {
  components: {
    Input
  },
  props: {
    data: {
      type: [String, Number],
      default: ""
    },
    placeholder: String
  },
  // model: {
  //   prop: "data",
  //   event: "update"
  // },
  watch: {
    data() {
      this.value = this.data;
    }
  },
  data() {
    return {
      value: this.data,
      editable: false
    };
  },
  mounted() {},
  methods: {
    handleChange(e) {
      const value = e.target.value;
      this.value = value;
    },
    check() {
      if (this.editable) {
        return;
      }
      this.editable = true;
      this.$nextTick(function() {
        this.$refs.input.focus();
      });
    },
    edit() {
      this.editable = true;
    },
    updateData() {
      //修改前回调
      this.$emit("change", this.value);
      // this.$emit("update", this.value);
      this.editable = false;
    },
    onBlur() {
      this.updateData();
    },
    onEnter() {
      this.$emit("onEnter");
    }
  }
};
</script>

<style lang="less" scoped>
.data-grid-cell-edit {
  display: flex;
  margin: -5px;
  .text-wrapper {
    height: 32px;
    line-height: 32px;
    flex: 1;
    overflow: hidden;
    text-overflow: ellipsis;
    white-space: nowrap;
  }
  .input-wrapper {
    flex: 1;
  }
}
</style>
