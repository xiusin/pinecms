package crud

const (
	controllerDir = "src/application/controllers/backend/"
	tableDir      = "src/application/models/tables/"
	feModuleDir   = "src/cool/modules/"
	routerFile    = "src/router/module.go"
	theme         = "vim"
	goExt         = ".go"
	controllerTpl = `package backend
import (
	"github.com/xiusin/pinecms/src/application/models/tables"
)

type [ctrl] struct {
	BaseController
}

func (c *[ctrl]) Construct() {
    c.SearchFields = map[string]searchFieldDsl{
[searchFieldDsl]
	}
	c.Group = "[table]管理"
  	c.ApiEntityName = "[table]"
	
	c.Table = &tables.[table]{}
	c.Entries = &[]tables.[table]{}
}`
	tableTpl = `package tables

[struct]
`

	indexVueTpl = `<template>
	<cl-crud @load="onLoad">
		<el-row type="flex">
			<cl-refresh-btn />
			<cl-add-btn />
			<cl-multi-delete-btn />
			<cl-flex1 />
			<cl-search-key />
		</el-row>

		<el-row>
			<cl-table v-bind="table" />
		</el-row>

		<el-row type="flex">
			<cl-flex1 />
			<cl-pagination />
		</el-row>

		<cl-upsert v-model="form" v-bind="upsert" />
	</cl-crud>
</template>

<script lang="ts">
import { CrudLoad, Table, Upsert } from "cl-admin-crud-vue3/types";
import { defineComponent, inject, reactive } from "vue";

export default defineComponent({
	name: "sys-[table]",

	setup() {
		const service = inject<any>("service");

		const form = reactive<any>({});

		const upsert = reactive<Upsert>({
			items: [formDSL]
		});

		const table = reactive<Table>({
			columns: [tableDSL]
		});

		function onLoad({ ctx, app }: CrudLoad) {
			ctx.service(service.[table]).done();
			app.refresh();
		}

		return {
			form,
			upsert,
			table,
			onLoad
		};
	}
});
</script>`

	serviceTsTpl = `import Router from "./router"; export default {[table]: new Router()};`

	serviceRouterTpl = `import { BaseService, Service, Permission } from "/@/core";
@Service("[table]")
class Sys[table] extends BaseService {}
export default Sys[table];`

	serviceIndexTsTpl = `import service from "./service"; export default { service };`
)
