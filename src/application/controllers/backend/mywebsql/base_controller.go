package mywebsql

import (
	"encoding/json"
	"fmt"
	"github.com/jmoiron/sqlx"
	"sync"

	"github.com/xiusin/pine"
	"github.com/xiusin/pine/di"
	"github.com/xiusin/pinecms/src/application/controllers/backend/mywebsql/common"
	"github.com/xiusin/pinecms/src/application/controllers/backend/mywebsql/render"

	// "github.com/xiusin/pinecms/src/common"
	"github.com/xiusin/pinecms/src/common/helper"
)

var l sync.Once


type MyWebSql struct {
	pine.Controller
	plush *render.Plush
}

func (c *MyWebSql) Construct() {
	l.Do(func() {
		plushEngine := render.New(helper.GetRootPath("mywebsql/modules/views"), true)
		plushEngine.AddFunc("T", common.T)

		plushEngine.AddFunc("getServerList", func() map[string]common.Server {
			return common.SERVER_LIST
		})

		pine.RegisterViewEngine(plushEngine)

		di.Set(common.RenderService, func(builder di.AbstractBuilder) (interface{}, error) {
			return plushEngine, nil
		}, true)
	})

	if c.Session().Get("theme_path") == "" {
		c.Session().Set("theme_path", "default")
	}
	c.ViewData("THEME_PATH", c.Session().Get("theme_path"))
	c.ViewData("MAX_TEXT_LENGTH_DISPLAY", common.MAX_TEXT_LENGTH_DISPLAY)
	c.ViewData("APP_VERSION", common.APP_VERSION)
	c.ViewData("EXTERNAL_PATH", "/mywebsql/")

	c.plush = di.MustGet(common.RenderService).(*render.Plush)

}

func (c *MyWebSql) saveAuthSession(serve common.Server) {
	sess, _ := json.Marshal(&serve)
	c.Session().Set("auth", string(sess))
}

func (c *MyWebSql) getAuthSession(serve common.Server) {
	sess, _ := json.Marshal(&serve)
	c.Session().Set("auth", string(sess))
}

func (c *MyWebSql) clearAuthSession() {
	c.Session().Remove("auth")
}

func (c *MyWebSql) GetSQLX() *sqlx.DB {
	var serve common.Server
	sess := c.Session().Get("auth")

	if len(sess) == 0 {
		return nil
	}
	json.Unmarshal([]byte(sess), &serve)
	db, _ := sqlx.Open(serve.Driver, fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8&parseTime=true", serve.User, serve.Password, serve.Host, serve.Port))
	return db
}
