<template>
	<div>
		<cl-crud :ref="setRefs('crud')" @load="onLoad" :on-refresh="onRefresh">
			<el-row type="flex">
				<cl-refresh-btn />
				<cl-add-btn />
				<el-button
					size="mini"
					icon="el-icon-price-tag"
					type="success"
					@click="showTagManager"
					:disabled="appid === ''"
					>标签管理</el-button
				>
				<el-button
					size="mini"
					icon="el-icon-sort"
					type="warning"
					@click="syncFans"
					:disabled="appid === ''"
					>同步粉丝</el-button
				>
				<el-button
					type="primary"
					size="mini"
					@click="userTagging('tagging')"
					:disabled="dataListSelections.length <= 0 || appid === ''"
					>绑定标签</el-button
				>
				<el-button
					type="primary"
					size="mini"
					@click="userTagging('untagging')"
					:disabled="dataListSelections.length <= 0 || appid === ''"
					>解绑标签</el-button
				>
				<cl-flex1 />
				<account-select v-model="appid" />
			</el-row>

			<el-row>
				<cl-table v-bind="table" @selection-change="onSelectionChange">
					<template #column-headimgurl="{ scope }">
						<el-image :src="scope.row.headimgurl" />
					</template>
				</cl-table>
			</el-row>

			<el-row type="flex">
				<cl-flex1 />
				<cl-pagination />
			</el-row>

			<cl-upsert v-bind="upsert" />
		</cl-crud>

		<wx-user-tags-manager
			:ref="setRefs('wxUserTagsEditor')"
			:visible="showWxUserTagsEditor"
			:appid="appid"
			@close="showWxUserTagsEditor = false"
		/>
		<wx-user-tagging
			:ref="setRefs('wxUserTagging')"
			:wxUsers="dataListSelections"
			:appid="appid"
		/>
	</div>
</template>

<script lang="ts">
import { defineComponent, inject, reactive, ref, watch } from "vue";
import { useRefs } from "/@/cool";
import { CrudLoad, Table, Upsert } from "@cool-vue/crud/types";
import { ElMessage } from "element-plus";
import WxUserTagsManager from "../components/wx-user-tags-manager.vue";
import WxUserTagging from "./wx-user-tagging.vue";
import AccountSelect from "../components/account-select.vue";

export default defineComponent({
	name: "wechat-user",
	components: {
		WxUserTagsManager,
		WxUserTagging,
		AccountSelect
	},
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
					span: 12,
					component: {
						name: "el-input",
						props: {
							placeholder: "昵称"
						}
					}
				},
				{
					prop: "phone",
					label: "手机号",
					span: 12,
					component: {
						name: "el-input",
						props: {
							placeholder: "请输入手机号"
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
					label: "功能",
					component: {
						name: "slot-btns"
					}
				},
				{
					prop: "remark",
					label: "备注",
					component: {
						name: "el-input",
						props: {
							type: "textarea",
							rows: 4,
							placeholder: "请输入备注"
						}
					}
				}
			]
		});

		const table = reactive<Table>({
			columns: [
				{
					type: "selection",
					width: 40
				},
				{
					prop: "openid",
					label: "OpenId",
					width: 230
				},
				{
					prop: "nickname",
					label: "昵称",
					align: "left"
				},
				{
					prop: "phone",
					label: "手机号",
					width: 150
				},
				{
					prop: "sex",
					label: "性别",
					width: 80,
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
					prop: "province",
					label: "省份",
					width: 100
				},
				{
					prop: "city",
					label: "城市",
					width: 100
				},
				{
					prop: "headimgurl",
					label: "头像",
					width: 80
				},
				{
					prop: "subscribe_time",
					label: "关注时间",
					width: 140
				},
				{
					prop: "subscribe",
					label: "是否关注",
					width: 90,
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
					prop: "subscribe_scene",
					label: "订阅场景",
					width: 200
				},
				{
					type: "op",
					width: 140,
					buttons: ["edit", "delete"]
				}
			]
		});

		function onLoad({ ctx, app }: CrudLoad) {
			ctx.service(service.wechat.user).done();
			app.refresh();
		}

		const showWxUserTagsEditor = ref(false);

		const appid = ref("");

		function showTagManager() {
			if (!form.value.appid) {
				ElMessage.error("请先选择一个公众号");
				return;
			}
			appid.value = form.value.appid;
			showWxUserTagsEditor.value = true;
			refs.value.wxUserTagsEditor.init(appid);
		}

		const form = ref({ appid: "" });

		function change(formData: any) {
			form.value = formData;
		}

		function syncFans() {
			if (!form.value.appid) {
				ElMessage.error("请先选择一个公众号");
				return;
			}
			service.wechat.user
				.sync({ appid: form.value.appid })
				.then(() => {
					refs.value.crud.refresh();
				})
				.catch((e: any) => {
					ElMessage.error(e);
				});
		}

		const dataListSelections = ref([]);

		function onSelectionChange(selection: any[]) {
			dataListSelections.value = selection;
		}

		function userTagging(action: string) {
			refs.value.wxUserTagging.init(action, appid.value);
		}

		watch(appid, (newValue, oldValue) => {
			//直接监听
			refs.value.crud.refresh();
		});

		// 刷新列表
		function refresh(params: any) {
			refs.value.crud.refresh(params);
		}

		// 刷新监听
		async function onRefresh(params: any, { next, render }: any) {
			params["appid"] = appid.value;
			const { list } = await next(params);
			render(list);
		}

		return {
			form,
			change,
			appid,
			userTagging,
			refresh,
			onRefresh,
			onSelectionChange,
			dataListSelections,
			showWxUserTagsEditor,
			showTagManager,
			syncFans,
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
