package main

import (
    "fmt"
    "net/http"
    "os"
)

func main() {
    message := os.Getenv("MESSAGE")
     message2 := os.Getenv("MESSAGE2")

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, message)
        fmt.Fprintf(w, message2)
        fmt.Fprintf(w, "hi i've added some go code")
        
    })

    fmt.Println("Starting server on port 8080.")
    http.ListenAndServe(":8080", nil)
}
