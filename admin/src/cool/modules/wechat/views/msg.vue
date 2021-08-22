<template>
	<div>
		<cl-crud :ref="setRefs('crud')" @load="onLoad">
			<el-row type="flex">
				<cl-refresh-btn />
				<cl-filter label="时间">
					<el-select placeholder="时间" size="mini">
						<el-option
							v-for="(name, key) in timeSelections"
							:key="key"
							:value="name"
							:label="key"
						/>
					</el-select>
				</cl-filter>
				<cl-filter label="消息类型">
					<el-select placeholder="消息类型" size="mini">
						<el-option value="" label="不限类型" />
						<el-option
							value="text,image,voice,shortvideo,video,news,music,location,link"
							label="消息"
						/>
						<el-option value="event,transfer_customer_service" label="事件" />
					</el-select>
				</cl-filter>
				<cl-flex1 />
				<cl-search-key />
			</el-row>

			<el-row>
				<cl-table v-bind="table">
					<template #column-nickname="{ scope }">
						{{ scope.row.fans_info.nickname }}
					</template>
					<template #column-headimgurl="{ scope }">
						<el-avatar :src="scope.row.fans_info.headimgurl" shape="square" />
					</template>
					<template #column-detail="{ scope }">
						{{ scope.row.detail.Content }}
					</template>
					<template #column-event="{ scope }">
						<el-tooltip
							v-if="scope.row.detail.Event"
							effect="dark"
							:content="scope.row.detail.EventKey"
							placement="top-start"
						>
							<el-tag size="mini" type="success">{{ scope.row.detail.Event }}</el-tag>
						</el-tooltip>
					</template>

					<template #slot-reply="{ scope }">
						<el-button
							v-if="canReply(scope.row.created_at)"
							size="mini"
							type="success"
							icon="el-icon-s-promotion"
							@click="setReply(scope.row.openid, scope.row.appid)"
							>回复
						</el-button>
					</template>
				</cl-table>
			</el-row>

			<el-row type="flex">
				<cl-flex1 />
				<cl-pagination />
			</el-row>
		</cl-crud>
		<cl-dialog title="消息回复" :close-on-click-modal="false" v-model="visible">
			<el-form v-model="dataForm" :ref="setRefs('dataForm')" :rules="dataRule">
				<el-form-item>
					<el-input
						v-model="dataForm.replyContent"
						type="textarea"
						:rows="5"
						placeholder="回复内容"
						maxlength="600"
						show-word-limit
						:autosize="{ minRows: 5, maxRows: 30 }"
						autocomplete
					/>
					<el-button type="text" v-show="'text' === dataForm.replyType" @click="addLink"
						>插入链接
					</el-button>
				</el-form-item>
			</el-form>
			<span slot="footer" class="dialog-footer">
				<el-button @click="visible = false" size="mini">取消</el-button>
				<el-button type="success" @click="dataFormSubmit()" size="mini">{{
					uploading ? "发送中..." : "发送"
				}}</el-button>
			</span>
		</cl-dialog>
	</div>
</template>

<script lang="ts">
import { defineComponent, inject, reactive, ref } from "vue";
import { useRefs } from "/@/core";
import { CrudLoad, Table } from "cl-admin-crud-vue3/types";
import moment from "moment";
import { ElMessage } from "element-plus";

export default defineComponent({
	name: "wechat-msg",

	setup() {
		const TIME_FORMAT = "YYYY/MM/DD hh:mm:ss";
		const service = inject<any>("service");
		const { refs, setRefs }: any = useRefs();

		const timeSelections = ref({
			近24小时: moment().subtract(1, "days").format(TIME_FORMAT),
			近3天: moment().subtract(3, "days").format(TIME_FORMAT),
			近7天: moment().subtract(7, "days").format(TIME_FORMAT),
			近30天: moment().subtract(30, "days").format(TIME_FORMAT)
		});

		const table = reactive<Table>({
			columns: [
				{
					prop: "nickname",
					label: "昵称",
					align: "left",
					width: 150
				},
				{
					prop: "headimgurl",
					label: "头像",
					width: 80
				},
				{
					prop: "created_at",
					label: "消息时间",
					width: 160
				},
				{
					prop: "in_out",
					label: "动作",
					width: 140,
					dict: [
						{
							type: "success",
							label: "来自用户的消息",
							value: 0
						},
						{
							type: "primary",
							label: "公众号发出的消息",
							value: 1
						}
					]
				},
				{
					prop: "msg_type",
					label: "消息类型",
					width: 100
				},
				{
					prop: "event",
					label: "事件",
					width: 100
				},
				{
					prop: "detail",
					label: "消息内容",
					align: "left"
				},
				{
					type: "op",
					width: 140,
					buttons: ["slot-reply"]
				}
			]
		});

		const visible = ref(false);

		function onLoad({ ctx, app }: CrudLoad) {
			ctx.service(service.wechat.msg).done();
			app.refresh();
		}

		const uploading = ref(false);

		function canReply(d: any) {
			return new Date(d).getTime() > new Date().getTime() - 24 * 60 * 60 * 1000;
		}

		const dataForm = ref({
			appid: "",
			openid: "",
			replyType: "text",
			replyContent: ""
		});

		const dataRule = {
			replyContent: [{ required: true, message: "回复内容不能为空", trigger: "blur" }]
		};

		function addLink() {
			dataForm.value.replyContent += '<a href="链接地址">链接文字</a>';
		}

		function setReply(openid: string, appid: string) {
			dataForm.value.appid = appid;
			dataForm.value.openid = openid;
			visible.value = true;
		}

		function dataFormSubmit() {
			uploading.value = true;
			service.wechat.msg
				.reply(dataForm.value)
				.then(({ data }: any) => {
					ElMessage.success("回复成功");
					uploading.value = false;
					visible.value = false;
					data.dataForm.value.openid = "";
					data.dataForm.value.replyContent = "";
				})
				.catch((e) => {
					ElMessage.error(e);
				});
		}

		return {
			canReply,
			uploading,
			dataRule,
			dataFormSubmit,
			setReply,
			addLink,
			dataForm,
			timeSelections,
			service,
			visible,
			refs,
			table,
			setRefs,
			onLoad
		};
	}
});
</script>
