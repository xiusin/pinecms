<template>
	<div style="padding: 10px">
		<el-tabs v-model="activeTab" @tab-click="handleTabClick">
			<el-tab-pane :label="'图片素材（'+assetsCount.imageCount+')'" name="image">
				<material-file fileType="image" ref="imagePanel" @change="materialCount"/>
			</el-tab-pane>
			<el-tab-pane :label="'语音素材（'+assetsCount.voiceCount+')'" name="voice" lazy>
				<material-file fileType="voice" ref="voicePanel" @change="materialCount"/>
			</el-tab-pane>
			<el-tab-pane :label="'视频素材（'+assetsCount.videoCount+')'" name="video" lazy>
				<material-file fileType="video" ref="videoPanel" @change="materialCount"/>
			</el-tab-pane>
			<el-tab-pane :label="'图文素材（'+assetsCount.newsCount+')'" name="news" lazy>
				<material-news ref="newsPanel" @change="materialCount"/>
			</el-tab-pane>
		</el-tabs>
	</div>
</template>
<script>

import MaterialFile from './assets/material-file.vue'
import MaterialNews from './assets/material-news.vue'

export default {
	data() {
		return {
			activeTab: 'image',
			assetsCount: {imageCount: 0, videoCount: 0, voiceCount: 0, newsCount: 0}
		};
	},
	components: {
		MaterialFile,
		MaterialNews
	},
	mounted() {
		this.materialCount();
	},
	methods: {
		handleTabClick(tab, event) {
			this.$nextTick(() => {
				this.$refs[tab.paneName + 'Panel'].init();
			})
		},
		materialCount() {
			this.service.wechat.material.total().then((data) => {
				this.assetsCount.imageCount = data.image_count
				this.assetsCount.videoCount = data.video_count
				this.assetsCount.voiceCount = data.voice_count
				this.assetsCount.newsCount = data.news_count
			}).catch((e) => {
				this.$message.error(e);
			})
		}
	}
};
</script>
