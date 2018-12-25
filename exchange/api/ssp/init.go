package ssp

var SSPList map[string]SSP

func init() {
	SSPList = make(map[string]SSP)
	SSPList["1"] = SSP{
		ID:   "1",
		Name: "SSP Pertama",
	}
}
