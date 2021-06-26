<script>
import {
  Menu,
  Tag,
  Input,
  Select,
  Icon,
  Button,
  Tooltip
} from "ant-design-vue";
import cloneDeep from "lodash/cloneDeep";

// 是否满足过滤条件
function hasKeyword(item, keyword, tags = []) {
  let hasTag = false;
  if (tags.length) {
    hasTag = tags.some(value => {
      return item.tag && item.tag.indexOf(value) > -1;
    });
  }
  const hasKeyword =
    keyword &&
    (item.title.indexOf(keyword) > -1 ||
      (item.url && item.url.indexOf(keyword) > -1));

  if (keyword && tags.length && hasKeyword && hasTag) {
    return true;
  } else if (keyword && !tags.length && hasKeyword) {
    return true;
  } else if (!keyword && tags.length && hasTag) {
    return true;
  }

  return false;
}

// 过滤菜单
function filterMenu(menus, keyword, _vm) {
  const newMenus = menus.filter(item => {
    let res = hasKeyword(item, keyword, _vm.currentTags);
    if (item.children && item.children.length) {
      item.children = filterMenu(item.children, keyword, _vm);
      _vm.currentOpenKeys.push(item.menu_key);
      if (item.children && item.children.length) {
        res = true;
      }
    }
    return res;
  });
  return newMenus;
}

