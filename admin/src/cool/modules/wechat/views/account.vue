<template>
	<cl-crud :ref="setRefs('crud')" @load="onLoad">
		<el-row type="flex">
			<cl-add-btn />
			<cl-refresh-btn />
			<cl-flex1 />
			<cl-search-key />
		</el-row>

		<el-row>
			<cl-table v-bind="table">
				<template #column-copy="{ scope }">
						<el-button v-copy="buildURL(scope.row.appid)" size="mini">
							{{ buildURL(scope.row.appid) }}
						</el-button>
				</template>
			</cl-table>
		</el-row>

		<el-row type="flex">
			<cl-flex1 />
			<cl-pagination />
		</el-row>

		<cl-upsert v-model="form" v-bind="upsert" />
	</cl-crud>
</template>

<script lang="ts">
import { defineComponent, inject, reactive, ref } from "vue";
import { useRefs } from "/@/core";
import { CrudLoad, Table, Upsert } from "cl-admin-crud-vue3/types";

export default defineComponent({
	name: "wechat-account",

	setup() {
		const service = inject<any>("service");

		const { refs, setRefs }: any = useRefs();

		const upsert = reactive<Upsert>({
			items: [
				{
					prop: "name",
					label: "公众号名称",
					span: 24,
					component: {
						name: "el-input",
						props: {
							placeholder: "公众号名称"
						}
					},
					rules: {
						required: true,
						message: "公众号名称不能为空"
					}
				},

				{
					prop: "type",
					label: "公众号类型",
					value: 1,
					component: {
						name: "el-radio-group",
						options: [
							{
								label: "订阅号",
								value: 1
							},
							{
								label: "服务号",
								value: 2
							}
						]
					}
				},
				{
					prop: "verified",
					label: "是否认证",
					value: true,
					component: {
						name: "el-radio-group",
						options: [
							{
								label: "是",
								value: true
							},
							{
								label: "否",
								value: false
							}
						]
					}
				},
				{
					prop: "appid",
					label: "appid",
					component: {
						name: "el-input",
						props: {
							placeholder: "请输入appid"
						}
					},
					rules: {
						required: true,
						message: "appid不能为空"
					}
				},
				{
					prop: "secret",
					label: "secret",
					component: {
						name: "el-input",
						props: {
							placeholder: "请输入secret"
						}
					},
					rules: {
						required: true,
						message: "secret不能为空"
					}
				},
				{
					prop: "token",
					label: "token",
					component: {
						name: "el-input",
						props: {
							placeholder: "请输入token"
						}
					},
					rules: {
						required: true,
						message: "token不能为空"
					}
				},
				{
					prop: "aesKey",
					label: "aesKey",
					component: {
						name: "el-input"
					}
				}
			]
		});

		const table = reactive<Table>({
			columns: [
				{
					type: "index",
					label: "#",
					width: 60
				},
				{
					prop: "name",
					label: "公众号名称",
					width: 250,
					align: "left"
				},
				{
					prop: "appid",
					label: "APPID",
					width: 300
				},
				{
					prop: "type",
					label: "类型",
					width: 140,
					dict: [
						{
							label: "公众号",
							value: 1,
							type: "success"
						},
						{
							label: "服务号",
							value: 2,
							type: "warning"
						}
					]
				},
				{
					prop: "verified",
					label: "是否认证",
					width: 140,
					dict: [
						{
							label: "是",
							value: 1,
							type: "success"
						},
						{
							label: "否",
							value: 0,
							type: "danger"
						}
					]
				},
				{
					prop: "copy",
					label: "通知地址"
				},
				{
					type: "op",
					buttons: ["edit", "delete"]
				}
			]
		});

		function onLoad({ ctx, app }: CrudLoad) {
			ctx.service(service.wechat.account).done();
			app.refresh();
		}

		function buildURL(appid: string) {
			return window.location.protocol + "//" + window.location.host + `/api/wechat/msg/${appid}`
		}

		return {
			service,
			buildURL,
			refs,
			table,
			setRefs,
			onLoad,
			upsert
		};
	}
});
</script>
