<template>
	<div>
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
				<cl-table v-bind="table" />
			</el-row>

			<el-row type="flex">
				<cl-flex1 />
				<cl-pagination />
			</el-row>

			<cl-upsert v-bind="upsert" :on-info="onInfo">
				<template #slot-appid="{ scope }">
					<account-select v-model="scope.appid" />
				</template>

				<template #slot-content="{ scope }">
					<el-input
						v-model="scope.replyContent"
						size="mini"
						type="textarea"
						:rows="5"
						placeholder="文本、图文ID、media_id、json配置"
					/>
					<el-button
						v-if="'text' === scope.replyType"
						type="text"
						size="mini"
						@click="
							() => {
								if (scope.replyContent) {
									scope.replyContent += '<a href=\'链接地址\'>链接文字</a>';
								} else {
									scope.replyContent = '<a href=\'链接地址\'>链接文字</a>';
								}
							}
						"
						>插入链接
					</el-button>
					<el-button
						type="text"
						size="mini"
						v-if="checkShowMaterial(scope.replyType)"
						@click="assetsSelectorVisible = true"
					>
						从素材库中选择<span
							v-if="
								'miniprogrampage' === scope.replyType || 'music' === scope.replyType
							"
							>缩略图</span
						>
					</el-button>
				</template>
			</cl-upsert>
		</cl-crud>

		<cl-dialog title="选择素材" v-model="assetsSelectorVisible" width="50%">
			<material-news v-if="assetsType === 'news'" selectMode @selected="onAssetsSelect" />
			<material-file v-else :fileType="assetsType" selectMode @selected="onAssetsSelect" />
		</cl-dialog>
	</div>
</template>

<script lang="ts">
import { defineComponent, inject, onMounted, reactive, ref } from "vue";
import { useRefs } from "/@/core";
import { CrudLoad, Table, Upsert } from "cl-admin-crud-vue3/types";
import AccountSelect from "../components/account-select.vue";
import MaterialFile from "./assets/material-file.vue";
import MaterialNews from "./assets/material-news.vue";

