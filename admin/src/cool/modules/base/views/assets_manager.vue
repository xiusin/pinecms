<template>
	<div class="system-user">
		<div class="pane">
			<div class="dir">
				<div class="container">
				</div>
			</div>

			<div class="editor">
				<div class="container">
					<component is="cl-codemirror" mode="htmlmixed" height="500px" />
				</div>
			</div>
		</div>
	</div>
</template>

<script lang="ts">
import { defineComponent, inject, reactive, ref } from "vue";
import { useRefs } from "/@/core";
import { QueryList, Table, Upsert } from "cl-admin-crud-vue3/types";
import { ElMessage } from "element-plus";

export default defineComponent({
	name: "sys-dict",

	setup() {
		const service = inject<any>("service");

		const { refs, setRefs } = useRefs();

		return {
			service,
			refs,
			setRefs,
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
