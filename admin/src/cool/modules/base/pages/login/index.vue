<template>
	<div class="login-bg">
		<div class="page-login">
			<div class="box">
				<div class="login-wrapper">
					<div class="login-screen">
						<div class="well">
							<div class="login-head">
								<img src="/login-head.png" style="width: 100%" />
							</div>
							<div class="login-form">
								<img id="profile-img" class="profile-img-card" src="/avatar.png" />
								<p id="profile-name" class="profile-name-card"></p>
								<el-form class="form" size="medium" :disabled="saving">
									<el-form-item>
										<el-input
											v-model="form.username"
											placeholder="请输入用户名"
											maxlength="32"
											size="small"
											auto-complete="off"
										>
											<template #prepend>
												<i class="el-icon-s-custom"></i>
											</template>
										</el-input>
									</el-form-item>

									<el-form-item>
										<el-input
											v-model="form.password"
											type="password"
											placeholder="请输入密码"
											maxlength="16"
											size="small"
											auto-complete="off"
										>
											<template #prepend>
												<i class="el-icon-key"></i>
											</template>
										</el-input>
									</el-form-item>

									<el-form-item label="验证码" class="captcha">
										<el-input
											v-model="form.verifyCode"
											placeholder="请输入图片验证码"
											maxlength="5"
											auto-complete="off"
											@keyup.enter="toLogin"
										/>

										<captcha
											:ref="setRefs('captcha')"
											v-model="form.captchaId"
											class="value"
											@change="
												() => {
													form.verifyCode = '';
												}
											"
										/>
									</el-form-item>
								</el-form>

								<el-button
									style="width: 100%"
									size="small"
									class="submit-btn"
									:loading="saving"
									@click="toLogin"
									>登录
								</el-button>
							</div>
						</div>
					</div>
				</div>
			</div>
		</div>
	</div>
</template>

<script lang="ts">
import { defineComponent, reactive, ref } from "vue";
import { ElMessage } from "element-plus";
import { useRouter } from "vue-router";
import { useStore } from "vuex";
import Captcha from "./components/captcha.vue";
import { useRefs } from "/@/core";

export default defineComponent({
	components: {
		Captcha
	},

	setup() {
		const router = useRouter();
		const store = useStore();
		const { refs, setRefs }: any = useRefs();

		const saving = ref<boolean>(false);

		// 登录表单数据
		const form = reactive({
			username: "admin",
			password: "123456",
			captchaId: "",
			verifyCode: ""
		});

		// 登录
		async function toLogin() {
			if (!form.username) {
				return ElMessage.warning("用户名不能为空");
			}

			if (!form.password) {
				return ElMessage.warning("密码不能为空");
			}

			// if (!form.verifyCode) {
			// 	return ElMessage.warning("图片验证码不能为空");
			// }

			saving.value = true;

			try {
				// 登录
				await store.dispatch("userLogin", form);

				// 用户信息
				// await store.dispatch("userInfo");

				// 权限菜单
				const [first] = await store.dispatch("permMenu");

				if (!first) {
					ElMessage.error("该账号没有权限");
				} else {
					router.push("/");
				}
			} catch (err) {
				ElMessage.error(err);
				refs.value.captcha.refresh();
			}

			saving.value = false;
		}

		return {
			refs,
			form,
			saving,
			toLogin,
			setRefs
		};
	}
});
</script>

<style scoped>
form .el-input-group__prepend {
	padding: 0 5px;
}
</style>

<style lang="scss">
body {
	color: #999;
	background-color: #f1f4fd;
	background-size: cover;
}

a {
	color: #444;
}

.login-bg {
	padding-top: 200px;
	background: url("/loginbg.png");
	height: 100%;
}

.login-screen {
	max-width: 430px;
	padding: 0;
	margin: 0 auto;
}

.login-screen .well {
	border-radius: 3px;
	-webkit-box-shadow: 0 0 30px rgba(0, 0, 0, 0.1);
	box-shadow: 0 0 30px rgba(0, 0, 0, 0.1);
	background: rgba(255, 255, 255, 1);
	border: none;
	overflow: hidden;
	padding: 0;
}

@media (max-width: 767px) {
	.login-screen {
		padding: 0 20px;
	}
}

.profile-img-card {
	width: 100px;
	height: 100px;
	display: block;
	-moz-border-radius: 50%;
	-webkit-border-radius: 50%;
	border-radius: 50%;
	margin: -93px auto 30px;
	border: 5px solid #fff;
}

.profile-name-card {
	text-align: center;
}

.login-head {
	background: #899fe1;
}

.login-form {
	padding: 40px 30px;
	position: relative;
	z-index: 99;
}

#login-form {
	margin-top: 20px;
}

#login-form .input-group {
	margin-bottom: 15px;
}

#login-form .form-control {
	font-size: 13px;
}
</style>
