<template>
	<el-checkbox-group v-model="checkList" size="mini" @change="changeVal">
		<el-checkbox v-for="item in props.options" :key="item.key + item.label" :label="item.key"
			>{{ item.label }}
		</el-checkbox>
	</el-checkbox-group>
</template>

<script lang="ts">
import { defineComponent, ref } from "vue";

export default defineComponent({
	name: "cms-checkbox",
	props: {
		modelValue: {
			type: String,
			default: () => ""
		},
		options: {
			type: Array,
			default: () => []
		}
	},
	emits: ["update:modelValue"],

	setup(props, { emit }) {
		const checkList = ref([]);
		try {
			checkList.value = props.modelValue.split(",");
		} catch (e) {
			checkList.value = [];
		}

		function changeVal(val) {
			emit("update:modelValue", val.join(","));
		}

		return { checkList, props, changeVal };
	}
});
</script>
