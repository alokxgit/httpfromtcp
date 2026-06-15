package main

import (
	"fmt"

	"github.com/alokxgit/httpfromtcp"
)

func mains() {
	router := httpfromtcp.NewServer()

	router.Handle("POST /", func(r *httpfromtcp.Req, w *httpfromtcp.ResponseWriter) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(httpfromtcp.StatusOk)

		w.Header().Set("Connection", "keep-alive")
		w.Write(r.Body)
	})

	router.Handle("GET /about/{id}/{profile}", func(r *httpfromtcp.Req, w *httpfromtcp.ResponseWriter) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(httpfromtcp.StatusOk)

		w.Header().Set("Connection", "keep-alive")
		fmt.Printf("| PathValue : %v , %v|", r.PathValue["id"], r.PathValue["profile"])
		w.Write(r.Body)
	})

	router.Handle("GET /service", func(r *httpfromtcp.Req, w *httpfromtcp.ResponseWriter) { w.Write("service page") })
	router.ListenAndServe(":4000")
}
