<template>
	<div>
		<cl-crud :ref="setRefs('crud')" @load="onLoad">
			<el-row type="flex">
				<cl-add-btn />
				<cl-refresh-btn />
				<cl-flex1 />
				<cl-search-key />
			</el-row>

			<el-row>
				<cl-table v-bind="table">
					<template #column-copy="{ scope }">
						<el-button v-copy="buildURL(scope.row.appid)" size="mini">
							{{ buildURL(scope.row.appid) }}
						</el-button>
					</template>
					<template #slot-map="{ scope }">
						<el-button size="mini" type="text" @click="showMapData(scope.row.appid)">
							会员概况
						</el-button>
					</template>
					<template #column-important="{ scope }">
						<el-dropdown size="mini" type="success" trigger="click">
							<span class="el-dropdown-link" style="font-size: 12px">
								操作<i class="el-icon-arrow-down el-icon--right"></i>
							</span>
							<template #dropdown>
								<el-dropdown-menu>
									<el-dropdown-item
										icon="el-icon-tip"
										@click="clearQuota(scope.row.appid)"
										>重置调用频次限制</el-dropdown-item
									>
								</el-dropdown-menu>
							</template>
						</el-dropdown>
					</template>
				</cl-table>
			</el-row>

			<el-row type="flex">
				<cl-flex1 />
				<cl-pagination />
			</el-row>

			<cl-upsert v-model="form" v-bind="upsert" />
		</cl-crud>

		<cl-dialog title="会员地域分布" width="80%" :close-on-click-modal="false" v-model="showMap">
			<el-row :gutter="20">
				<el-col :span="16">
					<v-chart
						:option="mapOptions"
						:init-options="{ height: 700, renderer: 'canvas' }"
						autoresize
					/>
				</el-col>
				<el-col :span="8">
					<v-chart :option="userChartsOptions" autoresize />
				</el-col>
			</el-row>
		</cl-dialog>
	</div>
</template>

<script lang="ts">
import { defineComponent, inject, reactive, ref } from "vue";
import { useRefs } from "/@/cool";
import { CrudLoad, Table, Upsert } from "@cool-vue/crud/types";
import { ElMessage, ElMessageBox } from "element-plus";

import chinaMap from "/public/china.json";

import { registerMap } from "echarts/core";

registerMap("china", chinaMap);

