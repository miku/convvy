package main

import (
	"github.com/codegangsta/negroni"
	"github.com/unrolled/render"
	"net/http"
)

func Index(w http.ResponseWriter, req *http.Request) {
	r := render.New(render.Options{})
	r.HTML(w, http.StatusOK, "index", nil)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", Index)

	n := negroni.Classic()
	n.UseHandler(mux)
	n.Run(":3000")
}
