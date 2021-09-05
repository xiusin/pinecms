<template>

</template>
<script>
import MaterialFile from "./material-file.vue";
import MaterialNews from "./material-news.vue";
export default {
	name: "assets-selector",
	components: {
		MaterialFile,
		MaterialNews
	},
	props: {
		selectType: {
			type: String,
			default: "image"
		},
		visible: {
			type: Boolean,
			default: true
		}
	},
	emits: ["selected", "onClose"],
	data: function () {
	  console.log("init");
		return {
		};
	},
	methods: {
		onSelect(itemInfo) {
			this.$emit("selected", itemInfo);
		},
		onClose() {
			this.$emit("onClose");
		}
	}
};
</script>
