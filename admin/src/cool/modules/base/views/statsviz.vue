<template>
	<iframe style="width:100%;height:100%;" src="http://localhost:2019/debug/statsviz/" frameborder="0" scrolling="no" id="bdIframe">
	</iframe>
</template>

<script lang="ts">
import {defineComponent} from "vue";

export default defineComponent({
	name: "sys-statsviz",

	setup() {
		return {

		};
	}
});
</script>
