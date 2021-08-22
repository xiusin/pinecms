<template>
	<cl-crud :ref="setRefs('crud')" @load="onLoad">
		<el-row type="flex">
			<cl-refresh-btn />
			<cl-add-btn />
			<el-button size="mini" icon="el-icon-price-tag" type="success">标签管理</el-button>
			<el-button size="mini" icon="el-icon-sort" type="warning">同步粉丝</el-button>
			<cl-flex1 />
			<cl-search-key />
		</el-row>

		<el-row>
			<cl-table v-bind="table" />
		</el-row>

		<el-row type="flex">
			<cl-flex1 />
			<cl-pagination />
		</el-row>

		<cl-upsert v-model="form" v-bind="upsert">
			<template #slot-btns="{ scope }">
				<el-button>查询</el-button>
				<el-button type="primary" size="mini">绑定标签</el-button>
				<el-button type="primary" size="mini">解绑标签</el-button>
				<el-button type="danger" size="mini">批量删除</el-button>
			</template>
		</cl-upsert>
	</cl-crud>
</template>

<script lang="ts">
import { defineComponent, inject, reactive } from "vue";
import { useRefs } from "/@/core";
import { CrudLoad, Table, Upsert } from "cl-admin-crud-vue3/types";

export default defineComponent({
	name: "wechat-user",

	setup() {
		const service = inject<any>("service");

		const { refs, setRefs }: any = useRefs();

		const upsert = reactive<Upsert>({
			items: [
				{
					prop: "tagid",
					label: "用户标签",
					span: 24,
					component: {
						name: "el-select",
						props: {
							placeholder: "请选择用户标签"
						}
					}
				},
				{
					prop: "nickname",
					label: "昵称",
					span: 24,
					component: {
						name: "el-input",
						props: {
							placeholder: "昵称"
						}
					}
				},
				{
					prop: "city",
					label: "城市",
					span: 12,
					component: {
						name: "el-input",
						props: {
							placeholder: "请输入城市"
						}
					}
				},
				{
					prop: "province",
					label: "省份",
					span: 12,
					component: {
						name: "el-input",
						props: {
							placeholder: "请输入省份"
						}
					}
				},
				{
					prop: "qrSceneStr",
					label: "关注场景值",
					component: {
						name: "el-input",
						props: {
							placeholder: "请输入关注场景值"
						}
					}
				},
				{
					prop: "qrSceneStr",
					label: "关注场景值",
					component: {
						name: "slot-btns"
					}
				}
			]
		});

		const table = reactive<Table>({
			columns: [
				{
					type: "index",
					label: "#",
					width: 60
				},
				{
					prop: "openid",
					label: "OpenId",
					width: 300
				},
				{
					prop: "nickname",
					label: "昵称"
				},
				{
					prop: "sex",
					label: "性别",
					width: 140,
					align: "left",
					dict: [
						{
							label: "未知",
							value: 0,
							type: "primary"
						},
						{
							label: "男",
							value: 1,
							type: "success"
						},
						{
							label: "女",
							value: 2,
							type: "warning"
						}
					]
				},
				{
					prop: "city",
					label: "城市",
					width: 140
				},
				{
					prop: "province",
					label: "省份",
					width: 140
				},
				{
					prop: "headimgurl",
					label: "头像",
					width: 140
				},
				{
					prop: "subscribe_time",
					label: "关注时间",
					width: 140
				},
				{
					prop: "subscribe",
					label: "是否关注",
					width: 140,
					dict: [
						{
							label: "否",
							value: false,
							type: "danger"
						},
						{
							label: "是",
							value: true,
							type: "success"
						}
					]
				},
				{
					type: "op",
					buttons: ["delete"]
				}
			]
		});

		function onLoad({ ctx, app }: CrudLoad) {
			ctx.service(service.wechat.user).done();
			app.refresh();
		}

		return {
			service,
			refs,
			table,
			setRefs,
			onLoad,
			upsert
		};
	}
});
</script>
