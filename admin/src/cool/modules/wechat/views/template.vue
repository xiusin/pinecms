<template>
	<div>
		<cl-crud :ref="setRefs('crud')" @load="onLoad">
			<el-row type="flex">
				<el-button type="warning" size="mini" @click="sync()" icon="el-icon-sort"
					>同步消息模板</el-button
				>
				<cl-refresh-btn />
				<el-button size="mini">
					<el-link
						type="primary"
						icon="el-icon-link"
						size="mini"
						style="font-size: 12px"
						target="_blank"
						href="https://kf.qq.com/faq/170209E3InyI170209nIF7RJ.html"
						>模板管理指引</el-link
					>
				</el-button>
				<cl-flex1 />
				<cl-search-key />
			</el-row>

			<el-row>
				<cl-table v-bind="table">
					<template #slot-config="{ scope }">
						<el-button type="text" @click="templateMsgTaskHandle(scope.row)" size="mini"
							>推送</el-button
						>
						<el-button size="mini" @click="addOrUpdateHandle(scope.row.id)" type="text"
							>配置</el-button
						>
					</template>
					<template #slot-copy="{ scope }">
						<el-button size="mini" @click="copyHandle(scope.row)" type="text"
							>复制</el-button
						>
					</template>
				</cl-table>
			</el-row>

			<el-row type="flex">
				<cl-flex1 />
				<cl-pagination />
			</el-row>
		</cl-crud>

		<add-or-update :visible="addOrUpdateVisible" :ref="setRefs('addOrUpdate')" />
		<template-msg-task :visible="templateMsgTaskVisible" :ref="setRefs('templateMsgTask')" />
	</div>
</template>

<script lang="ts">
import { defineComponent, inject, reactive, ref } from "vue";
import { useRefs } from "/@/core";
import AddOrUpdate from "./msg-template-add-or-update.vue";
import TemplateMsgTask from "../components/template-msg-task.vue";
import { CrudLoad, Table } from "cl-admin-crud-vue3/types";
import { ElMessage, ElMessageBox } from "element-plus";

export default defineComponent({
	name: "wechat-template",
	components: {
		AddOrUpdate,
		TemplateMsgTask
	},
	setup() {
		const service = inject<any>("service");

		const { refs, setRefs }: any = useRefs();

		const table = reactive<Table>({
			columns: [
				{
					prop: "name",
					label: "模版名称",
					width: 140,
					align: "left"
				},
				{
					prop: "title",
					label: "模板标题",
					width: 140
				},
				{
					prop: "templateId",
					label: "模板ID",
					width: 120,
					align: "left"
				},
				{
					prop: "content",
					label: "模板字段",
					showOverflowTooltip: true
				},
				{
					prop: "primary_industry",
					label: "一级行业",
					width: 100
				},
				{
					prop: "deputy_industry",
					label: "二级行业",
					width: 120
				},
				{
					prop: "status",
					label: "是否有效",
					width: 80,
					dict: [
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
					]
				},
				{
					type: "op",
					width: 200,
					buttons: ["slot-config", "slot-copy", "delete"]
				}
			]
		});

		function onLoad({ ctx, app }: CrudLoad) {
			ctx.service(service.wechat.template).done();
			app.refresh();
		}

		const addOrUpdateVisible = ref(false);
		const templateMsgTaskVisible = ref(false);

		function addOrUpdateHandle(id: any) {
			addOrUpdateVisible.value = true;
			refs.value.addOrUpdate.init(id);
		}

		function templateMsgTaskHandle(row: any) {
			templateMsgTaskVisible.value = true;
			refs.value.templateMsgTask.init(row);
		}

		function sync() {
			ElMessageBox.confirm("是否要同步微信模板, 此操作为增量更新？", "提示", {
				type: "warning"
			}).then(() => {
				service.wechat.template
					.sync()
					.then(() => {
						ElMessage.success("同步成功");
					})
					.catch((e: any) => {
						ElMessage.error(e);
					});
			});
		}

		function copyHandle(row: any) {
			service.wechat.template
				.info({ id: row.id })
				.then((data: any) => {
					data.name += "_COPY";
					data.id = 0;
					service.wechat.template.add(data).then(() => {
						ElMessage.success("复制模板成功");
						refs.crud.value.reload();
					});
				})
				.catch((e: any) => {
					ElMessage.error(e);
				});
		}

		return {
			sync,
			copyHandle,
			addOrUpdateHandle,
			templateMsgTaskHandle,
			templateMsgTaskVisible,
			addOrUpdateVisible,
			service,
			refs,
			table,
			setRefs,
			onLoad
		};
	}
});
</script>
