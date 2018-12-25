package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/ariefsam/godotenv"
	"github.com/ariefsam/openrtb"
	"github.com/ariefsam/pure"
	"github.com/ariefsam/rtb/exchange/api"
	"github.com/ariefsam/rtb/exchange/api/dsp"
	"github.com/ariefsam/rtb/exchange/api/dsp/service"
	"github.com/ariefsam/rtb/exchange/api/ssp"
	"github.com/ariefsam/rtb/exchange/router"
)

func MockBidRequest(openrtb.BidRequest) openrtb.BidResponse {
	var response openrtb.BidResponse
	response.ID = "sdfres"
	response.SeatBid = []openrtb.SeatBid{}
	return response
}

func MockBidRequestService(delay int, id string) func(openrtb.BidRequest) openrtb.BidResponse {

	Ret := func(openrtb.BidRequest) openrtb.BidResponse {
		var response openrtb.BidResponse
		response.ID = id
		fmt.Println("id response ", response.ID)
		response.SeatBid = []openrtb.SeatBid{}
		return response
	}
	return Ret
}
func main() {

	dspList := make(map[string]dsp.DSP)

	dspList["1"] = dsp.DSP{
		ID:                "1",
		BidRequestService: MockBidRequestService(10, "dsp 1"),
	}
	dspList["2"] = dsp.DSP{
		ID:                "2",
		BidRequestService: service.Ariefsam,
	}
	dspList["3"] = dsp.DSP{
		ID:                "3",
		BidRequestService: MockBidRequest,
	}
	dsp.DSPList = dspList

	var req openrtb.BidRequest
	req.ID = "sdfa"
	var sspEntity ssp.SSP
	_ = api.Inbound(req, sspEntity)

	var config map[string]string
	config, err := godotenv.Read()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("APP Name:", config["APP_NAME"])

	p := pure.New()

	p.Get("/", home)
	router.Register(p)

	srv := &http.Server{
		Handler: p.Serve(),
		Addr:    "localhost:8011",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 150 * time.Second,
		ReadTimeout:  150 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome")
}
