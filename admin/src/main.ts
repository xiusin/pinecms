import { createApp } from "vue";
import App from "./App.vue";

// cool
import { bootstrap } from "./core";

// router
import router from "./router";

// store
import store from "./store";

import "./mock";

// element-plus
import ElementPlus from "element-plus";
import "./assets/css/element-variables.scss";
// mitt
import mitt from "mitt";

// echarts
import VueECharts from "vue-echarts";

import VueUeditorWrap from "vue-ueditor-wrap";

const app = createApp(App);

bootstrap(app)
	.then(() => {
		// echarts 可视图表
		app.component("v-chart", VueECharts);
		// 事件通讯
		app.provide("mitt", mitt());

		app.use(store).use(router).use(VueUeditorWrap).use(ElementPlus).mount("#app");
	})
	.catch((err: string) => {
		console.error(`COOL-ADMIN 启动失败`, err);
	});

store.dispatch("appLoad");

// @ts-ignore
window.__app__ = app;
