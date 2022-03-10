
<template >
  <el-container class="container">
    <Modal v-if="showModal" />
    <template v-if="showCropper">
      <Cropper :showCropper="showCropper" :imgSrc="imgSrc" />
    </template>

<!--    <el-header height="62px" class="navbar">-->
<!--      <el-row :gutter="20">-->
<!--        <el-col :xs="3" :sm="3" :md="3" :lg="3" :xl="3" class="logo">-->
<!--          <img src="../assets/cloudDisk.svg" />-->
<!--          <span>Pine云盘</span>-->
<!--        </el-col>-->
<!--        <el-col :xs="6" :sm="6" :md="10" :lg="16" :xl="18" class="drive">-->
<!--&lt;!&ndash;          <el-radio-group v-model="currentDrive" @change="setCurrentDrive">&ndash;&gt;-->
<!--&lt;!&ndash;            <el-radio-button label="public">公盘</el-radio-button>&ndash;&gt;-->
<!--&lt;!&ndash;            <el-radio-button label="private">私盘</el-radio-button>&ndash;&gt;-->
<!--&lt;!&ndash;          </el-radio-group>&ndash;&gt;-->
<!--          <div id="tp-weather-widget"></div>-->
<!--        </el-col>-->
<!--        <el-col :xs="15" :sm="15" :md="11" :lg="5" :xl="3" class="user">-->

<!--        </el-col>-->
<!--      </el-row>-->
<!--    </el-header>-->

    <el-container class="content">
      <el-aside width="200px">
        <div class="content-a" @click="refreshAll">
          <i class="el-icon-folder-opened">&nbsp;全部文件</i>
        </div>
        <FolderTree :parent-id="0" />
        <div class="info">
          <InfoBlock />
        </div>
      </el-aside>
      <el-main>
        <el-container style="position:relative;height: 100%;">
          <el-header
            height="100px"
            style="margin-left: -20px;padding-left: 20px;border-bottom: 1px solid #f2f6fd;"
          >
            <div class="content-h">
              <Navbar manager="left" />
              <el-row>
                <Breadcrumb manager="left" />
              </el-row>
              <el-row type="flex" justify="space-between">
                <el-col
                  :span="15"
                  @click.native="sortBy('name')"
                  :style="{background:isCheckedAll?'#fff':''}"
                >
                  <el-checkbox
                    :indeterminate="isIndeterminate"
                    v-model="isCheckedAll"
                    @change="handleCheckAllChange"
                  >
                    <span v-show="!selectedFilesCount">
                      文件名
                      <template v-if="sortSettings.field === 'name'">
                        <i
                          class="fas fa-sort-amount-down"
                          v-show="sortSettings.direction === 'down'"
                        />
                        <i class="fas fa-sort-amount-up" v-show="sortSettings.direction === 'up'" />
                      </template>
                    </span>
                    <span
                      v-show="selectedFilesCount"
                    >{{ `已选中${selectedFilesCount}个文件/文件夹,size:${selectedFilesSize}` }}</span>
                  </el-checkbox>
                </el-col>
                <template v-if="!selectedFilesCount">
                  <el-col :span="3" @click.native="sortBy('size')">
                    大小
                    <template v-if="sortSettings.field === 'size'">
                      <i
                        class="fas fa-sort-amount-down"
                        v-show="sortSettings.direction === 'down'"
                      />
                      <i class="fas fa-sort-amount-up" v-show="sortSettings.direction === 'up'" />
                    </template>
                  </el-col>
                  <el-col :span="2" @click.native="sortBy('type')">
                    类型
                    <template v-if="sortSettings.field === 'type'">
                      <i
                        class="fas fa-sort-amount-down"
                        v-show="sortSettings.direction === 'down'"
                      />
                      <i class="fas fa-sort-amount-up" v-show="sortSettings.direction === 'up'" />
                    </template>
                  </el-col>
                  <el-col :span="4" @click.native="sortBy('date')">
                    修改日期
                    <template v-if="sortSettings.field === 'date'">
                      <i
                        class="fas fa-sort-amount-down"
                        v-show="sortSettings.direction === 'down'"
                      />
                      <i class="fas fa-sort-amount-up" v-show="sortSettings.direction === 'up'" />
                    </template>
                  </el-col>
                </template>
              </el-row>
            </div>
          </el-header>
          <el-main class="all-files">
            <!-- <Notification /> -->
            <ContextMenu />

            <keep-alive>
              <TableView v-if="viewType === 'table'" manager="left" />
              <GridView v-else manager="left" />
            </keep-alive>
          </el-main>
        </el-container>
      </el-main>
    </el-container>
  </el-container>
