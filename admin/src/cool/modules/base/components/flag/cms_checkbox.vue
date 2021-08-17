<template>
	<el-checkbox-group v-model="checkList" size="mini" @change="changeVal">
		<el-checkbox v-for="item in flags" :key="item.key + item.label" :label="item.key">{{item.label}}</el-checkbox>
	</el-checkbox-group>
</template>

<script lang="ts">
import {defineComponent, ref} from "vue";

export default defineComponent({
	name: "cms-checkbox",
	emits: ["update:modelValue"],
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

	setup(props, {emit}) {
		const checkList = ref(props.modelValue.split(','))
		const flags = ref(props.options);
		function changeVal(val) {
			emit('update:modelValue', val.join(','))
		}
		return {
			checkList,
			props,
			flags,
			changeVal
		};
	}
});
</script>
