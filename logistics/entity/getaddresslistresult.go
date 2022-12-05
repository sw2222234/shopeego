package entity

import (
	"github.com/wjp-letgo/letgo/lib"
	"github.com/wjp-letgo/shopeego/commonentity"
)

//GetAddressListResult
type GetAddressListResult struct {
	commonentity.Result
	Response GetAddressListResponse `json:"response"`
}

//String
func (g GetAddressListResult) String() string {
	return lib.ObjectToString(g)
}

//GetAddressListResponse
type GetAddressListResponse struct {
	ShowPickupAddress bool            `json:"show_pickup_address"`
	AddressList       []AddressEntity `json:"address_list"`
}

//String
func (g GetAddressListResponse) String() string {
	return lib.ObjectToString(g)
}
