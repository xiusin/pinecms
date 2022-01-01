package export

import "github.com/xiusin/pine"

type ExportIntf interface {
	SetContext(*pine.Context)
	Export()
}
