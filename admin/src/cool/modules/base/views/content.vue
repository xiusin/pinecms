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
					<cl-crud :ref="setRefs('crud')" @load="onLoad" v-if="!catType">
						<el-row type="flex">

								<cl-filter label="状态">
									<el-select size="mini" >
										<el-option value="" label="全部"></el-option>
										<el-option :value="0" label="禁用"></el-option>
										<el-option :value="1" label="启用"></el-option>
									</el-select>
								</cl-filter>

								<cl-filter label="姓名">
									<el-input placeholder="请输入姓名" clearable size="mini"></el-input>
								</cl-filter>

							<cl-filter label="日期">
								<el-date-picker
									size="mini"
									type="datetimerange"
									range-separator="至"
									start-placeholder="开始日期"
									end-placeholder="结束日期">
								</el-date-picker>
							</cl-filter>

							<cl-filter label="日期">

							</cl-filter>
								<cl-flex1 />
								<cl-search-key />
								<cl-refresh-btn />
								<cl-add-btn />
						</el-row>

						<el-row>
							<cl-table :ref="setRefs('table')" v-bind="table" :props="{
								height: '700px',
								fit: true,
								'highlight-current-row': true,
								stripe: true,
								'max-height': 900,
								}" :autoHeight="false" />
						</el-row>

						<el-row type="flex">
							<cl-flex1 />
							<cl-pagination />
						</el-row>

						<cl-upsert :ref="setRefs('upsert')" :items="upsert.items" />
					</cl-crud>
					<iframe v-if="catType === 2" src="http://www.baidu.com">
					</iframe>
					<template v-else>
						<cl-form inner>
							<el-form-item label="用户名">
								<el-input
									placeholder="请输入用户名"
									maxlength="20"
									auto-complete="off"
								/>
							</el-form-item>

							<el-form-item label="密码">
								<el-input
									type="password"
									placeholder="请输入密码"
									maxlength="20"
									auto-complete="off"
								/>
							</el-form-item>
						</cl-form>
					</template>
				</div>
			</div>
		</div>
	</div>
</template>

<script lang="ts">
import { computed, defineComponent, inject, onBeforeMount, reactive, ref, watch } from "vue";
import { useRefs } from "/@/core";
import { deepTree } from "/@/core/utils";
import {FormRef, QueryList, Table, Upsert} from "cl-admin-crud-vue3/types";

export default defineComponent({
	name: "sys-content",
	setup() {
		const service = inject<any>("service");
		const { refs, setRefs } = useRefs();
		const data = ref([]);
		const modelValue = ref("");

		// 树形列表
		const menuList = ref<any[]>([]);
		// 展开值
		const expandedKeys = ref<any[]>([]);
		// el-tree 组件
		const treeRef = ref<any>({});
		// 绑定值回调
		function onCurrentChange({ id, catname, type, model_id }: any) {
			catId.value = id;
			catName.value = catname;
			catType.value = type
			if (catType.value == 0) {
				midRef.value = model_id;
				refresh({ cid: catId.value });
			}
		}
		// 表格配置
		const table = ref<Table>();
		// 树形列表
		const treeList = computed(() => deepTree(menuList.value));

		const catId = ref<any>(0);
		const catType = ref<any>(0);
		const midRef = ref<any>(0);
		const catName = ref<string>("");
		const catKey = ref<string>("");

		onBeforeMount(async function () {
			const ret = await service.system.category.list();
			menuList.value = ret.list.filter((e: any) => e.type != 2);
			catId.value = menuList.value[0].id

			catName.value = menuList.value[0].catname
			catType.value = menuList.value[0].type
			if (catType.value == 0) {
				midRef.value = menuList.value[0].model_id;
				refresh({ cid: catId.value });
			}
		});

		const list = ref<QueryList[]>([]);

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
			if (!catType.value) {
				refs.value.crud.refresh(params);
			}
		}

		// crud 加载
		async function onLoad({ ctx }: any) {
			ctx.service(service.system.content).done();
		}

		watch(midRef, (newValue, oldValue)=> {
			service.system.model.modelTable({ mid: newValue }).then((data: any) => {
				console.log(data)
				data.columns.map((item: any) => {
					if (item.component) {
						item.component = typeof item.component == "string" ? Function("return " + item.component)() : item.component;
					}
					return item;
				})
				table.value = data
				table.value?.columns.push({
					label: "操作",
					type: "op",
					buttons: ["edit", "delete"]
				});
			})

		})

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
			treeList,
			refresh,
			treeRef,
			catType,
			onCurrentChange
		};
	},
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
		width: 180px;
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
		width: calc(100% - 190px);
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
