package dsp

var DSPList map[string]DSP

func init() {
	DSPList = make(map[string]DSP)
	DSPList["1"] = DSP{
		ID:   "1",
		Name: "DSP satu",
	}
	DSPList["2"] = DSP{
		ID:   "1",
		Name: "DSP dua",
	}

}
