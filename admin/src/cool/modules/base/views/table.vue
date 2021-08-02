<template>
	<cl-crud @load="onLoad" :on-refresh="onRefresh">
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
		<cl-upsert v-bind="upsert" />
	</cl-crud>

	<!-- 表单 -->
	<cl-form :ref="setRefs('form')" />
</template>

<script lang="ts">
import { defineComponent, inject, reactive, ref } from "vue";
import { useRefs } from "/@/core";
import { CrudLoad, Table, Upsert } from "cl-admin-crud-vue3/types";
import { useRoute } from "vue-router";

export default defineComponent({
	name: "sys-table",
	setup() {
		const service = inject<any>("service");
		const { refs, setRefs } = useRefs();
		const fields = ref<any[]>([]);

		service.system.table.fields().then((data: any) => {
			for (const idx in data) {
				fields.value.push({
					label: data[idx].name,
					value: data[idx].id,
					type: ["", "success", "info", "danger", "warning"][data[idx].id % 5]
				});
			}
		});

		const dict = [
			{
				label: "是",
				value: true,
				type: "success"
			},
			{
				label: "否",
				value: false,
				type: "danger"
			}
		];

		const table = reactive<Table>({
			props: {
				"default-sort": {
					prop: "listorder",
					order: "ascending"
				}
			},
			columns: [
				{
					prop: "form_name",
					label: "表单名称",
					width: 150,
					align: "left"
				},
				{
					prop: "table_field",
					label: "字段名称",
					width: 150,
					align: "left"
				},
				{
					prop: "field_type",
					label: "字段类型",
					width: 100,
					dict: fields
				},
				{
					label: "必填",
					prop: "required",
					width: 60,
					dict: dict
				},
				{
					label: "搜索",
					prop: "searchable",
					width: 60,
					dict: dict
				},
				{
					label: "排序",
					prop: "sortable",
					width: 60,
					dict: dict
				},
				{
					label: "列表",
					prop: "list_visible",
					width: 60,
					dict: dict
				},
				{
					label: "表单",
					prop: "visible",
					width: 60,
					dict: dict
				},
				// {
				// 	prop: "required_tips",
				// 	label: "验证提醒",
				// 	align: "left"
				// },
				// {
				// 	prop: "",
				// 	label: "验证规则",
				// 	width: 150,
				// 	align: "left"
				// },
				{
					prop: "default",
					label: "默认值",
					align: "left"
				},
				{
					label: "状态",
					prop: "status",
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
					buttons: ["edit"]
				}
			]
		});
		const route = useRoute();
		const upsert = reactive<Upsert>({
			width: "1000px",
			items: [
				{
					prop: "form_name",
					label: "表单名称",
					span: 12,
					component: {
						name: "el-input",
						props: {
							placeholder: "名称"
						}
					},
					rules: {
						required: true,
						message: "请填写模型名称"
					}
				},
				{
					prop: "field_type",
					label: "字段类型",
					span: 12,
					component: {
						name: "el-select",
						props: {
							placeholder: "请选择字段类型"
						},
						options: fields
					},
					rules: {
						required: true,
						message: "字段类型不能为空"
					}
				},
				{
					prop: "table_field",
					label: "表字段",
					span: 18,
					component: {
						name: "el-input",
						props: {
							placeholder: "请输入模型表名"
						}
					},
					rules: {
						required: true,
						message: "请输入模型表名"
					}
				},
				{
					prop: "main_table_field",
					label: "主表字段",
					span: 6,
					value: true,
					flex: false,
					component: {
						name: "el-switch",
						"active-value": true,
						"inactive-value": false
					}
				},
				{
					prop: "required",
					label: "必填",
					span: 6,
					value: true,
					flex: false,
					component: {
						name: "el-switch",
						"active-value": true,
						"inactive-value": false
					}
				},
				{
					prop: "required_tips",
					label: "必填提醒",
					span: 18,
					component: {
						name: "el-input",
						props: {
							placeholder: "请输入必填提醒"
						}
					}
				},
				{
					prop: "default",
					label: "默认值",
					span: 24,
					hidden: ({ scope }: any) => scope.type == 2,
					component: {
						name: "el-input",
						placeholder: "请输入默认值"
					}
				},
				{
					prop: "span",
					label: "表单宽度",
					span: 12,
					value: 12,
					component: {
						name: "el-slider",
						step: 4,
						"show-stops": true,
						max: 24,
						min: 4
					}
				},
				{
					prop: "list_width",
					label: "列表宽度",
					span: 12,
					value: 12,
					component: {
						name: "el-slider",
						step: 4,
						"show-stops": true,
						max: 24,
						min: 4
					}
				},

				{
					prop: "component",
					label: "自定义渲染",
					span: 24,
					component: {
						name: "el-input",
						type: "textarea",
						row: 4,
						placeholder: "自定义渲染, 用于替换默认渲染信息"
					}
				},

				{
					prop: "sortable",
					label: "允许排序",
					span: 24,
					value: true,
					flex: false,
					component: {
						name: "el-switch",
						"active-value": true,
						"inactive-value": false
					}
				},
				{
					prop: "searchable",
					label: "允许搜索",
					span: 24,
					value: true,
					flex: false,
					component: {
						name: "el-switch",
						"active-value": true,
						"inactive-value": false
					}
				},
				{
					prop: "list_visible",
					label: "列表显示",
					span: 24,
					value: true,
					flex: false,
					component: {
						name: "el-switch",
						"active-value": true,
						"inactive-value": false
					}
				},
				{
					prop: "visible",
					label: "表单显示",
					span: 24,
					value: true,
					flex: false,
					component: {
						name: "el-switch",
						"active-value": true,
						"inactive-value": false
					}
				},
				{
					prop: "datasource",
					label: "数据源",
					span: 24,
					hidden: ({ scope }: any) => scope.type == 2,
					component: {
						name: "el-input",
						type: "textarea",
						placeholder: "仅在下拉 级联， 字典类型接口可用"
					}
				},
				{
					prop: "validator",
					label: "验证规则",
					span: 24,
					hidden: ({ scope }: any) => scope.type == 2,
					component: {
						name: "el-input",
						type: "textarea",
						placeholder: "字段校验规则"
					}
				},
				{
					prop: "status",
					label: "状态",
					span: 24,
					value: true,
					flex: false,
					component: {
						name: "el-switch",
						"active-value": true,
						"inactive-value": false
					}
				}
			]
		});

		async function onRefresh(params: any, { next, render }: any) {
			let { list } = await next({
				...params,
				mid: parseInt(route.query.mid)
			});
			render(list);
		}

		function onLoad({ ctx, app }: CrudLoad) {
			ctx.service(service.system.table).done();
			app.refresh();
		}

		return {
			refs,
			table,
			upsert,
			onRefresh,
			setRefs,
			onLoad
		};
	}
});
</script>

<style lang="scss" scoped>
.change-btn {
	display: flex;
	position: absolute;
	right: 10px;
	bottom: 10px;
	z-index: 9;
}

.editor {
	transition: all 0.3s;
}
</style>
