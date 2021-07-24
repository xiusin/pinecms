<template>
	<div class="view-plugin">
		<cl-crud ref="crud" :on-refresh="onRefresh" @load="onLoad">
			<el-row type="flex" align="middle">
				<!-- 刷新按钮 -->
				<cl-refresh-btn />
				<cl-flex1 />
				<!-- 关键字搜索 -->
				<cl-search-key />
			</el-row>

			<el-row>
				<!-- 数据表格 -->
				<cl-table v-bind="table">
					<template #column-enable="{ scope }">
						<el-switch
							v-model="scope.row._enable"
							size="mini"
							@change="onEnableChange($event, scope.row)"
						/>
					</template>

					<!-- 配置按钮 -->
					<template #slot-conf="{ scope }">
						<el-button
							v-if="!scope.row.no_install"
							type="text"
							size="mini"
							@click="openConf(scope.row)"
							>配置</el-button
						>

						<el-button v-else type="text" size="mini" @click="installPlugin(scope.row)"
							>安装</el-button
						>
					</template>
				</cl-table>
			</el-row>

			<el-row type="flex">
				<cl-flex1 />
				<!-- 分页控件 -->
				<cl-pagination layout="total" />
			</el-row>
		</cl-crud>

		<el-drawer title="安装说明" v-model="visible" size="70%">
			<article class="markdown-body">
				<div v-html="pageHTML"></div>
			</article>

			<span slot="footer" class="drawer-footer">
				<el-button type="primary" @click="onInstallClick" size="mini">确 定</el-button>
				<el-button @click="onDrawClose" size="mini">取 消</el-button>
			</span>
		</el-drawer>

		<!-- 表单 -->
		<cl-form :ref="setRefs('form')" />
	</div>
</template>

<script lang="ts">
import { ElMessage } from "element-plus";
import { defineComponent, inject, reactive, ref } from "vue";
import { checkPerm } from "/$/base";
import { useRefs } from "/@/core";
import { CrudLoad, RefreshOp, Table } from "cl-admin-crud-vue3/types";

export default defineComponent({
	name: "plugin",

	setup() {
		const service = inject<any>("service");
		const { refs, setRefs } = useRefs();

		// 编辑权限
		const { config, getConfig, enable } = service.plugin.info.permission;

		const visible = ref<boolean>(false);
		const pageHTML = ref("");
		const installPath = ref("");

		const perms = reactive<any>({
			edit: checkPerm({
				and: [config, getConfig]
			}),
			enable: checkPerm(enable)
		});

		// crud 加载
		function onLoad({ ctx, app }: CrudLoad) {
			ctx.service(service.plugin.info)
				.set("dict", {
					api: {
						page: "list"
					}
				})
				.done();
			app.refresh();
		}

		// 刷新钩子
		function onRefresh(params: any, { next, render }: RefreshOp) {
			next(params).then((res: any) => {
				const list = res.list.map((e: any) => {
					e._enable = e.enable;
					return e;
				});

				render(list, {
					total: res.list.length
				});
			});
		}

		// 开启、关闭
		function onEnableChange(val: boolean, item: any) {
			service.plugin.info
				.enable({
					path: item.path,
					enable: val
				})
				.then(() => {
					ElMessage.success(val ? "开启成功" : "关闭成功");
				})
				.catch((err: string) => {
					ElMessage.error(err);
				});
		}

		async function installPlugin({ page, path }: any) {
			visible.value = true;
			pageHTML.value = page;
			installPath.value = path;
		}

		function onInstallClick() {
			service.plugin.info
				.install({ path: installPath.value })
				.then(() => {
					ElMessage.success("安装插件成功");
					close();
				})
				.catch((err: string) => {
					ElMessage.error(err);
				});
		}

		function onDrawClose() {
			visible.value = false;
			pageHTML.value = "";
			installPath.value = "";
		}

		// 打开配置
		async function openConf({ name, path, view }: any) {
			const form = await service.plugin.info.getConfig({ path });

			let items: any[];

			try {
				items = typeof view == "string" ? JSON.parse(view) : view;
			} catch (e) {
				items = [];
			}

			refs.value.form.open({
				title: `${name}配置`,
				items,
				form,
				on: {
					submit: (data: any, { close, done }: any) => {
						service.plugin.info
							.config({
								path,
								config: data
							})
							.then(() => {
								ElMessage.success("保存成功");
								close();
							})
							.catch((err: string) => {
								ElMessage.error(err);
								done();
							});
					}
				}
			});
		}

		// 表格配置
		const table = reactive<Table>({
			props: {
				"default-sort": {
					prop: "id",
					order: "descending"
				}
			},
			columns: [
				{
					label: "插件路径",
					prop: "path",
					width: 180,
					align: "left"
				},
				{
					label: "名称",
					prop: "name",
					showOverflowTooltip: true,
					width: 140
				},
				{
					label: "作者",
					prop: "author",
					width: 120
				},
				{
					label: "联系方式",
					prop: "contact",
					showOverflowTooltip: true,
					width: 180
				},
				{
					label: "功能描述",
					prop: "description",
					showOverflowTooltip: true
				},
				{
					label: "版本号",
					prop: "version",
					width: 110
				},
				{
					label: "是否启用",
					prop: "enable",
					width: 110
				},
				{
					label: "状态",
					prop: "status",
					width: 150,
					dict: [
						{
							label: "缺少配置",
							value: 0,
							type: "warning"
						},
						{
							label: "可用",
							value: 1,
							type: "success"
						},
						{
							label: "配置错误",
							value: 2,
							type: "danger"
						},
						{
							label: "未知错误",
							value: 3,
							type: "danger"
						}
					]
				},
				{
					label: "创建时间",
					prop: "created_at",
					width: 150,
					sortable: "custom"
				},
				{
					type: "op",
					width: 120,
					buttons: ["slot-conf"]
				}
			]
		});

		return {
			refs,
			perms,
			table,
			visible,
			onDrawClose,
			onInstallClick,
			pageHTML,
			setRefs,
			installPlugin,
			onLoad,
			onRefresh,
			onEnableChange,
			openConf
		};
	}
});
</script>

