package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("run")
	http.HandleFunc("/", helloHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello!\n")
}
