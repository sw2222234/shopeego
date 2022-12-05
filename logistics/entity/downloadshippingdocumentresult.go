package entity

import (
	"github.com/wjp-letgo/letgo/lib"
	"github.com/wjp-letgo/shopeego/commonentity"
)

//DownloadShippingDocumentResult
type DownloadShippingDocumentResult struct {
	commonentity.Result
	File []byte
}

//String
func (d DownloadShippingDocumentResult) String() string {
	return lib.ObjectToString(d)
}
