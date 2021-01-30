package routers

import (
	"fmt"
	"net/http"

	"github.com/Bundy-Mundi/chartcrawler/routers/people"
	"github.com/Bundy-Mundi/chartcrawler/routers/room"
)

// NewRouters - Handle all routers from here
func NewRouters() http.Handler {

	mux := http.NewServeMux()

	/* Home */
	mux.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "Home")
	})

	/* Example Routers */
	mux.Handle("/room/", http.StripPrefix("/room", room.NewRouter()))
	mux.Handle("/people/", http.StripPrefix("/people", people.NewRouter()))

	/* Static File Server */
	/* ABOUT FILE PATH:  https://stackoverflow.com/questions/52141282/http-fileserverhttp-dir-not-working-in-separate-package */
	fs := http.StripPrefix("/static/", http.FileServer(http.Dir("./routers/public")))
	mux.Handle("/static/", fs)

	return mux
}
