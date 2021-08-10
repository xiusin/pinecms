<template>
	<cl-crud @load="onLoad">
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
import { defineComponent, inject, reactive } from "vue";
import { useRefs } from "/@/core";
import { CrudLoad, Table, Upsert } from "cl-admin-crud-vue3/types";

export default defineComponent({
	name: "database-backup-list",

	setup() {
		const service = inject<any>("service");
		const { refs, setRefs } = useRefs();

		let file_size_format = function ($size = 0, $dec = 2) {
			const unit = ["B", "KB", "MB", "GB", "TB", "PB"];
			let pos = 0;
			while ($size >= 1024) {
				$size /= 1024;
				pos++;
			}
			return $size.toFixed(2) + unit[pos];
		};

		// 表格配置
		const table = reactive<Table>({
			columns: [
				{
					label: "源名称",
					prop: "original",
					minWidth: 150,
					align: "left"
				},
				{
					label: "图片",
					prop: "url",
					component: ({ h, scope }: any) => {
						return h("img", {
							src: scope.url,
							height: 40
						});
					}
				},
				{
					label: "文件大小",
					prop: "size",
					minWidth: 150,
					component: ({ h, scope }: any) => {
						return file_size_format(scope.size);
					}
				},
				{
					label: "类型",
					prop: "type",
					minWidth: 50
				},
				{
					label: "上传时间",
					prop: "upload_time",
					minWidth: 200,
					showOverflowTooltip: true
				},
				{
					label: "操作",
					type: "op",
					width: 100,
					buttons: ["delete"]
				}
			]
		});

		// crud 加载
		function onLoad({ ctx, app }: CrudLoad) {
			ctx.service(service.system.databaseBackupList).done();
			app.refresh();
		}

		return {
			refs,
			table,
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
