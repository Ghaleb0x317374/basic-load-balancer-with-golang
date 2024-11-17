package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Request served by %s at %s \n", r.Host, time.Now().Format(time.RFC3339))
	fmt.Println(r.URL)
}

var wg sync.WaitGroup

func main() {

	http.HandleFunc("/", handler)

	wg.Add(3)
	go func() {

		fmt.Println("Server started on port 8080")
		if err := http.ListenAndServe(":8080", nil); err != nil {
			fmt.Println("Error starting server on port 8080:", err)
		}
	}()
	go func() {

		fmt.Println("Server started on port 8081")
		if err := http.ListenAndServe(":8081", nil); err != nil {
			fmt.Println("Error starting server on port 8081:", err)
		}
	}()
	go func() {

		fmt.Println("Server started on port 8082")
		if err := http.ListenAndServe(":8082", nil); err != nil {
			fmt.Println("Error starting server on port 8082:", err)
		}
	}()

	wg.Wait()
}
