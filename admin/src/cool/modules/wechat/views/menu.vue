<template>
	<div style="padding: 10px; background: #fff">
		<div id="app-menu">
			<div class="weixin-preview">
				<div class="weixin-bd">
					<div class="weixin-header">公众号菜单
						<account-select v-model="appid"/>
					</div>
					<ul class="weixin-menu" id="weixin-menu">
						<li
							v-for="(btn, i) in menu.buttons"
							:key="i"
							class="menu-item"
							:class="{ current: selectedMenuIndex === i && selectedMenuLevel === 1 }"
							@click="selectMenu(i)"
						>
							<div class="menu-item-title">
								<span>{{ btn.name }}</span>
							</div>
							<ul class="weixin-sub-menu">
								<li
									v-for="(sub, i2) in (btn.sub_button || [])"
									:key="i2"
									class="menu-sub-item"
									:class="{
										current:
											selectedMenuIndex === i &&
											selectedSubMenuIndex === i2 &&
											selectedMenuLevel === 2,
										'on-drag-over': onDragOverMenu === i + '_' + i2
									}"
									@click.stop="selectSubMenu(i, i2)"
									draggable="true"
									@dragstart="selectSubMenu(i, i2)"
									@dragover.prevent="onDragOverMenu = i + '_' + i2"
									@drop="onDrop(i, i2)"
								>
									<div class="menu-item-title">
										<span>{{ sub.name }}</span>
									</div>
								</li>
								<li
									v-if="(btn.sub_button || []).length < 5"
									class="menu-sub-item"
									:class="{
										'on-drag-over':
											onDragOverMenu === i + '_' + (btn.sub_button || []).length
									}"
									@click.stop="addMenu(2, i)"
									@dragover.prevent="
										onDragOverMenu = i + '_' + btn.sub_button.length
									"
									@drop="onDrop(i, btn.sub_button.length)"
								>
									<div class="menu-item-title">
										<i class="el-icon-plus"></i>
									</div>
								</li>
								<i class="menu-arrow arrow_out"></i>
								<i class="menu-arrow arrow_in"></i>
							</ul>
						</li>
						<li class="menu-item" v-if="menu.buttons.length < 3" @click="addMenu(1)">
							<i class="el-icon-plus"></i>
						</li>
					</ul>
				</div>
			</div>
			<!-- 菜单编辑器 -->
			<div class="weixin-menu-detail" v-if="selectedMenuLevel > 0">
				<wx-menu-button-editor
					:button="selectedButton"
					:selectedMenuLevel="selectedMenuLevel"
					@delMenu="delMenu"
				/>
			</div>
		</div>
		<div class="weixin-btn-group" @click="updateWxMenu">
			<el-button type="success" :disabled="appid === ''" icon="el-icon-upload">发布</el-button>
			<el-button type="warning" :disabled="appid === ''" icon="el-icon-delete" @click="delMenu">清空</el-button>
		</div>
	</div>
</template>
<script>

import wxMenuButtonEditor from "./wx-menu-button-editor.vue";
import AccountSelect from "../components/account-select.vue";

