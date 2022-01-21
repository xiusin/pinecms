<template>
	<cl-crud :ref="setRefs('crud')" @load="onLoad" :on-refresh="onRefresh">
		<el-row type="flex">
			<cl-refresh-btn />
			<el-button size="mini" type="primary" @click="syncRemoteDB">导入远程数据库</el-button>
			<cl-flex1 />
		</el-row>

		<el-row>
			<cl-table :ref="setRefs('table')" v-bind="table" @row-click="onRowClick" />
		</el-row>

		<el-row type="flex">
			<cl-flex1 />
			<cl-pagination />
		</el-row>

		<cl-upsert v-model="form" v-bind="upsert" />
	</cl-crud>
</template>

<script lang="ts">
import { useRefs } from "/@/cool";
import { deepTree } from "/@/cool/utils";
import { CrudLoad, RefreshOp, Table, Upsert } from "@cool-vue/crud/types";
import { defineComponent, inject, reactive } from "vue";
import { ElMessage, ElMessageBox } from "element-plus";

export default defineComponent({
	name: "sys-district",

	setup() {
		const { refs, setRefs } = useRefs();
		const service = inject<any>("service");
		const form = reactive<any>({
			relevance: 1
		});

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
							placeholder: "请填写电话"
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
							placeholder: "请填写邮箱"
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
					prop: "order",
					order: "ascending"
				},
				"row-key": "id"
			},
			columns: [
				{
					prop: "order",
					label: "排序值",
					width: 80
				},
				{
					prop: "name",
					label: "名称",
					width: 200,
					align: "left"
				},
				{
					prop: "code",
					label: "编码",
					width: 200,
					align: "left"
				},
				{
					prop: "area_code",
					label: "区域编码",
					width: 200,
					align: "left"
				},
				{
					prop: "pinyin",
					label: "拼音",
					align: "left"
				},
				{
					prop: "initials",
					label: "简拼",
					align: "left"
				},
				{
					prop: "suffix",
					label: "后缀",
					align: "left"
				},
				{
					label: "操作",
					type: "op",
					width: 120
				}
			]
		});

		function onRefresh(pag: any, { render }: RefreshOp) {
			pag.size = 5; //todo 异步加载吧
			service.system.district.list(pag).then(({ list }: any) => {
				render(deepTree(list), {
					total: list.length,
					size: 5
				});
			});
		}

		function onRowClick(row: any, column: any) {
			if (column.property && row.children) {
				refs.value.table.toggleRowExpansion(row);
			}
		}

		function upsertAppend({ type, id }: any) {
			refs.value.crud.rowAppend({
				parentId: id,
				type: type + 1
			});
		}

		// crud 加载
		function onLoad({ ctx, app }: CrudLoad) {
			ctx.service(service.system.district).done();
			app.refresh();
		}

		function syncRemoteDB() {
			ElMessageBox.confirm(
				`是否从Github上同步最新的数据库内容,此操作将会覆盖原始数据`,
				"重要提示",
				{
					type: "warning"
				}
			)
				.then(() => {
					service.system.district
						.import()
						.then(() => {
							ElMessage.success("导入数据成功");
						})
						.catch((err: any) => {
							ElMessage.error(err);
						});
				})
				.catch(() => null);
		}

		return {
			syncRemoteDB,
			form,
			upsert,
			table,
			onLoad,
			onRefresh,
			onRowClick,
			upsertAppend,
			setRefs
		};
	}
});
</script>
