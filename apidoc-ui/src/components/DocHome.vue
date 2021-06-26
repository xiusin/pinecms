<template>
  <div>
    <div class="api-home-content">
      <h1>{{ config.title }}</h1>
      <div v-if="config.desc" class="desc">{{ config.desc }}</div>
      <div>{{ config.copyright }}</div>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    apiData: {
      type: Object,
      default: () => {}
    },

    config: {
      type: Object,
      default: () => {}
    }
  }
};
</script>

<style lang="less" scoped>
.api-home-content {
  text-align: center;
  padding: 100px 0;
  h1 {
    font-size: 40px;
  }
  .desc {
    margin-bottom: 20px;
  }
}
</style>
