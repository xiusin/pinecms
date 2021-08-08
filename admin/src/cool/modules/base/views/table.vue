<template>
	<div>
		<cl-crud @load="onLoad" :on-refresh="onRefresh">
			<el-row type="flex">
				<cl-refresh-btn />
				<cl-add-btn />
				<cl-upload-space accept=".jpg,.png,.txt" />
				<vue-ueditor-wrap :config="{ UEDITOR_HOME_URL: '/UEditor/' }" />
			</el-row>
			<el-row>
				<cl-table v-bind="table" />
			</el-row>
			<el-row type="flex">
				<cl-flex1 />
				<cl-pagination />
			</el-row>
			<cl-upsert v-bind="upsert">
				<template #slot-dictKey="{ scope }">
					<el-select
						:filterable="true"
						placeholder="请选择字典类型"
						:automatic-dropdown="true"
						size="mini"
					>
						<el-option
							v-for="(item, idx) in dictCatRef"
							:key="idx + '-' + item.value"
							:value="item.value"
							:label="item.label"
						/>
					</el-select>
				</template>
			</cl-upsert>
		</cl-crud>

		<!-- 表单 -->
		<cl-form :ref="setRefs('form')" />
	</div>
</template>

<script lang="ts">
import { defineComponent, inject, reactive, ref } from "vue";
import { useRefs } from "/@/core";
import { CrudLoad, Table, Upsert } from "cl-admin-crud-vue3/types";
import { useRoute } from "vue-router";
import ClUploadSpace from "../../upload/components/space/index.vue";
import VueUeditorWrap from "vue-ueditor-wrap";

export default defineComponent({
	name: "sys-table",
	components: { ClUploadSpace, VueUeditorWrap },
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

		const dictCatRef = ref([]);

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
					label: "名称",
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
					label: "类型",
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
					label: "字段",
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
						name: "el-checkbox",
						"true-label": true,
						"false-label": false
					}
				},
				{
					prop: "required",
					label: "必填",
					span: 6,
					value: true,
					flex: false,
					component: {
						name: "el-checkbox",
						"true-label": true,
						"false-label": false
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
					value: 150,
					component: {
						name: "el-slider",
						step: 10,
						"show-stops": true,
						max: 500,
						min: 80
					}
				},
				{
					prop: "center",
					label: "居中",
					span: 4,
					value: true,
					flex: false,
					component: {
						name: "el-checkbox",
						"true-label": true,
						"false-label": false
					}
				},
				{
					prop: "sortable",
					label: "排序",
					span: 4,
					value: true,
					flex: false,
					component: {
						name: "el-checkbox",
						"true-label": true,
						"false-label": false
					}
				},
				{
					prop: "searchable",
					label: "搜索",
					span: 4,
					value: true,
					flex: false,
					component: {
						name: "el-checkbox",
						"true-label": true,
						"false-label": false
					}
				},
				{
					prop: "list_visible",
					label: "列表",
					span: 4,
					value: true,
					flex: false,
					component: {
						name: "el-checkbox",
						"true-label": true,
						"false-label": false
					}
				},
				{
					prop: "visible",
					label: "表单",
					span: 4,
					value: true,
					flex: false,
					component: {
						name: "el-checkbox",
						"true-label": true,
						"false-label": false
					}
				},
				{
					prop: "show_component",
					label: "自定义渲染",
					span: 4,
					value: true,
					flex: false,
					component: {
						name: "el-checkbox",
						"true-label": true,
						"false-label": false
					}
				},
				{
					prop: "component",
					label: "渲染配置",
					span: 24,
					hidden: ({ scope }: any) => !scope.show_component,
					component: {
						name: "el-input",
						type: "textarea",
						row: 4,
						placeholder: "自定义渲染, 用于替换默认渲染信息, Render函数, json配置等"
					}
				},
				{
					prop: "is_dict",
					label: "字典",
					span: 4,
					value: true,
					flex: false,
					hidden: ({ scope }: any) => fn(scope),
					component: {
						name: "el-checkbox",
						"true-label": true,
						"false-label": false
					}
				},
				{
					prop: "dict_key",
					label: "类型",
					span: 20,
					hidden: ({ scope }: any) => {
						if (!fn(scope)) {
							return !scope.is_dict;
						} else {
							return true;
						}
					},
					component: {
						name: "el-select",
						options: dictCatRef
					}
				},
				{
					prop: "datasource",
					label: "数据源",
					span: 20,
					hidden: ({ scope }: any) => {
						if (!fn(scope)) {
							return scope.is_dict;
						} else {
							return true;
						}
					},
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

		const fn = (scope: any) => {
			return (
				scope.field_type != 5 &&
				scope.field_type != 6 &&
				scope.field_type != 7 &&
				scope.field_type != 8
			);
		};

		// 拉取字典
		service.system.dictCategory.select().then((data: any[]) => {
			dictCatRef.value = data;
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
			dictCatRef,
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
