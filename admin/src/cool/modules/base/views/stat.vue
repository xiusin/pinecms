<template>
	<div>
		<el-row :gutter="15" class="system_state">
			<el-col :span="12">
				<el-card v-if="state.os" class="card_item">
					<div slot="header" class="clearfix">
						<span>运行时</span>
						<el-tag type="success" size="mini" style="float: right; border-radius: 2px"
							>运行时长:
							{{ state.running_time + "分钟" }}
						</el-tag>
					</div>
					<div style="padding-top: 8px">
						<table class="table">
							<tr>
								<td>系统</td>
								<td>{{ state.os.goos }}</td>
								<td>cpu核心数</td>
								<td>{{ state.os.numCpu }}</td>
							</tr>
							<tr>
								<td>编译器</td>
								<td>{{ state.os.compiler }}</td>
								<td>go版本</td>
								<td>{{ state.os.goVersion }}</td>
							</tr>
							<tr>
								<td>Pine</td>
								<td>{{ state.pine_version }}</td>
								<td>PineCMS</td>
								<td>{{ state.pine_cms_version }}</td>
							</tr>
							<tr>
								<td>Xorm</td>
								<td>{{ state.xorm_version }}</td>
								<td>Mysql</td>
								<td>{{ state.mysql_version }}</td>
							</tr>
							<tr>
								<td>协程数</td>
								<td>{{ state.os.numGoroutine }}</td>
								<td>启动时间</td>
								<td>{{ state.start_time }}</td>
							</tr>
						</table>
					</div>
				</el-card>
			</el-col>
			<el-col :span="12">
				<el-card v-if="state.cpu" class="card_item">
					<div slot="header" class="clearfix">
						<span>核心使用率</span>
					</div>
					<div>
						<el-row :gutter="10">
							<el-col :span="24">
								<v-chart
									:option="chartPreCpuOption"
									autoresize
									:initOptions="{
										height: 350
									}"
								/>
							</el-col>
						</el-row>
					</div>
				</el-card>
			</el-col>
		</el-row>
		<el-row :gutter="15" class="system_state">
			<el-col :span="6">
				<el-card v-if="state.disk" shadow="hover">
					<div slot="header">
						<span>硬盘</span>
						<el-tag type="success" size="mini" style="float: right; border-radius: 2px"
							>总大小:{{ state.disk.totalGb + "G" }} 已用:{{
								state.disk.usedGb + "G"
							}}
						</el-tag>
					</div>
					<div>
						<el-row :gutter="10">
							<el-col :span="24">
								<v-chart
									:option="chartDiskOption"
									autoresize
									:initOptions="{
										height: 350
									}"
								/>
							</el-col>
						</el-row>
					</div>
				</el-card>
			</el-col>
			<el-col :span="6">
				<el-card v-if="state.ram" shadow="hover">
					<div slot="header" class="clearfix">
						<span>内存</span>
						<el-tag type="success" size="mini" style="float: right; border-radius: 2px"
							>总内存:{{ state.ram.totalMb + "M" }} 已用:{{ state.ram.usedMb + "M" }}
						</el-tag>
					</div>
					<div>
						<el-row :gutter="10">
							<el-col :span="24">
								<v-chart
									:option="chartRamOption"
									autoresize
									:initOptions="{
										height: 350
									}"
								/>
							</el-col>
						</el-row>
					</div>
				</el-card>
			</el-col>
			<el-col :span="6">
				<el-card v-if="state.cpu" shadow="hover">
					<div slot="header" class="clearfix">
						<span>CPU</span>
						<el-tag type="success" size="mini" style="float: right; border-radius: 2px">
							使用率:{{ parseFloat(state.cpu.cpu_percent).toFixed(2) + "%" }}
						</el-tag>
					</div>
					<div>
						<el-row :gutter="10">
							<el-col :span="24">
								<v-chart
									:option="chartCpuOption"
									autoresize
									:initOptions="{
										height: 350
									}"
								/>
							</el-col>
						</el-row>
					</div>
				</el-card>
			</el-col>
			<el-col :span="6">
				<el-card v-if="state.nets" shadow="hover">
					<div slot="header" class="clearfix">
						<span>网络IO</span>
						<el-tag type="success" size="mini" style="float: right; border-radius: 2px">
							上传:0 下载:0
						</el-tag>
					</div>
					<div>
						<el-row :gutter="10">
							<el-col :span="24">
								<v-chart
									:option="netIoOption"
									autoresize
									:initOptions="{
										height: 350
									}"
								/>
							</el-col>
						</el-row>
					</div>
				</el-card>
			</el-col>
		</el-row>
	</div>
</template>

<script lang="ts">
import { defineComponent, inject, onBeforeUnmount, reactive, ref } from "vue";
import Table from "../../demo/components/crud/table.vue";

