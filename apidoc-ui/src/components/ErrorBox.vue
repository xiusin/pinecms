<template>
  <div class="error-box">
    <div class="error-status">
      <div><Icon type="warning" /></div>
      <div>{{ error.status }}</div>
    </div>
    <div class="error-message">{{ error.message }}</div>
  </div>
</template>

<script>
export default {
  props: {
    error: {
      type: Object,
      default: () => {}
    }
  }
};
</script>
<style lang="less" scoped>
.error-box {
  padding: 50px 100px;
  text-align: center;
  .error-status {
    font-size: 50px;
    color: #ff4d4f;
  }
  .error-message {
    font-size: 16px;
    padding: 10px;
    background: #282c34;
    border-radius: 6px;
    color: #ddd;
    margin-bottom: 20px;
  }
}
</style>
