import axios from "axios";
import store from "/@/store";
import { isDev } from "/@/config/env";
import { href } from "/@/core/utils";
import NProgress from "nprogress";
import "nprogress/nprogress.css";
import { ElMessage } from "element-plus";
import {del} from "@vue/composition-api";

axios.defaults.timeout = 30000;
axios.defaults.withCredentials = true;
NProgress.configure({
	showSpinner: false
});

// 请求队列
const requests: Array<Function> = [];

// Token 是否刷新中
const isRefreshing = false;

// 忽略规则
const ignore = {
	NProgress: ["/sys/info/record"],
	token: ["/login", "/captcha"]
};

// Request
axios.interceptors.request.use(
	(config: any) => {
		const token = store.getters.token || "";

		if (config.url) {
			if (!ignore.token.some((e) => config.url.includes(e))) {
				config.headers["Authorization"] = token;
			}

			if (!ignore.NProgress.some((e) => config.url.includes(e))) {
				NProgress.start();
			}
		}

		// 请求信息
		if (isDev) {
			// console.log(config);
			// console.group(config.url);
			// console.log("method:", config.method);
			// console.table("data:", config.method == "get" ? config.params : config.data);
			// console.groupEnd();
		}
		if (
			config.method == "post" &&
			config.data !== undefined &&
			Object.keys(config.data).length
		) {
			config.data["params"] = {};
			for (const configKey in config.data) {
				if (config.data.hasOwnProperty(configKey) && configKey.startsWith("params.")) {
					config.data.params[configKey.replace("params.", "")] = config.data[configKey];
					delete config.data[configKey];
				}
			}
			if (Object.keys(config.data.params).length === 0) {
				delete config.data.params
			}
		}

		// 验证 token
		// if (token) {
		// 	if (config.url.includes("refreshToken")) {
		// 		return config;
		// 	}
		//
		// 	// 判断 token 是否过期
		// 	if (storage.isExpired("token")) {
		// 		// 判断 refreshToken 是否过期
		// 		if (storage.isExpired("refreshToken")) {
		// 			store.dispatch("userRemove");
		// 			return href("/login");
		// 		}
		//
		// 		// 是否在刷新中
		// 		if (!isRefreshing) {
		// 			isRefreshing = true;
		//
		// 			store.dispatch("refreshToken").then((token: string) => {
		// 				requests.forEach((cb) => cb(token));
		// 				requests = [];
		// 				isRefreshing = false;
		// 			});
		// 		}
		//
		// 		return new Promise((resolve) => {
		// 			// 继续请求
		// 			requests.push((token: string) => {
		// 				// 重新设置 token
		// 				config.headers["Authorization"] = token;
		// 				resolve(config);
		// 			});
		// 		});
		// 	}
		// }

		return config;
	},
	(error) => {
		return Promise.reject(error);
	}
);

// Response
axios.interceptors.response.use(
	(res) => {
		NProgress.done();
		const { code, data, message } = res.data;

		if (!res.data) {
			return res;
		}

		switch (code) {
			case 1000:
				return data;
			default:
				return Promise.reject(message);
		}
	},
	async (error) => {
		NProgress.done();

		if (error.response) {
			const { status, config } = error.response;

			switch (status) {
				case 401:
					await store.dispatch("userRemove");
					href("/login");
					break;

				case 403:
					if (isDev) {
						ElMessage.error(`${config.url} 无权限访问！！`);
					} else {
						href("/403");
					}
					break;

				case 404:
					break;

				case 500:
					if (!isDev) {
						href("/500");
					}
					break;

				case 502:
					if (isDev) {
						ElMessage.error(`${config.url} 服务异常！！`);
					} else {
						href("/502");
					}
					break;

				default:
					console.error(status, config.url);
			}
		}

		return Promise.reject(error.message);
	}
);

export default axios;
