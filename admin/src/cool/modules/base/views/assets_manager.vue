<template>
	<div class="system-user">
		<div class="pane">
			<div class="dir">
				<el-row style="padding: 5px; text-align: center; width: 100%">
					<el-button size="mini" type="waring" @click="addDir"> 新增目录 </el-button>
					<el-button size="mini" type="waring" @click="addAsset"> 新增 </el-button>
					<el-button size="mini" type="success" @click="save"> 保存 </el-button>
				</el-row>
				<div class="container">
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
						:modelValue="modelValue"
						mode="htmlmixed"
						height="700px"
					/>
				</div>
			</div>
		</div>
	</div>
	<cl-form ref="formRef" />
</template>

<script lang="ts">
import { defineComponent, inject, ref } from "vue";
import { useRefs } from "/@/cool";

export default defineComponent({
	name: "sys-assets-manager",

	setup() {
		const service = inject<any>("service");
		const { refs, setRefs } = useRefs();
		const data = ref([]);
		const modelValue = ref("");
		const formRef = ref();

		function reload() {
			service.system.assets.list().then((list: never[]) => {
				data.value = list;
			});
		}

		reload();

		function handleNodeClick(data: any) {
			service.system.assets.info({ path: data.full_path }).then((data: string) => {
				modelValue.value = data;
			});
		}

		function addDir() {
			formRef.value?.open({
				width: "600px",
				props: {
					labelWidth: "140px"
				},
				items: [
					{
						label: "目录名称",
						prop: "name",
						component: {
							name: "el-input"
						},
						rules: {
							required: true,
							message: "目录名称不能为空"
						}
					},
					{
						label: "上级目录",
						prop: "parent_name",
						component: {
							name: "el-select",
							options: kvGroups
						},
						rules: {
							required: true,
							message: "昵称不能为空"
						}
					}
				],
				on: {
					submit(data: any, { done }: any) {
						console.log(data);
						reload();
						done();
					}
				}
			});
		}

		function addAsset() {}

		function save() {}

		return {
			service,
			data,
			refs,
			setRefs,
			addDir,
			addAsset,
			save,
			modelValue,
			formRef,
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
		width: 250px;
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
		width: calc(100% - 260px);
		flex: 1;
		background-color: #fff;

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
