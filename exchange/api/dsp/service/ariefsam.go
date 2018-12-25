package service

import (
	"github.com/ariefsam/openrtb"
)

func Ariefsam(openrtb.BidRequest) openrtb.BidResponse {
	var response openrtb.BidResponse
	response.SeatBid = []openrtb.SeatBid{}
	response.ID = "988"
	return response
}
