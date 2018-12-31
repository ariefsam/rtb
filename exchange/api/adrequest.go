package api

import (
	"time"

	"github.com/ariefsam/openrtb"
	"github.com/ariefsam/rtb/exchange/api/dsp"
	"github.com/ariefsam/rtb/exchange/api/ssp"
)

func InboundBid(req openrtb.BidRequest, ssp ssp.SSP) openrtb.BidResponse {
	var response openrtb.BidResponse
	responseList := outboundBid(req)
	for _, val := range responseList {
		for _, seatBid := range val.SeatBid {
			response.SeatBid = append(response.SeatBid, seatBid)
		}
	}
	return response
}

func outboundBid(req openrtb.BidRequest) map[string]openrtb.BidResponse {
	x := make(map[string]openrtb.BidResponse)
	responseList := make(map[string]openrtb.BidResponse)
	for _, v := range dsp.DSPList {
		go func(dspEntity dsp.DSP) {
			response := dspEntity.BidRequestService(req)
			x[dspEntity.ID] = response
		}(v)
	}

	time.Sleep(1000 * time.Millisecond)
	for k, val := range x {
		responseList[k] = val
	}
	return responseList
}
