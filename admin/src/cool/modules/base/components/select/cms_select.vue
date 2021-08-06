<template>
	<el-select
		:filterable="props.filterable"
		:placeholder="props.placeholder"
		:multiple="props.multiple"
		:multipleLimit="props.multipleLimit"
		:automatic-dropdown="true"
		:size="props.size"
		:value="props.value"
	>
		<el-option
			v-for="(item, idx) in props.options"
			:key="idx + '-' + item.value"
			:value="item.value"
			:label="item.label"
		/>
	</el-select>
</template>

<script lang="ts">
import { defineComponent, onMounted } from "vue";

export default defineComponent({
	name: "cms-select",

	props: {
		multiple: {
			type: Boolean,
			default: false
		},
		collapseTags: {
			type: Boolean,
			default: false
		},
		filterable: {
			type: Boolean,
			default: true
		},
		placeholder: {
			type: String,
			default: "请选择"
		},
		multipleLimit: {
			type: Number,
			default: 10
		},
		url: {
			type: String,
			default: ""
		},
		size: {
			type: String,
			default: "mini"
		},
		autocomplete: {
			type: String,
			default: "off"
		},
		options: {
			type: Array
		}
	},

	setup(props) {
		const remoteMethod = (query: string) => {
			if (query !== "") {
			}
		};

		// TODO 远程加载资源
		if (props.url) {
			// todo 分为字典类型 或 完整api地址
			onMounted(() => {
				// 远程加载字典 或 其他表字典
			});
		}

		return {
			props,
			remoteMethod
		};
	}
});
</script>
