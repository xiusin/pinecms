<template>
	<cl-crud @load="onLoad">
		<el-row type="flex">
			<cl-refresh-btn />
			<cl-add-btn />
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

		<cl-upsert v-model="form" v-bind="upsert" />
	</cl-crud>
</template>

<script lang="ts">
import { CrudLoad, Table, Upsert } from "cl-admin-crud-vue3/types";
import { defineComponent, inject, reactive } from "vue";

export default defineComponent({
	name: "sys-position",

	setup() {
		const service = inject<any>("service");

		// 表单值
		const form = reactive<any>({
			relevance: 1
		});

		// 新增、编辑配置
		const upsert = reactive<Upsert>({
			items: [
				{
					prop: "name",
					label: "岗位名称",
					span: 24,
					component: {
						name: "el-input",
						props: {
							placeholder: "请输入岗位名称"
						}
					},
					rules: {
						required: true,
						message: "岗位名称不能为空"
					}
				},
				{
					prop: "code",
					label: "岗位编码",
					span: 24,
					component: {
						name: "el-input",
						props: {
							placeholder: "请输入岗位编码"
						}
					},
					rules: {
						required: true,
						message: "岗位编码不能为空"
					}
				},
				{
					prop: "remark",
					label: "备注",
					span: 24,
					component: {
						name: "el-input",
						props: {
							placeholder: "请输入岗位备注",
							type: "textarea"
						}
					}
				},
				{
					prop: "status",
					label: "状态",
					value: true,
					component: {
						name: "el-radio-group",
						options: [
							{
								label: "启用",
								value: true
							},
							{
								label: "禁用",
								value: false
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
					width: 180,
					align:"left"
				},
				{
					prop: "code",
					label: "编码",
					width: 150,
					align:"left"
				},

				{
					prop: "remark",
					showOverflowTooltip: true,
					label: "描述",
					align:"left"
				},
				{
					prop: "status",
					label: "状态",
					width: 80,
					dict: [
						{
							label: "启用",
							value: true,
							type: "success"
						},
						{
							label: "禁用",
							value: false,
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

		// crud 加载
		function onLoad({ ctx, app }: CrudLoad) {
			ctx.service(service.system.position).done();
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
