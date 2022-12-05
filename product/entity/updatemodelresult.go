package entity

import (
	"github.com/wjp-letgo/letgo/lib"
	"github.com/wjp-letgo/shopeego/commonentity"
)

//UpdateModelResult
type UpdateModelResult struct {
	commonentity.Result
	Warning string `json:"warning"`
}

//String
func (g UpdateModelResult) String() string {
	return lib.ObjectToString(g)
}
