<!-- 会员列表 -->
<template>
	<div class="system-user">
		<div class="pane">
			<div class="user">
				<div class="container">
					<cl-crud :ref="setRefs('crud')" :on-refresh="onRefresh" @load="onLoad">
						<el-row type="flex">
							<cl-refresh-btn />
							<cl-add-btn />
							<cl-flex1 />
							<cl-search-key />
						</el-row>

						<el-row>
							<cl-table :ref="setRefs('table')" v-bind="table">
								<!-- 头像 -->
								<template #column-avatar="{ scope }">
									<cl-avatar
										shape="square"
										size="30px"
										:src="scope.row.avatar"
										:style="{ margin: 'auto' }"
									/>
								</template>

								<template #column-group_id="{ scope }">
									{{ getGroupName(scope.row.group_id) }}
								</template>
							</cl-table>
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
</template>

<script lang="ts">
import { computed, defineComponent, inject, reactive, ref, watch } from "vue";
import { useRefs } from "/@/core";
import { Table, Upsert } from "cl-admin-crud-vue3/types";

export default defineComponent({
	name: "sys-member",

	setup() {
		const service = inject<any>("service");
		const { refs, setRefs } = useRefs();

		// 表格配置
		const table = reactive<Table>({
			columns: [
				{
					prop: "avatar",
					label: "头像"
				},
				{
					prop: "account",
					label: "账号",
					width: 150
				},
				{
					prop: "nickname",
					label: "昵称",
					width: 150
				},
				{
					prop: "email",
					label: "邮箱",
					width: 150
				},
				{
					prop: "telphone",
					label: "电话",
					width: 100
				},
				{
					prop: "sex",
					label: "性别",
					width: 80,
					dict: [
						{
							label: "女",
							value: 2,
							type: "success"
						},
						{
							label: "男",
							value: 1,
							type: "warning"
						},
						{
							label: "保密",
							value: 0,
							type: "danger"
						}
					]
				},
				{
					prop: "description",
					label: "个人描述"
				},
				{
					prop: "group_id",
					label: "会员分组",
					width: 100
				},
				{
					prop: "status",
					label: "状态",
					width: 80,
					dict: [
						{
							label: "正常",
							value: 2,
							type: "success"
						},
						{
							label: "待验证",
							value: 1,
							type: "warning"
						},
						{
							label: "禁用",
							value: 0,
							type: "danger"
						}
					]
				},
				{
					prop: "login_time",
					label: "最后登录时间",
					width: 150
				},
				{
					prop: "login_ip",
					label: "最后登录IP",
					width: 100
				},
				{
					prop: "created",
					label: "注册时间",
					width: 150
				},
				{
					type: "op",
					buttons: ["edit", "delete"],
					width: 160
				}
			]
		});

		let groups = {};

		let kvGroups = [{ label: "请选择分组", value: "" }];

		// 新增、编辑配置
		const upsert = reactive<Upsert>({
			items: [
				{
					prop: "avatar",
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
					prop: "email",
					label: "邮箱",
					span: 12,
					component: {
						name: "el-input",
						props: {
							placeholder: "邮箱"
						}
					},
					rules: {
						required: true,
						pattern: /^(\w-*\.*)+@(\w-?)+(\.\w{2,})+$/,
						message: "邮箱格式错误"
					}
				},
				{
					prop: "account",
					label: "用户名",
					span: 12,
					component: {
						name: "el-input",
						props: {
							placeholder: "用户名"
						}
					},
					rules: {
						required: true,
						pattern: /^[A-Za-z][-_!@#$%^&*a-zA-Z0-9]{4,15}$/,
						message: "用户名不能为空"
					}
				},
				{
					prop: "nickname",
					label: "昵称",
					span: 12,
					component: {
						name: "el-input",
						props: {
							placeholder: "请填写昵称"
						}
					}
				},
				{
					prop: "password",
					label: "密码",
					span: 12,
					component: {
						name: "el-input",
						props: {
							placeholder: "请填写密码",
							type: "password"
						}
					}
				},

				{
					prop: "qq",
					label: "QQ",
					span: 12,
					component: {
						name: "el-input"
					}
				},
				{
					prop: "telphone",
					label: "手机号",
					span: 12,
					component: {
						name: "el-input"
					},
					rules: {
						required: true,
						pattern: /^1\d{10}$/,
						message: "手机号格式错误或不能为空"
					}
				},
				{
					prop: "description",
					label: "介绍",
					span: 24,
					component: {
						name: "el-input",
						props: {
							placeholder: "请填写介绍",
							type: "textarea",
							rows: 4
						}
					}
				},
				{
					prop: "group_id",
					label: "分组",
					span: 24,
					component: {
						name: "el-select",
						options: kvGroups
					}
				},
				{
					prop: "sex",
					label: "性别",
					value: 0,
					span: 24,
					component: {
						name: "el-radio-group",
						options: [
							{
								label: "女",
								value: 2
							},
							{
								label: "男",
								value: 1
							},
							{
								label: "保密",
								value: 0
							}
						]
					}
				},
				{
					prop: "status",
					label: "状态",
					value: 2,
					component: {
						name: "el-radio-group",
						options: [
							{
								label: "正常",
								value: 2
							},
							{
								label: "待验证",
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

		service.system.memberGroup.select({}).then((data) => {
			for (const key in data) {
				groups[data[key].id] = data[key].name;
				kvGroups.push({
					label: data[key].name,
					value: parseInt(data[key].id)
				});
			}
		});

		function getGroupName(groupId) {
			console.log(groups);
			return groups[groupId];
		}

		// crud 加载
		function onLoad({ ctx, app }: any) {
			ctx.service(service.system.member).done();
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

		return {
			service,
			refs,
			table,
			upsert,
			setRefs,
			getGroupName,
			onLoad,
			refresh,
			onRefresh
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
		width: 300px;
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

	.user {
		width: calc(100% - 310px);
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
