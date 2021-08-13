<template>
	<cl-crud :ref="setRefs('crud')" @load="onLoad">
		<el-row type="flex">
			<cl-refresh-btn />

			<el-button
				v-permission="service.system.errorLog.permission.clear"
				size="mini"
				type="danger"
				@click="clear"
			>
				清空
			</el-button>
			<cl-flex1 />
			<cl-query field="params.method" :list="methods" />
			<cl-search-key placeholder="请输入请求地址, 参数，ip地址" />
		</el-row>

		<el-row>
			<cl-table v-bind="table" />
		</el-row>

		<el-row type="flex">
			<cl-flex1 />
			<cl-pagination />
		</el-row>
	</cl-crud>
</template>

<script lang="ts">
import { defineComponent, inject, reactive, ref } from "vue";
import { ElMessage, ElMessageBox } from "element-plus";
import { useRefs } from "/@/core";
import { CrudLoad, Table } from "cl-admin-crud-vue3/types";

export default defineComponent({
	name: "sys-errlog",

	setup() {
		const service = inject<any>("service");
		const { refs, setRefs }: any = useRefs();

		const day = ref<number>(1);

		// cl-table 配置
		const table = reactive<Table>({
			props: {
				"default-sort": {
					prop: "id",
					order: "descending"
				}
			},
			columns: [
				{
					type: "index",
					label: "#",
					width: 60
				},
				{
					prop: "time",
					label: "日志时间",
					width: 200
				},
				{
					prop: "level",
					label: "日志级别",
					width: 100
				},
				{
					prop: "message",
					label: "信息",
					showOverflowTooltip: true
				},
			]
		});

		// crud 加载
		function onLoad({ ctx, app }: CrudLoad) {
			ctx.service(service.system.errorLog).done();
			app.refresh();
		}

		// 清空日志
		function clear() {
			ElMessageBox.confirm("是否要清空日志", "提示", {
				type: "warning"
			})
				.then(() => {
					service.system.errorLog
						.clear()
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
			service,
			refs,
			day,
			table,
			setRefs,
			onLoad,
			clear
		};
	}
});
</script>
