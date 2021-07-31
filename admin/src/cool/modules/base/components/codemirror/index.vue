<template>
	<div ref="editorRef" class="cl-codemirror">
		<textarea id="editor" class="cl-code" :height="height" :width="width"></textarea>
	</div>
</template>

<script lang="ts">
import { defineComponent, nextTick, onMounted, ref, watch } from "vue";
import CodeMirror from "codemirror";

import "codemirror/theme/idea.css";
import "codemirror/lib/codemirror.css";
import "codemirror/addon/hint/show-hint.css";
import "codemirror/addon/hint/javascript-hint";
import "codemirror/mode/javascript/javascript";
import "codemirror/mode/htmlmixed/htmlmixed";

export default defineComponent({
	name: "cl-codemirror",

	props: {
		modelValue: null,
		height: String,
		width: String,
		mode: String,
		options: Object
	},

	emits: ["update:modelValue", "load"],

	setup(props, { emit }) {
		const editorRef = ref<any>(null);

		let editor: any = null;

		let mode: String | undefined = props.mode
		if (mode == "") {
			mode = "htmlmixed"
		}

		// 获取内容
		function getValue() {
			return editor ? editor.getValue() : "";
		}

		// 设置内容
		function setValue(val?: string) {
			if (editor) {
				editor.setValue(val || getValue());
			}
		}

		// 监听内容变化
		watch(
			() => props.modelValue,
			(val: string) => {
				if (editor) {
					if (val != getValue().replace(/\s/g, "")) {
						setValue(val);
					}
				}
			}
		);

		onMounted(function () {
			nextTick(() => {
				// 实例化
				editor = CodeMirror.fromTextArea(editorRef.value.querySelector("#editor"), {
					mode: mode,
					theme: "idea",
					styleActiveLine: true,
					lineNumbers: true,
					lineWrapping: true,
					indentUnit: 4,
					...props.options
				});

				// 输入监听
				editor.on("change", (e: any) => {
					emit("update:modelValue", e.getValue().replace(/\s/g, ""));
				});

				// 设置内容
				setValue(props.modelValue);

				// 加载回调
				emit("load", editor);

				// 设置编辑框大小
				editor.setSize(props.width || "auto", props.height || "auto");

				// shift + alt + f 格式化
				editor.display.wrapper.onkeydown = (e: any) => {
					const keyCode = e.keyCode || e.which || e.charCode;
					const altKey = e.altKey || e.metaKey;
					const shiftKey = e.shiftKey || e.metaKey;

					if (altKey && shiftKey && keyCode == 70) {
						setValue();
					}
				};
			});
		});

		return {
			editorRef
		};
	}
});
</script>

<style lang="scss" scoped>
.CodeMirror {
	border-radius: 3px;
	border: 1px solid #dcdfe6;
	box-sizing: border-box;
	font-family: monospace;
}

</style>
