package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func notify(w http.ResponseWriter, r *http.Request) {
	defer closeOrLog(r.Body)

	resp, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err) // nolint forbidigo: print is used on purpose
	}

	fmt.Println(string(resp)) // nolint forbidigo: print is used on purpose
}

func main() {
	http.HandleFunc("/notify", notify)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}

func closeOrLog(c io.Closer) {
	err := c.Close()
	if err != nil {
		log.Printf("can't close: %v", err)
	}
}
