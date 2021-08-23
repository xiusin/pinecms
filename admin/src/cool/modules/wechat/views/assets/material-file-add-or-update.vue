<template>
	<cl-dialog
		:title="!dataForm.id ? '新增' : '修改'"
		:close-on-click-modal="false"
		v-model="visible"
	>
		<el-form :model="dataForm" :rules="dataRule" ref="dataForm" label-width="80px">
			<el-form-item label="媒体文件">
				<el-button type="primary" size="mini">
					选择文件
					<input
						type="file"
						style="opacity: 0; height: 100%; position: absolute; left: 0; top: 0"
						@change="onFileChange"
					/>
				</el-button>
				<div>{{ dataForm.file.name }}</div>
			</el-form-item>
			<el-form-item label="媒体类型" prop="mediaType">
				<el-select
					v-model="dataForm.mediaType"
					placeholder="媒体类型"
					style="width: 100%"
					size="mini"
				>
					<el-option label="图片（2M以内，支持PNG\JPEG\JPG\GIF）" value="image" />
					<el-option label="视频（10M以内，只支持MP4）" value="video" />
					<el-option label="语音（2M、60s以内，支持AMR\MP3）" value="voice" />
					<el-option label="缩略图（64K以内JPG）" value="thumb" />
				</el-select>
			</el-form-item>
			<el-form-item label="素材名称" prop="fileName">
				<el-input
					size="mini"
					v-model="dataForm.fileName"
					placeholder="为便于管理建议按用途分类+素材内容命名"
				/>
			</el-form-item>
			<el-form-item label="视频标题" prop="title">
				<el-input
					size="mini"
					v-model="dataForm.title"
					placeholder="为便于管理建议按用途分类+素材内容命名"
				/>
			</el-form-item>
			<el-form-item label="视频标题" prop="title" v-if="dataForm.mediaType !== 'video'">
				<el-input size="mini" v-model="dataForm.title" placeholder="视频标题" />
			</el-form-item>
			<el-form-item
				label="视频描述"
				prop="introduction"
				v-if="dataForm.mediaType !== 'video'"
			>
				<el-input
					size="mini"
					v-model="dataForm.introduction"
					placeholder="为便于管理建议按用途分类+素材内容命名"
				/>
			</el-form-item>
		</el-form>
		<span class="dialog-footer">
			<el-button @click="visible = false" size="mini">取消</el-button>
			<el-button type="primary" @click="dataFormSubmit()" size="mini">{{
				uploading ? "提交中..." : "提交"
			}}</el-button>
		</span>
	</cl-dialog>
</template>

<script>
export default {
	data() {
		return {
			visible: false,
			uploading: false,
			dataForm: {
				mediaId: "",
				file: "",
				fileName: "",
				title: "",
				introduction: "",
				mediaType: "image"
			},
			dataRule: {
				fileName: [{ required: true, message: "素材名称不能为空", trigger: "blur" }],
				mediaType: [{ required: true, message: "素材类型不能为空", trigger: "blur" }]
			}
		};
	},
	methods: {
		init(fileType) {
			if (fileType) this.dataForm.mediaType = fileType;
			this.visible = true;
		},
		// 表单提交
		dataFormSubmit() {
			if (this.uploading) return;
			this.$refs["dataForm"].validate((valid) => {
				if (valid) {
					this.uploading = true;
					let form = new FormData();
					form.append("mediaId", this.dataForm.mediaId || "");
					form.append("file", this.dataForm.file);
					form.append("fileName", this.dataForm.fileName);
					form.append("mediaType", this.dataForm.mediaType);
					if (this.dataForm.mediaType === "video") {
						form.append("title", this.dataForm.title);
						form.append("introduction", this.dataForm.introduction);
					}
					this.service.wechat.material
						.upload(form)
						.then(({ data }) => {
							this.$message({
								message: "操作成功",
								type: "success",
								duration: 1500,
								onClose: () => {
									this.visible = false;
									this.$emit("refreshDataList");
								}
							});
							this.uploading = false;
						})
						.catch((e) => {
							this.uploading = false;
							this.$message.error(e);
						});
				} else {
					this.$message.error("验证失败");
				}
			});
		},
		onFileChange(e) {
			let file = event.currentTarget.files[0];
			this.dataForm.file = file;
			this.dataForm.fileName = file.name.substring(0, file.name.lastIndexOf("."));
			let mediaType = file.type.substring(0, file.type.lastIndexOf("/"));
			if (mediaType === "audio") mediaType = "voice";
			this.dataForm.mediaType = mediaType;
		}
	}
};
</script>
