<template>
	<iframe
		style="width: 100%; height: 100%"
		src="http://localhost:2019/debug/pprof/"
		frameborder="0"
		id="pprofIframe"
	>
	</iframe>
</template>

<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
	name: "sys-statsviz",

	setup() {
		return {};
	}
});
</script>
