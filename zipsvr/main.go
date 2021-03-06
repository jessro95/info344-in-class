package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
)

type zip struct { //single quote is for loon? and for quoted words, double quote. backtick = structure or preserve line breaks
	Zip   string `json:"zip"`
	City  string `json:"city"`
	State string `json:"state"`
}

type zipSlice []*zip
type zipIndex map[string]zipSlice

//has to have two params
func helloHandler(w http.ResponseWriter, r *http.Request) { //* means pointer not a value; // arrays are passed by val and is shallow copied. // pointer = reference
	name := r.URL.Query().Get("name")

	w.Header().Add("Content-Type", "text/plain")

	w.Write([]byte("hello " + name)) //respond writer that lets you write to client //slice of byte = dynamic vector/arrayList
}

//zi is receiver
func (zi zipIndex) zipsForCityHandler(w http.ResponseWriter, r *http.Request) {
	// /zips/city/seattle
	_, city := path.Split(r.URL.Path)
	lcity := strings.ToLower(city)

	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.Header().Add("Access-Control-Allow-Origin", "*") //originnanngsndgsknksgnskdgnklsdng

	encoder := json.NewEncoder(w)
	if err := encoder.Encode(zi[lcity]); err != nil {
		http.Error(w, "error encoding json: "+err.Error(), http.StatusInternalServerError)
	} //return slice
	//closure = function that returns toehr function

}

//where program starts
func main() {
	addr := os.Getenv("ADDR") // same as (var addr string = )
	if len(addr) == 0 {
		log.Fatal("please set ADDR environment var")
	}

	f, err := os.Open("./zips.csv")
	if err != nil {
		log.Fatal("error opening zips file: " + err.Error())
	}

	zips := make(zipSlice, 0, 43000)
	decoder := json.NewDecoder(f)

	if err := decoder.Decode(&zips); err != nil {
		log.Fatal("error decoding zips json: ")
	}
	fmt.Printf("loaded %d zips\n", len(zips))

	zi := make(zipIndex)

	for _, z := range zips {
		lower := strings.ToLower(z.City)
		zi[lower] = append(zi[lower], z)
	}

	fmt.Printf("there are %d zips in Seattle\n", len(zi["seattle"]))

	http.HandleFunc("/hello", helloHandler) //passing pointer rn

	http.HandleFunc("/zips/city/", zi.zipsForCityHandler) //linking!!!!!!

	fmt.Printf("server is listening at %s...\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil)) //nil is same as null //use default router
}
