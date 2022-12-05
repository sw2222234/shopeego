package entity

import (
	"github.com/wjp-letgo/letgo/lib"
	"github.com/wjp-letgo/shopeego/commonentity"
)

//SetNoteResult
type SetNoteResult struct {
	commonentity.Result
}

//String
func (s SetNoteResult) String() string {
	return lib.ObjectToString(s)
}
