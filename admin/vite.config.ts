import path from "path";
import type { UserConfig } from "vite";
import vue from "@vitejs/plugin-vue";
import vueJsx from "@vitejs/plugin-vue-jsx";
// import viteCompression from "vite-plugin-compression";
import { svgBuilder } from "./src/core/utils/svg";
import Components from "unplugin-vue-components/vite";

function resolve(dir: string) {
	return path.resolve(__dirname, ".", dir);
}

// https://vitejs.dev/config/
// viteCompression(),
export default (): UserConfig => {
	return {
		base: "/",
		plugins: [vue(), Components(), vueJsx(), svgBuilder("./src/icons/svg/")],
		resolve: {
			alias: {
				"/@": resolve("src"),
				"/#": resolve("types"),
				"/$": resolve("src/cool/modules")
			}
		},
		css: {
			preprocessorOptions: {
				scss: {
					additionalData: `@use "./src/assets/css/element.scss" as *;`
				}
			}
		},
		server: {
			port: 9090,
			hmr: {
				overlay: true
			},
			proxy: {
				"/dev": {
					target: "http://localhost:2019/v2",
					changeOrigin: true,
					rewrite: (path) => path.replace(/^\/dev/, "")
				},

				"/pro": {
					target: "https://pinecms.xiusin.cn/v2",
					changeOrigin: true,
					rewrite: (path) => path.replace(/^\/pro/, "/api")
				}
			}
		},
		build: {
			sourcemap: false,
			polyfillDynamicImport: false // 必须为false
		},
		optimizeDeps: {
			exclude: ["vue-demi"]
		}
	};
};