export default {
	components: {
		AccountSelect,
		wxMenuButtonEditor
	},
	data() {
		return {
			appid: "",
			menu: {buttons: [], menuid: 0}, //当前菜单
			selectedMenuIndex: 0, //当前选中菜单索引
			selectedSubMenuIndex: 0, //当前选中子菜单索引
			selectedMenuLevel: 0, //选中菜单级别
			selectedButton: 0, //选中的菜单按钮
			onDragOverMenu: 0 //当前鼠标拖动到的位置
		};
	},
	watch: {
		appid: function () {
			this.init();
		}
	},
	methods: {
		init() {
			this.service.wechat.menu.info({
				appid: this.appid,
			}).then((data) => {
				this.menu.buttons = data.button;
				this.menu.menuid = data.menuid;
			}).catch((e) => {
				this.$message.error(e)
				this.menu = {buttons: [], menuid: 0}
			});
		},

		//选中主菜单
		selectMenu(i) {
			this.selectedMenuLevel = 1;
			this.selectedSubMenuIndex = "";
			this.selectedMenuIndex = i;
			this.selectedButton = this.menu.buttons[i];
		},
		//选中子菜单
		selectSubMenu(i, i2) {
			this.selectedMenuLevel = 2;
			this.selectedMenuIndex = i;
			this.selectedSubMenuIndex = i2;
			this.selectedButton = this.menu.buttons[i].sub_button[i2];
		},
		//添加菜单
		addMenu(level, i) {
			if (this.appid === '') {
				this.$message.error('请先选择公众号');
				return;
			}
			try {
				if (level === 1 && this.menu.buttons.length < 3) {
					this.menu.buttons.push({
						type: "view",
						name: "菜单名称",
						sub_button: [],
						url: ""
					});
					this.selectMenu(this.menu.buttons.length - 1);
				}
				if (typeof this.menu.buttons[i].sub_button === "undefined") {
					this.menu.buttons[i].sub_button = [];
				}
				if (level === 2 && this.menu.buttons[i].sub_button.length < 5) {
					this.menu.buttons[i].sub_button.push({
						type: "view",
						name: "子菜单名称",
						url: ""
					});
					this.selectSubMenu(i, this.menu.buttons[i].sub_button.length - 1);
				}
			} catch (e) {
				console.error(e);
			}
		},
		//删除菜单
		delMenu() {
			if (this.selectedMenuLevel === 1 && confirm("删除后菜单下设置的内容将被删除")) {
				this.menu.buttons.splice(this.selectedMenuIndex, 1);
				this.unSelectMenu();
			} else if (this.selectedMenuLevel === 2) {
				this.menu.buttons[this.selectedMenuIndex].sub_button.splice(
					this.selectedSubMenuIndex,
					1
				);
				this.unSelectMenu();
			}
		},
		unSelectMenu() {
			this.selectedMenuLevel = 0;
			this.selectedMenuIndex = "";
			this.selectedSubMenuIndex = "";
			this.selectedButton = "";
		},
		updateWxMenu() {
			let btns = [];
			for (const idx in this.menu.buttons) {
				let subBtn =  this.menu.buttons[idx].sub_button || [];
				btns[idx] = this.menu.buttons[idx];
				btns[idx].sub_button = subBtn;
			}
			this.service.wechat.menu.update({
				appid: this.appid,
				menu: {
					button: this.menu.buttons || [],
					menuid: this.menu.menuid || 0
				},
			}).then((data) => {
				this.$message.success('发布菜单成功');
			}).catch(e => {
				this.$message.error(e);
			})
		},
		onDrop(i, i2) {
			//拖拽移动位置
			this.onDragOverMenu = "";
			if (i === this.selectedMenuIndex && i2 === this.selectedSubMenuIndex)
				//拖拽到了原位置
				return;
			if (i !== this.selectedMenuIndex && this.menu.buttons[i].sub_button.length >= 5) {
				this.$message.error("目标组已满");
				return;
			}
			this.menu.buttons[i].sub_button.splice(i2, 0, this.selectedButton);
			let delSubIndex = this.selectedSubMenuIndex;
			if (i === this.selectedMenuIndex && i2 < this.selectedSubMenuIndex) delSubIndex++;
			this.menu.buttons[this.selectedMenuIndex].sub_button.splice(delSubIndex, 1);
			this.unSelectMenu();
		}
	}
};
</script>
<style>
* {
	box-sizing: border-box;
}

#app-menu ul {
	padding: 0;
}

#app-menu li {
	list-style: none;
}

#app-menu {
	overflow: hidden;
	width: 100%;
}

.weixin-preview {
	position: relative;
	width: 320px;
	height: 540px;
	float: left;
	margin-right: 10px;
	border: 1px solid #e7e7eb;
}