export default defineComponent({
	name: "sys-stat",
	components: { Table },
	setup() {
		const service = inject<any>("service");
		let timer: NodeJS.Timeout;
		let state = ref({});

		const CpuCores = ref([]);

		const CpuSeries = ref([]);

		const CpuTimePos = ref([]);

		let isInit = false;

		function reload() {
			service.system.stat.data().then((data: any) => {
				state.value = data;
				if (!isInit) {
					isInit = true;
					CpuCores.value = data.cpu.cpus.map((item, idx) => "CPU" + idx);
					let series = [];

					for (const cpuCoresKey in CpuCores.value) {
						series.push({
							name: CpuCores.value[cpuCoresKey],
							type: "line",
							stack: "总量",
							data: [data.cpu.cpus[cpuCoresKey]]
						});
					}
					CpuSeries.value = series;
				} else {
					data.cpu.cpus.map((item, idx) => {
						CpuSeries.value[idx].data.push(item);
					});
				}

				try {
					chartDiskOption.series[0].data[0].value = state.value.disk.usedPercent;
					chartRamOption.series[0].data[0].value = state.value.ram.usedPercent;
					chartCpuOption.series[0].data[0].value = state.value.cpu.cpu_percent[0].toFixed(
						2
					);
				} catch (e) {}
			});
		}

		reload();
		timer = setInterval(() => {
			reload();
		}, 1000 * 10);

		onBeforeUnmount(() => {
			clearInterval(timer);
		});

		const chartDiskOption = reactive({
			tooltip: {
				formatter: "{a} <br/>{b} : {c}%"
			},
			series: [
				{
					name: "Pressure",
					type: "gauge",
					detail: {
						formatter: "{value}"
					},
					data: [
						{
							value: 0,
							name: "使用率"
						}
					]
				}
			]
		});

		const chartRamOption = reactive({
			tooltip: {
				formatter: "{a} <br/>{b} : {c}%"
			},
			series: [
				{
					name: "Pressure",
					type: "gauge",
					detail: {
						formatter: "{value}"
					},
					data: [
						{
							value: 0,
							name: "使用率"
						}
					]
				}
			]
		});

		const chartCpuOption = reactive({
			tooltip: {
				formatter: "{a} <br/>{b} : {c}%"
			},
			series: [
				{
					name: "Pressure",
					type: "gauge",
					detail: {
						formatter: "{value}"
					},
					data: [
						{
							value: 0,
							name: "使用率"
						}
					]
				}
			]
		});

		const colors = [
			"#8EB6D3",
			"#E9C1EB",
			"#97D9E1",
			"#D9AFD9",
			"#D5A8D1",
			"#C7EAFD",
			"#BBD5FF",
			"#E0C3EF",
			"#00DDFF",
			"#37A2FF",
			"#FF0087",
			"#FFBF00",
			"#ED6EA0",
			"#A6CE11"
		];

		const chartPreCpuOption = reactive({
			color: colors,
			tooltip: {
				trigger: "axis",
				axisPointer: {
					type: "cross",
					label: {
						backgroundColor: "#6a7985"
					}
				}
			},
			legend: {
				data: CpuCores //['Line 1', 'Line 2', 'Line 3', 'Line 4', 'Line 5']
			},
			grid: {
				left: "3%",
				right: "4%",
				bottom: "3%",
				containLabel: true
			},
			xAxis: [
				{
					type: "category",
					boundaryGap: false,
					data: CpuTimePos //['周一', '周二', '周三', '周四', '周五', '周六', '周日']
				}
			],
			yAxis: [
				{
					type: "value"
				}
			],
			series: CpuSeries
		});

		const netIoOption = reactive({
			title: {
				text: "Step Line"
			},
			tooltip: {
				trigger: "axis"
			},
			legend: {
				data: ["Step Start", "Step Middle", "Step End"]
			},
			grid: {
				left: "3%",
				right: "4%",
				bottom: "3%",
				containLabel: true
			},
			toolbox: {
				feature: {
					saveAsImage: {}
				}
			},
			xAxis: {
				type: "category",
				data: ["Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"]
			},
			yAxis: {
				type: "value"
			},
			series: [
				{
					name: "Step Start",
					type: "line",
					step: "start",
					data: [120, 132, 101, 134, 90, 230, 210]
				},
				{
					name: "Step Middle",
					type: "line",
					step: "middle",
					data: [220, 282, 201, 234, 290, 430, 410]
				},
				{
					name: "Step End",
					type: "line",
					step: "end",
					data: [450, 432, 401, 454, 590, 530, 510]
				}
			]
		});

		return {
			netIoOption,
			chartDiskOption,
			chartRamOption,
			chartCpuOption,
			chartPreCpuOption,
			timer,
			state,
			colors
		};
	}
});
</script>
<style scoped>
.system_state {
	padding: 10px;
}

.card_item {
	height: 400px;
}

.card_item .el-row {
	padding: 3px 0 3px 0;
}

.table {
	width: 100%;
	color: #666;
	border-collapse: collapse;
	background-color: #fff;
}

.table td:nth-child(odd) {
	width: 20%;
	text-align: right;
	background-color: #f7f7f7;
}

.table td {
	position: relative;
	padding: 9px 15px;
	overflow: hidden;
	font-size: 14px;
	line-height: 20px;
	text-overflow: ellipsis;
	white-space: nowrap;
	border: 1px solid #e6e6e6;
}
</style>
