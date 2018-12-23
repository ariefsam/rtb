package router

import (
	"fmt"
	"net/http"

	"github.com/ariefsam/pure"
)

func Register(p *pure.Mux) {
	p.Get("/ssp/add", addSSP)
}

func addSSP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "add ssp")
}
