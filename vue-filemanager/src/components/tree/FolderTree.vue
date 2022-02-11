<template >
  <div class="tree">
    <ul class="tree-branch">
      <li v-for="(directory, index) in subDirectories" :key="index">
        <p
          class="unselectable"
          :class="{'selected': isDirectorySelected(directory.path)}"
          @click="selectDirectory(directory.path)"
        >
          <i
            v-if="directory.props.hasSubdirectories"
            @click.stop="showSubdirectories(
                        directory.path,
                        directory.props.showSubdirectories
                      )"
            :class="[arrowState(index)
                    ? 
                    'el-icon-folder-opened':
                    'el-icon-folder'
                   ]"
          />
          {{ directory.basename }}
        </p>

        <transition name="fade-tree">
          <folder-tree
            v-show="arrowState(index)"
            v-if="directory.props.hasSubdirectories"
            :parent-id="directory.id"
          ></folder-tree>
        </transition>
      </li>
    </ul>
  </div>
</template>
<script>
export default {
  name: "FolderTree",
  props: {
    parentId: {
      type: Number,
      required: true
    }
  },
  data() {
    return {};
  },
  methods: {
    /**
     * 是否选择了此目录
     * @param path
     * @returns {Boolean}
     */
    isDirectorySelected(path) {
      return this.$store.state.fm.left.selectedDirectory === path;
    },

    /**
     * 显示子目录
     * @returns {Boolean}
     * @param index
     */
    arrowState(index) {
      return this.subDirectories[index].props.showSubdirectories;
    },

    /**
     * 显示/隐藏 子目录
     * @param path
     * @param showState
     */
    showSubdirectories(path, showState) {
      if (showState) {
        // 隐藏
        this.$store.dispatch("fm/tree/hideSubdirectories", path);
      } else {
        // 显示
        this.$store.dispatch("fm/tree/showSubdirectories", path);
      }
    },

    /**
     * 加载选中的目录，显示文件
     * @param path
     */
    selectDirectory(path) {
      // 仅当未选择此路径时
      if (!this.isDirectorySelected(path)) {
        this.$store.dispatch("fm/left/selectDirectory", {
          path,
          history: true
        });
      }
    }
  },
  computed: {
    /**
     * 所选目录的所有子目录
     * @returns {*}
     */
    subDirectories() {
      let a = this.$store.getters["fm/tree/directories"].filter(
        item => item.parentId === this.parentId
      );
      return a;
    }
  }
};
</script>
<style lang="scss" scoped>
.tree {
  overflow: auto;
  height: 87%;
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
  .tree-branch {
    display: inline-block;
    font-size: 0.95rem;
    margin: 0;
    height: calc(100% - 110px);
    padding-left: 1.4rem;
    list-style: none;
    color: lightslategrey;
    li > p {
      margin-bottom: 0.1rem;
      padding: 0.4rem 0.4rem;
      white-space: nowrap;
      cursor: pointer;

      &:hover,
      &.selected {
        background-color: #f8f9fa;
      }
    }
  }
}
.fade-tree-enter-active,
.fade-tree-leave-active {
  transition: all 0.3s ease;
}
.fade-tree-enter,
.fade-tree-leave-to {
  transform: translateX(20px);
  opacity: 0;
}
</style>