export default defineComponent({
	name: "wechat-account",
	components: {
		MaterialFile,
		MaterialNews,
		AccountSelect
	},
	setup() {
		const service = inject<any>("service");

		const { refs, setRefs }: any = useRefs();

		const form = ref({ appid: "" });

		const currentRuleForm = ref({});

		const accounts = ref([]);

		const KefuMsgType: any = {
			text: "文本消息",
			image: "图片消息",
			voice: "语音消息",
			video: "视频消息",
			music: "音乐消息",
			news: "文章链接",
			mpnews: "公众号图文消息",
			wxcard: "卡券消息",
			miniprogrampage: "小程序消息",
			msgmenu: "菜单消息"
		};

		let msgTypeOptions = [];

		for (const kefuMsgTypeKey in KefuMsgType) {
			msgTypeOptions.push({
				label: KefuMsgType[kefuMsgTypeKey],
				value: kefuMsgTypeKey
			});
		}

		const upsert = reactive<Upsert>({
			items: [
				{
					prop: "ruleName",
					label: "规则名称",
					component: {
						name: "el-input",
						props: {
							placeholder: "请输入规则名称"
						}
					},
					rules: {
						required: true,
						message: "生效起始时间不能为空"
					}
				},
				{
					prop: "matchValue",
					label: "匹配关键词",
					component: {
						name: "el-input",
						props: {
							placeholder: "请输入规则名称"
						}
					},
					rules: {
						required: true,
						message: "生效起始时间不能为空"
					}
				},
				{
					prop: "appid",
					label: "作用范围",
					component: {
						name: "slot-appid"
					}
				},
				{
					prop: "exactMatch",
					label: "精确匹配",
					value: true,
					component: {
						name: "el-switch",
						props: {
							"active-value": true,
							"inactive-value": false
						}
					}
				},
				{
					prop: "replyType",
					label: "消息类型",
					component: {
						name: "el-select",
						options: msgTypeOptions
					},
					rules: {
						required: true
					}
				},
				{
					prop: "replyContent",
					label: "回复内容",
					component: {
						name: "slot-content"
					}
				},
				{
					prop: "status",
					label: "是否启用",
					value: true,
					component: {
						name: "el-switch",
						props: {
							"active-value": true,
							"inactive-value": false
						}
					}
				},
				{
					prop: "effectTimeStart",
					label: "生效时间",
					span: 12,
					component: {
						name: "el-time-picker",
						props: {
							format: "HH:mm:ss",
							valueFormat: "HH:mm:ss"
						}
					}
				},
				{
					prop: "effectTimeEnd",
					label: "失效时间",
					span: 12,
					component: {
						name: "el-time-picker",
						props: {
							format: "HH:mm:ss",
							valueFormat: "HH:mm:ss"
						}
					}
				}
			]
		});

		const table = reactive<Table>({
			columns: [
				{
					prop: "ruleName",
					label: "规则名称",
					width: 100
				},
				{
					prop: "matchValue",
					label: "匹配关键词",
					width: 100
				},
				{
					prop: "replyType",
					label: "消息类型",
					width: 100,
					formatter: replyTypeFormat
				},
				{
					prop: "replyContent",
					label: "回复内容",
					showOverflowTooltip: true,
					align: "left"
				},
				{
					prop: "appid",
					label: "公众号名称",
					dict: accounts,
					align: "left",
					width: 230
				},
				{
					prop: "exactMatch",
					label: "精确匹配",
					width: 80,
					align: "left",
					dict: [
						{
							label: "是",
							value: true,
							type: "success"
						},
						{
							label: "否",
							value: false,
							type: "warning"
						}
					]
				},
				{
					prop: "status",
					label: "有效",
					width: 70,
					align: "left",
					dict: [
						{
							label: "是",
							value: true,
							type: "success"
						},
						{
							label: "否",
							value: false,
							type: "warning"
						}
					]
				},
				{
					prop: "effectTimeStart",
					label: "生效时间",
					width: 100
				},
				{
					prop: "effectTimeEnd",
					label: "失效时间",
					width: 100
				},
				{
					type: "op",
					buttons: ["edit", "delete"]
				}
			]
		});

		function replyTypeFormat(row: string, column: string, cellValue: string) {
			return KefuMsgType[cellValue];
		}

		function onLoad({ ctx, app }: CrudLoad) {
			ctx.service(service.wechat.rule).done();
			app.refresh();
		}

		onMounted(() => {
			service.wechat.account.select().then((data: any) => {
				data.unshift({ label: "全部公众号", value: "" });
				accounts.value = data;
			});
		});

		const assetsType = ref("");

		const assetsSelectorVisible = ref(false);

		function onAssetsSelect(assetsInfo: any) {
			if (assetsInfo.replyType == "miniprogrampage" || assetsInfo.replyType == "music") {
				let data = JSON.parse(assetsInfo.replyContent);
				if (data && data.thumb_media_id) data.thumb_media_id = assetsInfo.mediaId;
				currentRuleForm.value.replyContent = JSON.stringify(data, null, 4);
			} else {
				currentRuleForm.value.replyContent = assetsInfo.media_id;
			}
			console.log(currentRuleForm.value);
			assetsSelectorVisible.value = false;
		}

		// 检查展示素材 并显示打开类型
		function checkShowMaterial(type: string) {
			const config = {
				image: "image",
				voice: "voice",
				video: "video",
				mpnews: "news",
				miniprogrampage: "image",
				music: "image"
			};
			assetsType.value = config[type] || "";
			return assetsType.value;
		}
		function showSelector() {
			assetsSelectorVisible.value = true;
		}

		function onInfo(data: any, { done }) {
			currentRuleForm.value = data;
			done(data);
		}
		return {
			onInfo,
			service,
			refs,
			form,
			table,
			setRefs,
			assetsSelectorVisible,
			onAssetsSelect,
			showSelector,
			checkShowMaterial,
			assetsType,
			onLoad,
			upsert
		};
	}
});
</script>
