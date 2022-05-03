package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func notify(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	resp, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(resp))
}

func main() {
	http.HandleFunc("/notify", notify)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
