<template>
	<div style="padding: 10px; background: #fff">
		<div style="padding: 10px 0">
			<account-select v-model="appid" />
			<el-button style="margin-left: 5px;" size="mini" type="warning" @click="sync()" :disabled="appid === ''">同步素材</el-button>
		</div>
		<el-tabs v-model="activeTab" @tab-click="handleTabClick">
			<el-tab-pane :label="'图片素材（' + assetsCount.imageCount + ')'" name="image">
				<material-file fileType="image" ref="imagePanel" @change="materialCount" />
			</el-tab-pane>
			<el-tab-pane :label="'语音素材（' + assetsCount.voiceCount + ')'" name="voice" lazy>
				<material-file fileType="voice" ref="voicePanel" @change="materialCount" />
			</el-tab-pane>
			<el-tab-pane :label="'视频素材（' + assetsCount.videoCount + ')'" name="video" lazy>
				<material-file fileType="video" ref="videoPanel" @change="materialCount" />
			</el-tab-pane>
			<el-tab-pane :label="'图文素材（' + assetsCount.newsCount + ')'" name="news" lazy>
				<material-news ref="newsPanel" @change="materialCount" />
			</el-tab-pane>
		</el-tabs>
	</div>
</template>
<script>
import MaterialFile from "./assets/material-file.vue";
import MaterialNews from "./assets/material-news.vue";
import AccountSelect from "../components/account-select.vue";

export default {
	components: {
		AccountSelect,
		MaterialFile,
		MaterialNews
	},
	data() {
		return {
			appid: "",
			activeTab: "image",
			assetsCount: { imageCount: 0, videoCount: 0, voiceCount: 0, newsCount: 0 }
		};
	},
	mounted() {
		this.materialCount();
	},
	methods: {
		sync() {
			this.service.wechat.material.sync().then(() => {
				this.$message.success("同步完成");
			}).catch((e) => {
				this.$message.error(e);
			})
		},
		handleTabClick(tab) {
			this.$nextTick(() => {
				this.$refs[tab.paneName + "Panel"].init();
			});
		},
		materialCount() {
			this.service.wechat.material
				.total()
				.then((data) => {
					this.assetsCount.imageCount = data.image_count;
					this.assetsCount.videoCount = data.video_count;
					this.assetsCount.voiceCount = data.voice_count;
					this.assetsCount.newsCount = data.news_count;

					this.$refs[this.activeTab + "Panel"].init();
				})
				.catch((e) => {
					this.$message.error(e);
				});
		}
	}
};
</script>
