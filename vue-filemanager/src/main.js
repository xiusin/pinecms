import Vue from 'vue'
import Vuex from 'vuex'
import router from './router'
import App from './App.vue'
import fm from './store'
import HTTP from './http/axios'

import {
  Button,
  Container,
  Header,
  Aside,
  Main,
  Row,
  Col,
  Radio,
  RadioGroup,
  RadioButton,
  Checkbox,
  CheckboxGroup,
  Breadcrumb,
  BreadcrumbItem,
  Form,
  FormItem,
  Input,
  Dialog,
  MessageBox,
  Progress,
  InfiniteScroll,
  Notification,
  Message,
  Upload,
  Avatar,
  Dropdown,
  DropdownMenu,
  DropdownItem
} from 'element-ui'
import 'image-preview-vue/lib/imagepreviewvue.css'
Vue.use(Vuex);
Vue.component(Button.name, Button); // 或者 VUe.use(Button)
Vue.component(Container.name, Container)
Vue.component(Header.name, Header)
Vue.component(Main.name, Main)
Vue.component(Aside.name, Aside)
Vue.component(Row.name, Row)
Vue.component(Col.name, Col)
Vue.component(Radio.name, Radio)
Vue.component(RadioGroup.name, RadioGroup)
Vue.component(RadioButton.name, RadioButton)
Vue.component(Checkbox.name, Checkbox)
Vue.component(CheckboxGroup.name, CheckboxGroup)
Vue.component(Breadcrumb.name, Breadcrumb)
Vue.component(BreadcrumbItem.name, BreadcrumbItem)
Vue.component(Form.name, Form)
Vue.component(FormItem.name, FormItem)
Vue.component(Input.name, Input)
Vue.component(Progress.name, Progress)
Vue.component(Dialog.name, Dialog)
Vue.component(Upload.name, Upload)
Vue.component(Avatar.name, Avatar)
Vue.component(Dropdown.name, Dropdown)
Vue.component(DropdownMenu.name, DropdownMenu)
Vue.component(DropdownItem.name, DropdownItem)
Vue.directive('infinite-scroll', InfiniteScroll)
Vue.prototype.$confirm = MessageBox.confirm;
Vue.prototype.$notify = Notification;
Vue.prototype.$message = Message;
const store = new Vuex.Store({
  strict: process.env.NODE_ENV !== 'production',
  modules: {
    fm
  }
})

Vue.config.productionTip = process.env.NODE_ENV === "production";

store.commit("fm/settings/initAxiosSettings");
const lang = store.state.fm.settings.translations["zh-CN"]
HTTP.interceptors.request.use(
  config => {
    // 重写公共url和请求头
    config.baseURL = store.getters["fm/settings/baseUrl"];
    config.headers = store.getters["fm/settings/headers"];
    store.commit("fm/messages/addLoading");
    return config;
  },
  error => {
    store.commit("fm/messages/subtractLoading");
    return Promise.reject(error);
  }
);
HTTP.interceptors.response.use(
  response => {
    store.commit("fm/messages/subtractLoading");
    // 如果有消息文本，创建提醒
    if (Object.prototype.hasOwnProperty.call(response.data, "result")) {
      if (response.data.result.message) {
        const message = {
          status: response.data.result.status,
          message: Object.prototype.hasOwnProperty.call(
              lang.response,
              response.data.result.message
            ) ?
            lang.response[response.data.result.message] : response.data.result.message
        };
        // 显示提醒
        // EventBus.$emit("addNotification", message);
        if (message.status == "danger") {
          Notification.error({
            title: "失败",
            message: message.message
          });
        } else {
          Notification.success({
            title: "成功",
            message: message.message
          });
        }
        // 提示消息
        store.commit("fm/messages/setActionResult", message);
      }
    }
    return response;
  },
  error => {
    store.commit("fm/messages/subtractLoading");
    const errorMessage = {
      status: 0,
      message: ""
    };
    const errorNotificationMessage = {
      status: "error",
      message: ""
    };
    // 添加消息
    if (error.response) {
      errorMessage.status = error.response.status;
      // 赋值回复的错误信息
      if (error.response.data.message) {
        const trMessage = Object.prototype.hasOwnProperty.call(
            this.lang.response,
            error.response.data.message
          ) ?
          this.lang.response[error.response.data.message] :
          error.response.data.message;
        errorNotificationMessage.message = errorMessage.message = trMessage;
      } else {
        errorNotificationMessage.message = errorMessage.message =
          error.response.statusText;
      }
    } else if (error.request) {
      errorMessage.status = error.request.status;
      errorNotificationMessage.message = errorMessage.message =
        error.request.statusText || "网络错误!!!";
    } else {
      errorNotificationMessage.message = errorMessage.message =
        error.message;
    }
    // 设置错误消息
    store.commit("fm/messages/setError", errorMessage);
    // 显示提示
    this.$notify.error({
      title: "错误",
      message: errorNotificationMessage.message
    });
    // EventBus.$emit("addNotification", errorNotificationMessage);

    return Promise.reject(error);
  }
);

//全局前置守卫，判断用户是否登陆
router.beforeEach((to, from, next) => {
  if (window.localStorage) {
    const auto = JSON.parse(window.localStorage.getItem("auto"))
    if (auto) {
      store.commit(`fm/setAutoLogin`, auto.a);
    }
  }

  if (to.path === '/login') {
    store.dispatch('fm/checkLoginStatus').then(resp => {
      if (resp.data.result.status === 'success') {
        if (store.state.fm.autoLogin && store.state.fm.isLogin) {
          next('/file')
        } else {
          next();
        }
      }
    })
  } else if (to.path === '/register') {
    next();
  } else {
    store.dispatch('fm/checkLoginStatus').then(resp => {
      if (resp.data.result.status === 'success') {
        store.commit(
          "fm/setUserName", {
            name: resp.data.name,
            nickname: resp.data.nickname
          }
        );
        if (store.state.fm.isLogin) {
          next();
        } else {
          next('/login')
        }
      } else {
        next('/login');
      }
    })

    // let token = localStorage.getItem('Authorization');
    // console.log('登录token',token)
    // if (token === null || token === '' || token === undefined) {
    // router.replace('/login')
    // next('/login');
    // } else {
    // document.title = to.meta.title ? to.meta.title : '轻享云盘'
    // next()
    // }
  }
});
new Vue({
  router,
  store,
  render: h => h(App),
}).$mount('#app')