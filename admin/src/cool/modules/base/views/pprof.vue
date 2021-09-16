<template>
	<iframe style="width: 100%; height: 100%" :src="url" frameborder="0" id="pprofIframe"> </iframe>
</template>

<script lang="ts">
import { defineComponent, inject, onMounted, ref } from "vue";

export default defineComponent({
	name: "sys-pprof",
	setup() {
		const url = ref("");

		const service = inject<any>("service");

		onMounted(() => {
			service.system.common.pprof().then((data: any) => {
				url.value = data;
			});
		});

		return {
			url
		};
	}
});
</script>
