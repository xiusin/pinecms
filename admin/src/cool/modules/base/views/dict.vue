<template>
	<div class="system-user">
		<div class="pane">
			<div class="dept">
				<div class="container">
					<cl-crud :ref="setRefs('categoryCrud')" @load="onCategoryLoad">
						<el-row type="flex">
							<cl-refresh-btn />
							<cl-add-btn />
						</el-row>

						<el-row>
							<cl-table
								:ref="setRefs('categoryTable')"
								v-bind="categoryTable"
								@selection-change="onSelectionChange"
							>
								<template #slot-btn="{ scope }">
									<el-button
										@click="
											changeCategory(
												scope.row.id,
												scope.row.key,
												scope.row.name
											)
										"
										type="text"
										size="mini"
										>属性
									</el-button>
								</template>
							</cl-table>
						</el-row>
						<cl-upsert :ref="setRefs('categoryUpsert')" :items="categoryUpsert.items" />
					</cl-crud>
				</div>
			</div>

			<!-- 成员列表 -->
			<div class="user">
				<div class="container">
					<cl-crud :ref="setRefs('crud')" :on-refresh="onRefresh" @load="onLoad">
						<el-row type="flex">
							<el-button size="small" v-if="catId > 0"
								>当前分类: {{ catName }}({{ catKey }})</el-button
							>
							<cl-refresh-btn />
							<cl-add-btn />
							<cl-flex1 />
							<cl-search-key />
						</el-row>

						<el-row>
							<cl-table
								:ref="setRefs('table')"
								v-bind="table"
								@selection-change="onSelectionChange"
							>
								<!-- 头像 -->
								<template #column-headImg="{ scope }">
									<cl-avatar
										shape="square"
										size="medium"
										:src="scope.row.headImg"
										:style="{ margin: 'auto' }"
									/>
								</template>

								<!-- 权限 -->
								<template #column-roleName="{ scope }">
									<el-tag
										v-for="(item, index) in scope.row.roleNameList"
										:key="index"
										disable-transitions
										size="small"
										effect="dark"
										style="margin: 2px"
										>{{ item }}
									</el-tag>
								</template>
							</cl-table>
						</el-row>

						<el-row type="flex">
							<cl-flex1 />
							<cl-pagination />
						</el-row>

						<cl-upsert
							:ref="setRefs('upsert')"
							:items="upsert.items"
							:on-open="onOpenUpsert"
						/>
					</cl-crud>
				</div>
			</div>
		</div>
	</div>
</template>

<script lang="ts">
import { defineComponent, inject, reactive, ref } from "vue";
import { useRefs } from "/@/core";
import { QueryList, Table, Upsert } from "cl-admin-crud-vue3/types";
import { ElMessage } from "element-plus";

export default defineComponent({
	name: "sys-dict",

	setup() {
		const service = inject<any>("service");

		const { refs, setRefs } = useRefs();

		// 选择项
		const selects = reactive<any>({
			dept: {},
			ids: []
		});

		// 表格配置
		const table = reactive<Table>({
			props: {
				"default-sort": {
					prop: "id",
					order: "descending"
				}
			},
			columns: [
				{
					prop: "name",
					label: "名称",
					width: 170,
					align: "left"
				},
				{
					prop: "value",
					align: "left",
					label: "值"
				},
				{
					prop: "cat_name",
					label: "所属分类",
					align: "left"
				},
				{
					prop: "status",
					label: "状态",
					minWidth: 50,
					dict: [
						{
							label: "正常",
							value: true,
							type: "success"
						},
						{
							label: "禁用",
							value: 0,
							type: "danger"
						}
					]
				},
				{
					type: "op",
					buttons: ["edit", "delete"],
					width: 120
				}
			]
		});

		const categoryTable = reactive<Table>({
			props: {
				"default-sort": {
					prop: "id",
					order: "descending"
				}
			},
			columns: [
				{
					prop: "name",
					label: "字典名称",
					width: 150,
					align: "left"
				},
				{
					prop: "key",
					label: "字典标识",
					align: "left",
				},
				{
					prop: "status",
					label: "状态",
					minWidth: 50,
					dict: [
						{
							label: "正常",
							value: true,
							type: "success"
						},
						{
							label: "禁用",
							value: 0,
							type: "danger"
						}
					]
				},
				{
					type: "op",
					buttons: ["slot-btn", "edit", "delete"],
					width: 150
				}
			]
		});

		const categoryUpsert = reactive<Upsert>({
			items: [
				{
					prop: "name",
					label: "字典名称",
					span: 12,
					component: {
						name: "el-input",
						props: {
							placeholder: "请填写字典名称"
						}
					},
					rules: {
						required: true,
						message: "字典名称不能为空"
					}
				},
				{
					prop: "key",
					label: "字典标识",
					span: 12,
					component: {
						name: "el-input",
						props: {
							placeholder: "请填写字典标识"
						}
					},
					rules: {
						required: true,
						message: "字典标识不能为空"
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
					value: true,
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
						options: list
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
					value: true,
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

		function changeCategory(categoryId: any, categoryKey: string, categoryName: string) {
			catId.value = categoryId;
			catKey.value = categoryKey;
			catName.value = categoryName;
			list.value = [{ label: categoryName, value: categoryId }];
			refresh({ cid: categoryId });
		}

		// crud 加载
		async function onLoad({ ctx, app }: any) {
			ctx.service(service.system.dict).done();
			const cats = await service.system.dictCategory.list({ size: 1 });
			if (cats.list.length) {
				catId.value = cats.list[0].id;
				catKey.value = cats.list[0].key;
				catName.value = cats.list[0].name;
				app.refresh({ cid: catId });
			}
		}

		// crud 加载
		function onCategoryLoad({ ctx, app }: any) {
			ctx.service(service.system.dictCategory).done();
			app.refresh();
		}

		// 刷新列表
		function refresh(params: any) {
			refs.value.crud.refresh(params);
		}

		// 刷新监听
		async function onRefresh(params: any, { next, render }: any) {
			const { list } = await next(params);

			render(
				list.map((e: any) => {
					if (e.roleName) {
						e.roleNameList = e.roleName.split(",");
					}

					e.status = Boolean(e.status);

					return e;
				})
			);
		}

		// 多选监听
		function onSelectionChange(selection: any[]) {
			selects.ids = selection.map((e) => e.id);
		}

		function onOpenUpsert() {
			if (catId.value == 0) {
				ElMessage.error("请先选择要添加的目标分类");
				arguments[2].close();
			}
		}

		return {
			service,
			refs,
			selects,
			categoryUpsert,
			table,
			upsert,
			setRefs,
			changeCategory,
			onLoad,
			onCategoryLoad,
			refresh,
			onRefresh,
			catId,
			catName,
			catKey,
			onSelectionChange,
			categoryTable,
			onOpenUpsert
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

	.dept {
		height: 100%;
		width: 450px;
		max-width: calc(100% - 50px);
		background-color: #fff;
		transition: width 0.3s;
		margin-right: 10px;
		flex-shrink: 0;

		& ._collapse {
			margin-right: 0;
			width: 0;
		}
	}

	.user {
		width: calc(100% - 460px);
		flex: 1;

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
