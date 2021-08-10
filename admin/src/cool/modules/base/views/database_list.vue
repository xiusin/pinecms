<template>
	<cl-crud @load="onLoad">
		<el-row>
			<el-button size="mini" type="success" @click="backup"> 备份 </el-button>
			<el-button size="mini" type="danger"> 优化 </el-button>
			<el-button size="mini" type="danger"> 修复 </el-button>
		</el-row>
		<el-row>
			<cl-table v-bind="table" @selection-change="onSelectionChange" />
		</el-row>
	</cl-crud>

	<cl-form ref="sqlFormRef"></cl-form>
</template>

<script lang="ts">
import { defineComponent, inject, reactive, ref } from "vue";
import { useRefs } from "/@/core";
import { CrudLoad, FormRef, Table, Upsert } from "cl-admin-crud-vue3/types";
import { ElMessage, ElMessageBox } from "element-plus";

export default defineComponent({
	name: "database-list",

	setup() {
		const service = inject<any>("service");

		const { refs, setRefs } = useRefs();

		const sqlFormRef = ref<FormRef>();

		// 选择项
		const selects = ref<any>([]);
		// 表格配置
		const table = reactive<Table>({
			columns: [
				{
					type: "selection",
					width: 60
				},
				{
					label: "表名",
					prop: "id",
					width: 220,
					align: "left"
				},
				{
					label: "存储引擎",
					prop: "engine",
					width: 150,
					align: "left"
				},
				{
					label: "记录数",
					prop: "total",
					width: 100
				},
				{
					label: "备注",
					prop: "comment",
					align: "left"
				}
			]
		});

		// 新增编辑配置
		const upsert = reactive<Upsert>({
			width: "1000px",
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
					}
				},
				{
					prop: "content",
					label: "内容",
					component: {
						name: "slot-content"
					}
				}
			]
		});

		// crud 加载
		function onLoad({ ctx, app }: CrudLoad) {
			ctx.service(service.system.databaseList).done();
			app.refresh();
		}
		// 多选监听
		function onSelectionChange(selection: any[]) {
			selects.value = selection.map((e) => e.id);
		}

		function backup() {
			ElMessageBox.confirm("是否要备份数据库？", "提示", {
				type: "warning"
			})
				.then(() => {
					service.system.databaseList
						.backup({
							ids: selects.value
						})
						.then(() => {
							ElMessage.success("清空成功");
							refs.value.crud.refresh();
						})
						.catch((err: string) => {
							ElMessage.error(err);
						});
				})
				.catch(() => null);
		}

		return {
			refs,
			table,
			upsert,
			backup,
			setRefs,
			onLoad,
			sqlFormRef,
			onSelectionChange
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
