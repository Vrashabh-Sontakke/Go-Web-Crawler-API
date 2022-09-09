package crawler

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func engineCrawler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	vin := vars["vin"]
	fmt.Fprintf(w, "VIN: %s", vin)
}
