package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/Misha/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Mishas Code")
	})

	http.ListenAndServe(":8087", nil)
}
