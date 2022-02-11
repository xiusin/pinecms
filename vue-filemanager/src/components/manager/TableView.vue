<template >
  <div class="table">
    <el-checkbox-group v-model="checkedFiles" @change="handleCheckedFilesChange">
      <ul class="infinite-list" @mouseleave="isshow($event,false)">
        <li
          v-for="(directory, i) in directories"
          :key="`d-${i}`"
          :data-operate-id="`d-${i}`"
          @mouseenter="isshow"
          @contextmenu.prevent="contextMenu(directory, $event)"
          @dblclick.stop="selectDirectory(directory.path)"
          class="infinite-list-item"
        >
          <span class="mutli" :key="`md-${i}`" @click="mutliSelected('directories', directory)"></span>
          <el-checkbox :label="directory" :key="`d-${i}`">
            <el-row
              type="flex"
              class="file-info"
              justify="space-between"
              @click.native.capture="selectItem('directories', directory, $event)"
            >
              <el-col :span="14" style="position:relative;" :title="directory.basename">
                <i class="fa fa-folder fa-lg" style="color: #ffd659;" />
                {{ directory.basename }}
                <div
                  class="operate"
                  :style="{'display':!menuVisible&&checkedFiles.length<2&&operateId==`d-${i}`?'block':'none'}"
                >
                  <el-button title="分享" type="text">
                    <i class="el-icon-share"></i>
                  </el-button>
                  <!-- <el-button type="text">
                    <i class="el-icon-download"></i>
                  </el-button>-->
                  <el-button title="更多" type="text" @click.native="showMore(`d-${i}`)">
                    <i class="el-icon-more"></i>
                  </el-button>
                  <div class="more" v-if="moreId==`d-${i}`">
                    <ul>
                      <li @click="copyAction">复制</li>
                      <li
                        @click="cutAction"
                        :class="{'eventnone':acl&&selectedItemAcl?true:false}"
                      >剪切</li>
                      <li
                        @click="renameAction"
                        :class="{'eventnone':acl&&selectedItemAcl?true:false}"
                      >重命名</li>
                      <li
                        @click="deleteAction"
                        :class="{'eventnone':acl&&selectedItemAcl?true:false}"
                      >删除</li>
                    </ul>
                  </div>
                </div>
              </el-col>
              <el-col :span="2">
                <p style="margin-left: 56%;">-</p>
              </el-col>
              <el-col :span="3">
                <p style="margin-left: 65%;">目录</p>
              </el-col>
              <el-col :span="5">
                <p style="margin-left: 14%;">{{ timestampToDate(directory.timestamp) }}</p>
              </el-col>
            </el-row>
          </el-checkbox>
        </li>
        <li
          v-for="(file, i) in files"
          :key="`f-${i}`"
          :data-operate-id="`f-${i}`"
          @mouseenter="isshow"
          @contextmenu.prevent="contextMenu(file, $event)"
          @dblclick.stop="selectAction(file.path, file.extension)"
          class="infinite-list-item"
        >
          <span class="mutli" :key="`mf-${i}`" @click="mutliSelected('files', file)"></span>
          <el-checkbox :label="file" :key="`f-${i}`">
            <el-row
              type="flex"
              class="file-info"
              justify="space-between"
              @click.native.capture="selectItem('files', file, $event)"
            >
              <el-col :span="14" style="position:relative;" :title="file.filename">
                <i
                  class="fa fa-lg"
                  :class="extensionToIcon(file.extension)"
                  :style="{color:extensionToColor(file.extension)}"
                />
                {{ file.filename ? file.filename : file.basename }}
                <div
                  class="operate"
                  :style="{'display':!menuVisible&&checkedFiles.length<2&&operateId==`f-${i}`?'block':'none'}"
                >
                  <el-button title="分享" type="text">
                    <i class="el-icon-share"></i>
                  </el-button>
                  <el-button title="下载" type="text" @click.native="downloadAction">
                    <i class="el-icon-download"></i>
                  </el-button>
                  <el-button type="text" title="更多" @click.native="showMore(`f-${i}`)">
                    <i class="el-icon-more"></i>
                  </el-button>
                  <div class="more" v-if="moreId==`f-${i}`">
                    <ul>
                      <li @click="copyAction">复制</li>
                      <li
                        @click="cutAction"
                        :class="{'eventnone':acl&&selectedItemAcl?true:false}"
                      >剪切</li>
                      <li
                        @click="renameAction"
                        :class="{'eventnone':acl&&selectedItemAcl?true:false}"
                      >重命名</li>
                      <li
                        @click="deleteAction"
                        :class="{'eventnone':acl&&selectedItemAcl?true:false}"
                      >删除</li>
                    </ul>
                  </div>
                </div>
              </el-col>
              <el-col :span="3">
                <p style="margin-left: 24%;">{{ bytesToHuman(file.size) }}</p>
              </el-col>
              <el-col :span="2">
                <p style="margin-left: 50%;">{{ file.extension }}</p>
              </el-col>
              <el-col :span="5">
                <p style="margin-left: 14%;">{{ timestampToDate(file.timestamp) }}</p>
              </el-col>
            </el-row>
          </el-checkbox>
        </li>
      </ul>
    </el-checkbox-group>
  </div>
