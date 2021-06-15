<template>
	<cl-crud @load="onLoad">
		<el-row type="flex">
			<cl-refresh-btn/>
			<cl-add-btn/>
			<cl-flex1/>
			<cl-query field="param" :list="groupsList"/>
			<cl-search-key/>
		</el-row>

		<el-row>
			<cl-table v-bind="table"/>
		</el-row>

		<el-row type="flex">
			<cl-flex1/>
		</el-row>

		<cl-upsert :ref="setRefs('upsert')" v-bind="upsert" @open="onUpsertOpen"></cl-upsert>
	</cl-crud>
</template>

<script lang="ts">
import {defineComponent, inject, onBeforeMount, reactive, ref} from "vue";
import {useRefs} from "/@/core";
import {CrudLoad, QueryList, Table, Upsert} from "cl-admin-crud-vue3/types";

export default defineComponent({
	name: "sys-setting",
	setup() {
		const service = inject<any>("service");
		const {refs, setRefs} = useRefs();

		let groupsList = ref<QueryList[]>([])

		onBeforeMount(() => {
			service.system.setting.groupList().then((list: QueryList[]) => {
				groupsList.value?.push(...list)
			})
		})

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
					label: "排序",
					prop: "listorder",
					width: 50,
					align: "left",
				},
				{
					label: "名称",
					prop: "form_name",
					width: 120,
					align: "left",
				},
				{
					label: "KEY",
					prop: "key",
					width: 200,
					align: "left"
				},
				{
					label: "分组",
					prop: "group",
					width: 100,
				},
				{
					label: "VALUE",
					prop: "value",
					showOverflowTooltip: true,
					align: "left"
				},
				{
					label: "备注",
					prop: "remark",
					width: 150,
					align: "left",
					showOverflowTooltip: true
				},
				{
					label: "操作",
					type: "op",
					width: 80,
					buttons: ["edit"]
				}
			]
		});
		// 新增编辑配置
		const upsert = reactive<Upsert>({
			width: "1000px",
			items: [
				{
					prop: "form_name",
					label: "名称",
					component: {
						name: "el-input",
						props: {
							placeholder: "请输入名称"
						}
					},
					rules: {
						required: true,
						message: "名称不能为空"
					}
				},
				{
					prop: "key",
					label: "KEY",
					component: {
						name: "el-input",
						props: {
							placeholder: "请输入Key"
						}
					},
					rules: {
						required: true,
						message: "Key不能为空"
					}
				},
				{
					prop: "value",
					label: "VALUE",
					component: ({ h, scope }) => {
						return h("input", {
							props: {
								type: "textarea"
							},
							attrs: {
								placeholder: "请填写内容"
							}
						});
					}
				},
				{
					prop: "listorder",
					label: "排序",
					component: {
						name: "el-input",
						props: {
							placeholder: "请输入Key"
						}
					}
				},
				{
					prop: "remark",
					label: "备注",
					component: {
						name: "el-input",
						props: {
							placeholder: "请输入备注",
							rows: 3,
							type: "textarea"
						}
					}
				}
			]
		});

		// crud 加载
		function onLoad({ctx, app}: CrudLoad) {
			ctx.service(service.system.setting).done();
			app.refresh();
		}

		// 监听打开
		function onUpsertOpen(isEdit: boolean, data: any) {

		}

		return {
			refs,
			table,
			upsert,
			setRefs,
			onLoad,
			groupsList,
			onUpsertOpen
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
