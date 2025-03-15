package main

import (
	"fmt"
	"net/http"
)

// Handler function for the root endpoint "/"
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

func main() {
	// Register the handler function
	http.HandleFunc("/", helloHandler)

	// Define the port
	port := "8090"
	fmt.Println("Server is running on http://localhost:" + port)

	// Start the server
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
