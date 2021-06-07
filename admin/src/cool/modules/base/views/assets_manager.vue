<template>
	<cl-crud @load="onLoad">
		<el-row type="flex">
			<cl-refresh-btn />
<!--			<cl-add-btn />-->
		</el-row>

		<el-row>
			<cl-table v-bind="table" />
		</el-row>


		<cl-upsert :ref="setRefs('upsert')" v-bind="upsert" @open="onUpsertOpen">
			<template #slot-content="{ scope }">
				<div v-for="(item, index) in tab.list" :key="index" class="editor">
					<template v-if="tab.index == index">
						<el-button class="change-btn" size="mini" @click="changeTab(item.to)">{{
							item.label
						}}</el-button>

						<component :is="item.component" v-model="scope.data" height="300px" />
					</template>
				</div>
			</template>
		</cl-upsert>
	</cl-crud>
</template>

<script lang="ts">
import { ElMessageBox } from "element-plus";
import { defineComponent, inject, nextTick, reactive } from "vue";
import { useRefs } from "/@/core";
import { CrudLoad, Table, Upsert } from "cl-admin-crud-vue3/types";

export default defineComponent({
	name: "sys-assets-manager",

	setup() {
		const service = inject<any>("service");
		const { refs, setRefs } = useRefs();

		// 选项卡
		const tab = reactive<any>({
			index: null,

			list: [
				{
					label: "切换富文本编辑器",
					to: 1,
					component: "cl-codemirror"
				},
				{
					label: "切换代码编辑器",
					to: 0,
					component: "cl-editor-quill"
				}
			]
		});

		// 表格配置
		const table = reactive<Table>({
			columns: [
				{
					label: "文件名",
					prop: "name",
					minWidth: 150
				},
				{
					label: "文件大小",
					prop: "size",
					minWidth: 150
				},
				{
					label: "更新时间",
					prop: "updated",
					minWidth: 200,
					showOverflowTooltip: true
				},
				{
					label: "操作",
					type: "op",
					width: 200,
					buttons: ["edit"]
				}
			]
		});

		// 新增编辑配置
		const upsert = reactive<Upsert>({
			width: "1000px",
			items: [
				{
					prop: "name",
					label: "名称"
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
			ctx.service(service.system.assets).done();
			app.refresh();
		}

		// 切换编辑器
		function changeTab(i: number) {
			ElMessageBox.confirm("切换编辑器会清空输入内容，是否继续？", "提示", {
				type: "warning"
			})
				.then(() => {
					tab.index = i;
					refs.value.upsert.setForm("data", "");
				})
				.catch(() => null);
		}

		// 监听打开
		function onUpsertOpen(isEdit: boolean, data: any) {
			tab.index = null;

			nextTick(() => {
				if (isEdit) {
					tab.index = /<*>/g.test(data.data) ? 1 : 0;
				} else {
					tab.index = 1;
				}
			});
		}

		return {
			refs,
			tab,
			table,
			upsert,
			setRefs,
			onLoad,
			changeTab,
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
