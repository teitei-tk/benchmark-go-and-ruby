package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type (
	response struct {
		Hoge string `json:"hoge"`
		Foo  int    `json:"foo"`
	}
)

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	result := response{
		Hoge: "fuga",
		Foo:  1000000000,
	}

	json.NewEncoder(w).Encode(result)
}

func hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func main() {
	router := httprouter.New()
	router.GET("/", index)
	router.GET("/hello/:name", hello)

	log.Fatal(http.ListenAndServe(":3102", router))
}
