<template>
	<cl-crud :ref="setRefs('crud')" @load="onLoad" :on-refresh="onRefresh">
		<el-row type="flex">
			<cl-refresh-btn/>
			<cl-add-btn/>
			<cl-flex1/>
			<cl-search-key/>
		</el-row>

		<el-row>
			<cl-table :ref="setRefs('table')" v-bind="table" @row-click="onRowClick"/>
		</el-row>

		<el-row type="flex">
			<cl-flex1/>
			<cl-pagination/>
		</el-row>

		<cl-upsert v-model="form" v-bind="upsert"/>
	</cl-crud>
</template>

<script lang="ts">
import {useRefs} from "/@/core";
import {deepTree} from "/@/core/utils";
import {CrudLoad, RefreshOp, Table, Upsert} from "cl-admin-crud-vue3/types";
import {defineComponent, inject, reactive} from "vue";

export default defineComponent({
	name: "sys-department",

	setup() {

		const {refs, setRefs} = useRefs();

		const service = inject<any>("service");

		// 表单值
		const form = reactive<any>({
			relevance: 1
		});

		// 新增、编辑配置
		const upsert = reactive<Upsert>({
			items: [
				{
					prop: "parent_id",
					label: "上级部门",
					span: 24,
					component: {
						name: "el-department-tree"
					}
				},
				{
					prop: "name",
					label: "部门名称",
					span: 12,
					component: {
						name: "el-input",
						props: {
							placeholder: "请输入部门名称"
						}
					},
					rules: {
						required: true,
						message: "部门名称不能为空"
					}
				},
				{
					prop: "leader_name",
					label: "负责人",
					span: 12,
					component: {
						name: "el-input",
						props: {
							placeholder: "请填写负责人"
						}
					}
				},
				{
					prop: "leader_phone",
					label: "联系电话",
					span: 12,
					component: {
						name: "el-input",
						props: {
							placeholder: "请填写电话",
						}
					}
				},
				{
					prop: "email",
					label: "邮箱",
					span: 12,
					component: {
						name: "el-input",
						props: {
							placeholder: "请填写邮箱",
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
							min: 0
						}
					}
				},
				{
					prop: "status",
					label: "状态",
					value: true,
					span: 12,
					component: {
						name: "el-radio-group",
						options: [
							{
								label: "正常",
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
				},
				"row-key": "id"
			},
			columns: [
				{
					prop: "listorder",
					label: "排序值",
					width: 80
				},
				{
					prop: "name",
					label: "名称",
					width: 200,
					align: "left",
				},
				{
					prop: "leader_name",
					label: "负责人",
					width: 200,
					align: "left",
				},
				{
					prop: "leader_phone",
					label: "联系电话",
					width: 200,
					align: "left",
				},

				{
					prop: "email",
					label: "邮箱",
					align: "left",
				},
				{
					prop: "status",
					label: "启用",
					width: 80,
					dict: [
						{
							label: "正常",
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
					type: "op",
					width: 120
				}
			]
		});

		function onRefresh(_: any, {render}: RefreshOp) {
			service.system.department.list().then((list: any[]) => {
				render(deepTree(list), {
					total: list.length
				});
			});
		}

		function onRowClick(row: any, column: any) {
			if (column.property && row.children) {
				refs.value.table.toggleRowExpansion(row);
			}
		}

		function upsertAppend({type, id}: any) {
			refs.value.crud.rowAppend({
				parentId: id,
				type: type + 1
			});
		}

		// crud 加载
		function onLoad({ctx, app}: CrudLoad) {
			ctx.service(service.system.department).done();
			app.refresh();
		}

		return {
			form,
			upsert,
			table,
			onLoad,
			onRefresh,
			onRowClick,
			upsertAppend,
			setRefs,
		};
	}
});
</script>
