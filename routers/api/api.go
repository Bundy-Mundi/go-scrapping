package api

import (
	"fmt"
	"net/http"

	"github.com/Bundy-Mundi/chartcrawler/routers/api/users"
	"github.com/Bundy-Mundi/chartcrawler/scrapper"
	"github.com/gorilla/mux"
)

func APIRouter(mux *mux.Router) {

	mux.HandleFunc("/get", func(res http.ResponseWriter, req *http.Request) {
		var urls = []string{"https://www.billboard.com/charts/hot-100"}
		s := scrapper.NewScrapper(urls)
		result := s.Scrape()
		req.Header.Set("Content-Type", "application/json")

		for _, v := range result["https://www.billboard.com/charts/hot-100"] {
			fmt.Println(v)
		}
	})
	mux.HandleFunc("/post", func(res http.ResponseWriter, req *http.Request) {

	})
	mux.HandleFunc("/put", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "Put")
	})
	mux.HandleFunc("/delete", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "Delete")
	})

	/* Users API */
	sub := mux.PathPrefix("/users").Subrouter()
	sub.HandleFunc("/{id:[0-9]+}", users.GetUserInfo).Methods("GET")
	sub.HandleFunc("/", users.CreatUserHandler).Methods("POST")
}
