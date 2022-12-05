package entity

import (
	"github.com/wjp-letgo/letgo/lib"
	"github.com/wjp-letgo/shopeego/commonentity"
)

//UpdateGlobalPriceResult
type UpdateGlobalPriceResult struct {
	commonentity.Result
	Warning string `json:"warning"`
}

//String
func (g UpdateGlobalPriceResult) String() string {
	return lib.ObjectToString(g)
}
