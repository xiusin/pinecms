<template>
	<cl-crud @load="onLoad">
		<el-row type="flex">
			<cl-refresh-btn />
			<cl-add-btn />
			<cl-multi-delete-btn />
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

		<cl-upsert v-model="form" v-bind="upsert" />
	</cl-crud>
</template>

<script lang="ts">
import { CrudLoad, Table, Upsert } from "cl-admin-crud-vue3/types";
import { defineComponent, inject, reactive } from "vue";

export default defineComponent({
	name: "sys-todo",

	setup() {
		const service = inject<any>("service");

		const form = reactive<any>({});

		const upsert = reactive<Upsert>({
			items: [
				{
					"component": {
						"name": "el-input",
						"props": {
							"options": [
								{
									"key": "1",
									"label": "外部链接"
								},
								{
									"key": "2",
									"label": "内部链接"
								},
								{
									"key": "3",
									"label": "通用链接"
								}
							],
							"size": "mini"
						}
					},
					"label": "字符串多选",
					"prop": "type",
					"value": "'1'"
				},
				{
					"component": {
						"name": "el-input",
						"props": {
							"size": "mini"
						}
					},
					"label": "普通输入框",
					"prop": "name"
				},
				{
					"component": {
						"name": "cms-textarea",
						"props": {
							"size": "mini"
						}
					},
					"label": "普通多行输入框",
					"prop": "introduce",
					"value": "''"
				},
				{
					"component": {
						"name": "el-input-number",
						"props": {
							"controls-position": "right",
							"size": "mini",
							"step": 1,
							"step-strictly": true
						}
					},
					"label": "不可为空数字",
					"prop": "listorder"
				},
				{
					"component": {
						"name": "cms-radio",
						"props": {
							"controls-position": "right",
							"options": [
								{
									"key": 0,
									"label": "待审核"
								},
								{
									"key": 1,
									"label": "通过"
								},
								{
									"key": 2,
									"label": "拒绝"
								}
							],
							"size": "mini",
							"step": 1,
							"step-strictly": true
						}
					},
					"label": "tinyint单选",
					"prop": "status",
					"value": 0
				},
				{
					"component": {
						"name": "el-date-picker",
						"props": {
							"size": "mini",
							"type": "daterange"
						}
					},
					"label": "日期",
					"prop": "put_date"
				},
				{
					"component": {
						"name": "el-date-picker",
						"props": {
							"size": "mini",
							"type": "daterange"
						}
					},
					"label": "时间日期",
					"prop": "put_datetime"
				},
				{
					"component": {
						"name": "el-date-picker",
						"props": {
							"size": "mini",
							"type": "daterange"
						}
					},
					"label": "开始时间$end=end_time",
					"prop": "start_time"
				},
				{
					"component": {
						"name": "el-date-picker",
						"props": {
							"size": "mini",
							"type": "daterange"
						}
					},
					"label": "结束时间被引用隐藏到代码区间选择器",
					"prop": "end_time"
				},
				{
					"component": {
						"name": "cl-upload-space",
						"props": {
							"accept": ".jpg,.png,.jpeg,.bmp,.gif",
							"drag": true,
							"icon": "el-icon-picture",
							"listType": "picture-card",
							"size": [
								45,
								45
							],
							"text": "请选择图片"
						}
					},
					"label": "单图上传",
					"prop": "logo"
				},
				{
					"component": {
						"name": "cl-upload-space",
						"props": {
							"accept": ".jpg,.png,.jpeg,.bmp,.gif",
							"drag": true,
							"icon": "el-icon-picture",
							"listType": "picture-card",
							"multiple": true,
							"size": [
								45,
								45
							],
							"text": "请选择图片"
						}
					},
					"label": "多图上传",
					"prop": "logos"
				}
			]
		});

		const table = reactive<Table>({
			columns: [
				{
					"label": "id",
					"prop": "id"
				},
				{
					"label": "字符串多选",
					"prop": "type"
				},
				{
					"label": "普通输入框",
					"prop": "name"
				},
				{
					"label": "普通多行输入框",
					"prop": "introduce"
				},
				{
					"label": "不可为空数字",
					"prop": "listorder"
				},
				{
					"label": "tinyint单选",
					"prop": "status"
				},
				{
					"label": "日期",
					"prop": "put_date"
				},
				{
					"label": "时间日期",
					"prop": "put_datetime"
				},
				{
					"label": "开始时间$end=end_time",
					"prop": "start_time"
				},
				{
					"label": "结束时间被引用隐藏到代码区间选择器",
					"prop": "end_time"
				}
			]
		});

		function onLoad({ ctx, app }: CrudLoad) {
			ctx.service(service.todo).done();
			app.refresh();
		}

		return {
			form,
			upsert,
			table,
			onLoad
		};
	}
});
</script>