</template>
<script>
import { mapState } from "vuex";
import EventBus from "../eventBus";
// Components
import Navbar from "@/components/blocks/Navbar.vue";
// import User from "@/components/blocks/User.vue";
import FolderTree from "@/components/tree/FolderTree.vue";
import InfoBlock from "@/components/blocks/InfoBlock.vue";
import ContextMenu from "@/components/blocks/ContextMenu.vue";
import Modal from "@/components/modals/Modal.vue";
// import Notification from "@/components/blocks/Notification.vue";
import Breadcrumb from "@/components/manager/Breadcrumb.vue";
import TableView from "@/components/manager/TableView.vue";
import GridView from "@/components/manager/GridView.vue";
import Cropper from "@/components/modals/views/Cropper.vue";
// Mixins
// import translate from "../mixins/translate";
import helper from "@/mixins/helper";
export default {
  name:'FileManager',
  mixins: [helper],
  data() {
    return {
      currentDrive: "public",
      selectedFilesCount: 0,
      showCropper: false,
      imgSrc: ""
    };
  },
  components: {
    FolderTree,
    InfoBlock,
    // Notification,
    ContextMenu,
    Navbar,
    Breadcrumb,
    TableView,
    GridView,
    Modal,
    Cropper,
    // User
  },
  methods: {
    /**
     * 选择磁盘
     * @param disk
     */
    selectDisk(disk) {
      if (this.selectedDisk !== disk) {
        this.$store.dispatch("fm/selectDisk", {
          disk,
          manager: this.activeManager
        });
      }
    },
    setCurrentDrive(currentDrive) {
      console.log("当前磁盘：", currentDrive);
      this.selectDisk(currentDrive);
    },
    handleCheckAllChange(val) {
      this.isIndeterminate = false;
      if (val) {
        this.$store.commit(`fm/${this.activeManager}/setAllSelected`, {
          dir: this.directories,
          file: this.files
        });
      } else {
        this.$store.commit(`fm/${this.activeManager}/removeAllSelected`);
      }
    },
    /**
     * 按字段排序
     * @param field
     */
    sortBy(field) {
      this.$store.dispatch(`fm/${this.activeManager}/sortBy`, {
        field,
        direction: null
      });
    },
    /**
     * 刷新
     */
    refreshAll() {
      this.$store.dispatch("fm/refreshAll");
    }
    /**
     * 添加axios请求拦截器
     */
    // requestInterceptor() {

    // },
    /**
     * 添加axios回复拦截器
     */
    // responseInterceptor() {

    // }
  },
  created() {
    // 初始化 Axios
    //this.$store.commit("fm/settings/initAxiosSettings");
    // this.requestInterceptor();
    // this.responseInterceptor();
    // 初始化设置
    this.$store.dispatch("fm/initializeApp");
  },
  mounted() {
    EventBus.$on("showCropper", (isshow, imgSrc) => {
      this.showCropper = isshow;
      this.imgSrc = imgSrc;
    });
    window.localStorage.setItem(
      "auto",
      JSON.stringify({
        a: this.$store.state.fm.autoLogin
      })
    );
    /**
     * 天气
     */
    (function(a, h, g, f, e, d, c, b) {
      b = function() {
        d = h.createElement(g);
        c = h.getElementsByTagName(g)[0];
        d.src = e;
        d.charset = "utf-8";
        d.async = 1;
        c.parentNode.insertBefore(d, c);
      };
      a["SeniverseWeatherWidgetObject"] = f;
      a[f] ||
        (a[f] = function() {
          (a[f].q = a[f].q || []).push(arguments);
        });
      a[f].l = +new Date();
      if (a.attachEvent) {
        a.attachEvent("onload", b);
      } else {
        a.addEventListener("load", b, false);
      }
    })(
      window,
      document,
      "script",
      "SeniverseWeatherWidget",
      "//cdn.sencdn.com/widget2/static/js/bundle.js?t=" +
        parseInt((new Date().getTime() / 100000000).toString(), 10)
    );
    window.SeniverseWeatherWidget("show", {
      flavor: "slim",
      location: "WTMKQ069CCJ7",
      geolocation: true,
      language: "zh-Hans",
      unit: "c",
      theme: "auto",
      token: "fd7ef8c0-c3c9-4c09-aab8-7b61c7cec928",
      hover: "enabled",
      container: "tp-weather-widget"
    });
  },
  watch: {
    selectedCount(newV) {
      this.selectedFilesCount = newV;
    }
  },
  computed: {
    ...mapState("fm", {
      // windowsConfig: state => state.settings.windowsConfig,
      activeManager: state => state.activeManager,
      showModal: state => state.modal.showModal
      // fullScreen: state => state.settings.fullScreen
    }),
    /**
     * 磁盘列表
     * @returns {Array}
     */
    disks() {
      return this.$store.getters["fm/diskList"];
    },
    /**
     * 选择的磁盘
     */
    selectedDisk() {
      return this.$store.state.fm[this.activeManager].selectedDisk;
    },
    /**
     * 当前所选目录下的所有文件
     * @returns {*}
     */
    files() {
      return this.$store.getters[`fm/${this.activeManager}/files`];
    },

    /**
     * 当前所选目录的文件夹列表
     * @returns {*}
     */
    directories() {
      return this.$store.getters[`fm/${this.activeManager}/directories`];
    },
    isCheckedAll: {
      get: function() {
        return this.$store.getters[`fm/${this.activeManager}/isCheckedAll`];
      },
      set: function(val) {
        this.$store.commit(`fm/${this.activeManager}/setIsCheckedAll`, val);
      }
    },
    isIndeterminate: {
      get: function() {
        return this.$store.getters[`fm/${this.activeManager}/isIndeterminate`];
      },
      set: function(val) {
        this.$store.commit(`fm/${this.activeManager}/setIsIndeterminate`, val);
      }
    },
    // lang() {
    //   return this.$store.state.fm.settings.translations["zh-CN"];
    // },
    /**
     * 排序设置
     * @returns {*}
     */
    sortSettings() {
      return this.$store.state.fm[`${this.activeManager}`].sort;
    },

    /**
     * 当前路径下选择的文件和文件夹的数量总和
     * @returns {*}
     */
    selectedCount() {
      return this.$store.getters[`fm/${this.activeManager}/selectedCount`];
    },
    /**
     * 计算选中的文件大小
     * @returns {*|string}
     */
    selectedFilesSize() {
      return this.bytesToHuman(
        this.$store.getters[`fm/${this.activeManager}/selectedFilesSize`]
      );
    },
    /**
     * 视图类型 - 网格 或 表格
     * @returns {String}
     */
    viewType() {
      return this.$store.state.fm[`${this.activeManager}`].viewType;
    }
  },
  destroyed() {
    // 重置状态
    this.$store.dispatch("fm/resetState");

    // 删除events
    EventBus.$off(["contextMenu", "showCropper", "addNotification"]);
  }
};
</script>
<style lang="scss" scoped>
.el-container.container {
  position: relative;
  height: 100vh;
  width: 100vw;

  .el-header,
  .el-main {
    background-color: #fff;
  }
  .el-aside {
    background: #f7f7f7;
  }
}

