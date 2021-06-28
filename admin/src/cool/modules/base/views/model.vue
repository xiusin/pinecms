<template>
	<cl-crud @load="onLoad">
		<el-row>
			<cl-table v-bind="table"/>
		</el-row>
		<el-row type="flex">
			<el-button @click="open()">添加行</el-button>
			<cl-flex1/>
			<cl-pagination/>
		</el-row>
	</cl-crud>

	<el-drawer v-model="drawer" title="模型编辑" size="100%"></el-drawer>

	<cl-form :ref="formRef" inner>
		<!-- 动态增减表单验证 -->
		<template #slot-validate="{ scope }">
			<el-form-item
				v-for="(item, index) in scope.vads"
				:key="index"
				:prop="'vads.' + index + '.val'"
				:rules="{ required: true, message: '请输入' }"
			>
				<el-input v-model="item.val" />
			</el-form-item>
		</template>
	</cl-form>
</template>

<script lang="ts">
import {defineComponent, h, inject, reactive, ref, resolveComponent} from "vue";
import {useRefs} from "/@/core";
import {CrudLoad, FormItem, FormRef, Table, Upsert} from "cl-admin-crud-vue3/types";

export default defineComponent({
	name: "sys-model",

	setup() {
		const service = inject<any>("service");
		const {refs, setRefs} = useRefs();

		const formRef = ref<FormRef>();

		function renderDivider(label: string) {
			const el: any = resolveComponent("el-divider");

			return h(
				el,
				{
					"content-position": "left"
				},
				{
					default: () => label
				}
			);
		}

		const items: FormItem[] = [
			{
				props: {
					labelWidth: "0px"
				},
				component: () => {
					return renderDivider("测试组件渲染");
				}
			},
			{
				label: ".vue 组件",
				value: 10,
				prop: "vue"
			},
			{
				label: "tsx",
				prop: "tsx",
				value: "Hello!"
			},
			{
				props: {
					labelWidth: "0px"
				},
				component: () => {
					return renderDivider("测试内嵌CRUD");
				}
			},
			{
				props: {
					labelWidth: "0px"
				},
				component: {
					name: "slot-crud"
				}
			},
			{
				props: {
					labelWidth: "0px"
				},
				component: () => {
					return renderDivider("测试验证规则");
				}
			},
			{
				prop: "vads",
				value: [],
				label: "动态增减表单验证",
				component: {
					name: "slot-validate"
				}
			},
			{
				props: {
					labelWidth: "0px"
				},
				component: () => {
					return renderDivider("测试显隐");
				}
			},
			{
				label: "奇术",
				prop: "qs",
				value: [],
				component: {
					name: "el-select",
					props: {
						placeholder: "请选择奇术",
						multiple: true
					},
					options: [
						{
							label: "烟水还魂",
							value: 1
						},
						{
							label: "雨恨云愁",
							value: 2
						}
					]
				}
			},
			{
				label: "技能",
				prop: "jn",
				value: 1,
				component: {
					name: "el-select",
					props: {
						placeholder: "请选择技能"
					},
					options: [
						{
							label: "飞羽箭",
							value: 1
						},
						{
							label: "落星式",
							value: 2
						}
					]
				}
			},
			{
				label: "五行",
				prop: "wx",
				value: 0,
				hidden: ({ scope }: any) => {
					return scope.jn == 1;
				},
				component: {
					name: "el-radio-group",
					options: [
						{
							label: "水",
							value: 0
						},
						{
							label: "火",
							value: 1
						},
						{
							label: "雷",
							value: 2
						},
						{
							label: "风",
							value: 3
						},
						{
							label: "土",
							value: 4
						}
					]
				}
			},
			{
				label: "雨润",
				prop: "s1",
				hidden: ({ scope }: any) => {
					return scope.wx != 0;
				},
				component: ({ h }: any) => {
					return h("p", "以甘甜雨露的滋润使人精力充沛");
				}
			},
			{
				label: "风雪冰天",
				prop: "s2",
				hidden: ({ scope }: any) => {
					return scope.wx != 0;
				},
				component: ({ h }: any) => {
					return h("p", "召唤漫天风雪，对敌方造成巨大的杀伤力");
				}
			},
			{
				label: "三昧真火",
				prop: "h",
				hidden: ({ scope }: any) => {
					return scope.wx != 1;
				},
				component: ({ h }: any) => {
					return h("p", "召唤三昧真火焚烧敌方的仙术");
				}
			},
			{
				label: "惊雷闪",
				prop: "l",
				hidden: ({ scope }: any) => {
					return scope.wx != 2;
				},
				component: ({ h }: any) => {
					return h("p", "召唤惊雷无数，对敌方全体进行攻击，是十分强力的仙术");
				}
			},
			{
				label: "如沐春风",
				prop: "f",
				hidden: ({ scope }: any) => {
					return scope.wx != 3;
				},
				component: ({ h }: any) => {
					return h("p", "温暖柔和的复苏春风，使人回复活力");
				}
			},
			{
				label: "艮山壁障",
				prop: "t",
				hidden: ({ scope }: any) => {
					return scope.wx != 4;
				},
				component: ({ h }: any) => {
					return h("p", "以艮山之灵形成一道壁障，受此壁障守护者刀枪不入");
				}
			}
		];

		function open() {
			drawer.value = true
			console.log(formRef.value?.create({
				width: "1000px",
				props: {
					labelWidth: "140px"
				},
				items,
				on: {
					submit(data, { done }) {
						done();
					}
				}
			}).open());
		}


		let drawer = ref<boolean>(false)

		// 表格配置
		const table = reactive<Table>({
			columns: [
				{
					prop: "table_name",
					label: "模型名称",
					width: 150,
					align: "left"
				},
				{
					prop: "table",
					label: "表名",
					width: 150,
					align: "left"
				},
				{
					label: "类型",
					prop: "model_type",
					width: 50,
					align:"left"
				},
				{
					label: "列表模板",
					prop: "fe_tpl_list",
					width: 150,
					align: "left"
				},
				{
					label: "详情模板",
					prop: "fe_tpl_detail",
					width: 150,
					align:"left"
				},
				{
					label: "备注",
					prop: "remark",
					align: "left",
				},
				{
					label: "操作",
					type: "op",
					width: 100,
					buttons: ["edit", "delete"]
				}
			]
		});

		// 新增编辑配置
		const upsert = reactive<Upsert>({
			width: "1000px",
			items: [
				{
					prop: "table_name",
					label: "模型名称",
					span: 24,
					component: {
						name: "el-input",
						props: {
							placeholder: "名称"
						}
					},
				},
				{
					prop: "table",
					label: "模型名称",
					span: 24,
					component: {
						name: "el-input",
						props: {
							placeholder: "名称"
						}
					},
				},
				{
					prop: "content",
					label: "内容",
					component: {
						name: "slot-content"
					}
				}
			]
		});

		// crud 加载
		function onLoad({ctx, app}: CrudLoad) {
			ctx.service(service.system.model).done();
			app.refresh();
		}

		return {
			refs,
			table,
			upsert,
			setRefs,
			onLoad,
			drawer,
			open,
			formRef
		};
	}
});
</script>

<style lang="scss" scoped>
.change-btn {
	display: flex;
	position: absolute;
	right: 10px;
	bottom: 10px;
	z-index: 9;
}

.editor {
	transition: all 0.3s;
}
</style>
