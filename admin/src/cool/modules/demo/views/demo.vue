<template>
	<div class="demo scroller1">
		<el-row :gutter="10">
			<el-col v-for="(item, index) in list" :key="index" :xs="24" :sm="12" :md="8" :lg="6">
				<component :is="item" />
			</el-col>
		</el-row>


	</div>
</template>

<script lang="ts">
import BClUpload from "../components/demo/b-cl-upload.vue";
import BVCopy from "../components/demo/b-v-copy.vue";
import BIconSvg from "../components/demo/b-icon-svg.vue";
import BClCrud from "../components/demo/b-cl-crud.vue";
import BClContextMenu from "../components/demo/b-cl-context-menu.vue";
import BErrorPage from "../components/demo/b-error-page.vue";
import BClEditorQuill from "../components/demo/b-cl-editor-quill.vue";

export default {
	name: "demo",

	components: {
		BClUpload,
		BVCopy,
		BIconSvg,
		BClCrud,
		BClContextMenu,
		BErrorPage,
		BClEditorQuill
	},

	setup() {
		return {
			list: [
				"b-cl-upload",
				"b-cl-crud",
				"b-icon-svg",
				"b-v-copy",
				"b-cl-context-menu",
				"b-error-page",
				"b-cl-editor-quill"
			],
			data: [
				{
					"name": "README.md",
					"size": 132,
					"mode": "-rw-r--r--",
					"modTime": "2021-03-17 14:32:28",
					"isDir": false
				},
				{
					"name": "src",
					"size": 4096,
					"mode": "drwxr-xr-x",
					"modTime": "2021-03-17 14:32:28",
					"isDir": true
				},
				{
					"name": "build.gradle",
					"size": 2132,
					"mode": "-rw-r--r--",
					"modTime": "2021-03-17 14:32:28",
					"isDir": false
				},
				{
					"name": "var",
					"size": 4096,
					"mode": "drwxr-xr-x",
					"modTime": "2021-11-05 10:35:10",
					"isDir": true
				},
				{
					"name": "tmp",
					"size": 4096,
					"mode": "dtrwxrwxrwx",
					"modTime": "2022-01-01 20:06:04",
					"isDir": true
				},
				{
					"name": "Cache.go",
					"size": 1231,
					"mode": "-rw-r--r--",
					"modTime": "2021-09-29 15:45:24",
					"isDir": false
				},
				{
					"name": "lost+found",
					"size": 16384,
					"mode": "drwx------",
					"modTime": "2020-09-04 00:54:29",
					"isDir": true
				},
				{
					"name": "pom.xml",
					"size": 2631,
					"mode": "-rw-r--r--",
					"modTime": "2021-03-17 14:32:28",
					"isDir": false
				},
				{
					"name": "etc",
					"size": 4096,
					"mode": "drwxr-xr-x",
					"modTime": "2021-09-27 15:06:53",
					"isDir": true
				},
				{
					"name": "usr",
					"size": 4096,
					"mode": "drwxr-xr-x",
					"modTime": "2021-02-09 16:54:20",
					"isDir": true
				},
				{
					"name": "run",
					"size": 800,
					"mode": "drwxr-xr-x",
					"modTime": "2021-03-04 04:30:57",
					"isDir": true
				},
				{
					"name": "mnt",
					"size": 4096,
					"mode": "drwxr-xr-x",
					"modTime": "2020-11-03 23:22:25",
					"isDir": true
				},
				{
					"name": "sys",
					"size": 0,
					"mode": "dr-xr-xr-x",
					"modTime": "2020-11-03 23:22:25",
					"isDir": true
				},
				{
					"name": ".gitignore",
					"size": 34,
					"mode": "-rw-r--r--",
					"modTime": "2021-03-17 14:32:28",
					"isDir": false
				},
				{
					"name": "root",
					"size": 4096,
					"mode": "dr-xr-x---",
					"modTime": "2021-12-31 23:55:52",
					"isDir": true
				},
				{
					"name": "boot",
					"size": 4096,
					"mode": "dr-xr-xr-x",
					"modTime": "2021-02-09 16:59:15",
					"isDir": true
				},
				{
					"name": "data",
					"size": 4096,
					"mode": "drwxr-xr-x",
					"modTime": "2021-02-15 01:56:37",
					"isDir": true
				},
				{
					"name": "home",
					"size": 4096,
					"mode": "drwxr-xr-x",
					"modTime": "2020-11-03 23:22:25",
					"isDir": true
				},
				{
					"name": "dev",
					"size": 3040,
					"mode": "drwxr-xr-x",
					"modTime": "2020-11-03 23:22:25",
					"isDir": true
				},
				{
					"name": "srv",
					"size": 4096,
					"mode": "drwxr-xr-x",
					"modTime": "2020-11-03 23:22:25",
					"isDir": true
				},
				{
					"name": "media",
					"size": 4096,
					"mode": "drwxr-xr-x",
					"modTime": "2020-11-03 23:22:25",
					"isDir": true
				},
				{
					"name": "proc",
					"size": 0,
					"mode": "dr-xr-xr-x",
					"modTime": "2020-11-03 23:22:25",
					"isDir": true
				},
				{
					"name": "opt",
					"size": 4096,
					"mode": "drwxr-xr-x",
					"modTime": "2021-02-14 22:13:05",
					"isDir": true
				}
			]
		};
	}
};
</script>

<style lang="scss">
.demo {
	.scope {
		background-color: #fff;
		border-radius: 3px;
		margin-bottom: 10px;

		.h {
			height: 30px;
			display: flex;
			align-items: center;
			padding: 10px;
			font-size: 12px;

			span {
				background-color: $color-primary;
				color: #fff;
				border-radius: 3px;
				padding: 2px 5px;
				margin-right: 10px;
				font-size: 14px;
				letter-spacing: 1px;
			}
		}

		.c {
			padding: 10px;
			height: 50px;
			box-sizing: border-box;

			&._svg {
				.icon-svg {
					margin-right: 15px;
				}
			}

			a {
				font-size: 13px;
				color: #666;
				position: relative;

				&:hover {
					&:after {
						content: "";
						width: 100%;
						height: 1px;
						position: absolute;
						bottom: -2px;
						left: 0;
						background-color: $color-primary;
					}
				}
			}
		}

		.f {
			height: 30px;
			display: flex;
			align-items: center;
			justify-content: space-between;
			padding: 10px;
			font-size: 12px;

			.date {
				color: #999;
			}
		}
	}
}
</style>
