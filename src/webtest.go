package main

import (
	"fmt"
	"net/http"
)
import "log"

// ABC import "net/http"

var _ int = 0x0001

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))

}
func handler(w http.ResponseWriter, r *http.Request) {
	fprintf, err := fmt.Fprintf(w, "URL.PATH=%q /n", r.URL.Path)
	if err != nil {
		return
	}
	fmt.Println(fprintf)
}
