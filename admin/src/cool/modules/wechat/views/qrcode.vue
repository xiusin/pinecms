<template>
	<cl-crud :ref="setRefs('crud')" @load="onLoad">
		<el-row type="flex">
			<cl-add-btn />
			<cl-refresh-btn />
			<cl-flex1 />
			<cl-filter-group v-model="form">
				<account-select v-model="form.appid" />
			</cl-filter-group>
		</el-row>

		<el-row>
			<cl-table v-bind="table">
				<template #column-ticket="{ scope }">
					<el-button
						@click="
							showQrcode(
								'https://mp.weixin.qq.com/cgi-bin/showqrcode?ticket=' +
									scope.row.ticket
							)
						"
						size="mini"
						>点击查看</el-button
					>
				</template>
			</cl-table>
		</el-row>

		<el-row type="flex">
			<cl-flex1 />
			<cl-pagination />
		</el-row>

		<cl-dialog v-model="visible" title="二维码" width="30%">
			<div style="text-align: center">
				<el-image :src="qrcodeUrl" />
			</div>
		</cl-dialog>

		<cl-upsert v-model="form" v-bind="upsert">
			<template #slot-appid="{ scope }">
				<account-select v-model="scope.appid" />
			</template>
		</cl-upsert>
	</cl-crud>
</template>

<script lang="ts">
import { defineComponent, inject, onMounted, reactive, ref } from "vue";
import { useRefs } from "/@/core";
import { CrudLoad, Table, Upsert } from "cl-admin-crud-vue3/types";
import AccountSelect from "../components/account-select.vue";

export default defineComponent({
	name: "wechat-qrcode",

	components: {
		AccountSelect
	},

	setup() {
		const service = inject<any>("service");

		const { refs, setRefs }: any = useRefs();

		const form = ref({ appid: "" });

		const accounts = ref([]);

		const upsert = reactive<Upsert>({
			items: [
				{
					prop: "appid",
					label: "选择公众号",
					component: {
						name: "slot-appid"
					},
					rules: {
						required: true
					}
				},
				{
					prop: "is_temp",
					label: "二维码类型",
					value: true,
					component: {
						name: "el-radio-group",
						options: [
							{
								label: "临时",
								value: true
							},
							{
								label: "永久",
								value: false
							}
						]
					}
				},
				{
					prop: "scene_str",
					label: "场景值",
					component: {
						name: "el-input",
						props: {
							placeholder: "任意字符串"
						}
					},
					rules: {
						required: true
					}
				},
				{
					prop: "expire_time",
					label: "过期时间",
					hidden: ({ scope }: any) => !scope.is_temp,
					component: {
						name: "el-date-picker",
						props: {
							type: "datetime",
							placeholder: "最多30天",
							valueFormat: "YYYY-MM-DD HH:mm:ss",
							format: "YYYY-MM-DD HH:mm:ss"
						}
					},
					rules: {
						required: true
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
					prop: "appid",
					label: "所属公众号",
					dict: accounts
				},
				{
					prop: "is_temp",
					label: "类型",
					width: 80,
					dict: [
						{
							label: "临时",
							value: true,
							type: "success"
						},
						{
							label: "永久",
							value: false,
							type: "warning"
						}
					]
				},
				{
					prop: "scene_str",
					label: "场景值",
					align: "left",
					showOverflowTooltip: true,
					minWidth: 130
				},
				{
					prop: "ticket",
					label: "二维码图片",
					width: 130
				},
				{
					prop: "url",
					label: "解析后的地址",
					width: 300,
					showOverflowTooltip: true
				},
				{
					prop: "expire_time",
					label: "失效时间",
					width: 160
				},
				{
					type: "op",
					buttons: ["delete"]
				}
			]
		});

		function onLoad({ ctx, app }: CrudLoad) {
			ctx.service(service.wechat.qrcode).done();
			app.refresh();
		}

		const visible = ref(false);

		const qrcodeUrl = ref("");

		function showQrcode(url: string) {
			visible.value = true;
			qrcodeUrl.value = url;
		}

		onMounted(() => {
			service.wechat.account.select().then((data: any) => {
				data.unshift({ label: "全部公众号", value: "" });
				accounts.value = data;
			});
		});
		return {
			accounts,
			qrcodeUrl,
			showQrcode,
			visible,
			service,
			refs,
			table,
			setRefs,
			form,
			onLoad,
			upsert
		};
	}
});
</script>
