<template>
	<div class="system-user" v-loading="loading">
		<div class="pane">
			<div class="dir">
				<el-row
					style="
						padding: 15px 8px;
						text-align: center;
						width: 100%;
						border-bottom: 1px solid #eeeeee;
					"
				>
					<el-button size="mini" type="waring" @click="addScript"> 新增 </el-button>
					<el-button size="mini" type="success" @click="save"> 保存 </el-button>
				</el-row>
				<div class="container" style="padding: 8px">
					<el-tree
						:data="data"
						:props="{ children: 'children', label: 'label' }"
						@node-click="handleNodeClick"
					/>
				</div>
			</div>

			<div class="editor">
				<div class="container">
					<component
						is="cl-codemirror"
						class="htmleditor"
						style="font-family: 'monospace'"
						:modelValue="modelValue"
						mode="go"
						height="800px"
					/>
				</div>
			</div>
		</div>
	</div>
</template>

<script lang="ts">
import { defineComponent, inject, ref } from "vue";
import { useRefs } from "/@/core";
import { ElMessage, ElMessageBox } from "element-plus";

export default defineComponent({
	name: "script",

	setup() {
		const loading = ref(false);
		const service = inject<any>("service");
		const { refs, setRefs } = useRefs();
		const data = ref([]);
		const modelValue = ref("");
		const formRef = ref();
		const currentPath = ref("");

		function reload() {
			service.task.info.scriptList({}).then((list: never[]) => {
				data.value = list;
			});
		}
		reload();
		function handleNodeClick(data: any) {
			if (data.full_path == "") {
				ElMessage.error("参数错误");
				return;
			}
			loading.value = true;
			service.task.info.scriptInfo({ path: data.full_path }).then((d: string) => {
				loading.value = false;
				currentPath.value = data.full_path;
				modelValue.value = d;
			});
		}

		function addScript() {
			ElMessageBox.prompt("", "新脚本名称", {
				confirmButtonText: "确定",
				cancelButtonText: "取消",
				inputPattern: /\w+\.gsh/,
				inputErrorMessage: "不可用的脚本名称"
			})
				.then(({ value }) => {
					loading.value = true;
					service.task.info
						.saveInfo({
							path: value,
							content: "",
							edit: false
						})
						.then((data: string) => {
							reload();
							loading.value = false;
							currentPath.value = data;
							modelValue.value = "";
							ElMessage.success("新增脚本成功");
						});
				})
				.catch(() => {
					loading.value = false;
				});
		}

		function save() {
			if (currentPath.value == "") {
				ElMessage.error("请先选择编辑一个脚本");
				return;
			}
			loading.value = true;
			service.task.info
				.saveInfo({
					path: currentPath.value,
					content: modelValue.value,
					edit: true
				})
				.then(() => {
					reload();
					loading.value = false;
					ElMessage.success("保存脚本成功");
				});
		}

		return {
			service,
			data,
			refs,
			setRefs,
			addScript,
			save,
			modelValue,
			formRef,
			loading,
			handleNodeClick
		};
	}
});
</script>

<style lang="scss" scoped>
.system-user {
	.pane {
		display: flex;
		height: 100%;
		width: 100%;
		position: relative;
	}

	.dir {
		height: 100%;
		width: 350px;
		max-width: calc(100% - 50px);
		background-color: #fff;
		transition: width 0.3s;
		margin-right: 10px;
		flex-shrink: 0;

		&._collapse {
			margin-right: 0;
			width: 0;
		}
	}

	.editor {
		width: calc(100% - 360px);
		flex: 1;
		background-color: #f8f9fa;

		.header {
			display: flex;
			align-items: center;
			justify-content: center;
			height: 40px;
			position: relative;
			background-color: #fff;

			span {
				font-size: 14px;
				white-space: nowrap;
				overflow: hidden;
			}

			.icon {
				position: absolute;
				left: 0;
				top: 0;
				font-size: 18px;
				cursor: pointer;
				background-color: #fff;
				height: 40px;
				width: 80px;
				line-height: 40px;
				padding-left: 10px;
			}
		}
	}

	.dept,
	.user {
		overflow: hidden;

		.container {
			height: calc(100% - 40px);
		}
	}

	@media only screen and (max-width: 768px) {
		.dept {
			width: calc(100% - 100px);
		}
	}
}
</style>
