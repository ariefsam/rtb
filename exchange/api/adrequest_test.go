package api

import (
	"reflect"
	"testing"

	"github.com/ariefsam/openrtb"
	"github.com/ariefsam/rtb/exchange/api/dsp"
	"github.com/ariefsam/rtb/exchange/api/ssp"
)

func MockBidRequest(openrtb.BidRequest) openrtb.BidResponse {
	var response openrtb.BidResponse
	response.SeatBid = []openrtb.SeatBid{}
	return response
}

func Test_outbound(t *testing.T) {
	type args struct {
		req openrtb.BidRequest
		dsp dsp.DSP
	}

	var test1 args

	test1.dsp.BidRequestService = MockBidRequest

	tests := []struct {
		name string
		args args
		want openrtb.BidResponse
	}{
		{
			name: "Mock bid",
			args: args{
				req: openrtb.BidRequest{
					ID: "1a",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := outbound(tt.args.req, tt.args.dsp); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("outbound() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_inbound(t *testing.T) {
	type args struct {
		req openrtb.BidRequest
		ssp ssp.SSP
	}
	tests := []struct {
		name string
		args args
		want openrtb.BidResponse
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := inbound(tt.args.req, tt.args.ssp); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inbound() = %v, want %v", got, tt.want)
			}
		})
	}
}
