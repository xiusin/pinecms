<template>
	<cl-crud @load="onLoad">
		<el-row type="flex">
			<cl-refresh-btn />
			<cl-add-btn />
		</el-row>

		<el-row>
			<cl-table v-bind="table" />
		</el-row>

		<el-row type="flex">
			<cl-flex1 />
			<cl-pagination />
		</el-row>

		<cl-upsert v-model="form" v-bind="upsert" />
	</cl-crud>
</template>

<script lang="ts">
import { CrudLoad, Table, Upsert } from "cl-admin-crud-vue3/types";
import { defineComponent, inject, reactive } from "vue";

export default defineComponent({
	name: "sys-member-group",

	setup() {
		const service = inject<any>("service");

		const form = reactive<any>({ relevance: 1 });

		const upsert = reactive<Upsert>({
			items: [
				{
					prop: "name",
					label: "分组名称",
					span: 24,
					component: {
						name: "el-input",
						props: {
							placeholder: "分组名称"
						}
					},
					rules: {
						required: true,
						message: "分组名称不能为空"
					}
				},
				{
					prop: "description",
					label: "描述",
					span: 24,
					component: {
						name: "el-input",
						props: {
							placeholder: "分组名称",
							type: "textarea",
							rows: 4
						}
					}
				},
				{
					prop: "listorder",
					label: "排序",
					span: 8,
					value: 0,
					component: {
						name: "el-input-number",
						props: {
							placeholder: "排序"
						}
					},
					rules: {
						required: true,
						message: "排序不能为空"
					}
				},
				{
					prop: "status",
					label: "状态",
					value: 1,
					component: {
						name: "el-radio-group",
						options: [
							{
								label: "启用",
								value: 1
							},
							{
								label: "禁用",
								value: 0
							}
						]
					}
				}
			]
		});

		// 表格配置
		const table = reactive<Table>({
			props: {
				"default-sort": {
					prop: "listorder",
					order: "descending"
				}
			},
			columns: [
				{
					prop: "name",
					label: "名称",
					width: 200
				},
				{
					prop: "listorder",
					label: "排序",
					width: 80
				},
				{
					prop: "description",
					label: "描述",
					align: "left",
				},
				{
					prop: "status",
					label: "状态",
					width: 100,
					dict: [
						{
							label: "正常",
							value: 1,
							type: "success"
						},
						{
							label: "禁用",
							value: 0,
							type: "danger"
						}
					]
				},
				{
					label: "操作",
					type: "op"
				}
			]
		});

		function onLoad({ ctx, app }: CrudLoad) {
			ctx.service(service.system.memberGroup).done();
			app.refresh();
		}

		return {
			form,
			upsert,
			table,
			onLoad
		};
	}
});
</script>
