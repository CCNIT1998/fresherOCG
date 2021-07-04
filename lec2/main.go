package main

import (
	"fmt"
	"net/http"

)
const (
	nameKey = "your_name"
)

func Hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello guys")
}

func HiWithParam(w http.ResponseWriter, req *http.Request) {
	params := req.URL.Query()
	// get "name" query

	if yourName, ok := params[nameKey]; ok {
		fmt.Fprintf(w, "Hi %s", yourName[0])
	} else {
		fmt.Fprintln(w, `Hi guys. I don't know your name because you don't enter the your_name query param`)
	}
}

func RunServer() {
	fmt.Println("Starting server. Please open http://localhost:8090/hello")
	defer func() {
		fmt.Println("Server is stopped")
	}()

	http.HandleFunc("/hello", Hello) // simple hello
	// hi handler will read query param
	// --> browse to http://localhost:8090/hi?name=Trung. will return Hi Trung
	http.HandleFunc("/hi", HiWithParam)
	//////////////////////////////////////////////////////////////////////////////////
	// run server
	if err := http.ListenAndServe(":8090", nil); err != nil {
		panic("Error when running server")
	}
}

func main() {
	// Just acll RunServer function on server package
	RunServer()
}
