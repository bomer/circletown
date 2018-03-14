package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func areaOfCircle(r int) float32 {
	var pi float32 = 3.14
	return float32(pi * float32(r) * float32(r))
}

func handler(w http.ResponseWriter, r *http.Request) {

	num := r.URL.Query().Get("number")
	if num != "" {
		convertnum, _ := strconv.Atoi(num)
		area = areaOfCircle(convertnum)
	}

	// fmt.Fprintf(w, "Hi there. Seems you hit a dummy end point! Bumblebay tuna!")
	fmt.Fprintf(w, "%f", area)
}

var area float32

func main() {

	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/area", handler)

	http.ListenAndServe(":8009", nil)
}
