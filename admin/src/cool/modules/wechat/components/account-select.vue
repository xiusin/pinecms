<template>
	<el-select
		filterable
		placeholder="选择公众号"
		:size="size"
		v-model="value"
		clearable
		@change="onCurrentChange"
	>
		<el-option
			v-for="(item, idx) in options"
			:key="idx + '-' + item.value"
			:value="item.value"
			:label="item.label"
		/>
	</el-select>
</template>

<script lang="ts">
import { defineComponent, inject, onMounted, reactive, ref } from "vue";
import { ElMessage } from "element-plus";

export default defineComponent({
	name: "account-select",

	props: {
		modelValue: [Number, String],
		size: {
			type: String,
			default: "mini",
		}
	},

	emits: ["update:modelValue"],

	setup(props, { emit }) {
		const service = inject<any>("service");

		const options = ref([]);

		const value = ref(props.modelValue);

		const size = ref(props.size)

		onMounted(() => {
			service.wechat.account
				.select()
				.then((data: any) => {
					// data.unshift({ label: "请选择公众号", value: "" });
					options.value = data;
				})
				.catch((e: any) => {
					ElMessage.error(e);
				});
		});

		// 绑定值回调
		function onCurrentChange(value: any) {
			emit("update:modelValue", value);
		}

		return {
			value,
			size,
			options,
			onCurrentChange
		};
	}
});
</script>
