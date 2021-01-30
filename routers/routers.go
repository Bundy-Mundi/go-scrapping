package routers

import (
	"net/http"

	"github.com/Bundy-Mundi/chartcrawler/routers/home"
)

// NewRouters - Handle all routers from here
func NewRouters() http.Handler {
	mux := http.NewServeMux()
	mux.Handle("/", &home.Router{})
	return mux
}