export default {
  components: {
    Tag,
    Menu,
    MenuSubMenu: Menu.SubMenu,
    MenuItem: Menu.Item,
    MenuItemGroup: Menu.ItemGroup,
    InputSearch: Input.Search,
    Select,
    SelectOption: Select.Option,
    SelectOptGroup: Select.OptGroup,
    Icon,
    Button,
    Tooltip
  },
  props: {
    apiData: {
      type: Array,
      default: () => []
    },
    groups: {
      type: Array,
      default: () => []
    },
    tags: {
      type: Array,
      default: () => []
    },
    sideSize: {
      type: String,
      default: ""
    },
    docs: {
      type: Array,
      default: () => []
    },
    config: {
      type: Object,
      default: () => {}
    },
    device: {
      type: String,
      default: "xl"
    }
  },
  data() {
    return {
      currentGroupName: 0,
      menuData: [],
      openKeys: [],
      currentOpenKeys: [],
      currentTags: [],
      keyword: "",
      selectedKeys: [],
      envCofnig: {}
    };
  },
  watch: {
    apiData() {
      this.onSearch();
    },
    docs() {
      this.onSearch();
    }
  },

  created() {
    // eslint-disable-next-line no-undef
    this.envCofnig = config;
    this.onSearch();
  },
  methods: {
    onMenuClick(data) {
      this.selectedKeys = [data.menu_key];
      this.$emit("change", data);
    },
    setSelectedKeys(keys) {
      this.selectedKeys = keys;
    },
    getMenuData() {
      return this.menuData;
    },
    renderItem(menu) {
      if (!menu.hidden) {
        if (menu.items) {
          // 分组
          return this.renderGroup(menu);
        }
        if (menu.children) {
          return this.renderSubMenu(menu);
        }
        return this.renderMenuItem(menu);
      }
      return null;
    },
    renderGroup(menu) {
      const itemArr = [];
      if (menu.items && menu.items.length) {
        menu.items.forEach(item => itemArr.push(this.renderItem(item)));
        const title = menu.title == "全部" ? "未分组" : menu.title;
        return (
          <MenuItemGroup title={title} {...{ key: menu.name }}>
            {itemArr}
          </MenuItemGroup>
        );
      }
      return null;
    },
    renderSubMenu(menu) {
      const itemArr = [];
      if (menu.children && menu.children.length) {
        menu.children.forEach(item => itemArr.push(this.renderItem(item)));
      }
      let controller = "";
      if (menu.controller && this.envCofnig.MENU.SHOW_CONTROLLER_CLASS) {
        controller = <b style="margin-right:10px;">{menu.controller}</b>;
      }
      let apiActionButton = "";

      return (
        <MenuSubMenu {...{ key: menu.menu_key }}>
          <div slot="title" class="menu-sub">
            <Icon type="folder-open" />
            {controller}
            <span>{menu.title}</span>
            {apiActionButton}
          </div>
          {itemArr}
        </MenuSubMenu>
      );
    },
    renderMenuItem(menu) {
      // eslint-disable-next-line no-undef
      if (menu && menu.url && menu.method) {
        // 接口
        let method = "";
        if (this.envCofnig.MENU.SHOW_API_METHOD) {
          let tagColor = "";
          switch (menu.method) {
            case "GET":
              tagColor = "#87d068";
              break;
            case "POST":
              tagColor = "#2db7f5";
              break;
            case "PUT":
              tagColor = "#ff9800";
              break;
            case "DELETE":
              tagColor = "#ff4d4f";
              break;
            default:
              tagColor = "#ccc";
              break;
          }
          let methodTags = "";
          if (menu.method.indexOf(",") > -1) {
            const tags = menu.method.split(",");
            methodTags = tags.map((p, i) => {
              if (i < 4) {
                return <span class={`method-item ${p} method-item-${i}`} />;
              }
            });
            const methodText = `${tags.length}`;
            method = (
              <span
                title={menu.method}
                class={`action-title-tag method-multiple-tag method-num-${
                  tags.length > 4 ? 4 : tags.length
                }`}
              >
                {methodTags}
                <span class={`method-item empty`} />
                {methodText}
              </span>
            );
          } else {
            method = (
              <Tag class="action-title-tag" color={tagColor}>
                <span>{menu.method}</span>
              </Tag>
            );
          }
        }
        let url = "";
        if (this.envCofnig.MENU.SHOW_API_URL) {
          url = <span style="margin-left:10px;">{menu.url}</span>;
        }

        return (
          <MenuItem
            {...{
              key: menu.menu_key,
              on: {
                click: () => {
                  this.onMenuClick(menu);
                }
              }
            }}
          >
            <span class="action-title">
              <div class="action-title_wraper">
                {method}
                {menu.title}
                {url}
              </div>
            </span>
          </MenuItem>
        );
      } else if (menu && menu.type === "md") {
        //doc文档
        return (
          <MenuItem
            {...{
              key: menu.menu_key,
              on: {
                click: () => {
                  this.onMenuClick(menu);
                }
              }
            }}
          >
            <span class="action-title">
              <div class="action-title_wraper">
                <Icon type="file-text" />
                {menu.title}
              </div>
            </span>
          </MenuItem>
        );
      }
      return (
        <MenuItem {...{ key: menu.menu_key }}>
          <span>{menu.title}</span>
        </MenuItem>
      );
    },
    handleSort(list) {
      return list.sort(function(a, b) {
        a.sort = a.sort ? a.sort : 999;
        b.sort = b.sort ? b.sort : 999;
        if (a.sort < b.sort) {
          return -1;
        } else if (a.sort == b.sort) {
          return 0;
        } else {
          return 1;
        }
      });
    },
    handleGroupMenuData(data) {
      const { groups, currentGroupName } = this;
      if (!(groups && groups.length)) {
        return data;
      }
      const apiData = cloneDeep(data);
      if (currentGroupName) {
        // 指定分组
        let list = apiData.filter(p => p.group == currentGroupName);
        if (list && list.length) {
          list = this.handleSort(list);
        }
        return list;
      }
      const groupNames = groups.map(p => p.name);
      let groupData = groups.map(item => {
        if (item.name === 0) {
          item.items = apiData.filter(p => !groupNames.includes(p.group));
        } else {
          item.items = apiData.filter(p => p.group == item.name);
        }
        if (item.items && item.items.length) {
          item.items = this.handleSort(item.items);
        }
        return item;
      });
      return groupData;
    },
    handleDocsData(docsData, key = "") {
      const { config, currentGroupName } = this;
      let data = null;
      if (
        docsData &&
        docsData.length &&
        (!currentGroupName || currentGroupName === "markdown_doc")
      ) {
        let items = [];
        if (key) {
          items = filterMenu(docsData, key, this);
        } else {
          items = docsData;
        }
        // 过滤分组
        data = {
          title:
            config.docs && config.docs.menu_title
              ? config.docs.menu_title
              : "文档",
          items: items
        };
      }
      return data;
    },
    onSearch() {
      const apiData = cloneDeep(this.apiData);
      const docsData = cloneDeep(this.docs);
      let menuData = [];
      this.currentOpenKeys = [];
      if (this.keyword || this.currentTags.length) {
        const filterData = filterMenu(apiData, this.keyword, this);
        // 分组
        const groupData = this.handleGroupMenuData(filterData);
        menuData = groupData;
      } else {
        // 无搜索条件,显示所有
        const groupData = this.handleGroupMenuData(apiData);
        menuData = groupData;
      }

      const docsList = this.handleDocsData(docsData, this.keyword);
      if (docsList) {
        this.menuData = [docsList, ...menuData];
        this.openKeys = this.currentOpenKeys;
      } else {
        this.menuData = menuData;
      }
    },
    renderGroupsSelect() {
      const { groups, currentGroupName } = this;
      if (!(groups && groups.length)) {
        return null;
      }
      const that = this;
      const selectOptions = groups.map(item => {
        return <SelectOption value={item.name}>{item.title}</SelectOption>;
      });
      const selectProps = {
        props: {
          value: currentGroupName,
          allowClear: true,
          placeholder: "选择分组"
        },
        on: {
          change: val => {
            that.currentGroupName = val;
            that.onSearch();
          }
        },
        style: {
          width: "140px",
          minWidth: "120px",
          marginRight: "10px"
        }
      };
      return <Select {...selectProps}>{selectOptions}</Select>;
    },
    renderTagsSelect() {
      const { tags, currentTags } = this;
      if (!(tags && tags.length)) {
        return null;
      }
      const that = this;
      const selectOptions = tags.map(item => {
        return <SelectOption value={item}>{item}</SelectOption>;
      });
      const selectProps = {
        props: {
          value: currentTags,
          allowClear: true,
          mode: "multiple",
          maxTagCount: 3,
          placeholder: "Tags筛选"
        },
        on: {
          change: val => {
            that.currentTags = val;
            that.onSearch();
          }
        },
        style: {
          width: "100%",
          minWidth: "120px",
          marginRight: "10px"
        }
      };
      return (
        <div class="flex" style="margin-top:5px;overflow:hidden;">
          <div style="line-height:32px;">Tags：</div>
          <div class="flex-item">
            <Select {...selectProps}>{selectOptions}</Select>
          </div>
        </div>
      );
    },
    onCrudClick() {
      this.$emit("showCrud");
    },
    onOpenChange(openKeys) {
      this.openKeys = openKeys;
    }
  },
  render() {
    const {
      renderGroupsSelect,
      onCrudClick,
      config,
      device,
      onOpenChange,
      openKeys,
      renderTagsSelect,
      onSearch,
      keyword,
      selectedKeys
    } = this;
    const menuTree = this.menuData.map(item => {
      return this.renderItem(item);
    });
    let createCrudButton = "";
    if (
      config.crud &&
      config.crud.model &&
      config.debug &&
      device !== "mobile"
    ) {
      createCrudButton = (
        <Tooltip placement="top">
          <template slot="title">快速创建Crud接口</template>
          <Button
            style={{ padding: "0 8px", marginLeft: "10px" }}
            {...{ on: { click: onCrudClick } }}
          >
            <Icon type="plus" />
          </Button>
        </Tooltip>
      );
    }

    const menuProps = {
      props: {
        openKeys,
        selectedKeys
      },
      on: {
        openChange: onOpenChange
      }
    };

    const searchInputProps = {
      props: {
        value: keyword
      },
      on: {
        search: onSearch,
        change: e => {
          const { value } = e.target;
          this.keyword = value;
          if (!value) {
            onSearch();
          }
        }
      }
    };

    return (
      <div class="doc-menu">
        <div class="doc-menu-header">
          <div class="header-search">
            {renderGroupsSelect()}
            <InputSearch
              allowClear={true}
              placeholder="请输入关键词"
              style="flex"
              {...searchInputProps}
            />

            {createCrudButton}
          </div>
          {renderTagsSelect()}
        </div>
        <div class="doc-menu-box">
          <Menu style="width: 100%" mode="inline" {...menuProps}>
            {menuTree}
          </Menu>
        </div>
      </div>
    );
  }
};
</script>

