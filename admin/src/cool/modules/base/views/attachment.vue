<template>
	<cl-crud @load="onLoad">
		<el-row type="flex">
			<cl-refresh-btn/>
		</el-row>

		<el-row>
			<cl-table v-bind="table"/>
		</el-row>

		<el-row type="flex">
			<cl-flex1/>
			<cl-pagination/>
		</el-row>
	</cl-crud>

	<cl-dialog v-model="preview.visible" title="图片预览" :props="{ width: previewWidth }">
		<img style="width: 100%" :src="preview.url" alt=""/>
	</cl-dialog>
</template>

<script lang="ts">
import {defineComponent, inject, onMounted, reactive} from "vue";
import {ElImage} from "element-plus"
import {useRefs} from "/@/core";
import {CrudLoad, Table} from "cl-admin-crud-vue3/types";

export default defineComponent({
	name: "attachment",
	components: {
		ElImage
	},
	setup() {
		const service = inject<any>("service");
		const {refs, setRefs} = useRefs();
		let preview = {
			visible: false,
			url: ""
		};

		let urls: string[] = [];

		let file_size_format = function ($size = 0, $dec = 2) {
			const unit = ["B", "KB", "MB", "GB", "TB", "PB"];
			let pos = 0;
			while ($size >= 1024) {
				$size /= 1024;
				pos++;
			}
			return $size.toFixed(2) + unit[pos];
		};

		let previewWidth = {
			type: String,
			default: "500px"
		};

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
					component: {
						name: "el-image",
						props: {
							style: {
								width: 40,
								height: 40
							},
							fit: "contain",
						}
					}
				},
				{
					label: "文件大小",
					prop: "size",
					width: 100,
					component: ({h, scope}: any) => {
						return file_size_format(scope.size);
					}
				},
				{
					label: "类型",
					prop: "type",
					width: 100
				},
				{
					label: "所属分类",
					prop: "classifyId",
					component: {
						name: "el-image",
						props: {
							style: {
								width: 40,
								height: 40
							},
							fit: "contain",
						}
					}
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

		onMounted(() => {

		})

		// crud 加载
		function onLoad({ctx, app}: CrudLoad) {
			ctx.service(service.system.attachment).done();
			app.refresh();
		}

		return {
			refs,
			table,
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
