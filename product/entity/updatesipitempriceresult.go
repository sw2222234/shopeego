package entity

import (
	"github.com/wjp-letgo/letgo/lib"
	"github.com/wjp-letgo/shopeego/commonentity"
)

//UpdateSipItemPriceResult
type UpdateSipItemPriceResult struct {
	commonentity.Result
	Warning string `json:"warning"`
}

//String
func (r UpdateSipItemPriceResult) String() string {
	return lib.ObjectToString(r)
}
