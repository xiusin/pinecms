package models

import (
	"xorm.io/builder"
)

type BaseModel struct {
	builder.Between
}
