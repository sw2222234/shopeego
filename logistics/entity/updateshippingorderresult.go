package entity

import (
	"github.com/wjp-letgo/letgo/lib"
	"github.com/wjp-letgo/shopeego/commonentity"
)

//UpdateShippingOrderResult
type UpdateShippingOrderResult struct {
	commonentity.Result
}

//String
func (u UpdateShippingOrderResult) String() string {
	return lib.ObjectToString(u)
}
