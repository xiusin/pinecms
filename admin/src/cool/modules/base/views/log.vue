<template>
	<cl-crud :ref="setRefs('crud')" @load="onLoad">
		<el-row type="flex">
			<cl-refresh-btn />

			<el-button
				v-permission="service.system.log.permission.clear"
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
	name: "sys-log",

	setup() {
		const service = inject<any>("service");
		const { refs, setRefs }: any = useRefs();

		const methods = [
			{
				label: "GET",
				value: "GET"
			},
			{
				label: "POST",
				value: "POST"
			}
		];

		// 天数
		const day = ref<number>(1);

		// cl-table 配置
		const table = reactive<Table>({
			"context-menu": ["refresh"],
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
					prop: "userid",
					label: "用户ID",
					width: 60
				},
				{
					prop: "username",
					label: "昵称",
					width: 150
				},
				{
					prop: "uri",
					label: "请求地址",
					width: 140,
					showOverflowTooltip: true,
					align: "left"
				},
				{
					prop: "method",
					label: "Method",
					width: 140,
					showOverflowTooltip: true,
					align: "center"
				},
				{
					prop: "params",
					label: "参数",
					align: "left",
					minWidth: 200,
					showOverflowTooltip: true
				},
				{
					prop: "ip",
					label: "ip",
					width: 120
				},
				{
					prop: "ipAddr",
					label: "ip地址",
					width: 150,
					component: ({ h, scope }: any) => {
						if (scope.ip == "127.0.0.1") {
							return "本机";
						} else {
							return "未知";
						}
					}
				},
				{
					prop: "time",
					label: "请求时间",
					width: 200,
					sortable: true
				}
			]
		});

		// crud 加载
		function onLoad({ ctx, app }: CrudLoad) {
			ctx.service(service.system.log).done();
			app.refresh();
		}

		// 保存天数
		function saveDay() {
			service.system.log.setKeep(day.value).then(() => {
				ElMessage.success("保存成功");
			});
		}

		// 清空日志
		function clear() {
			ElMessageBox.confirm("是否要清空日志", "提示", {
				type: "warning"
			})
				.then(() => {
					service.system.log
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

		// 获取天数
		service.system.log.getKeep().then((res: number) => {
			day.value = Number(res);
		});

		return {
			service,
			refs,
			day,
			table,
			setRefs,
			methods,
			onLoad,
			saveDay,
			clear
		};
	}
});
</script>
