package router

import (
	"github.com/xiusin/pine"
	requestLog "github.com/xiusin/pine/middlewares/request-log"
	"github.com/xiusin/pinecms/src/application/controllers"
	"github.com/xiusin/pinecms/src/application/controllers/backend"
	"github.com/xiusin/pinecms/src/application/controllers/middleware"
	"github.com/xiusin/pinecms/src/application/controllers/middleware/apidoc"
	"github.com/xiusin/pinecms/src/common/helper"
	"github.com/xiusin/pinecms/src/config"
)

type RouteGroup struct {
	Prefix  string
	Handler pine.IController
}

type Interceptor struct {
	Debug       bool
	Interceptor pine.Handler
}

const CasbinModelConf = "resources/configs/rbac_models.conf"

func InitApiRouter(app *pine.Application) {
	if config.IsDebug() {
		app.Use(middleware.Demo(), middleware.Cors(), requestLog.RequestRecorder())
	}
	app.Use(middleware.Pprof(), middleware.SetGlobalConfigData(), apidoc.New(app, nil), middleware.StatesViz(app))

	casbin :=  middleware.Casbin(config.InitDB(), CasbinModelConf)

	admin := app.Group("/v2", middleware.VerifyJwtToken("/v2/public", "/v2/api"), casbin)

	InitModuleRouter(admin, app)

	adminRouteGroups := []RouteGroup{
		{Prefix: "/user", Handler: new(backend.UserController)},
		{Prefix: "/role", Handler: new(backend.AdminRoleController)},
		{Prefix: "/menu", Handler: new(backend.MenuController)},
		{Prefix: "/link", Handler: new(backend.LinkController)},
		{Prefix: "/log", Handler: new(backend.LogController)},
		{Prefix: "/errlog", Handler: new(backend.ErrorLogController)},
		{Prefix: "/assets", Handler: new(backend.AssetsManagerController)},
		{Prefix: "/attachment", Handler: new(backend.AttachmentController)},
		{Prefix: "/attachment/type", Handler: new(backend.AttachmentTypeController)},
		{Prefix: "/setting", Handler: new(backend.SettingController)},
		{Prefix: "/dict/category", Handler: new(backend.DictCategoryController)},
		{Prefix: "/dict", Handler: new(backend.DictController)},
		{Prefix: "/model", Handler: new(backend.DocumentController)},
		{Prefix: "/category", Handler: new(backend.CategoryController)},
		{Prefix: "/district", Handler: new(backend.DistrictController)},
		{Prefix: "/ad", Handler: new(backend.AdController)},
		{Prefix: "/ad/space", Handler: new(backend.AdSpaceController)},
		{Prefix: "/department", Handler: new(backend.DepartmentController)},
		{Prefix: "/position", Handler: new(backend.PositionController)},
		{Prefix: "/level", Handler: new(backend.LevelController)},
		{Prefix: "/stat", Handler: new(backend.StatController)},
		{Prefix: "/plugin", Handler: new(backend.PluginController)},
		{Prefix: "/tags", Handler: new(backend.TagsController)},
		{Prefix: "/member", Handler: new(backend.MemberController)},
		{Prefix: "/member/group", Handler: new(backend.MemberGroupController)},
		{Prefix: "/table", Handler: new(backend.TableController)},
		{Prefix: "/content", Handler: new(backend.ContentController)},
		{Prefix: "/public", Handler: new(backend.PublicController)},
		{Prefix: "/api", Handler: new(backend.PublicController)},
		{Handler: new(backend.ImSessionController)},
		{Handler: new(backend.LoginController)},
		{Handler: new(backend.IndexController)},
		{Handler: new(backend.DatabaseController)},
		{Handler: new(backend.DatabaseBackupController)},
	}

	for _, group := range adminRouteGroups {
		admin.Handle(group.Handler, group.Prefix)
	}
	helper.Inject(controllers.ServiceBackendRouter, admin)

	InitSubModuleRouter(app, admin)
}