</template>
<script>
import helper from "@/mixins/helper";
import managerHelper from "@/components/manager/mixins/manager";
import EventBus from "@/eventBus.js";
export default {
  props: {
    manager: { type: String, required: true }
  },
  mixins: [helper, managerHelper],
  data() {
    return {
      operateId: -1,
      checkdId: "",
      menuVisible: false,
      moreId: ""
    };
  },
  computed: {
    check() {
      return this.checkdId;
    },
    /**
     * 上下文菜单
     * @returns {*}
     */
    menu() {
      return this.$store.state.fm.settings.contextMenu;
    },
    /**
     * 选中的文件或文件夹
     * @returns {*}
     */
    selectedItems() {
      return this.$store.getters["fm/selectedItems"];
    },
    /**
     * ACL on/off
     */
    acl() {
      return this.$store.state.fm.settings.acl;
    },
    /**
     * 判断选中的文件的acl
     * @returns
     */
    selectedItemAcl() {
      return this.selectedItems.every(function(item) {
        return item.acl === 1;
      });
    }
  },
  // watch: {
  // isCheckedAll: {
  // handler: function(newV) {
  // if (newV) {
  // 全选，添加所有文件
  // } else {
  // 不全选则清空
  // this.$store.commit(`fm/${this.manager}/resetSelected`);
  // }
  // this.checkedFiles = newV ? [...this.directories, ...this.files] : [];
  // if (
  //   newV &&
  //   (this.selected["directories"].length == 0 ||
  //     this.selected["files"].length == 0)
  // ) {
  //   this.setAllSelected(true);
  // }
  // }
  // }
  // allChange: {
  // handler: function(newV) {
  //   this.setAllSelected(newV);
  // }
  // immediate: true
  // }
  // },
  mounted() {
    EventBus.$on("menuVisible", visible => {
      this.menuVisible = visible;
    });
  },
  methods: {
    handleCheckedFilesChange(value) {
      let checkedCount = value.length;
      let allFilesCount = this.files.length + this.directories.length;
      let isCheckedAll = checkedCount === allFilesCount;
      let isIndeterminate = checkedCount > 0 && checkedCount < allFilesCount;
      this.$store.commit(`fm/${this.manager}/setChAndIn`, {
        isIndeterminate,
        isCheckedAll
      });
    },
    // load() {
    // this.count += 2;
    // },
    isshow(e, show = true) {
      if (!show) {
        this.operateId = -1;
        return;
      }
      this.operateId = e.target.dataset.operateId;
      this.moreId = "";
    },
    /**
     * 显示"更多"菜单
     */
    showMore(moreId) {
      this.moreId = moreId;
    },
    /**
     * 按字段排序
     * @param field
     */
    sortBy(field) {
      this.$store.dispatch(`fm/${this.manager}/sortBy`, {
        field,
        direction: null
      });
    },
    /**
     * 复制选中的文件项，添加到剪切板
     */
    copyAction() {
      this.$store.dispatch("fm/toClipboard", "copy");
    },
    /**
     * 剪切选中的文件项，添加到剪切板
     */
    cutAction() {
      this.$store.dispatch("fm/toClipboard", "cut");
    },

    /**
     * 重命名选中的文件项
     */
    renameAction() {
      // 显示命名模块
      this.$store.commit("fm/modal/setModalState", {
        modalName: "Rename",
        show: true
      });
    },
    /**
     * 删除选中的文件项
     */
    deleteAction() {
      // 显示删除模块
      this.$store.commit("fm/modal/setModalState", {
        modalName: "Delete",
        show: true
      });
    }
  }
};
</script>
<style lang="scss" scoped>
.table {
  height: 100%;

  .el-checkbox-group {
    height: 100%;
    ul {
      height: 100%;
      overflow-x: hidden !important;
      overflow-y: auto;
      padding: 0;
      // overflow-y: scroll !important;
      &::-webkit-scrollbar {
        width: 5px;
      }
      &::-webkit-scrollbar-track {
        background-color: #f5f5f5;
        -webkit-box-shadow: inset 0 0 3px rgba(0, 0, 0, 0.1);
        border-radius: 5px;
      }
      &::-webkit-scrollbar-thumb {
        background-color: rgba(140, 199, 181, 0.8);
        border-radius: 5px;
      }
      &::-webkit-scrollbar-button {
        display: none;
        background-color: #eee;
      }
      &::-webkit-scrollbar-corner {
        background-color: black;
      }
      li {
        position: relative;
        height: 45px;
        line-height: 45px;
        border-bottom: 1px solid #f2f6fd;
        white-space: nowrap;
        text-overflow: ellipsis;
        .mutli {
          position: absolute;
          left: 0px;
          top: 0px;
          z-index: 9999;
          width: 45px;
          height: 45px;
        }
        ::v-deep .el-checkbox {
          width: calc(100% - 20px);
          padding-left: 20px;
          &:hover {
            background: #f4fbff;
            border-color: #cbedff;
            border-bottom: 1px solid #daebfe;
            // border-top: 1px solid #daebfe;
          }
          .el-checkbox__label {
            width: 98%;
            height: 90%;
            line-height: 3;
            .operate {
              position: absolute;
              right: 6px;
              top: 0px;
              .el-button {
                font-size: 1.2rem;
              }
              .more {
                border: 1px solid rgba(9, 170, 255, 0.2);
                text-align: center;
                position: absolute;
                z-index: 9999;
                right: -32px;
                top: 32px;
                width: 84px;
                border-radius: 5px;
                cursor: pointer;
                ul {
                  list-style: none;
                  color: #09aaff;
                  li {
                    background: #fff;
                    height: 27px !important;
                    line-height: 30px;
                    &:hover {
                      background: rgb(228, 238, 254, 0.5);
                    }
                  }
                }
              }
            }
          }
        }
        ::v-deep .el-checkbox.is-checked {
          background: #f4fbff;
          border-color: #cbedff;
          border-bottom: 1px solid #daebfe;
          &::before {
            content: "";
            border-top: 1px solid #daebfe;
            position: absolute;
            top: -1px;
            left: -20px;
            display: block;
            width: 100%;
            z-index: 1;
            visibility: visible;
          }
        }
        .file-info {
          .el-col:first-child {
            padding-right: 120px;
          }
          .el-col {
            overflow: hidden;
            white-space: nowrap;
            text-overflow: ellipsis;
          }
          .el-col:first-child {
            overflow: visible;
          }
          .el-col:not(:first-child) {
            text-align: center;
          }
          p {
            font-size: 14px;
            margin: 0;
            width: fit-content;
          }
        }
      }
    }
  }
}
.eventnone {
  pointer-events: none;
  color: lightgray;
}
</style>