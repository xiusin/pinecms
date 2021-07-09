<template>
	<cl-crud @load="onLoad">
		<el-row>
			<cl-table v-bind="table"/>
		</el-row>
		<el-row type="flex">
			<cl-flex1/>
			<cl-pagination/>
		</el-row>
		<cl-upsert v-bind="upsert" />
	</cl-crud>
</template>

<script lang="ts">
import {defineComponent, h, inject, reactive, ref, resolveComponent} from "vue";
import {useRefs} from "/@/core";
import {CrudLoad, FormItem, FormRef, Table, Upsert} from "cl-admin-crud-vue3/types";

export default defineComponent({
	name: "sys-model",

	setup() {
		const service = inject<any>("service");
		const {refs, setRefs} = useRefs();

		const models = ref([]);

		service.system.assets.select().then((data) => {
			models.value?.push(...data);
		})

		// 表格配置
		const table = reactive<Table>({
			columns: [
				{
					prop: "table_name",
					label: "模型名称",
					width: 150,
					align: "left"
				},
				{
					prop: "table",
					label: "模型表名",
					width: 150,
					align: "left"
				},
				{
					label: "列表模板",
					prop: "fe_tpl_list",
					width: 150,
					align: "left"
				},
				{
					label: "详情模板",
					prop: "fe_tpl_detail",
					width: 150,
					align:"left"
				},
				{
					label: "备注",
					prop: "remark",
					align: "left"
				},
				{
					label: "启用",
					prop: "enabled",
					width: 70,
					dict: [
						{
							label: "启用",
							value: 1,
							type: "primary"
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
					type: "op",
					width: 100,
					buttons: ["slot-model-define","edit", "delete"]
				}
			]
		});

		// 新增编辑配置
		const upsert = reactive<Upsert>({
			width: "1000px",
			items: [
				{
					prop: "table_name",
					label: "模型名称",
					span: 24,
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
					prop: "table",
					label: "模型表名",
					span: 24,
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
					prop: "fe_tpl_list",
					label: "列表模板",
					span: 24,
					component: {
						name: "el-select",
						props: {
							placeholder: "请选择列表模板"
						},
						options:models
					}
				},
				{
					prop: "fe_tpl_detail",
					label: "详情模板",
					span: 24,
					component: {
						name: "el-select",
						props: {
							placeholder: "请选择详情模板"
						},
						options:models
					}
				},
				{
					prop: "remark",
					label: "描述",
					span: 24,
					hidden: ({ scope }: any) => scope.type == 2,
					component: {
						name: "el-input",
						type: "textarea"
					}
				},
				{
					prop: "enabled",
					label: "是否启用",
					span: 24,
					value: 1,
					flex: false,
					component: {
						name: "el-switch",
						"active-value": 1,
						"inactive-value": 0
					}
				}
			]
		});

		// crud 加载
		function onLoad({ctx, app}: CrudLoad) {
			ctx.service(service.system.model).done();
			app.refresh();
		}

		return {
			refs,
			table,
			upsert,
			setRefs,
			onLoad,
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
