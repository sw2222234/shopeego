package entity

import (
	"github.com/wjp-letgo/letgo/lib"
	"github.com/wjp-letgo/shopeego/commonentity"
)

//GetGlobalItemIDResult
type GetGlobalItemIDResult struct {
	commonentity.Result
	Warning  string                        `json:"warning"`
	Response GetGlobalItemIDResultResponse `json:"response"`
}

//String
func (r GetGlobalItemIDResult) String() string {
	return lib.ObjectToString(r)
}

//GetGlobalItemIDResultResponse
type GetGlobalItemIDResultResponse struct {
	ItemIdMap []ItemIdMapEntity `json:"item_id_map"`
}

//String
func (r GetGlobalItemIDResultResponse) String() string {
	return lib.ObjectToString(r)
}
