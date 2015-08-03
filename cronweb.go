package main

import (
	"./core"
	"fmt"
	"log"
	"net/http"
)

type String string

func (s String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, s)
}

func main() {
	addr := "localhost:3881"
	ts := core.NewTaskStore(1)
	http.Handle("/", String(ts.Raw()))
	fmt.Println("listen at", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
