package gen

import (
	"github.com/yinyapeng/hb_temp/api/gen"
)

type ControllerV1 struct{}

func NewV1() gen.IGenV1 {
	return &ControllerV1{}
}
