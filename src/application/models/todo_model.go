package models

import (
	"github.com/go-xorm/xorm"
	"github.com/xiusin/pine/di"
)

type TodoModel struct {
	orm *xorm.Engine
}

func NewTodoModel() *TodoModel {
	return &TodoModel{orm: di.MustGet("*xorm.Engine").(*xorm.Engine)}
}