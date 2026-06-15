package main

import "github.com/alokxgit/httpfromtcp"

func main() {
	server := httpfromtcp.NewServer()

	server.Handle("GET /", func(req *httpfromtcp.Req, rw *httpfromtcp.ResponseWriter) {
		rw.WriteHeader(httpfromtcp.StatusOk)
		rw.Header().Set("Content-Type", "text/plain")
		rw.Header().Set("Connection", "Keep-alive")

		rw.Write("Hello, world!")
	})

	server.ListenAndServe(":8080")
}
