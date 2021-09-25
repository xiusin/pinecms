<template>
	<iframe style="width: 100%; height: 100%" :src="url" frameborder="0" id="bdIframe"> </iframe>
</template>

<script lang="ts">
import { defineComponent, inject, onMounted, ref } from "vue";

export default defineComponent({
	name: "sys-statsviz",

	setup() {
		const url = ref("");
		const service = inject<any>("service");
		onMounted(() => {
			service.common.statsviz().then((data: any) => {
				url.value = data;
			});
		});
		return {
			url
		};
	}
});
</script>
