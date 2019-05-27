package main

import (
	"net/http"
	"log"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/wasm")
		http.ServeFile(w, r, "go-wasm/main.wasm");
	})
	log.Fatal(http.ListenAndServe(":3000", nil))
}
