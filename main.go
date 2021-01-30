package main

import (
	"net/http"

	"github.com/Bundy-Mundi/chartcrawler/routers"
)

func main() {

	mux := routers.NewRouters()
	http.ListenAndServe(":3000", mux)
}
