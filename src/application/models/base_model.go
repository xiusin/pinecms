package models

import (
	"github.com/go-xorm/xorm"
	"xorm.io/builder"
)

type BaseModel struct {
	orm *xorm.Engine
	builder.Between
}
