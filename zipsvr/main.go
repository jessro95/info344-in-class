package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

//has to have two params
func helloHandler(w http.ResponseWriter, r *http.Request) { //* means pointer not a value; // arrays are passed by val and is shallow copied. // pointer = reference
	name := r.URL.Query().Get("name")

	w.Header().Add("Content-Type", "text/plain")

	w.Write([]byte("hello " + name)) //respond writer that lets you write to client //slice of byte = dynamic vector/arrayList
}

//where program starts
func main() {
	addr := os.Getenv("ADDR") // same as (var addr string = )
	if len(addr) == 0 {
		log.Fatal("please set ADDR environment var")
	}

	http.HandleFunc("/hello", helloHandler) //passing pointer rn

	fmt.Printf("server is listening at %s...\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil)) //nil is same as null //use default router
}
