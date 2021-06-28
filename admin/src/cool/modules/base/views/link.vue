<template>
	<cl-crud @load="onLoad">
		<el-row type="flex">
			<cl-refresh-btn/>
			<cl-add-btn/>
			<cl-multi-delete-btn/>
			<cl-flex1/>
			<cl-search-key/>
		</el-row>

		<el-row>
			<cl-table v-bind="table"/>
		</el-row>

		<el-row type="flex">
			<cl-flex1/>
			<cl-pagination/>
		</el-row>

		<cl-upsert v-model="form" v-bind="upsert"/>
	</cl-crud>
</template>

<script lang="ts">
import {CrudLoad, Table, Upsert} from "cl-admin-crud-vue3/types";
import {defineComponent, inject, reactive} from "vue";

export default defineComponent({
	name: "sys-link",

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
					label: "名称",
					span: 24,
					component: {
						name: "el-input",
						props: {
							placeholder: "名称"
						}
					},
					rules: {
						required: true,
						message: "名称不能为空"
					}
				},
				{
					prop: "logo",
					label: "Logo",
					span: 24,
					component: {
						name: "cl-upload",
						props: {
							text: "选择图片",
							icon: "el-icon-picture"
						}
					}
				},
				{
					prop: "url",
					label: "地址",
					span: 12,
					component: {
						name: "el-input",
						props: {
							placeholder: "请填写地址"
						}
					},
					rules: [
						{
							required: true,
							message: "地址不能为空"
						}
					]
				},
				{
					prop: "introduce",
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
					prop: "passed",
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
					type: "selection",
					width: 40
				},
				// {
				// 	prop: "linktype",
				// 	label: "类型",
				// 	minWidth: 60
				// },
				{
					prop: "name",
					label: "名称",
					minWidth: 80
				},
				{
					prop: "logo",
					label: "Logo",
					component: ({h, scope}: any) => {
						return h("img", {
							src: scope.logo,
							height: 40
						});
					},
					minWidth: 120
				},
				{
					prop: "url",
					label: "地址",
					showOverflowTooltip: true,
					minWidth: 150
				},

				{
					prop: "introduce",
					label: "描述",
					minWidth: 150
				},
				{
					prop: "passed",
					label: "启用",
					minWidth: 50,
					dict: [
						{
							label: "启用",
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

		// crud 加载
		function onLoad({ctx, app}: CrudLoad) {
			ctx.service(service.system.link).done();
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
