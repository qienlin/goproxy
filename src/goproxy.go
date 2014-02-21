package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func handler(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.URL)
	fmt.Println(req.UserAgent())
//	transport := &http.Transport{}
//	response, err := transport.RoundTrip(req)
//	if err != nil {
//		fmt.Println(err.Error())
//	}
//	w.WriteHeader(response.StatusCode)
	w.Write([]byte("goproxy test"))
}
