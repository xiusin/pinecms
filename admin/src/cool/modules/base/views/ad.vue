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
								:ref="setRefs('adSpaceTable')"
								v-bind="adSpaceTable"
								@selection-change="onSelectionChange"
							>
								<template #column-name="{ scope }">
									{{ scope.row.name
									}}<el-popover
										placement="top"
										:width="200"
										trigger="hover"
										:content="scope.row.remark"
									>
										<template #reference>
											<el-icon>
												<picture>查看</picture>
											</el-icon>
										</template>
									</el-popover>
								</template>
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
										>广告
									</el-button>
								</template>
							</cl-table>
						</el-row>
						<cl-upsert :ref="setRefs('categoryUpsert')" :items="categoryUpsert.items" />
					</cl-crud>
				</div>
			</div>

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
								<template #column-image="{ scope }">
									<el-image
										v-if="scope.row.image"
										lazy
										:preview-src-list="[scope.row.image]"
										:src="scope.row.image"
										fit="contain"
										style="max-height: 50px; max-width: 80px"
									>
										<template #error>
											<div class="image-slot" style="font-size: 45px">
												<icon-svg name="icon-wechat-material" />
											</div>
										</template>
									</el-image>
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
import { useRefs } from "/@/cool";
import { QueryList, Table, Upsert } from "@cool-vue/crud/types";
import { ElMessage, ElIcon } from "element-plus";
import IconSvg from "../components/icon-svg/index.vue";
import { Picture } from "@element-plus/icons-vue";

export default defineComponent({
	name: "sys-ad",

	components: {
		IconSvg,
		ElIcon,
		Picture
	},

	setup() {
		const service = inject<any>("service");

		const { refs, setRefs } = useRefs();

		const selects = reactive<any>({
			dept: {},
			ids: []
		});

		const table = reactive<Table>({
			props: {
				"default-sort": {
					prop: "id",
					order: "descending"
				}
			},
			columns: [
				{
					prop: "listorder",
					label: "排序",
					width: 170
				},
				{
					prop: "name",
					label: "名称",
					width: 170
				},
				{
					prop: "image",
					label: "图片"
					// component: {
					// 	name: "el-image",
					// 	props: {
					// 		fit: "contain",
					// 		height: 40
					// 	}
					// }
				},
				{
					prop: "start_time",
					label: "开始时间"
				},
				{
					prop: "end_time",
					label: "结束时间"
				},
				{
					prop: "status",
					label: "状态",
					minWidth: 50,
					dict: [
						{
							label: "正常",
							value: 1,
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

		const adSpaceTable = reactive<Table>({
			props: {
				"default-sort": {
					prop: "id",
					order: "descending"
				}
			},
			columns: [
				{
					prop: "name",
					label: "广告位名称",
					align: "left"
				},
				{
					prop: "key",
					label: "标识",
					align: "left",
					width: 80
				},
				{
					type: "op",
					buttons: ["slot-btn", "edit", "delete"],
					width: 120
				}
			]
		});

		const categoryUpsert = reactive<Upsert>({
			items: [
				{
					prop: "name",
					label: "广告位名称",
					span: 12,
					component: {
						name: "el-input",
						props: {
							placeholder: "请填写广告位名称"
						}
					},
					rules: {
						required: true,
						message: "广告位名称不能为空"
					}
				},
				{
					prop: "key",
					label: "标识",
					span: 12,
					component: {
						name: "el-input",
						props: {
							placeholder: "请填写标识"
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
					prop: "space_id",
					label: "广告位",
					component: {
						name: "el-select",
						props: {
							placeholder: "请选择分类"
						},
						options: list
					},
					span: 12,
					rules: {
						required: true,
						message: "分类必选"
					}
				},
				{
					prop: "data_range",
					label: "有效期",
					span: 12,
					component: {
						name: "el-date-picker",
						props: {
							type: "datetimerange",
							"range-separator": "至",
							"start-placeholder": "开始时间",
							"end-placeholder": "结束日期"
						}
					},
					rules: {
						required: true,
						message: "值不能为空"
					}
				},
				{
					prop: "image",
					label: "图片",
					span: 24,
					component: {
						name: "cl-upload"
					}
				},
				{
					prop: "linkurl",
					label: "链接地址",
					span: 24,
					component: {
						name: "el-input"
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
								value: 1
							},
							{
								label: "禁用",
								value: 0
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
			ctx.service(service.system.ad).done();
			const cats = await service.system.adSpace.list({ size: 1 });
			if (cats.list.length) {
				catId.value = cats.list[0].id;
				catKey.value = cats.list[0].key;
				catName.value = cats.list[0].name;
				app.refresh({ cid: catId });
			}
		}

		// crud 加载
		function onCategoryLoad({ ctx, app }: any) {
			ctx.service(service.system.adSpace).done();
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
			adSpaceTable,
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
		width: 350px;
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
		width: calc(100% - 610px);
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

	.image-slot {
		display: flex;
		justify-content: center;
		align-items: center;
		width: 100%;
		height: 100%;
		background: #f5f7fa;
		color: #909399;
	}
}
</style>
