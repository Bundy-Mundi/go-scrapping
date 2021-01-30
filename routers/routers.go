package routers

import (
	"fmt"
	"net/http"

	"github.com/Bundy-Mundi/chartcrawler/routers/api"
	g_mux "github.com/gorilla/mux"
)

// NewRouters - Handle all routers from here
func NewRouters() http.Handler {
	var dir string = "./routers/public"
	mux := g_mux.NewRouter()

	/* Home */
	mux.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "Home")
	})

	/* API */
	s := mux.PathPrefix("/api/v1/").Subrouter()
	api.APIRouter(s)

	/* Static File Server */
	// ABOUT FILE PATH:  https://stackoverflow.com/questions/52141282/http-fileserverhttp-dir-not-working-in-separate-package
	mux.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(dir))))

	return mux
}
