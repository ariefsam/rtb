package api

import (
	"encoding/json"
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

func Test_BidOneImpression(t *testing.T) {

	dspList := make(map[string]dsp.DSP)

	var resp openrtb.BidResponse
	var seat openrtb.SeatBid
	var bid openrtb.Bid

	resp.ID = "R1"
	resp.Currency = "IDR"

	bid.ID = "B1"
	bid.Price = 1
	bid.ImpID = "i1"
	seat.Bid = []openrtb.Bid{bid}

	resp.SeatBid = []openrtb.SeatBid{seat}
	dspList["d1"] = dsp.DSP{
		ID: "d1",
		BidRequestService: MockBidRequestService(
			10,
			resp,
		),
	}
	bid.ID = "B2"
	bid.Price = 2500
	bid.ImpID = "i1"
	seat.Bid = []openrtb.Bid{bid}
	resp.ID = "R2"
	resp.SeatBid = []openrtb.SeatBid{seat}
	dspList["d2"] = dsp.DSP{
		ID:   "d2",
		Name: "DSP 2",
		BidRequestService: MockBidRequestService(
			10,
			resp,
		),
	}

	bid.ID = "B3"
	bid.Price = 2600
	seat.Bid = []openrtb.Bid{bid}
	resp.ID = "R3"
	resp.SeatBid = []openrtb.SeatBid{seat}
	dspList["d3"] = dsp.DSP{
		ID: "d3",
		BidRequestService: MockBidRequestService(
			10,
			resp,
		),
	}

	var req openrtb.BidRequest
	req.ID = "sdfa"
	req.Imp = []openrtb.Impression{
		openrtb.Impression{
			ID: "i1",
			Banner: &openrtb.Banner{
				H: 250,
				W: 250,
			},
		},
	}
	var sspEntity ssp.SSP
	dsp.DSPList = dspList
	BidResponse := InboundBid(req, sspEntity)
	// if winnerID != "d3" {
	// 	t.Errorf("False winner, want d3")
	// }
	x, err := json.Marshal(BidResponse)
	if err != nil {
		t.Errorf("Failed to decode Bid response")
	}
	fmt.Println("Testing...", string(x))

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
