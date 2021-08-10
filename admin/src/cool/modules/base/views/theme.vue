<template>
	<el-row :gutter="20">
		<el-col :span="4" v-for="(o, index) in themeList" :key="index">
			<el-card :body-style="{ padding: '0px' }">
				<img :src="baseUrl + '/assets/thumb?id=' + o.dir" class="image" />
				<div style="padding: 14px">
					<span>{{ o.name }}</span>
					<div class="bottom clearfix">
						<time class="time">{{ o.description }}</time>
						<template v-if="o.is_default">
							<el-button type="text" class="button" disabled>√</el-button>
						</template>
						<template v-else>
							<el-button type="text" class="button" onclick="setTheme(o.dir)"
								>√</el-button
							>
						</template>
					</div>
				</div>
			</el-card>
		</el-col>
	</el-row>
</template>

<script lang="ts">
import { defineComponent, inject, ref } from "vue";
import { baseUrl } from "/@/config/env";

import { useRefs } from "/@/core";

export default defineComponent({
	name: "sys-theme",
	setup() {
		const service = inject<any>("service");
		const { refs, setRefs } = useRefs();

		let themeList = ref([]);

		function reloadThemes() {
			service.system.assets.themes().then((data) => {
				themeList.value?.push(...data);
			});
		}

		reloadThemes();

		function setTheme(dirname) {
			service.system.assets
				.theme({
					theme: dirname
				})
				.then((data) => {
					reloadThemes();
				});
		}

		const currentDate = new Date();
		return {
			refs,
			setRefs,
			currentDate,
			themeList,
			baseUrl,
			setTheme
		};
	}
});
</script>

<style scoped>
.time {
	font-size: 13px;
	color: #999;
}

.bottom {
	margin-top: 13px;
	line-height: 12px;
}

.button {
	padding: 0;
	float: right;
}

.image {
	width: 100%;
	display: block;
}

.clearfix:before,
.clearfix:after {
	display: table;
	content: "";
}

.clearfix:after {
	clear: both;
}
</style>