.el-row,
.el-col {
  height: 100%;
}

.el-container.container {
  height: inherit;
  &:after {
    content: "";
    display: block;
    clear: both;
  }
  .navbar {
    min-width: 1500px;
    position: absolute;
    top: 0;
    right: 0;
    left: 0;
    width: 100%;
    z-index: 41;
    box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);
    .logo {
      display: flex;
      align-items: center;
      justify-content: space-evenly;
      padding-left: 0 !important;
      padding: 0;
      cursor: pointer;
      span {
        font-size: 1.5rem;
      }
    }
    .drive{
      display: flex;
      justify-content: flex-start;
      align-items: center;
      #tp-weather-widget {
        margin-left: 40px;
      }
    }

  .user {
    justify-content: flex-start;
    align-items: center;
    text-align: right;
    #tp-weather-widget {
      margin-left: 40px;
    }
  }

  ::v-deep .el-radio-button {
      width: 50px;
      .el-radio-button__inner {
        background-color: transparent;
        color: #6c757d;
        border: 0px solid black;
        font-size: 1.1rem;
        font-weight: 500;
        box-shadow: none;
      }
    }
    ::v-deep .el-radio-button__orig-radio:checked + .el-radio-button__inner {
      background-color: transparent;
      color: #409eff;
      &::after {
        content: "";
        width: 20px;
        border-bottom: 2px solid #409eff;
        position: absolute;
        bottom: 0;
        left: 28px;
      }
    }
  }
  ::v-deep .el-container.content {
    /*height: calc(100% - 62px);*/
    /*margin-top: 62px;*/
    .el-aside {
      padding-top: 10px;
      position: relative;
      overflow: hidden;
      .content-a {
        display: flex;
        justify-content: center;
        align-items: center;
        font-size: 1.1rem;
        background: #eaeaea;
        color: #09aaff;
        height: 38px;
        cursor: pointer;
        .el-icon-folder-opened::before {
          font-size: 1.3rem;
        }
      }
      .info {
        height: 7%;
        display: grid;
        align-items: center;
        border-top: 1px solid rgba(188, 192, 194, 0.3);
      }
    }
    .el-main {
      padding-top: 10px;
      overflow-y: hidden;
      .el-header {
        padding: 0;

        .content-h {
          .el-row {
            color: #5b667b;
            white-space: nowrap;
          }
          & > .el-row:not(:last-child) {
            padding: 5px 0;
          }
          & > .el-row:last-child {
            margin-top: 5px 0;
            line-height: 28px;
            .el-col {
              font-size: 0.9rem;
            }
            .el-col:hover {
              background: #f4fbff;
              border-color: #cbedff;
              cursor: pointer;
            }
            i {
              color: #09aaff;
            }
          }
        }
      }
      .el-main.all-files {
        position: relative;
        padding: 0;
        margin-left: -20px;
      }
    }
  }
}
@media screen and (max-width: 768px) {
  .el-aside {
    display: none;
    transition: cubic-bezier(0.68, -0.55, 0.265, 1.55);
  }
  .el-container.content {
    .el-main {
      padding-left: 5px;
      padding-right: 5px;
    }
  }
}
</style>
