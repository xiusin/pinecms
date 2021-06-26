package tables

import "time"

type DictCategory struct {
	Id        uint      `xorm:"pk autoincr 'id'" json:"id" api:"remark:分类ID"`
	Key       string    `json:"key" xorm:"key unique notnull" api:"remark:分类标识，唯一|require:true"`
	Name      string    `json:"name" xorm:"notnull" api:"remark:分类名称|require:true"`
	Remark    string    `json:"remark" api:"remark:分类备注信息"`
	Status    bool      `json:"status" api:"remark:状态|require:true"`
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
}

type Dict struct {
	Id        uint      `xorm:"pk autoincr 'id'" json:"id" api:"remark:字典属性ID"`
	Cid       uint      `json:"cid" xorm:"notnull" api:"remark:分类ID"`
	CatName   string    `json:"cat_name" xorm:"-"`
	Name      string    `json:"name" xorm:"unique notnull" api:"remark:字典名称|require:true"`
	Value     string    `json:"value" api:"remark:字典值|require:true"`
	Remark    string    `json:"remark" api:"remark:分类备注信息"`
	Sort      uint      `json:"sort"`
	Status    bool      `json:"status"`
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
}
