package main

import "github.com/alokxcode/httpfromtcp"

func main() {
    server := httpfromtcp.NewServer()

    server.Handle("GET /", func(req *httpfromtcp.Req, rw *httpfromtcp.ResponseWriter) {
        rw.WriteHeader(httpfromtcp.StatusOk)
        rw.Header().Set("Content-Type", "text/plain")
		rw.Header().Set("Connection","keep-alive")

        rw.Write("Hello, world!")
    })

    server.ListenAndServe(":8080")
}