<style lang="less" scoped>
.doc-menu {
  .action-title {
    &_wraper {
      overflow: hidden;
      text-overflow: ellipsis;
      white-space: nowrap;
    }
  }
  .action-title-tag {
    width: 50px;
    text-align: center;
    padding: 0 3px;
  }
  .method-multiple-tag {
    position: relative;
    background: #ccc;
    width: 50px;
    text-align: center;
    padding: 0 3px;
    border-radius: 4px;
    display: inline-block;
    line-height: 22px;
    margin-right: 8px;
    &.method-num-2 {
      padding-left: 13px;
      .method-item.empty {
        left: 10px;
      }
    }
    &.method-num-3 {
      padding-left: 18px;
      .method-item.empty {
        left: 15px;
      }
    }
    &.method-num-4 {
      padding-left: 23px;
      .method-item.empty {
        left: 20px;
      }
    }
    .method-item {
      width: 10px;
      height: 100%;
      display: inline-block;
      border-radius: 4px;
      position: absolute;
      top: 0;

      &.GET {
        background: #87d068;
      }
      &.POST {
        background: #2db7f5;
      }
      &.PUT {
        background: #ff9800;
      }
      &.DELETE {
        background: #ff4d4f;
      }
      &.empty {
        background: #ccc;
      }
      &.method-item-0 {
        left: 0;
      }
      &.method-item-1 {
        left: 5px;
      }
      &.method-item-2 {
        left: 10px;
      }
      &.method-item-3 {
        left: 15px;
      }
    }
  }
  .doc-menu-url {
    padding: 2px 10px;
    background: #f1f1f1;
    border-radius: 4px;
    line-height: 1.6;
    margin-bottom: 5px;
  }
  .doc-menu-header {
    padding: 6px;
    border-bottom: 1px solid #ddd;
    .header-search {
      display: flex;
    }
  }
  .doc-menu-box {
    width: 100%;
    height: calc(100vh - 124px);
    overflow: hidden;
    overflow-y: auto;
    padding-bottom: 50px;
  }
}
/deep/ .ant-menu-sub.ant-menu-inline > .ant-menu-item {
  height: auto;
}
/deep/ .ant-menu-inline,
.ant-menu-vertical,
.ant-menu-vertical-left {
  border: none;
}
.menu-sub {
  .menu-sub-actions {
    position: absolute;
    right: 40px;
    button {
      padding: 0 5px;
      color: #666;
      display: none;
      &:hover {
        color: #1890ff;
      }
    }
    i {
      margin-right: 0;
      font-size: 12px;
    }
  }
  &:hover {
    .menu-sub-actions {
      button {
        display: inline-block;
      }
    }
  }
}
</style>
