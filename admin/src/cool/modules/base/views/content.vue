<template>
	<div class="system-user">
		<div class="pane">
			<div class="dir">
				<div class="container">
					<el-tree
						ref="treeRef"
						node-key="menuId"
						:data="treeList"
						:props="{ label: 'name', children: 'children' }"
						:highlight-current="true"
						:expand-on-click-node="false"
						:default-expanded-keys="expandedKeys"
						@current-change="onCurrentChange"
					/>
				</div>
			</div>

			<div class="editor">
				<div class="container">
					<cl-crud :ref="setRefs('crud')" @load="onLoad">
						<el-row type="flex">
							<el-button size="small" v-if="catId > 0"
								>当前栏目: {{ catName }}({{ catKey }})</el-button
							>
							<cl-refresh-btn />
							<cl-add-btn />
							<cl-flex1 />
							<cl-search-key />
						</el-row>

						<el-row>
							<cl-table :ref="setRefs('table')" v-bind="table" />
						</el-row>

						<el-row type="flex">
							<cl-flex1 />
							<cl-pagination />
						</el-row>

						<cl-upsert :ref="setRefs('upsert')" :items="upsert.items" />
					</cl-crud>
				</div>
			</div>
		</div>
	</div>
	<cl-form ref="formRef" />
</template>

<script lang="ts">
import { computed, defineComponent, inject, onMounted, reactive, ref } from "vue";
import { useRefs } from "/@/core";
import { deepTree } from "/@/core/utils";
import { QueryList, Table, Upsert } from "cl-admin-crud-vue3/types";

export default defineComponent({
	name: "sys-content",
	setup() {
		const service = inject<any>("service");
		const { refs, setRefs } = useRefs();
		const data = ref([]);
		const modelValue = ref("");
		const formRef = ref();
		// 树形列表
		const menuList = ref<any[]>([]);
		// 展开值
		const expandedKeys = ref<any[]>([]);
		// el-tree 组件
		const treeRef = ref<any>({});
		// 绑定值回调
		function onCurrentChange({ id, catname }: any) {
			catId.value = id;
			catName.value = catname;
			refresh({ cid: id });
		}
		// 树形列表
		const treeList = computed(() => deepTree(menuList.value));

		onMounted(async function () {
			const ret = await service.system.category.list();
			menuList.value = ret.list.filter((e: any) => e.type != 2);
			// 获取模型表结果
			table.value = await service.system.model.modelTable({ mid: 1 });
			table.value?.columns.push({
				label: "操作",
				type: "op",
				buttons: ["edit", "delete"]
			});
		});
		const list = ref<QueryList[]>([]);

		const catId = ref<any>(0);
		const catName = ref<string>("");
		const catKey = ref<string>("");

		const upsert = reactive<Upsert>({
			items: [
				{
					prop: "name",
					label: "名称",
					span: 24,
					component: {
						name: "el-input",
						props: {
							placeholder: "请填写名称"
						}
					},
					rules: {
						required: true,
						message: "名称不能为空"
					}
				},
				{
					prop: "cid",
					label: "字典分类",
					component: {
						name: "el-select",
						props: {
							placeholder: "请选择分类"
						},
						options: menuList
					},
					rules: {
						required: true,
						message: "分类必选"
					}
				},
				{
					prop: "value",
					label: "值",
					span: 24,
					component: {
						name: "el-input",
						props: {
							placeholder: "请填写值",
							type: "textarea",
							rows: 4
						}
					},
					rules: {
						required: true,
						message: "值不能为空"
					}
				},
				{
					prop: "remark",
					label: "备注",
					span: 24,
					component: {
						name: "el-input",
						props: {
							placeholder: "请填写备注",
							type: "textarea",
							rows: 4
						}
					}
				},
				{
					prop: "status",
					label: "状态",
					value: 1,
					component: {
						name: "el-radio-group",
						options: [
							{
								label: "正常",
								value: true
							},
							{
								label: "禁用",
								value: false
							}
						]
					}
				}
			]
		});

		// 刷新列表
		function refresh(params: any) {
			refs.value.crud.refresh(params);
		}

		// 表格配置
		const table = ref<Table>();
		// crud 加载
		async function onLoad({ ctx }: any) {
			ctx.service(service.system.content).done();
		}
		return {
			service,
			data,
			refs,
			expandedKeys,
			setRefs,
			modelValue,
			table,
			menuList,
			upsert,
			list,
			onLoad,
			catName,
			catKey,
			catId,
			formRef,
			treeList,
			refresh,
			treeRef,
			onCurrentChange
		};
	}
});
</script>

<style lang="scss" scoped>
.system-user {
	.pane {
		display: flex;
		height: 100%;
		width: 100%;
		position: relative;
	}

	.dir {
		height: 100%;
		width: 250px;
		padding: 10px;
		max-width: calc(100% - 50px);
		background-color: #fff;
		transition: width 0.3s;
		margin-right: 10px;
		flex-shrink: 0;

		&._collapse {
			margin-right: 0;
			width: 0;
		}
	}

	.editor {
		width: calc(100% - 260px);
		flex: 1;
		background-color: #fff;

		.header {
			display: flex;
			align-items: center;
			justify-content: center;
			height: 40px;
			position: relative;
			background-color: #fff;

			span {
				font-size: 14px;
				white-space: nowrap;
				overflow: hidden;
			}

			.icon {
				position: absolute;
				left: 0;
				top: 0;
				font-size: 18px;
				cursor: pointer;
				background-color: #fff;
				height: 40px;
				width: 80px;
				line-height: 40px;
				padding-left: 10px;
			}
		}
	}

	.dept,
	.user {
		overflow: hidden;

		.container {
			height: calc(100% - 40px);
		}
	}

	@media only screen and (max-width: 768px) {
		.dept {
			width: calc(100% - 100px);
		}
	}
}
</style>