.weixin-preview a {
	text-decoration: none;
	color: #616161;
}

.weixin-preview .weixin-hd .weixin-title {
	color: #fff;
	font-size: 15px;
	width: 100%;
	text-align: center;
	position: absolute;
	top: 33px;
	left: 0px;
}

.weixin-preview .weixin-header {
	text-align: center;
	padding: 10px 0;
	background-color: #616161;
	color: #ffffff;
}

.weixin-preview .weixin-menu {
	position: absolute;
	bottom: 0;
	left: 0;
	right: 0;
	border-top: 1px solid #e7e7e7;
	background-position: 0 0;
	background-repeat: no-repeat;
	margin-bottom: 0px;
}

/*一级*/
.weixin-preview .weixin-menu .menu-item {
	position: relative;
	float: left;
	line-height: 50px;
	height: 50px;
	text-align: center;
	width: 33.33%;
	border-left: 1px solid #e7e7e7;
	cursor: pointer;
	color: #616161;
}

/*二级*/
.weixin-preview .weixin-sub-menu {
	position: absolute;
	bottom: 60px;
	left: 0;
	right: 0;
	border-top: 1px solid #d0d0d0;
	margin-bottom: 0px;
	background: #fafafa;
	display: block;
	padding: 0;
}

.weixin-preview .weixin-sub-menu .menu-sub-item {
	line-height: 50px;
	height: 50px;
	text-align: center;
	width: 100%;
	border: 1px solid #d0d0d0;
	border-top-width: 0px;
	cursor: pointer;
	position: relative;
	color: #616161;
}

.weixin-preview .weixin-sub-menu .menu-sub-item.on-drag-over {
	border-top: 2px solid #44b549;
}

.weixin-preview .menu-arrow {
	position: absolute;
	left: 50%;
	margin-left: -6px;
}

.weixin-preview .arrow_in {
	bottom: -4px;
	display: inline-block;
	width: 0px;
	height: 0px;
	border-width: 6px 6px 0px;
	border-style: solid dashed dashed;
	border-color: #fafafa transparent transparent;
}

.weixin-preview .arrow_out {
	bottom: -5px;
	display: inline-block;
	width: 0px;
	height: 0px;
	border-width: 6px 6px 0px;
	border-style: solid dashed dashed;
	border-color: #d0d0d0 transparent transparent;
}

.weixin-preview .menu-item .menu-item-title,
.weixin-preview .menu-sub-item .menu-item-title {
	width: 100%;
	overflow: hidden;
	white-space: nowrap;
	text-overflow: ellipsis;
	box-sizing: border-box;
}

.weixin-preview .menu-item.current,
.weixin-preview .menu-sub-item.current {
	border: 1px solid #44b549;
	background: #fff;
	color: #44b549;
}

.weixin-preview .weixin-sub-menu.show {
	display: block;
}

.weixin-preview .icon_menu_dot {
	/* background: url(../images/index_z354723.png) 0px -36px no-repeat; */
	width: 7px;
	height: 7px;
	vertical-align: middle;
	display: inline-block;
	margin-right: 2px;
	margin-top: -2px;
}

.weixin-preview .icon14_menu_add {
	/* background: url(../images/index_z354723.png) 0px 0px no-repeat; */
	width: 14px;
	height: 14px;
	vertical-align: middle;
	display: inline-block;
	margin-top: -2px;
}

.weixin-preview li:hover .icon14_menu_add {
	/* background: url(../images/index_z354723.png) 0px -18px no-repeat; */
}

.weixin-preview .menu-item:hover {
	color: #000;
}

.weixin-preview .menu-sub-item:hover {
	background: #eee;
}

.weixin-preview li.current:hover {
	background: #fff;
	color: #44b549;
}

/*菜单内容*/
.weixin-menu-detail {
	width: calc(100% - 340px);
	padding: 0px 20px 5px;
	background-color: #f4f5f9;
	border: 1px solid #e7e7eb;
	float: left;
	min-height: 540px;
}

