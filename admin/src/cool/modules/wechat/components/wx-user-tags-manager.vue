<template>
	<cl-dialog title="公众号用户标签管理" :close-on-click-modal="false" v-model="visible">
		<div class="panel flex flex-wrap" v-loading="submitting" style="min-height: 165px">
			<el-tag
				size="small"
				v-for="tag in wxUserTags"
				closable
				@click="editTag(tag.id, tag.name)"
				@close="deleteTag(tag.id)"
				:disable-transitions="false"
				:key="tag.id"
			>
				{{ tag.id }} {{ tag.name }}
			</el-tag>
			<el-input
				class="input-new-tag"
				v-if="inputVisible"
				placeholder="回车确认"
				v-model="inputValue"
				ref="saveTagInput"
				size="mini"
				@keyup.enter.native="addTag"
			/>
			<el-button v-else class="button-new-tag" size="mini" @click="showInput"
				>+ 添加</el-button
			>
		</div>
		<span class="dialog-footer" >
			<el-button @click="visible = false" size="mini">关闭</el-button>
		</span>
	</cl-dialog>
</template>
<script>
export default {
	name: "wx-user-tags-manager",
	emits: ["change"],
	data() {
		return {
			appid: "",
			visible: false,
			wxUserTags: [],
			dialogVisible: false,
			inputVisible: false,
			inputValue: "",
			submitting: false
		};
	},
	methods: {
		init(appid) {
			console.log(this.appid);
			this.appid = appid;
			this.visible = true;
			this.submitting = false;
			this.getWxUserTags();
		},
		getWxUserTags() {
			this.service.wechat.tags
				.list({
					appid: this.appid
				})
				.then((data) => {
					this.wxUserTags = data || [];
				})
				.catch((e) => {
					this.$message.error(e);
				});
		},
		deleteTag(tagid) {
			if (this.submitting) {
				return;
			}
			this.$confirm(`确定删除标签?`, "提示", {
				confirmButtonText: "确定",
				cancelButtonText: "取消",
				type: "warning"
			}).then(() => {
				this.submitting = true;
				this.service.wechat.tags
					.delete({
						ids: [tagid],
						id: tagid,
						appid: this.appid
					})
					.then(() => {
						this.getWxUserTags();
						// this.$emit("change");
						this.submitting = false;
					})
					.catch((e) => {
						this.$message.error(e);
					});
			});
		},
		showInput() {
			this.inputVisible = true;
			this.$nextTick((_) => {
				this.$refs.saveTagInput.$refs.input.focus();
			});
		},
		addTag() {
			let newTagName = this.inputValue;
			this.saveTag(newTagName);
			this.inputVisible = false;
			this.inputValue = "";
		},
		editTag(tagid, orignName = "") {
			this.$prompt("请输入新标签名称", "提示", {
				confirmButtonText: "确定",
				cancelButtonText: "取消",
				inputValue: orignName,
				inputPattern: /^.{1,30}$/,
				inputErrorMessage: "名称1-30字符"
			}).then(({ value }) => {
				console.log(value);
				this.saveTag(value, tagid);
			});
		},
		saveTag(name, tagid) {
			if (this.submitting) {
				return;
			}
			this.submitting = true;
			this.service.wechat.tags[tagid ? "edit" : "add"]({
				id: tagid ? tagid : undefined,
				name: name
			})
				.then(() => {
					this.getWxUserTags();
					this.$emit("change");
					this.submitting = false;
				})
				.catch((e) => {
					this.$message.error(e);
				});
		}
	}
};
</script>
<style scoped>
.panel {
	flex: 1;
}
.el-tag,
.button-new-tag {
	margin: 5px;
}
.input-new-tag {
	width: inherit;
}
</style>
