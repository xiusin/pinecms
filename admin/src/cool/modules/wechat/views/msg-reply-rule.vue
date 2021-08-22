<template>
	<cl-crud :ref="setRefs('crud')" @load="onLoad">
		<el-row type="flex">
			<cl-add-btn />
			<cl-refresh-btn />
			<cl-flex1 />
			<cl-search-key />
		</el-row>

		<el-row>
			<cl-table v-bind="table">
				<template #column-appid="{ scope }">
					{{ scope.row.appid ? "当前公众号" : "全部公众号" }}
				</template>
			</cl-table>
		</el-row>

		<el-row type="flex">
			<cl-flex1 />
			<cl-pagination />
		</el-row>

		<cl-upsert v-model="form" v-bind="upsert">
			<template #slot-content="{ scope }">
				<el-input
					v-model="scope.replyContent"
					size="mini"
					type="textarea"
					:rows="5"
					placeholder="文本、图文ID、media_id、json配置"
				/>
				<el-button
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
					>插入链接</el-button
				>
				<el-button type="text" size="mini"> 从素材库中选择<span>缩略图</span> </el-button>
			</template>
		</cl-upsert>
	</cl-crud>

	<el-dialog
		title="选择素材"
		v-model="accessModalRef"
		:modal="true"
		append-to-body
		@close="onClose"
	>
		<material-news @selected="onSelect" selectMode />
		<!--		<material-file :fileType="selectType" @selected="onSelect" selectMode></material-file>-->
	</el-dialog>
</template>

<script lang="ts">
import { defineComponent, inject, reactive, ref } from "vue";
import { useRefs } from "/@/core";
import { CrudLoad, Table, Upsert } from "cl-admin-crud-vue3/types";

export default defineComponent({
	name: "wechat-account",
	components: {
		MaterialFile: () => import("./assets/material-file.vue"),
		MaterialNews: () => import("./assets/material-news.vue")
	},
	setup() {
		const service = inject<any>("service");

		const { refs, setRefs }: any = useRefs();

		const accessModalRef = ref(true);

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
						name: "el-select",
						options: [
							{
								label: "全部公众号",
								value: ""
							},
							{
								label: "当前公众号",
								value: "1"
							}
						]
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
					type: "index",
					label: "#",
					width: 60
				},
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
					width: 100
				},
				{
					prop: "appid",
					label: "作用范围",
					width: 100
				},
				{
					prop: "name",
					label: "公众号名称"
				},
				{
					prop: "exactMatch",
					label: "精确匹配",
					width: 140,
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
					label: "是否有效",
					width: 140,
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
					label: "生效时间"
				},
				{
					prop: "effectTimeEnd",
					label: "失效时间"
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

		return {
			service,
			accessModalRef,
			refs,
			table,
			setRefs,
			onLoad,
			upsert
		};
	}
});
</script>
