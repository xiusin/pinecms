<template >
  <el-dropdown @command="handleCommand">
    <span class="el-dropdown-link">
      <el-avatar src="https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png"></el-avatar>
      <span class="name">
        <em>{{ username }}</em>
        <i class="el-icon--right el-icon-arrow-down"></i>
      </span>
    </span>
    <el-dropdown-menu slot="dropdown">
      <el-dropdown-item command="personalData">个人资料</el-dropdown-item>
      <el-dropdown-item command="signOut">退出</el-dropdown-item>
    </el-dropdown-menu>
  </el-dropdown>
</template>
<script>
import GET from "@/http/get";
export default {
  data() {
    return {
      isEnter: false
    };
  },
  methods: {
    handleCommand(command) {
      switch (command) {
        case "personalData":
          this.$notify.info({
            title: "个人资料",
            message: "暂无"
          });
          break;
        case "signOut":
          GET.logout().then(resp => {
            if (resp.data.result.status == "success") {
              this.$store.commit("fm/setLoginStatus", false);
              console.log(
                "store.state.fm.isLogin",
                this.$store.state.fm.isLogin
              );
              this.$router.push({ path: "/login" });
            }
          });
          break;
        default:
          break;
      }
    }
  },
  computed: {
    username() {
      const name = this.$store.getters[`fm/getUsername`];
      console.log("name====", name);
      return name.nickname ?? name.username;
    }
  }
};
</script>
<style lang="scss" scoped>
.el-dropdown {
  ::v-deep span.el-dropdown-link {
    cursor: pointer;
    color: #409eff;
    &:hover {
      .name > i {
        transform: rotate(180deg);
        -webkit-transform: rotate(180deg);
        transform-origin: 50% 50% 0;
        transition: transform 0.3s ease-in 0s;
      }
    }
    .name {
      vertical-align: super;
      font-size: 1rem;
      font-weight: 500;
      color: #424e67;
    }
  }
}
</style>