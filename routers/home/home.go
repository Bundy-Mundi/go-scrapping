package home

import (
	"fmt"
	"net/http"
)

// Router - Serves /home
type Router struct{}

func (r *Router) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "Home")
}
