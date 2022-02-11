<template>
  <div
    ref="contextMenu"
    v-if="menuVisible"
    :style="menuStyle"
    class="fm-context-menu"
    @blur="closeMenu"
    tabindex="-1"
  >
    <ul v-for="(group, index) in menu" :key="`g-${index}`" class="list-unstyled">
      <template v-for="(item, index) in group">
        <li
          :key="`i-${index}`"
          v-if="showMenuItem(item.name)"
          :class="{'eventnone': !authorIsUser&&acl&&(((item.name=='cut'||item.name=='delete' || item.name=='rename')&&selectedItemAcl)||(item.name=='paste' && selectedDisk == 'public' && selectedDir !== null))?true:false}"
          @click="menuAction(item.name)"
        >
          <i class="fa-fw" :class="item.icon"></i>
          {{ lang.contextMenu[item.name] }}
        </li>
      </template>
    </ul>
  </div>
</template>

<script>
/* eslint-disable no-param-reassign */
import EventBus from "../../eventBus";
import translate from "../../mixins/translate";
import contextMenu from "./mixins/contextMenu";
import contextMenuRules from "./mixins/contextMenuRules";
import contextMenuActions from "./mixins/contextMenuActions";

export default {
  name: "ContextMenu",
  mixins: [translate, contextMenu, contextMenuRules, contextMenuActions],
  data() {
    return {
      menuVisible: false,
      menuStyle: {
        top: 0,
        left: 0
      }
    };
  },
  mounted() {
    /**
     * Listen events
     * 'contextMenu'
     */
    EventBus.$on("contextMenu", event => this.showMenu(event));
  },
  computed: {
    /**
     * 上下文菜单
     * @returns {*}
     */
    menu() {
      return this.$store.state.fm.settings.contextMenu;
    },
    /**
     * ACL on/off
     */
    acl() {
      return this.$store.state.fm.settings.acl;
    },
    /**
     * 当前路径
     */
    selectedDir() {
      return this.$store.getters["fm/selectedDirectory"];
    }
  },
  methods: {
    /**
     * 显示上下文菜单
     * @param event
     */
    showMenu(event) {
      if (this.selectedItems.length) {
        this.menuVisible = true;
        EventBus.$emit("menuVisible", true);
        // 菜单获得焦点
        this.$nextTick(() => {
          this.$refs.contextMenu.focus();
          // 设置上下文菜单位置
          this.setMenu(event.pageY, event.pageX);
        });
      }
    },

    /**
     * 设置上下文菜单位置
     * @param top
     * @param left
     */
    setMenu(top, left) {
      // 得到父结点 (.fm-body)
      const el = this.$refs.contextMenu.parentNode;

      // 得到父元素的大小及其相对于视口的位置
      const elSize = el.getBoundingClientRect();

      // 获得相对于整个网页左上角定位的属性值，给top、left属性值加上当前的滚动位置
      const elY = window.pageYOffset + elSize.top;
      const elX = window.pageXOffset + elSize.left;

      // 计算初始的坐标
      let menuY = top - elY;
      let menuX = left - elX;

      //计算最大 X轴 或 Y轴 的坐标
      const maxY =
        elY + (el.offsetHeight - this.$refs.contextMenu.offsetHeight - 25);
      const maxX =
        elX + (el.offsetWidth - this.$refs.contextMenu.offsetWidth - 25);

      if (top > maxY) menuY = maxY - elY;
      if (left > maxX) menuX = maxX - elX;

      // 设置坐标
      this.menuStyle.top = `${menuY}px`;
      this.menuStyle.left = `${menuX}px`;
    },

    /**
     * 关闭上下文菜单
     */
    closeMenu() {
      this.menuVisible = false;
      EventBus.$emit("menuVisible", false);
    },

    /**
     * 显示指定的菜单项
     * @param name
     * @returns {Boolean}
     */
    showMenuItem(name) {
      if (Object.prototype.hasOwnProperty.call(this, `${name}Rule`)) {
        return this[`${name}Rule`]();
      }

      return false;
    },

    /**
     * 点击上下文菜单时调用指定的 菜单项action
     * @param name
     */
    menuAction(name) {
      if (Object.prototype.hasOwnProperty.call(this, `${name}Action`)) {
        this[`${name}Action`]();
      }
      this.closeMenu();
    }
  }
};
</script>

<style lang="scss" scoped>
.fm-context-menu {
  position: absolute;
  z-index: 9997;
  border: 1px solid #dde0e4;
  border-radius: 5px;
  box-shadow: 0 0 8px #ccc;
  background: #fff;
  font-size: 14px;
  color: #5b667b;
  &:focus {
    outline: none;
  }

  .list-unstyled {
    padding: 0;
    margin: 0;
    list-style: none;
    margin-bottom: 0;
    border-bottom: 1px solid rgba(0, 0, 0, 0.125);
  }

  ul > li {
    padding: 0.4rem 1rem;
  }
  & > ul:first-child > li:first-child {
    border-top-right-radius: 5px;
    border-top-left-radius: 5px;
  }
  & > ul:last-child > li:last-child {
    border-bottom-left-radius: 5px;
    border-bottom-right-radius: 5px;
  }
  ul > li:not(.disabled) {
    cursor: pointer;

    &:hover {
      background-color: #4281f4;

      color: #fff;
    }
  }
}
.eventnone {
  pointer-events: none;
  color: lightgray;
}
</style>
