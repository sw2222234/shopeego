package entity

import (
	"github.com/wjp-letgo/letgo/lib"
	"github.com/wjp-letgo/shopeego/commonentity"
)

//UpdateGlobalSizeChartResult
type UpdateGlobalSizeChartResult struct {
	commonentity.Result
	Warning string `json:"warning"`
}

//String
func (r UpdateGlobalSizeChartResult) String() string {
	return lib.ObjectToString(r)
}
