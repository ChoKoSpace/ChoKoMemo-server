package main

import (
	"fmt"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ChoKoMemo-Server -> %s", r.URL.Path[1:])
}
func main() {
	http.HandleFunc("/", index)
	log.Fatal(http.ListenAndServe(":3901", nil))
}