.weixin-menu-detail .menu-name {
	float: left;
	height: 40px;
	line-height: 40px;
	font-size: 18px;
}

.weixin-menu-detail .menu-del {
	float: right;
	height: 40px;
	line-height: 40px;
	color: #459ae9;
	cursor: pointer;
}

.weixin-menu-detail .menu-input-group {
	width: 100%;
	margin: 10px 0 10px 0;
	overflow: hidden;
}

.weixin-menu-detail .menu-label {
	float: left;
	font-size: 13px;
	margin-right: 10px;
	height: 40px;
	line-height: 40px;
	width: 80px;
	text-align: right;
}

.weixin-menu-detail .menu-input {
	float: left;
	width: 380px;
}

.weixin-menu-detail .menu-input-text {
	border: 0px;
	outline: 0px;
	background: #fff;
	width: 300px;
	padding: 5px 0px 5px 0px;
	margin-left: 10px;
	text-indent: 10px;
	height: 35px;
}

.weixin-menu-detail .menu-tips {
	color: #1c1c1c;
	padding: 4px 10px;
	margin-bottom: 15px;
	border-left: 2px solid #459ae9;
	height: 30px;
	line-height: 25px;
	font-size: 14px;
}

.weixin-menu-detail .menu-tips.cursor {
	color: #459ae9;
	cursor: pointer;
}

.weixin-menu-detail .menu-input .menu-tips {
	margin: 0 0 0 10px;
}

.weixin-menu-detail .menu-content {
	padding: 16px 20px;
	margin-top: 30px;
	border: 1px solid #e7e7eb;
	background-color: #fff;
}


.weixin-menu-detail .menu-content .menu-label {
	text-align: left;
	width: 100px;
}

.weixin-menu-detail .menu-content .menu-input-text {
	border: 1px solid #e7e7eb;
}

.weixin-menu-detail .menu-content .menu-tips {
	padding-bottom: 10px;
}

.weixin-menu-detail .menu-msg-content {
	padding: 0;
	border: 1px solid #e7e7eb;
	background-color: #fff;
}

.weixin-menu-detail .menu-msg-content .menu-msg-head {
	overflow: hidden;
	border-bottom: 1px solid #e7e7eb;
	line-height: 38px;
	height: 38px;
	padding: 0 20px;
}

.weixin-menu-detail .menu-msg-content .menu-msg-panel {
	padding: 30px 50px;
}

.weixin-menu-detail .menu-msg-content .menu-msg-select {
	padding: 40px 20px;
	border: 2px dotted #d9dadc;
	text-align: center;
}

.weixin-menu-detail .menu-msg-content .menu-msg-select:hover {
	border-color: #b3b3b3;
}

.weixin-menu-detail .menu-msg-content strong {
	display: block;
	padding-top: 3px;
	font-weight: 400;
	font-style: normal;
}

.weixin-menu-detail .menu-msg-content .menu-msg-title {
	float: left;
	width: 310px;
	height: 30px;
	line-height: 30px;
	overflow: hidden;
	text-overflow: ellipsis;
	white-space: nowrap;
}

.icon36_common {
	width: 36px;
	height: 36px;
	vertical-align: middle;
	display: inline-block;
}

.icon_msg_sender {
	margin-right: 3px;
	margin-top: -2px;
	width: 20px;
	height: 20px;
	vertical-align: middle;
	display: inline-block;
	/* background: url(../images/msg_tab_z25df2d.png) 0 -270px no-repeat; */
}

.weixin-btn-group {
	text-align: center;
	width: 100%;
	margin: 30px 0px;
	overflow: hidden;
}

.weixin-btn-group .btn {
	width: 100px;
	border-radius: 0px;
}

#material-list {
	padding: 20px;
	overflow-y: scroll;
	height: 558px;
}

#news-list {
	padding: 20px;
	overflow-y: scroll;
	height: 558px;
}

#material-list table {
	width: 100%;
}
</style>
