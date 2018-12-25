package dsp

import (
	"github.com/ariefsam/openrtb"
)

type DSP struct {
	ID                string
	Name              string
	BidRequestService func(openrtb.BidRequest) openrtb.BidResponse
}
