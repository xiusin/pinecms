<template>
	<el-radio-group v-model="val" size="mini" @change="changeVal">
		<el-radio v-for="item in props.options" :label="item.key">{{ item.label }}</el-radio>
	</el-radio-group>
</template>

<script lang="ts">
import {defineComponent, ref} from "vue";

export default defineComponent({
	name: "cms-radio",
	emits: ["update:modelValue"],
	props: {
		modelValue: {
			type: Number,
			default: () => 0
		},
		options: {
			type: Array,
			default: () => []
		}
	},

	setup(props, {emit}) {
		console.log(props)
		const val = ref(props.modelValue)
		function changeVal(val) {
			emit('update:modelValue', parseInt(val))
		}
		return {val, props, changeVal};
	}
});
</script>
