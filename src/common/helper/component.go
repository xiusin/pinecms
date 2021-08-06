package helper

import "github.com/xiusin/pinecms/src/application/models/tables"

// baseItem 基础组件
type baseItem struct {
	TagName     string `json:"tag_name"`
	Name        string `json:"name"`
	Placeholder string `json:"placeholder"`
	Style       string `json:"style"`
	Label       string `json:"label"`
}

// FormItemDict 字典组件 el-autocomplete
type FormItemDict struct {
	baseItem
}

// FormItemInput 输入框
type FormItemInput struct {
	baseItem
	MixLength     uint
	Type          string
	ShowWordLimit bool
}

// FormItemSelect 下拉选择组件
type FormItemSelect struct {
}

// FormItemDateTime 时间日期组件
type FormItemDateTime struct {
}

// FormItemTags 标签组件
type FormItemTags struct {
}

// FormItemUpload 附件选择上传框
type FormItemUpload struct {
}

// FormItemUploadImage 图片上传
type FormItemUploadImage struct {
	FormItemUpload
}

// FormItemUeditor 富文本编辑器
type FormItemUeditor struct {
}

// FormItemMarkdownEditor markdown编辑器
type FormItemMarkdownEditor struct {
}

// FormItemCodeEditor 代码编辑器
type FormItemCodeEditor struct {
}

// FormItemAttr 文档属性编辑器
type FormItemAttr struct {
}

// FormItemNumberInput 数组输入框
type FormItemNumberInput struct {
	baseItem
	Min          float64 `json:"min"`
	Max          float64 `json:"max"`
	Step         float64 `json:"step"`
	StepStrictly bool    `json:"step-strictly"`
	Precision    float64 `json:"precision"`
	Size         string  `json:"size"`
	Controls     bool    `json:"controls"`
}

//FormItemCheckbox 多选框
type FormItemCheckbox struct {
}

// FormItemRadio 单选框
type FormItemRadio struct {
	baseItem
	Options  []tables.KV `json:"options"`
	Disabled bool        `json:"disabled"`
	Size     string      `json:"size"`
	Border   bool        `json:"border"`
	Label    interface{} `json:"label"`
}

// FormItemSwitch 开关按钮
type FormItemSwitch struct {
}

// FormItemSlider 滑块组件
type FormItemSlider struct {
}

// FormItemCascader 级联组件
type FormItemCascader struct {
}

// FormItemTransfer 穿梭器组件
type FormItemTransfer struct {
}

// FormItemColorPicker 颜色选择器
type FormItemColorPicker struct {
}
