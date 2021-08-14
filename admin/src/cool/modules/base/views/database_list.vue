<template>
	<cl-crud @load="onLoad">
		<el-row>
			<el-button size="mini" type="success" @click="backup"> 备份 </el-button>
			<el-button size="mini" type="danger" @click="optimize"> 优化 </el-button>
			<el-button size="mini" type="danger" @click="repair"> 修复 </el-button>
		</el-row>
		<el-row>
			<cl-table v-bind="table" @selection-change="onSelectionChange" />
		</el-row>
	</cl-crud>
</template>

<script lang="ts">
import { defineComponent, inject, reactive, ref } from "vue";
import { useRefs } from "/@/core";
import { CrudLoad, Table } from "cl-admin-crud-vue3/types";
import { ElMessage, ElMessageBox } from "element-plus";

export default defineComponent({
	name: "database-list",

	setup() {
		const service = inject<any>("service");

		const { refs, setRefs } = useRefs();

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
					align: "left",
					component: {
						"name": "el-input",
						"props": {
							"size": "mini",
							"clearable": true,
							onChange: (val, val1) => {
								console.log(val, val1)
							},
						}
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
							ElMessage.success("备份成功");
						})
						.catch((err: string) => {
							ElMessage.error(err);
						});
				})
				.catch(() => null);
		}


		function repair() {
			if (selects.value.length == 0) {
				ElMessage.warning("请先选择要修复的表")
				return
			}
			service.system.databaseList.repair({
				"tables": selects.value
			}).then((data) => {
				ElMessage.success(data)
			}).catch((e) => {
				ElMessage.error(e)
			})
		}

		function optimize() {
			if (selects.value.length == 0) {
				ElMessage.warning("请先选择要优化的表")
				return
			}
			service.system.databaseList.optimize({
				"tables": selects.value
			}).then((data) => {
				ElMessage.success(data)
			}).catch((e) => {
				ElMessage.error(e)
			})
		}

		return {
			refs,
			table,
			backup,
			repair,
			optimize,
			setRefs,
			onLoad,
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
