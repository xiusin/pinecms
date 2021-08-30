<template>
	<cl-dialog
		:title="modeDesc[mode] + '用户标签'"
		:close-on-click-modal="false"
		v-model="dialogVisible"
	>
		<div>
			<el-select
				v-model="selectedTagid"
				filterable
				size="mini"
				placeholder="请选择标签"
				style="width: 100%"
			>
				<el-option
					v-for="tagid in tagsInOpts"
					:key="tagid"
					:label="getTagName(tagid)"
					:value="tagid"
				/>
			</el-select>
			<div style="margin: 20px 0; font-size: 12px">已选择用户数：<font style="color: red;">{{ wxUsers.length }}</font>
			</div>
		</div>
		<span slot="footer" class="dialog-footer">
			<el-button @click="dialogVisible = false" size="mini">关闭</el-button>
			<el-button type="primary" @click="dataFormSubmit()" :disabled="submitting" size="mini">确定</el-button>
		</span>
	</cl-dialog>
</template>
<script>

export default {
	name: "wx-user-tagging",
	props: {
		wxUsers: Array,

	},
	data() {
		return {
			wxUserTags: [],
			tagsInOpts: [],
			mode: "tagging", //操作，tagging | untagging
			modeDesc: {
				tagging: "绑定",
				untagging: "解绑"
			},
			appid: "",
			selectedTagid: "",
			dialogVisible: false,
			submitting: false
		};
	},
	/**
	 * 返回下拉选择框中的选项列表
	 * 假设 all= 全部标签，intersection = 用户标签交集（即所有用户都有的） ，union=用户标签并集（即至少一个用户的）
	 * 那么绑定时可选：all-intersection的差集，即所有用户都有的就不列出来了
	 *     解绑时可选：，union ，即用户有的标签都列出来
	 */
	methods: {
		async init(mode, appid) {
			if ("tagging" === mode || "untagging" === mode) {
				this.mode = mode;
				this.dialogVisible = true;
				this.appid = appid;
			} else {
				throw "mode参数有误";
			}
			let userTags = this.wxUsers.map((u) => u.tagid_list || []); //示例：[[1,2],[],[1,3]]
			this.wxUserTags = await this.service.wechat.tags.list({appid: this.appid})

			if (this.mode === "tagging") {
				//绑定标签时可选：所有标签 - 用户标签交集
				let all = this.wxUserTags.map((item) => item.id);
				this.tagsInOpts = all.filter(
					(tagid) => !userTags.every((tagsIdArray) => tagsIdArray.indexOf(tagid) > -1)
				);
			} else if (this.mode === "untagging") {
				let unionSet = new Set();
				userTags.forEach((tagsIdArray) => {
					tagsIdArray.forEach((tagid) => unionSet.add(tagid));
				}); //将用户的标签放到unionSet中去重
				this.tagsInOpts = Array.from(unionSet); //unionSet转为数组
			}
		},
		getTagName(tagid) {
			let tag = this.wxUserTags.find((item) => item.id == tagid);
			return tag ? tag.name : "?";
		},
		dataFormSubmit() {
			if (this.submitting) return;
			if (!this.selectedTagid) {
				this.$message.error("未选择标签");
				return;
			}
			this.submitting = true;
			let openidList = this.wxUsers.map((u) => u.openid);
			this.service.wechat.tags.tagging({
				id: this.selectedTagid,
				openids: openidList,
				action: this.mode,
				appid: this.appid
			})
				.then(() => {
					this.submitting = false;
					this.$message({
						message: "操作成功,列表数据需稍后刷新查看",
						type: "success",
						onClose: () => (this.dialogVisible = false)
					});
				}).catch((e) => {
				this.submitting = false;
				this.$message.error(e);
			})
		}
	}
};
</script>
