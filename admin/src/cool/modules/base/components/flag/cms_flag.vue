<template>
	<el-checkbox-group v-model="checkList" size="mini">
		<el-checkbox v-for="item in flags" :label="item.label" :key="item.key"></el-checkbox>
	</el-checkbox-group>
</template>

<script lang="ts">
import { defineComponent, ref } from "vue";

export default defineComponent({
	name: "cms-flag",

	props: {
		modelValue: {
			type: String,
			default: () => ""
		}
	},

	setup(props) {
		const checkList = ref(props.modelValue.split(','))
		console.log(checkList.value)
		const flags =ref([
			{"label":"头条","key":"h"},
			{"label":"推荐","key":"c"},
			{"label": "图片", "key": "p"},
			{"label": "幻灯", "key": "f"},
			{"label": "跳转", "key": "j"},
			{"label": "图文", "key": "a"},
			{"label": "加粗", "key": "b"},
		]);

		return {
			checkList,
			props,
			flags
		};
	}
});
</script>
