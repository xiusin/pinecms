<template>
  <div ref="container" :class="['doc-wraper', device]">
    <error-box v-if="error.status" :error="error" />
    <div v-else :class="['doc-content']">
      <div v-if="device !== 'mobile'" class="affix-wraper">
        <div class="affix-box">
          <a-anchor :getContainer="pageContainer" :target-offset="targetOffset">
            <a-anchor-link
              v-for="(item, i) in navs"
              :key="i"
              :href="`#${item.title}`"
              :title="item.title"
            >
              <template v-if="item.children && item.children.length">
                <a-anchor-link
                  v-for="(cItem, j) in item.children"
                  :key="j"
                  :href="`#${cItem.title}`"
                  :title="cItem.title"
                />
              </template>
            </a-anchor-link>
          </a-anchor>
        </div>
      </div>
      <a-spin tip="Loading..." :spinning="loading">
        <div class=" markdown" v-html="mdHtml"></div>
      </a-spin>
    </div>
  </div>
</template>

<script>
import { Spin, Anchor, Button } from "ant-design-vue";
import { getMdDetail } from "@/api/app";
import marked from "marked";
import hljs from "highlight.js";
import "highlight.js/styles/atom-one-dark.css";
import "./markdown.less";
import { trim, setCurrentUrl, changeUrlArg, deleteUrlArg } from "@/utils/utils";
import ErrorBox from "./ErrorBox";
import PasswordModal from "./auth/passwordModal";
import { ls } from "@/utils/cache";
marked.setOptions({
  highlight: function(code) {
    return hljs.highlightAuto(code).value;
  }
});

export default {
  components: {
    [Spin.name]: Spin,
    [Anchor.name]: Anchor,
    [Anchor.Link.name]: Anchor.Link,
    [Button.name]: Button,
    ErrorBox
  },
  props: {
    docData: {
      type: Object,
      default: () => {}
    },
    appKey: {
      type: String,
      default: ""
    },
    device: {
      type: String,
      default: "xl"
    },
    cardRef: Function
  },
  data() {
    return {
      mdHtml: "",
      loading: false,
      targetOffset: undefined,
      navs: [],
      pageContainer: null,
      error: {}
    };
  },
  watch: {
    docData(val) {
      this.fetchMdContent(val);
    }
  },
  created() {
    if (this.docData && this.docData.path) {
      this.fetchMdContent(this.docData);
    }
  },
  mounted() {
    this.targetOffset = window.innerHeight / 2;
    this.pageContainer = () => {
      return window.document.getElementById("pageContainer");
    };
  },
  methods: {
    trim,
    fetchMdContent(item) {
      this.loading = true;
      getMdDetail({
        appKey: this.appKey,
        path: item.path
      })
        .then(res => {
          this.loading = false;
          if (res.data.code !== 0) {
            this.error = {
              status: res.data.code,
              message: res.data.msg
            };
            return false;
          }
          this.navs = this.handleNavTree(res.data.data);
          this.$nextTick(() => {
            this.mdHtml = marked(res.data.data);
          });
          // 更新url
          let url = changeUrlArg(window.location.href, "md", item.path);
          url = deleteUrlArg(url, "api");
          setCurrentUrl(url);
        })
        .catch(err => {
          const status =
            err.response && err.response.status ? err.response.status : 500;
          if (status === 401) {
            ls.remove("token");
            PasswordModal({
              appKey: this.currentAppKey,
              success: () => {
                window.location.reload();
              }
            });
          } else {
            this.error = {
              status: status,
              message:
                err.response && err.response.data && err.response.data.msg
                  ? err.response.data.msg
                  : err.message
            };
          }
          this.loading = false;
        });
    },
    getTitle(content) {
      let nav = [];

      let tempArr = [];
      content.replace(/(#+)[^#][^\n]*?(?:\n)/g, function(match, m1) {
        let title = match.replace("\n", "");
        let level = m1.length;
        title = title.replace(/^#+/, "").replace(/\([^)]*?\)/, "");
        tempArr.push({
          title: trim(title),
          level: level,
          children: []
        });
      });

      // 只处理一级二级标题，以及添加与id对应的index值
      nav = tempArr.filter(item => item.level === 2 || item.level === 3);
      let index = 0;
      return (nav = nav.map(item => {
        item.index = index++;
        return item;
      }));
    },
    // 将一级二级标题数据处理成树结构
    handleNavTree(mdContent) {
      let navs = this.getTitle(mdContent);
      let navLevel = [2, 3];
      let retNavs = [];
      let toAppendNavList;
      navLevel.forEach(level => {
        // 遍历一级二级标题，将同一级的标题组成新数组
        toAppendNavList = this.find(navs, {
          level: level
        });
        if (retNavs.length === 0) {
          // 处理一级标题
          retNavs = retNavs.concat(toAppendNavList);
        } else {
          // 处理二级标题，并将二级标题添加到对应的父级标题的children中
          toAppendNavList.forEach(item => {
            item = Object.assign(item);
            let parentNavIndex = this.getParentIndex(navs, item.index);
            return this.appendToParentNav(retNavs, parentNavIndex, item);
          });
        }
      });
      return retNavs;
    },
    find(arr, condition) {
      return arr.filter(item => {
        for (let key in condition) {
          if (condition.hasOwnProperty(key) && condition[key] !== item[key]) {
            return false;
          }
        }
        return true;
      });
    },
    getParentIndex(nav, endIndex) {
      for (var i = endIndex - 1; i >= 0; i--) {
        if (nav[endIndex].level > nav[i].level) {
          return nav[i].index;
        }
      }
    },
    appendToParentNav(nav, parentIndex, newNav) {
      let index = this.findIndex(nav, {
        index: parentIndex
      });
      nav[index].children = nav[index].children.concat(newNav);
    },
    findIndex(arr, condition) {
      let ret = -1;
      arr.forEach((item, index) => {
        for (var key in condition) {
          if (condition.hasOwnProperty(key) && condition[key] !== item[key]) {
            return false;
          }
        }
        ret = index;
      });
      return ret;
    }
  }
};
</script>