<style>
/*1.显示滚动条：当内容超出容器的时候，可以拖动：*/
.el-drawer__body {
	overflow: auto;
}
/*2.隐藏滚动条，太丑了*/
.el-drawer__container ::-webkit-scrollbar {
	display: none;
}

.drawer-footer {
	padding: 10px;
	position: absolute;
	bottom: 10px;
}
.markdown-body ol ol,
.markdown-body ul ol,
.markdown-body ol ul,
.markdown-body ul ul,
.markdown-body ol ul ol,
.markdown-body ul ul ol,
.markdown-body ol ul ul,
.markdown-body ul ul ul {
	margin-top: 0;
	margin-bottom: 0;
}
.markdown-body {
	font-family: "Helvetica Neue", Helvetica, "Segoe UI", Arial, freesans, sans-serif,
		"Apple Color Emoji", "Segoe UI Emoji", "Segoe UI Symbol";
	font-size: 16px;
	color: #333;
	line-height: 1.6;
	word-wrap: break-word;
	padding: 45px;
	background: #fff;
	border: 1px solid #ddd;
	-webkit-border-radius: 0 0 3px 3px;
	border-radius: 0 0 3px 3px;
}
.markdown-body > *:first-child {
	margin-top: 0 !important;
}
.markdown-body > *:last-child {
	margin-bottom: 0 !important;
}
.markdown-body * {
	-webkit-box-sizing: border-box;
	-moz-box-sizing: border-box;
	box-sizing: border-box;
}
.markdown-body h1,
.markdown-body h2,
.markdown-body h3,
.markdown-body h4,
.markdown-body h5,
.markdown-body h6 {
	margin-top: 1em;
	margin-bottom: 16px;
	font-weight: bold;
	line-height: 1.4;
}
.markdown-body p,
.markdown-body blockquote,
.markdown-body ul,
.markdown-body ol,
.markdown-body dl,
.markdown-body table,
.markdown-body pre {
	margin-top: 0;
	margin-bottom: 16px;
}
.markdown-body h1 {
	margin: 0.67em 0;
	padding-bottom: 0.3em;
	font-size: 2.25em;
	line-height: 1.2;
	border-bottom: 1px solid #eee;
}
.markdown-body h2 {
	padding-bottom: 0.3em;
	font-size: 1.75em;
	line-height: 1.225;
	border-bottom: 1px solid #eee;
}
.markdown-body h3 {
	font-size: 1.5em;
	line-height: 1.43;
}
.markdown-body h4 {
	font-size: 1.25em;
}
.markdown-body h5 {
	font-size: 1em;
}
.markdown-body h6 {
	font-size: 1em;
	color: #777;
}
.markdown-body ol,
.markdown-body ul {
	padding-left: 2em;
}
.markdown-body ol ol,
.markdown-body ul ol {
	list-style-type: lower-roman;
}
.markdown-body ol ul,
.markdown-body ul ul {
	list-style-type: circle;
}
.markdown-body ol ul ul,
.markdown-body ul ul ul {
	list-style-type: square;
}
.markdown-body ol {
	list-style-type: decimal;
}
.markdown-body ul {
	list-style-type: disc;
}
.markdown-body blockquote {
	margin-left: 0;
	margin-right: 0;
	padding: 0 15px;
	color: #777;
	border-left: 4px solid #ddd;
}
.markdown-body table {
	display: block;
	width: 100%;
	overflow: auto;
	word-break: normal;
	word-break: keep-all;
	border-collapse: collapse;
	border-spacing: 0;
}
.markdown-body table tr {
	background-color: #fff;
	border-top: 1px solid #ccc;
}
.markdown-body table tr:nth-child(2n) {
	background-color: #f8f8f8;
}
.markdown-body table th,
.markdown-body table td {
	padding: 6px 13px;
	border: 1px solid #ddd;
}
.markdown-body pre {
	word-wrap: normal;
	padding: 16px;
	overflow: auto;
	font-size: 85%;
	line-height: 1.45;
	background-color: #f7f7f7;
	-webkit-border-radius: 3px;
	border-radius: 3px;
}
.markdown-body pre code {
	display: inline;
	max-width: initial;
	padding: 0;
	margin: 0;
	overflow: initial;
	font-size: 100%;
	line-height: inherit;
	word-wrap: normal;
	white-space: pre;
	border: 0;
	-webkit-border-radius: 3px;
	border-radius: 3px;
	background-color: transparent;
}
.markdown-body pre code:before,
.markdown-body pre code:after {
	content: normal;
}
.markdown-body code {
	font-family: Consolas, "Liberation Mono", Menlo, Courier, monospace;
	padding: 0;
	padding-top: 0.2em;
	padding-bottom: 0.2em;
	margin: 0;
	font-size: 85%;
	background-color: rgba(0, 0, 0, 0.04);
	-webkit-border-radius: 3px;
	border-radius: 3px;
}
.markdown-body code:before,
.markdown-body code:after {
	letter-spacing: -0.2em;
	content: "\00a0";
}
.markdown-body a {
	color: #4078c0;
	text-decoration: none;
	background: transparent;
}
.markdown-body img {
	max-width: 100%;
	max-height: 100%;
	-webkit-border-radius: 4px;
	border-radius: 4px;
	-webkit-box-shadow: 0 0 10px #555;
	box-shadow: 0 0 10px #555;
}
.markdown-body strong {
	font-weight: bold;
}
.markdown-body em {
	font-style: italic;
}
.markdown-body del {
	text-decoration: line-through;
}
.task-list-item {
	list-style-type: none;
}
.task-list-item input {
	font: 13px/1.4 Helvetica, arial, nimbussansl, liberationsans, freesans, clean, sans-serif,
		"Apple Color Emoji", "Segoe UI Emoji", "Segoe UI Symbol";
	margin: 0 0.35em 0.25em -1.6em;
	vertical-align: middle;
}
.task-list-item input[disabled] {
	cursor: default;
}
.task-list-item input[type="checkbox"] {
	-webkit-box-sizing: border-box;
	-moz-box-sizing: border-box;
	box-sizing: border-box;
	padding: 0;
}
.task-list-item input[type="radio"] {
	-webkit-box-sizing: border-box;
	-moz-box-sizing: border-box;
	box-sizing: border-box;
	padding: 0;
}
</style>
