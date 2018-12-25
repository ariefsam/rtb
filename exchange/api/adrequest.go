package api

import (
	"fmt"

	"github.com/ariefsam/openrtb"
	"github.com/ariefsam/rtb/exchange/api/dsp"
	"github.com/ariefsam/rtb/exchange/api/ssp"
)

func inbound(req openrtb.BidRequest, ssp ssp.SSP) openrtb.BidResponse {
	var response openrtb.BidResponse
	for k, v := range dsp.DSPList {
		fmt.Println(k, v)
	}
	return response
}

func outbound(req openrtb.BidRequest, dsp dsp.DSP) openrtb.BidResponse {
	var response openrtb.BidResponse
	response.BidID = "sdf"
	response.Validate()
	return response
}
