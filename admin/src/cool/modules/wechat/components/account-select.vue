<template>
	<el-select
		filterable
		placeholder="选择公众号"
		size="mini"
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
import { defineComponent, inject, onMounted, ref } from "vue";
import { ElMessage } from "element-plus";

export default defineComponent({
	name: "account-select",

	props: {
		modelValue: {
			type: String,
			default: ""
		}
	},

	emits: ["update:modelValue"],

	setup(props, { emit }) {
		const service = inject<any>("service");

		const options = ref([]);
		console.log(props);
		const value = ref(props.modelValue);

		onMounted(() => {
			service.wechat.account
				.select()
				.then((data: any) => {
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
			options,
			value,
			onCurrentChange
		};
	}
});
</script>
