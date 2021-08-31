<template>
	<div>
		<div class="menu-input-group">
			<div class="menu-name">
				<el-tag>{{ button.name }}</el-tag>
			</div>
			<div class="menu-del" @click="$emit('delMenu')">
				<el-button size="mini">删除菜单</el-button>
			</div>
		</div>
		<el-divider />
		<div class="menu-input-group">
			<div class="menu-label">菜单名称</div>
			<div class="menu-input">
				<el-input
					v-model="button.name"
					placeholder="请输入菜单名称"
					size="medium"
					:maxlength="selectedMenuLevel === 1 ? 5 : 8"
					show-word-limit
					@input="checkMenuName(button.name)"
				>
				</el-input>
			</div>
		</div>
		<div v-show="!button.subButtons || button.subButtons.length === 0">
			<div class="menu-input-group">
				<div class="menu-label">菜单内容</div>
				<div class="menu-input">
					<el-select v-model="button.type" name="type" size="small" style="width: 380px">
						<el-option key="view" value="view">跳转网页(view)</el-option>
						<el-option key="media_id" value="media_id">发送消息(media_id)</el-option>
						<el-option key="view_limited" value="view_limited" disabled>跳转公众号图文消息链接(view_limited)</el-option>
						<el-option key="miniprogram" value="miniprogram">打开指定小程序(miniprogram)</el-option>
						<el-option key="click" value="click">自定义点击事件(click)</el-option>
						<el-option key="scancode_push" value="scancode_push">扫码上传消息(scancode_push)</el-option>
						<el-option key="scancode_waitmsg" value="scancode_waitmsg">扫码提示下发(scancode_waitmsg)</el-option>
						<el-option key="pic_sysphoto" value="pic_sysphoto">系统相机拍照(pic_sysphoto)</el-option>
						<el-option key="pic_photo_or_album" value="pic_photo_or_album">弹出拍照或者相册(pic_photo_or_album)</el-option>
						<el-option key="pic_weixin" value="pic_weixin">弹出微信相册(pic_weixin)</el-option>
						<el-option key="location_select" value="location_select">弹出地理位置选择器(location_select)</el-option>
					</el-select>
				</div>
			</div>
			<div class="menu-content" v-if="button.type === 'view'">
				<div class="menu-input-group">
					<p class="menu-tips">订阅者点击该子菜单会跳到以下链接</p>
					<div class="menu-label">页面地址</div>
					<div class="menu-input">
						<input
							type="text"
							placeholder=""
							class="menu-input-text"
							v-model="button.url"
						/>
					</div>
				</div>
			</div>
			<div class="menu-content" v-else-if="button.type === 'media_id'">
				<div class="menu-input-group">
					<p class="menu-tips">订阅者点击该菜单会收到以下图文消息</p>
					<div class="menu-label">media_id</div>
					<div class="menu-input">
						<el-input
							type="text"
							placeholder="图文消息media_id"
							v-model="button.mediaId"
						/>
					</div>
				</div>
			</div>
			<div class="menu-content" v-else-if="button.type === 'miniprogram'">
				<div class="menu-input-group">
					<p class="menu-tips">订阅者点击该子菜单会跳到以下小程序</p>
					<div class="menu-label">小程序appId</div>
					<div class="menu-input">
						<el-input
							size="small"
							type="text"
							placeholder="小程序的appId（仅认证公众号可配置）"
							v-model="button.appId"
						/>
					</div>
				</div>
				<div class="menu-input-group">
					<div class="menu-label">小程序路径</div>
					<div class="menu-input">
						<el-input
							size="small"
							type="text"
							placeholder="小程序的页面路径 pages/index/index"
							v-model="button.pagePath"
						/>
					</div>
				</div>
				<div class="menu-input-group">
					<div class="menu-label">备用网页</div>
					<div class="menu-input">
						<el-input
							size="small"
							type="text"
							placeholder="旧版微信客户端无法支持小程序，用户点击菜单时将会打开备用网页"
							v-model="button.url"
						/>
					</div>
				</div>
			</div>
			<div class="menu-content" v-else>
				<div class="menu-input-group">
					<p class="menu-tips">用于消息接口推送，不超过128字节</p>
					<div class="menu-label">菜单KEY值</div>
					<div class="menu-input">
						<input
							type="text"
							placeholder=""
							class="menu-input-text"
							v-model="button.key"
						/>
					</div>
				</div>
			</div>
		</div>
	</div>
</template>
<script>
export default {
	props: {
		selectedMenuLevel: {
			type: Number,
			default: 1
		},
		button: {
			type: Object,
			required: true
		}
	},
	emits: ["delMenu"],
	data() {
		return {
			menuNameBounds: false
		};
	},
	methods: {
		checkMenuName: function (val) {
			if (this.selectedMenuLevel === 1 && this.getMenuNameLen(val) <= 10) {
				this.menuNameBounds = false;
			} else this.menuNameBounds = !(this.selectedMenuLevel === 2 && this.getMenuNameLen(val) <= 16);
		},
		getMenuNameLen: function (val) {
			var len = 0;
			for (var i = 0; i < val.length; i++) {
				var a = val.charAt(i);
				a.match(/[^\x00-\xff]/gi) != null ? (len += 2) : (len += 1);
			}
			return len;
		}
	}
};
</script>
