package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		name := os.Getenv("NAME")
		job := os.Getenv("JOB")
		favAnimal := os.Getenv("FAV_ANIMAL")
		currentTime := time.Now().Format(time.RFC1123)
		fmt.Fprintf(w, "Hello from BND\nName: %s\nJob: %s\nFavourite Animal: %s\nTime: %s", name, job, favAnimal, currentTime)
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
