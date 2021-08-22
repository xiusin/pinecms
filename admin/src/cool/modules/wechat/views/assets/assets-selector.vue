<template>

</template>
<script>
export default {
	name: "assets-selector",
	emits: ["selected", "onClose"],
	data: function () {
		return {
			dataVisible: this.visible
		}
	},
	components: {
		MaterialFile: () => import("./material-file.vue"),
		MaterialNews: () => import("./material-news.vue")
	},
	props: {
		selectType: {
			type: String,
			default: "image"
		},
		visible: {
			type: Boolean,
			default: false
		}
	},
	methods: {
		onSelect(itemInfo) {
			this.$emit("selected", itemInfo);
		},
		onClose() {
			this.$emit("onClose");
		}
	}
}
</script>
