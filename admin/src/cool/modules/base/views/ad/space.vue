<template>
	<div class="system-user">
		<div class="pane">
			<div class="user">
				<div class="container">
					<cl-crud :ref="setRefs('space')" :on-refresh="onRefresh" @load="onLoad">
						<el-row type="flex">
							<cl-refresh-btn />
							<cl-add-btn />
							<cl-multi-delete-btn />
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
									>{{ item }}</el-tag
									>
								</template>

								<!-- 单个转移 -->
								<template #slot-move-btn="{ scope }">
									<el-button
										v-permission="service.system.user.permission.move"
										type="text"
										size="mini"
										@click="toMove(scope.row)"
									>转移</el-button
									>
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
							:on-submit="onUpsertSubmit"
						>
							<template #slot-tips>
								<div>
									<i class="el-icon-warning"></i>
									<span style="margin-left: 6px">新增用户默认密码为：123456</span>
								</div>
							</template>
						</cl-upsert>
					</cl-crud>
				</div>
			</div>
		</div>
	</div>
</template>

<script lang="ts">
import { computed, defineComponent, inject, reactive, watch } from "vue";
import { useStore } from "vuex";
import { useRefs } from "/@/core";
import { Table, Upsert } from "cl-admin-crud-vue3/types";

export default defineComponent({
	name: "space",

	setup() {
		const service = inject<any>("service");
		const store = useStore();
		const { refs, setRefs } = useRefs();
		// 选择项
		const selects = reactive<any>({ ids: [] });
		// 表格配置
		const table = reactive<Table>({
			props: {
				"default-sort": {
					prop: "createTime",
					order: "descending"
				}
			},
			columns: [
				{
					prop: "name",
					label: "广告名称",
					width: 100
				},
				{
					prop: "image",
					label: "Logo",
					// width: 95,
					// component: {
					// 	name: "cl-upload",
					// 	props: {
					// 		size: [70, 70],
					// 		text: "选择图片",
					// 		icon: "el-icon-picture"
					// 	}
					// }
				},
				{
					prop: "space_name",
					label: "广告位",
					span: 24
				},
				{
					prop: "link_url",
					label: "链接地址",
				},
				{
					type: "op",
					buttons: ["edit", "delete"],
					span: 14,
				}
			]
		});

		// 新增、编辑配置
		const upsert = reactive<Upsert>({
			items: [
				{
					prop: "headImg",
					label: "头像",
					span: 24,
					component: {
						name: "cl-upload",
						props: {
							text: "选择头像",
							icon: "el-icon-picture"
						}
					}
				},
				{
					prop: "name",
					label: "姓名",
					span: 24,
					component: {
						name: "el-input",
						props: {
							placeholder: "请填写姓名"
						}
					},
					rules: {
						required: true,
						message: "姓名不能为空"
					}
				},
				{
					prop: "nickName",
					label: "昵称",
					span: 12,
					component: {
						name: "el-input",
						props: {
							placeholder: "请填写昵称"
						}
					},
					rules: {
						required: true,
						message: "昵称不能为空"
					}
				},
				{
					prop: "username",
					label: "用户名",
					span: 12,
					component: {
						name: "el-input",
						props: {
							placeholder: "请填写用户名"
						}
					},
					rules: [
						{
							required: true,
							message: "用户名不能为空"
						}
					]
				},
				{
					prop: "password",
					label: "密码",
					span: 12,
					hidden: ":isAdd",
					component: {
						name: "el-input",
						props: {
							placeholder: "请填写密码",
							type: "password"
						}
					},
					rules: [
						{
							min: 6,
							max: 16,
							message: "密码长度在 6 到 16 个字符"
						}
					]
				},
				{
					prop: "roleIdList",
					label: "角色",
					span: 24,
					value: [],
					component: {
						name: "cl-role-select",
						props: {
							props: {
								"multiple-limit": 3
							}
						}
					},
					rules: {
						required: true,
						message: "角色不能为空"
					}
				},
				{
					prop: "phone",
					label: "手机号码",
					span: 12,
					component: {
						name: "el-input",
						props: {
							placeholder: "请填写手机号码"
						}
					}
				},
				{
					prop: "email",
					label: "邮箱",
					span: 12,
					component: {
						name: "el-input",
						props: {
							placeholder: "请填写邮箱"
						}
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
								label: "开启",
								value: 1
							},
							{
								label: "关闭",
								value: 0
							}
						]
					}
				},
				{
					prop: "tips",
					hidden: ":isEdit",
					component: {
						name: "slot-tips"
					}
				}
			]
		});

		// 浏览器信息
		const browser = computed(() => store.getters.browser);

		// 监听屏幕大小变化
		watch(
			() => browser.value.isMini,
			(val: boolean) => {
			},
			{
				immediate: true
			}
		);

		// crud 加载
		function onLoad({ ctx, app }: any) {
			ctx.service(service.system.ad).done();
			app.refresh();
		}

		// 刷新列表
		function refresh(params: any) {
			refs.value.crud.refresh(params);
		}

		// 提交钩子
		function onUpsertSubmit(_: boolean, data: any, { next }: any) {
			let departmentId = data.departmentId;

			if (!departmentId) {
				departmentId = selects.dept.id;

				if (!departmentId) {
					departmentId = dept.value[0].id;
				}
			}

			next({
				...data,
				departmentId
			});
		}

		// 刷新监听
		async function onRefresh(params: any, { next, render }: any) {
			const { list } = await next(params);
			console.log("list", list);
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

		return {
			service,
			refs,
			selects,
			table,
			upsert,
			browser,
			setRefs,
			refresh,
			onRefresh,
			onLoad,
			onUpsertSubmit,
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

.user {
	width: calc(100% - 310px);
	flex: 1;
}

.user {
	overflow: hidden;

	.container {
		height: calc(100% - 40px);
	}
}

}
</style>
