package entity

import (
	"github.com/wjp-letgo/letgo/lib"
	"github.com/wjp-letgo/shopeego/commonentity"
)

//UnSplitOrderResult
type AddInvoiceDataResult struct {
	commonentity.Result
}

//String
func (a AddInvoiceDataResult) String() string {
	return lib.ObjectToString(a)
}
