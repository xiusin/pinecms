package export

import "github.com/xiusin/pine"

type Swag struct {
	template string
}

func (s *Swag) SetContext(ctx *pine.Context) {
	s.template = ""
}

func (s *Swag) Export() {

}
