package room

import (
	"fmt"
	"net/http"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "Room")
	})
	mux.HandleFunc("/ben", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "Ben's Room")
	})
	mux.HandleFunc("/sam", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "Sam's Room")
	})
	return mux
}
