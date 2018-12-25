package service

import (
	"github.com/ariefsam/openrtb"
)

func BidRequest(openrtb.BidRequest) openrtb.BidResponse {
	var response openrtb.BidResponse
	response.SeatBid = []openrtb.SeatBid{}
	return response
}
