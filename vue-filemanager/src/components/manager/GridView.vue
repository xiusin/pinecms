<template >
  <div class="grid" @mouseleave="isshow($event,false)">
    <el-checkbox-group v-model="checkedFiles" @change="handleCheckedFilesChange">
      <div
        class="grid-item"
        v-for="(directory, index) in directories"
        :key="`d-${index}`"
        :data-operate-id="`d-${index}`"
        :title="directory.basename"
        :class="{'active': checkSelect('directories', directory.path)}"
        @mouseenter="isshow"
        @contextmenu.prevent="contextMenu(directory, $event)"
        @dblclick.stop="selectDirectory(directory.path)"
      >
        <el-checkbox
          :style="{'display':checkSelect('directories', directory.path)||operateId==`d-${index}`?'block':'none'}"
          :label="directory"
          :key="`d-${index}`"
          @change="mutli('directories', directory)"
        >
          <span></span>
        </el-checkbox>
        <div @click="selectGridItem('directories', directory, $event)">
          <div class="item-icon">
            <i class="fa fa-folder fa-5x" style="color: #ffd659;" />
          </div>
          <div class="item-info">{{ directory.basename }}</div>
        </div>
      </div>
      <div
        class="grid-item"
        v-for="(file, index) in files"
        :key="`f-${index}`"
        :title="file.basename"
        :data-operate-id="`f-${index}`"
        :class="{'active': checkSelect('files', file.path)}"
        @mouseenter="isshow"
        @dblclick="selectAction(file.path, file.extension)"
        @contextmenu.prevent="contextMenu(file, $event)"
      >
        <el-checkbox
          :style="{'display':checkSelect('files', file.path)||operateId==`f-${index}`?'block':'none'}"
          :label="file"
          :key="`f-${index}`"
          @change="mutli('files', file)"
        >
          <span></span>
        </el-checkbox>
        <div @click="selectGridItem('files', file, $event)" class="grid-file">
          <div class="item-icon">
            <thumbnail v-if="thisImage(file.extension)" :disk="disk" :file="file"></thumbnail>
            <i
              v-else
              class="fa fa-5x"
              :class="extensionToIcon(file.extension)"
              :style="{color:extensionToColor(file.extension)}"
            ></i>
          </div>
          <div class="item-info">
            {{ `${file.filename}.${file.extension}` }}
            <br />
            {{ bytesToHuman(file.size) }}
          </div>
        </div>
      </div>
    </el-checkbox-group>
  </div>
</template>
<script>
import helper from "@/mixins/helper.js";
import managerHelper from "./mixins/manager";
import Thumbnail from "./Thumbnail.vue";
export default {
  props: {
    manager: { type: String, required: true }
  },
  mixins: [helper, managerHelper],
  data() {
    return {
      disk: "",
      //   checkedFiles: [],
      operateId: -1
    };
  },
  components: {
    Thumbnail
  },
  mounted() {
    this.disk = this.selectedDisk;
  },
  beforeUpdate() {
    // 如果磁盘改变
    if (this.disk !== this.selectedDisk) {
      this.disk = this.selectedDisk;
    }
  },
  computed: {
    /**
     * 图片扩展列表
     * @returns {*}
     */
    imageExtensions() {
      return this.$store.state.fm.settings.imageExtensions;
    }
  },
  // watch: {
  // isCheckedAll: {
  //   handler: function(newV) {
  //     this.checkedFiles = newV ? [...this.directories, ...this.files] : [];
  //     if (
  //       newV &&
  //       (this.selected["directories"].length == 0 ||
  //         this.selected["files"].length == 0)
  //     ) {
  //       this.setAllSelected(true);
  //     }
  //   }
  //   immediate: true
  // }
  // allChange: {
  //   handler: function(newV) {
  //     this.setAllSelected(newV);
  //   }
  //   immediate: true
  // }
  // },
  methods: {
    /**
     * 检查文件扩展
     * @param extension
     * @returns {Boolean}
     */
    thisImage(extension) {
      // 未找到符合的扩展名
      if (!extension) return false;
      return this.imageExtensions.includes(extension.toLowerCase());
    },
    handleCheckedFilesChange(value) {
      let checkedCount = value.length;
      let allFilesCount = this.files.length + this.directories.length;
      let isCheckedAll = checkedCount === allFilesCount;
      let isIndeterminate = checkedCount > 0 && checkedCount < allFilesCount;
      this.$store.commit(`fm/${this.manager}/setChAndIn`, {
        isIndeterminate,
        isCheckedAll
      });
      //   if (isCheckAll) {
      //     this.$emit("changeInd", isindeterminate, isCheckAll);
      //   }
      //   this.$emit("changeInd", isindeterminate, undefined);
    },
    isshow(e, show = true) {
      if (!show) {
        this.operateId = -1;
        return;
      }
      this.operateId = e.target.dataset.operateId;
    },
    mutli(type, item) {
      this.mutliGridSelected(type, item);
    }
  }
};
</script>
<style lang="scss" scoped>
.grid {
  position: relative;
  height: 100%;
  overflow-x: hidden !important;
  overflow-y: auto;
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
  .grid-item {
    box-sizing: border-box;
    position: relative;
    display: inline-block;
    width: 120px;
    // height: 146px;
    padding: 0.5rem;
    margin: 4px 0 0 6px;
    border-radius: 5px;
    user-select: none;
    text-align: center;
    border: 1px solid transparent;

    &.active {
      background-color: #f1f5fa;
      box-shadow: 0px 0px 5px rgba(241, 245, 250, 1);
      border: 1px solid rgba(144, 216, 255, 0.5);
    }

    &:not(.active):hover {
      background-color: #f1f5fa;
      box-shadow: 0px 0px 5px rgba(241, 245, 250, 1);
    }
    .item-icon {
      cursor: pointer;
      i::before {
        font-size: 4.5rem;
      }
    }
    .item-icon > i,
    .item-icon > figure > i {
      color: #6d757d;
    }
    .item-info {
      font-size: 1rem;
      overflow: hidden;
      text-overflow: ellipsis;
      white-space: nowrap;
    }
    .grid-file {
      display: flex;
      flex-direction: column;
      justify-content: space-between;
      height: 100%;
    }
  }
}
::v-deep .el-checkbox-group {
  width: 100%;
  height: 100%;
  // 一行内每个元素的高度由最高的元素决定
  display: flex;
  flex-wrap: wrap;
  align-content: flex-start;
  .el-checkbox:last-of-type {
    position: absolute;
    z-index: 999;
    left: 5px;
    top: 5px;
  }
}

::v-deep .el-checkbox {
  .el-checkbox__input {
    .el-checkbox__inner {
      border-radius: 50%;
      width: 19px;
      height: 19px;
      background-color: rgba(9, 170, 255, 0.2);
      border-color: rgba(9, 170, 255, 0.2);
      &:hover {
        background-color: rgba(9, 170, 255, 0.5);
        border-color: rgba(9, 170, 255, 0.5);
      }
      &::after {
        border-right: 2px solid #fff;
        border-bottom: 2px solid #fff;
        height: 10px;
        width: 4px;
        left: 5px;
        transform: rotate(45deg) scaleY(1);
      }
    }
  }

  .el-checkbox__input.is-checked {
    .el-checkbox__inner {
      background-color: rgba(9, 170, 255, 1);
      border-color: rgba(9, 170, 255, 1);
    }
  }
}
</style>