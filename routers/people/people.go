package people

import (
	"fmt"
	"net/http"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "People")
	})
	mux.HandleFunc("/ben", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "Ben is awesom")
	})
	mux.HandleFunc("/sam", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "Sam is ... meh")
	})
	return mux
}
