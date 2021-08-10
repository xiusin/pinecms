<template>
	<cl-crud :ref="setRefs('crud')" :on-refresh="onRefresh" @load="onLoad">
		<el-row type="flex">
			<cl-refresh-btn />
			<cl-add-btn />
		</el-row>

		<el-row>
			<cl-table :ref="setRefs('table')" v-bind="table" @row-click="onRowClick">
				<template #column-name="{ scope }">
					<span>{{ scope.row.name }}</span>
					<el-tag
						v-if="!scope.row.ismenu"
						size="mini"
						effect="dark"
						type="danger"
						style="margin-left: 10px"
						>隐
					</el-tag>
				</template>

				<template #column-url="{ scope }">
					<el-button size="mini" round>
						<el-link
							:underline="false"
							:href="scope.row.url"
							target="_blank"
							icon="el-icon-share"
						/>
					</el-button>
				</template>
			</cl-table>
		</el-row>

		<el-row type="flex">
			<cl-flex1 />
			<cl-pagination :props="{ layout: 'total' }" />
		</el-row>

		<!-- 编辑 -->
		<cl-upsert v-bind="upsert" />
	</cl-crud>
</template>

<script lang="ts">
import { useRefs } from "/@/core";
import { deepTree } from "/@/core/utils";
import { defineComponent, inject, reactive, ref } from "vue";
import { CrudLoad, RefreshOp, Table, Upsert } from "cl-admin-crud-vue3/types";

export default defineComponent({
	name: "sys-category",
	setup() {
		const { refs, setRefs } = useRefs();
		const service = inject<any>("service");

		let models = ref([]);

		const catType = [
			{
				label: "栏目",
				value: 0,
				type: "success"
			},
			{
				label: "单页",
				value: 1,
				type: "warning"
			},
			{
				label: "链接",
				value: 2,
				type: "info"
			}
		];

		// crud 加载
		function onLoad({ ctx, app }: CrudLoad) {
			ctx.service(service.system.category).done();
			app.refresh();

			// 读取模型文档列表
			service.system.model.select().then((list: any[]) => {
				models.value.push(...list);
			});
		}

		function onRefresh(_: any, { render }: RefreshOp) {
			service.system.category.list().then(({ list }: any) => {
				console.log("list", list);
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

		function upsertAppend({ type, id }: any) {
			refs.value.crud.rowAppend({
				parentId: id,
				type: type + 1
			});
		}

		// 表格配置
		const table = reactive<Table>({
			props: {
				"row-key": "id",
				"default-sort": {
					prop: "listorder",
					order: "ascending"
				}
			},
			columns: [
				{
					prop: "name",
					label: "名称",
					align: "left",
					width: 200
				},

				{
					prop: "type",
					label: "类型",
					width: 100,
					dict: catType
				},
				{
					prop: "listorder",
					label: "排序号",
					width: 90
				},
				{
					prop: "dir",
					label: "静态目录",
					width: 150,
					align: "left",
					showOverflowTooltip: true
				},
				{
					prop: "url",
					label: "链接",
					width: 120
				},
				{
					prop: "keywords",
					label: "关键字",
					showOverflowTooltip: true,
					width: 200,
					align: "left"
				},
				{
					prop: "description",
					label: "描述",
					align: "left"
				},
				{
					label: "操作",
					type: "op",
					buttons: ["edit", "delete"]
				}
			]
		});

		// 新增、编辑配置
		const upsert = reactive<Upsert>({
			width: "800px",
			items: [
				{
					prop: "type",
					value: 0,
					label: "栏目类型",
					span: 24,
					component: {
						name: "el-radio-group",
						options: catType
					}
				},
				{
					prop: "name",
					label: "栏目名称",
					span: 24,
					component: {
						name: "el-input",
						props: {
							placeholder: "请输入分类名称"
						}
					},
					rules: {
						required: true,
						message: "分类名称不能为空"
					}
				},
				{
					prop: "parentId",
					label: "上级栏目",
					span: 24,
					component: {
						name: "el-category-tree"
					}
				},
				{
					prop: "model_id",
					label: "文档模型",
					hidden: ({ scope }: any) => scope.type == 2,
					span: 24,
					component: {
						name: "el-select",
						props: {
							placeholder: "请选择文档模型"
						},
						options: models
					},
					rules: {
						required: true,
						message: "请选择文档模型"
					}
				},
				{
					prop: "dir",
					label: "静态目录",
					hidden: ({ scope }: any) => scope.type == 2,
					span: 24,
					component: {
						name: "el-input",
						props: {
							placeholder: "请输入静态目录,仅支持0-9,字母和下划线"
						}
					},
					rules: {
						required: true,
						pattern: /^[A-Za-z0-9_-]+$/,
						message: "目录格式填写错误"
					}
				},
				{
					prop: "url",
					label: "地址",
					hidden: ({ scope }: any) => scope.type != 2,
					span: 24,
					component: {
						name: "el-input",
						props: {
							placeholder: "请输入链接地址"
						}
					}
				},
				{
					prop: "keywords",
					label: "关键字",
					span: 24,
					hidden: ({ scope }: any) => scope.type == 2,
					component: {
						name: "el-input",
						props: {
							placeholder: "请输入关键字"
						}
					}
				},
				{
					prop: "description",
					label: "描述",
					span: 24,
					hidden: ({ scope }: any) => scope.type == 2,
					component: {
						name: "el-input",
						type: "textarea"
					}
				},
				{
					prop: "listorder",
					label: "排序",
					span: 6,
					component: {
						name: "el-input-number",
						min: 0
					}
				},
				{
					prop: "ismenu",
					label: "是否显示",
					span: 24,
					value: true,

					flex: false,
					component: {
						name: "el-switch"
					}
				}
			]
		});

		return {
			refs,
			table,
			upsert,
			setRefs,
			onLoad,
			onRefresh,
			onRowClick,
			upsertAppend
		};
	}
});
</script>
