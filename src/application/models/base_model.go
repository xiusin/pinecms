package models

import (
	"xorm.io/builder"
	"xorm.io/xorm"
)

type BaseModel struct {
	orm *xorm.Engine
	builder.Between
}
