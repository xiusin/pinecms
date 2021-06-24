<template>
	<cl-crud @load="onLoad">
		<el-row type="flex">
			<cl-upload v-model="urls" multiple :limit="5" accept="*" list-type="text"/>
		</el-row>

		<el-row>
			<cl-table v-bind="table" />
		</el-row>

		<el-row type="flex">
			<cl-flex1 />
			<cl-pagination />
		</el-row>
	</cl-crud>


	<cl-dialog
		v-model="preview.visible"
		title="图片预览"
		:props="{width: previewWidth}"
	>
		<img style="width: 100%" :src="preview.url" alt="" />
	</cl-dialog>

</template>

<script lang="ts">
import { defineComponent, inject, reactive } from "vue";
import { useRefs } from "/@/core";
import { CrudLoad, Table, Upsert } from "cl-admin-crud-vue3/types";

export default defineComponent({
	name: "attachment",

	setup() {
		const service = inject<any>("service");
		const { refs, setRefs } = useRefs();
		let preview = {
			visible: false,
			url: "",
		}

		let urls = []

		let file_size_format = function ($size = 0, $dec = 2) {
			const unit = ["B", "KB", "MB", "GB", "TB", "PB"];
			let pos = 0;
			while ($size >= 1024) {
				$size /= 1024;
				pos++;
			}
			return  $size.toFixed(2)  + unit[pos];
		}

		let previewWidth = {
			type: String,
			default: "500px"
		}

		// 表格配置
		const table = reactive<Table>({
			columns: [
				{
					label: "源名称",
					prop: "original",
					minWidth: 150,
					align: "left",
				},
				{
					label: "图片",
					prop: "url",
					component: ({h, scope}) => {
						return h("img", {
							src: scope.url,
							height: 40
						});
					},
				},
				{
					label: "文件大小",
					prop: "size",
					minWidth: 150,
					component: ({h, scope}) => {
						return file_size_format(scope.size);
					},
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
					},
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
			ctx.service(service.system.attachment).done();
			app.refresh();
		}

		return {
			refs,
			table,
			upsert,
			setRefs,
			onLoad,
			preview,
			previewWidth,
			urls
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
