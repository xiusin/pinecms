<template>
	<div class="cl-filter-group">
		<div class="cl-filter-group__items"><slot></slot></div>
		<div class="cl-filter-group__op">
			<el-button type="primary" size="mini" @click="search()"> 搜索 </el-button>
			<el-button size="mini" @click="reset()"> 重置 </el-button>
		</div>
	</div>
</template>

<script>
import { cloneDeep } from "/@/cool/utils";

export default {
	name: "cl-filter-group",

	componentName: "ClFilterGroup",

	inject: ["crud"],

	props: {
		// 表单值
		modelValue: {
			type: Object,
			default: () => {
				return {};
			}
		},

		// 搜索时钩子, data, { next }
		onSearch: Function
	},

	data() {
		return {
			oldForm: cloneDeep(this.modelValue),
			form: {},
			loading: false
		};
	},

	watch: {
		modelValue: {
			immediate: true,
			deep: true,
			handler(val) {
				this.form = val;
				this.$emit("change", this.form);
			}
		}
	},

	methods: {
		search() {
			const next = (params) => {
				this.loading = true;

				this.crud.refresh({
					param: this.form,
					page: 1,
					...params
				});
				// .done(() => {
				// 	this.loading = false;
				// });
			};

			if (this.onSearch) {
				this.onSearch(this.form, { next });
			} else {
				next();
			}
		},

		reset() {
			for (let i in this.form) {
				this.form[i] = this.oldForm[i] === undefined ? undefined : this.oldForm[i];
			}

			this.search();
			this.$emit("reset");
		}
	}
};
</script>

<style scoped>
.cl-filter-group__items {
	margin-right: 10px;
}
</style>
