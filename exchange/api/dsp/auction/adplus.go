package auction

import (
	"github.com/ariefsam/openrtb"
)

func Adplus(openrtb.BidRequest) openrtb.BidResponse {
	var response openrtb.BidResponse
	response.SeatBid = []openrtb.SeatBid{}
	response.ID = "988"
	return response
}
