<template>
	<cl-form ref="settingFormRef" inner></cl-form>
</template>

<script lang="ts">
import {defineComponent, ref} from "vue";
import {useRefs} from "../../../../core";

export default defineComponent({
	name: "sys-setting",

	setup() {
		const { refs, setRefs } = useRefs();
		const formRef = refs.settingFormRef;
		console.log("formRef", formRef)
		formRef.open({
			items: [
				{
					type: "tabs",
					props: {
						labels: [
							{
								label: "基本信息",
								value: "base"
							},
							{
								label: "金融",
								value: "financial"
							}
						]
					}
				},
				{
					label: "昵称",
					prop: "name",
					component: {
						name: "el-input",

						attrs: {
							placeholder: "请填写昵称"
						}
					},
					rules: {
						required: true,
						message: "昵称不能为空"
					}
				},
				{
					label: "年纪",
					prop: "age",
					value: 18,
					component: {
						name: "el-input-number"
					}
				},
				{
					label: "支付宝",
					prop: "alipay",
					value: 300,
					group: "financial",
					component: {
						name: "el-input-number"
					},
					append: ({ h }) => {
						return h("p", "元");
					}
				},
				{
					label: "微信",
					prop: "wechat",
					value: 10,
					group: "financial",
					component: {
						name: "el-input-number"
					},
					append: ({ h }) => {
						return h("p", "元");
					}
				}
			],
			on: {
				submit: (data, { close }) => {
					close();
				}
			}
		});

		return {
			formRef
		};
	}
});
</script>

<style lang="scss" scoped>
.change-btn {
	display: flex;
	position: absolute;
	right: 10px;
	bottom: 10px;
	z-index: 9;
}

.editor {
	transition: all 0.3s;
}
</style>
