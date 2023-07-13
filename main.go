package main

import (
	"fmt"
	"net/http"
)

func mainPage(w http.ResponseWriter, req *http.Request) {
	response := "hello there"
	_, err := w.Write([]byte(response))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/mainPage", mainPage)
	fmt.Println("Listening port:9090")
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		panic(err)
	}
}
