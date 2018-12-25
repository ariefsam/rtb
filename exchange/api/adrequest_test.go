package api

import (
	"fmt"
	"testing"

	"github.com/ariefsam/openrtb"
	"github.com/ariefsam/rtb/exchange/api/dsp"
	"github.com/ariefsam/rtb/exchange/api/ssp"
)

func MockBidRequest(openrtb.BidRequest) openrtb.BidResponse {
	var response openrtb.BidResponse
	response.ID = "sdfres"
	response.SeatBid = []openrtb.SeatBid{}
	return response
}

func MockBidRequestService(delay int, response openrtb.BidResponse) func(openrtb.BidRequest) openrtb.BidResponse {
	r := response
	Ret := func(openrtb.BidRequest) openrtb.BidResponse {
		resp := r
		fmt.Println("calling service ", resp.ID)
		return resp
	}
	return Ret
}

func Test_Timeout(t *testing.T) {

	dspList := make(map[string]dsp.DSP)

	var resp openrtb.BidResponse
	var seat openrtb.SeatBid
	var bid openrtb.Bid
	bid.Price = 1000
	resp.ID = "R1"
	dspList["1"] = dsp.DSP{
		ID: "d1",
		BidRequestService: MockBidRequestService(
			10,
			resp,
		),
	}
	dspList["2"] = dsp.DSP{
		ID:                "d2",
		BidRequestService: MockBidRequestService(10, "2"),
	}
	dspList["3"] = dsp.DSP{
		ID:                "d3",
		BidRequestService: MockBidRequestService(10, "3"),
	}

	var req openrtb.BidRequest
	req.ID = "sdfa"
	var sspEntity ssp.SSP
	dsp.DSPList = dspList
	BidResponses := Inbound(req, sspEntity)
	fmt.Println("Testing...", BidResponses)

	// dsp.DSPList = dspList

	// var test1 args
	// test1.dsp.BidRequestService = MockBidRequestService(10, "1")

	// tests := []struct {
	// 	name string
	// 	args args
	// 	want openrtb.BidResponse
	// }{
	// 	{
	// 		name: "Mock bid",
	// 		args: test1,
	// 	},
	// }
	// for _, tt := range tests {
	// 	t.Run(tt.name, func(t *testing.T) {
	// 		if got := outbound(tt.args.req, tt.args.dsp); !reflect.DeepEqual(got, tt.want) {
	// 			//t.Errorf("outbound() = %v, want %v", got, tt.want)
	// 		}
	// 	})
	// }
}
