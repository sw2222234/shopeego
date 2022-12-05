package entity

import (
	"github.com/wjp-letgo/letgo/lib"
	"github.com/wjp-letgo/shopeego/commonentity"
)

//GetGlobalItemInfoResult
type GetGlobalItemInfoResult struct {
	commonentity.Result
	Response GetGlobalItemInfoResultResponse `json:"response"`
	Warning  string                          `json:"warning"`
}

//String
func (g GetGlobalItemInfoResult) String() string {
	return lib.ObjectToString(g)
}

//GetGlobalItemInfoResultResponse
type GetGlobalItemInfoResultResponse struct {
	GlobalItemList []GlobalItemEntity `json:"global_item_list"`
}

//String
func (g GetGlobalItemInfoResultResponse) String() string {
	return lib.ObjectToString(g)
}
