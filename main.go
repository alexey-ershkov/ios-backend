package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	PORT := ":3000"

	fmt.Println("Server started on port ", PORT)
	log.Fatal(http.ListenAndServe(PORT, nil))
}
