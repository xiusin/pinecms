<template>
	<cl-crud @load="onLoad">
		<el-row type="flex">
			<cl-refresh-btn />
			<cl-add-btn />
			<cl-multi-delete-btn />
			<cl-flex1 />
			<cl-search-key />
		</el-row>

		<el-row>
			<cl-table v-bind="table" />
		</el-row>

		<el-row type="flex">
			<cl-flex1 />
			<cl-pagination />
		</el-row>

		<cl-upsert v-model="form" :on-submit="upsertSubmit" v-bind="upsert" />
	</cl-crud>
</template>

<script lang="ts">
import { CrudLoad, Table, Upsert } from "cl-admin-crud-vue3/types";
import { defineComponent, inject, reactive } from "vue";

export default defineComponent({
	name: "sys-role",

	setup() {
		const service = inject<any>("service");

		// 表单值
		const form = reactive<any>({
			relevance: 1
		});

		// 新增、编辑配置
		const upsert = reactive<Upsert>({
			width: "800px",

			items: [
				{
					prop: "rolename",
					label: "名称",
					span: 24,
					component: {
						name: "el-input",
						props: {
							placeholder: "请填写名称"
						}
					},
					rules: {
						required: true,
						message: "名称不能为空"
					}
				},
				{
					prop: "description",
					label: "备注",
					span: 24,
					component: {
						name: "el-input",
						props: {
							placeholder: "请填写备注",
							type: "textarea",
							rows: 4
						}
					}
				},
				{
					prop: "listorder",
					label: "排序值",
					span: 24,
					component: {
						name: "el-input",
						type: "number"
					}
				},
				{
					label: "功能权限",
					prop: "menuIdList",
					value: [],
					component: {
						name: "cl-role-perms"
					}
				},
				// {
				// 	label: "数据权限",
				// 	prop: "departmentIdList",
				// 	value: [],
				// 	component: {
				// 		name: "cl-dept-check"
				// 	}
				// }
			]
		});

		// 表格配置
		const table = reactive<Table>({
			props: {
				"default-sort": {
					prop: "id",
					order: "ascending"
				}
			},
			columns: [
				{
					type: "selection",
					width: 60
				},
				{
					prop: "rolename",
					label: "名称",
					width: 150,
					align: "left"
				},
				{
					prop: "description",
					label: "备注",
					showOverflowTooltip: true,
					align: "left"
				},
				{
					label: "操作",
					type: "op"
				}
			]
		});

		// crud 加载
		function onLoad({ ctx, app }: CrudLoad) {
			ctx.service(service.system.role).done();
			app.refresh();
		}

		function upsertSubmit(isEdit: boolean, data: any, { next }: any) {
			data.listorder = parseInt(data.listorder);
			next(data);
		}

		return {
			form,
			upsert,
			table,
			onLoad,
			upsertSubmit
		};
	}
});
</script>