export default defineComponent({
	name: "wechat-account",

	setup() {
		function convertData(data) {
			const res = [];
			for (let i = 0; i < data.length; i++) {
				const geoCoord = geoCoordMap[data[i].name];
				if (geoCoord) {
					res.push({
						name: data[i].name,
						value: geoCoord.concat(data[i].value)
					});
				}
			}
			return res;
		}

		const service = inject<any>("service");

		const { refs, setRefs }: any = useRefs();

		const upsert = reactive<Upsert>({
			items: [
				{
					prop: "name",
					label: "公众号名称",
					span: 24,
					component: {
						name: "el-input",
						props: {
							placeholder: "公众号名称"
						}
					},
					rules: {
						required: true,
						message: "公众号名称不能为空"
					}
				},

				{
					prop: "type",
					label: "公众号类型",
					value: 1,
					component: {
						name: "el-radio-group",
						options: [
							{
								label: "订阅号",
								value: 1
							},
							{
								label: "服务号",
								value: 2
							}
						]
					}
				},
				{
					prop: "verified",
					label: "是否认证",
					value: true,
					component: {
						name: "el-radio-group",
						options: [
							{
								label: "是",
								value: true
							},
							{
								label: "否",
								value: false
							}
						]
					}
				},
				{
					prop: "appid",
					label: "appid",
					component: {
						name: "el-input",
						props: {
							placeholder: "请输入appid"
						}
					},
					rules: {
						required: true,
						message: "appid不能为空"
					}
				},
				{
					prop: "secret",
					label: "secret",
					component: {
						name: "el-input",
						props: {
							placeholder: "请输入secret"
						}
					},
					rules: {
						required: true,
						message: "secret不能为空"
					}
				},
				{
					prop: "token",
					label: "token",
					component: {
						name: "el-input",
						props: {
							placeholder: "请输入token"
						}
					},
					rules: {
						required: true,
						message: "token不能为空"
					}
				},
				{
					prop: "aesKey",
					label: "aesKey",
					component: {
						name: "el-input"
					}
				}
			]
		});

		const table = reactive<Table>({
			columns: [
				{
					type: "index",
					label: "#",
					width: 60
				},
				{
					prop: "name",
					label: "公众号名称",
					width: 250,
					align: "left"
				},
				{
					prop: "appid",
					label: "APPID",
					width: 300
				},
				{
					prop: "type",
					label: "类型",
					width: 140,
					dict: [
						{
							label: "公众号",
							value: 1,
							type: "success"
						},
						{
							label: "服务号",
							value: 2,
							type: "warning"
						}
					]
				},
				{
					prop: "verified",
					label: "是否认证",
					width: 140,
					dict: [
						{
							label: "是",
							value: 1,
							type: "success"
						},
						{
							label: "否",
							value: 0,
							type: "danger"
						}
					]
				},
				{
					prop: "copy",
					label: "通知地址"
				},
				{
					prop: "important",
					label: "重要操作",
					width: 130
				},
				{
					type: "op",
					width: 170,
					buttons: ["slot-map", "edit", "delete"]
				}
			]
		});

		const showMap = ref(false);

		const geoCoordMap = {
			海门: [121.15, 31.89],
			鄂尔多斯: [109.781327, 39.608266],
			招远: [120.38, 37.35],
			舟山: [122.207216, 29.985295],
			齐齐哈尔: [123.97, 47.33],
			盐城: [120.13, 33.38],
			赤峰: [118.87, 42.28],
			青岛: [120.33, 36.07],
			乳山: [121.52, 36.89],
			金昌: [102.188043, 38.520089],
			泉州: [118.58, 24.93],
			莱西: [120.53, 36.86],
			日照: [119.46, 35.42],
			胶南: [119.97, 35.88],
			南通: [121.05, 32.08],
			拉萨: [91.11, 29.97],
			云浮: [112.02, 22.93],
			梅州: [116.1, 24.55],
			文登: [122.05, 37.2],
			上海: [121.48, 31.22],
			攀枝花: [101.718637, 26.582347],
			威海: [122.1, 37.5],
			承德: [117.93, 40.97],
			厦门: [118.1, 24.46],
			汕尾: [115.375279, 22.786211],
			潮州: [116.63, 23.68],
			丹东: [124.37, 40.13],
			太仓: [121.1, 31.45],
			曲靖: [103.79, 25.51],
			烟台: [121.39, 37.52],
			福州: [119.3, 26.08],
			瓦房店: [121.979603, 39.627114],
			即墨: [120.45, 36.38],
			抚顺: [123.97, 41.97],
			玉溪: [102.52, 24.35],
			张家口: [114.87, 40.82],
			阳泉: [113.57, 37.85],
			莱州: [119.942327, 37.177017],
			湖州: [120.1, 30.86],
			汕头: [116.69, 23.39],
			昆山: [120.95, 31.39],
			宁波: [121.56, 29.86],
			湛江: [110.359377, 21.270708],
			揭阳: [116.35, 23.55],
			荣成: [122.41, 37.16],
			连云港: [119.16, 34.59],
			葫芦岛: [120.836932, 40.711052],
			常熟: [120.74, 31.64],
			东莞: [113.75, 23.04],
			河源: [114.68, 23.73],
			淮安: [119.15, 33.5],
			泰州: [119.9, 32.49],
			南宁: [108.33, 22.84],
			营口: [122.18, 40.65],
			惠州: [114.4, 23.09],
			江阴: [120.26, 31.91],
			蓬莱: [120.75, 37.8],
			韶关: [113.62, 24.84],
			嘉峪关: [98.289152, 39.77313],
			广州: [113.23, 23.16],
			延安: [109.47, 36.6],
			太原: [112.53, 37.87],
			清远: [113.01, 23.7],
			中山: [113.38, 22.52],
			昆明: [102.73, 25.04],
			寿光: [118.73, 36.86],
			盘锦: [122.070714, 41.119997],
			长治: [113.08, 36.18],
			深圳: [114.07, 22.62],
			珠海: [113.52, 22.3],
			宿迁: [118.3, 33.96],
			咸阳: [108.72, 34.36],
			铜川: [109.11, 35.09],
			平度: [119.97, 36.77],
			佛山: [113.11, 23.05],
			海口: [110.35, 20.02],
			江门: [113.06, 22.61],
			章丘: [117.53, 36.72],
			肇庆: [112.44, 23.05],
			大连: [121.62, 38.92],
			临汾: [111.5, 36.08],
			吴江: [120.63, 31.16],
			石嘴山: [106.39, 39.04],
			沈阳: [123.38, 41.8],
			苏州: [120.62, 31.32],
			茂名: [110.88, 21.68],
			嘉兴: [120.76, 30.77],
			长春: [125.35, 43.88],
			胶州: [120.03336, 36.264622],
			银川: [106.27, 38.47],
			张家港: [120.555821, 31.875428],
			三门峡: [111.19, 34.76],
			锦州: [121.15, 41.13],
			南昌: [115.89, 28.68],
			柳州: [109.4, 24.33],
			三亚: [109.511909, 18.252847],
			自贡: [104.778442, 29.33903],
			吉林: [126.57, 43.87],
			阳江: [111.95, 21.85],
			泸州: [105.39, 28.91],
			西宁: [101.74, 36.56],
			宜宾: [104.56, 29.77],
			呼和浩特: [111.65, 40.82],
			成都: [104.06, 30.67],
			大同: [113.3, 40.12],
			镇江: [119.44, 32.2],
			桂林: [110.28, 25.29],
			张家界: [110.479191, 29.117096],
			宜兴: [119.82, 31.36],
			北海: [109.12, 21.49],
			西安: [108.95, 34.27],
			金坛: [119.56, 31.74],
			东营: [118.49, 37.46],
			牡丹江: [129.58, 44.6],
			遵义: [106.9, 27.7],
			绍兴: [120.58, 30.01],
			扬州: [119.42, 32.39],
			常州: [119.95, 31.79],
			潍坊: [119.1, 36.62],
			重庆: [106.54, 29.59],
			台州: [121.420757, 28.656386],
			南京: [118.78, 32.04],
			滨州: [118.03, 37.36],
			贵阳: [106.71, 26.57],
			无锡: [120.29, 31.59],
			本溪: [123.73, 41.3],
			克拉玛依: [84.77, 45.59],
			渭南: [109.5, 34.52],
			马鞍山: [118.48, 31.56],
			宝鸡: [107.15, 34.38],
			焦作: [113.21, 35.24],
			句容: [119.16, 31.95],
			北京: [116.46, 39.92],
			徐州: [117.2, 34.26],
			衡水: [115.72, 37.72],
			包头: [110, 40.58],
			绵阳: [104.73, 31.48],
			乌鲁木齐: [87.68, 43.77],
			枣庄: [117.57, 34.86],
			杭州: [120.19, 30.26],
			淄博: [118.05, 36.78],
			鞍山: [122.85, 41.12],
			溧阳: [119.48, 31.43],
			库尔勒: [86.06, 41.68],
			安阳: [114.35, 36.1],
			开封: [114.35, 34.79],
			济南: [117, 36.65],
			德阳: [104.37, 31.13],
			温州: [120.65, 28.01],
			九江: [115.97, 29.71],
			邯郸: [114.47, 36.6],
			临安: [119.72, 30.23],
			兰州: [103.73, 36.03],
			沧州: [116.83, 38.33],
			临沂: [118.35, 35.05],
			南充: [106.110698, 30.837793],
			天津: [117.2, 39.13],
			富阳: [119.95, 30.07],
			泰安: [117.13, 36.18],
			诸暨: [120.23, 29.71],
			郑州: [113.65, 34.76],
			哈尔滨: [126.63, 45.75],
			聊城: [115.97, 36.45],
			芜湖: [118.38, 31.33],
			唐山: [118.02, 39.63],
			平顶山: [113.29, 33.75],
			邢台: [114.48, 37.05],
			德州: [116.29, 37.45],
			济宁: [116.59, 35.38],
			荆州: [112.239741, 30.335165],
			宜昌: [111.3, 30.7],
			义乌: [120.06, 29.32],
			丽水: [119.92, 28.45],
			洛阳: [112.44, 34.7],
			秦皇岛: [119.57, 39.95],
			株洲: [113.16, 27.83],
			石家庄: [114.48, 38.03],
			莱芜: [117.67, 36.19],
			常德: [111.69, 29.05],
			保定: [115.48, 38.85],
			湘潭: [112.91, 27.87],
			金华: [119.64, 29.12],
			岳阳: [113.09, 29.37],
			长沙: [113, 28.21],
			衢州: [118.88, 28.97],
			廊坊: [116.7, 39.53],
			菏泽: [115.480656, 35.23375],
			合肥: [117.27, 31.86],
			武汉: [114.31, 30.52],
			大庆: [125.03, 46.58]
		};

		function onLoad({ ctx, app }: CrudLoad) {
			ctx.service(service.wechat.account).done();
			app.refresh();
		}

		function buildURL(appid: string) {
			return (
				window.location.protocol + "//" + window.location.host + `/api/wechat/msg/${appid}`
			);
		}

		function clearQuota(appid: string) {
			ElMessageBox.confirm(
				"该操作每月只可执行10次, 确定要操作吗? 详细查看: https://developers.weixin.qq.com/doc/offiaccount/Message_Management/API_Call_Limits.html",
				"重要提示",
				{
					type: "warning"
				}
			).then(() => {
				service.wechat.account
					.clearQuota({
						appid: appid
					})
					.then((data: any) => {
						ElMessage.success(data);
					})
					.catch((e: any) => ElMessage.error(e));
			});
		}

		const mapOptions = ref({
			textStyle: {
				fontFamily: 'Inter, "Helvetica Neue", Arial, sans-serif'
			},
			backgroundColor: "#404a59",
			title: {
				text: "会员城市分布",
				left: "center",
				textStyle: {
					color: "#fff"
				}
			},
			tooltip: {
				trigger: "item"
			},
			legend: {
				orient: "vertical",
				y: "bottom",
				x: "right",
				data: ["pm2.5"],
				textStyle: {
					color: "#fff"
				}
			},
			geo: {
				map: "china",
				emphasis: {
					label: {
						show: false
					},
					itemStyle: {
						areaColor: "#2a333d"
					}
				},
				itemStyle: {
					areaColor: "#323c48",
					borderColor: "#111"
				}
			},
			series: [
				{
					name: "pm2.5",
					type: "scatter",
					coordinateSystem: "geo",
					data: [],
					symbolSize: (val) => val[2] / 10,
					tooltip: {
						formatter: function (val) {
							return val.name + ": " + val.value[2];
						}
					},
					itemStyle: {
						color: "#ddb926"
					}
				},
				{
					name: "Top 5",
					type: "effectScatter",
					coordinateSystem: "geo",
					data: [],
					symbolSize: (val) => val[2] / 10,
					showEffectOn: "render",
					rippleEffect: {
						brushType: "stroke"
					},
					emphasis: {
						scale: true
					},
					tooltip: {
						formatter: function (val) {
							return val.name + ": " + val.value[2];
						}
					},
					label: {
						formatter: "{b}",
						position: "right",
						show: true
					},
					itemStyle: {
						color: "#f4e925",
						shadowBlur: 10,
						shadowColor: "#333"
					},
					zlevel: 1
				}
			]
		});

		const userChartsOptions = ref({
			title: {
				text: "会员地域分布饼图",
				left: "center"
			},
			tooltip: {
				trigger: "item"
			},
			legend: {
				orient: "vertical",
				left: "left"
			},
			series: [
				{
					name: "地域分布",
					type: "pie",
					radius: "50%",
					data: [],
					emphasis: {
						itemStyle: {
							shadowBlur: 10,
							shadowOffsetX: 0,
							shadowColor: "rgba(0, 0, 0, 0.5)"
						}
					}
				}
			]
		});

		function showMapData(appid) {
			showMap.value = true;
			service.wechat.account
				.distribution({ appid: appid, type: 1 })
				.then((data: any) => {
					let dd = [];
					for (const key in data) {
						if (data[key]["province"] == "") {
							data[key]["province"] = "未知区域";
						}
						dd.push({
							name: data[key]["province"],
							value: parseInt(data[key]["total"])
						});
					}
					userChartsOptions.value.series[0].data = dd;
				})
				.catch((e) => {
					ElMessage.error(e);
				});

			service.wechat.account
				.distribution({ appid: appid, type: 2 })
				.then((data: any) => {
					let dd = [];
					for (const key in data) {
						dd.push({ name: data[key]["city"], value: parseInt(data[key]["total"]) });
					}
					mapOptions.value.series[0].data = convertData(dd);
					mapOptions.value.series[1].data = convertData(
						dd.sort((a, b) => b.value - a.value).slice(0, 6)
					);
				})
				.catch((e) => {
					ElMessage.error(e);
				});
		}

		return {
			service,
			buildURL,
			refs,
			table,
			setRefs,
			onLoad,
			upsert,
			clearQuota,
			mapOptions,
			showMapData,
			showMap,
			userChartsOptions
		};
	}
});
</script>
