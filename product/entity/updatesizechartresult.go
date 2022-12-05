package entity

import (
	"github.com/wjp-letgo/letgo/lib"
	"github.com/wjp-letgo/shopeego/commonentity"
)

//UpdateSizeChartResult
type UpdateSizeChartResult struct {
	commonentity.Result
	Warning string `json:"warning"`
}

//String
func (r UpdateSizeChartResult) String() string {
	return lib.ObjectToString(r)
}
