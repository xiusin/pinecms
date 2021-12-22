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

import "element-plus/theme-chalk/src/index.scss";
// mitt
import mitt from "mitt";

// echarts
import VueECharts from "vue-echarts";

import VueUeditorWrap from "vue-ueditor-wrap";

const app = createApp(App);

bootstrap(app)
	.then(() => {
		app.component("v-chart", VueECharts);
		app.provide("mitt", mitt());
		app.use(store).use(ElementPlus).use(router).use(VueUeditorWrap).mount("#app");
	})
	.catch((err: string) => {
		console.error(`启动失败`, err);
	});

store.dispatch("appLoad");

// @ts-ignore
window.__app__ = app;
