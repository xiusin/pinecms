<template>
	<div class="mod-menu">
		<el-form :inline="true" :model="dataForm">
			<el-form-item>
				<el-button size="mini" type="primary" @click="addOrUpdateHandle()">新增</el-button>
			</el-form-item>
		</el-form>
		<div v-loading="dataListLoading">
			<div class="card" v-for="item in dataList" :key="item.media_id" @click="onSelect(item)">
				<el-image
					v-if="fileType === 'image'"
					class="card-image"
					alt=""
					:src="item.url"
					fit="contain"
				/>
				<div v-else class="card-preview">
					<div
						v-if="fileType === 'voice'"
						class="card-preview-icon el-icon-microphone"
					></div>
					<div
						v-if="fileType === 'video'"
						class="card-preview-icon el-icon-video-camera-solid"
					></div>
					<div class="card-preview-text">管理后台不支持预览<br />微信中可正常播放</div>
				</div>
				<div class="card-footer">
					<div class="text-cut-name">{{ item.name }}</div>
					<!--                    <div>{{$moment(item.updateTime).calendar()}}</div>-->
					<div class="flex justify-between align-center" v-show="!selectMode">
						<el-button
							size="mini"
							type="text"
							v-copy="item.media_id"
							icon="el-icon-copy-document"
							>复制media_id
						</el-button>
						<el-button
							size="mini"
							type="text"
							icon="el-icon-delete"
							@click="deleteHandle(item.media_id)"
							>删除
						</el-button>
					</div>
				</div>
			</div>
		</div>
		<el-pagination
			@current-change="currentChangeHandle"
			:current-page="pageIndex"
			:page-sizes="[20]"
			:page-size="20"
			:total="totalCount"
			layout="total, prev,pager, next, jumper"
		/>
		<!-- 弹窗, 新增 / 修改 -->
		<add-or-update v-if="addOrUpdateVisible" ref="addOrUpdate" @refreshDataList="onChange" />
	</div>
</template>
<script>
import AddOrUpdate from "./material-file-add-or-update.vue";

export default {
	name: "material-file",
	components: {
		AddOrUpdate
	},
	props: {
		fileType: {
			type: String,
			default: "image"
		},
		selectMode: {
			// 是否选择模式，选择模式下点击素材选中，不可新增和删除
			type: Boolean,
			default: false
		}
	},
	data() {
		return {
			dataForm: {},
			addOrUpdateVisible: false,
			dataList: [],
			pageIndex: 1,
			pageSize: 20,
			totalCount: 0,
			dataListLoading: false
		};
	},
	mounted() {},
	methods: {
		init() {
			if (!this.dataList.length) {
				this.getDataList();
			}
		},
		getDataList() {
			if (this.dataListLoading) return;
			this.dataListLoading = true;
			this.service.wechat.material
				.page({
					page: this.pageIndex,
					type: this.fileType,
					size: this.pageSize,
					appid: "wxe43df03110f5981b"
				})
				.then((data) => {
					this.dataList = data.item;
					this.totalCount = data.total_count;
					this.pageIndex++;
					this.dataListLoading = false;
				})
				.catch((e) => {
					this.dataListLoading = false;
					this.$message.error(e);
				});
		},
		// 新增 / 修改
		addOrUpdateHandle() {
			this.addOrUpdateVisible = true;
			this.$nextTick(() => {
				this.$refs.addOrUpdate.init(this.fileType);
			});
		},
		onSelect(itemInfo) {
			if (!this.selectMode) return;
			this.$emit("selected", itemInfo);
		},
		//删除
		deleteHandle(id) {
			this.$confirm(`确定对[mediaId=${id}]进行删除操作?`, "提示", {
				confirmButtonText: "确定",
				cancelButtonText: "取消",
				type: "warning"
			})
				.then(() => {
					this.service.wechat.material.delete({ media_id: id }).then(() => {
						this.$message({
							message: "操作成功",
							type: "success",
							duration: 1500,
							onClose: () => this.onChange()
						});
					});
				})
				.catch((e) => {
					this.$message.error(e);
				});
		},
		// 当前页
		currentChangeHandle(val) {
			this.pageIndex = val;
			// this.getDataList()
		},
		onChange() {
			this.pageIndex = 1;
			// this.getDataList()
			this.$emit("change");
		}
	}
};
</script>
<style scoped>
.card {
	width: 170px;
	display: inline-block;
	background: #ffffff;
	border: 1px solid #ebeef5;
	box-shadow: 1px 1px 20px 0 rgba(0, 0, 0, 0.1);
	margin: 0 10px 10px 0;
	vertical-align: top;
	box-sizing: border-box;
}

.card:hover {
	border: 1px solid #66b1ff;
	margin-bottom: 6px;
}

.card-image {
	line-height: 170px;
	max-height: 170px;
	width: 100%;
}

.card-preview {
	padding: 20px 0;
	color: #d9d9d9;
	display: flex;
	justify-content: center;
	align-items: center;
}

.card-preview-icon {
	font-size: 30px;
	margin-right: 5px;
}

.card-preview-text {
	font-size: 12px;
}

.card-footer {
	color: #ccc;
	font-size: 12px;
	padding: 15px 10px;
}
</style>
