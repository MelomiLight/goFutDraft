package main

import (
    "encoding/json"
    "fmt"
    "net/http"
)

type Item struct {
    ID     int    `json:"id"`
    Name   string `json:"name"`
    League int    `json:"league"`
}

func main() {
    // Define some sample data
    items := []Item{
        {1, "Arsenal", 13},
        {2, "Aston Villa", 13},
        {3, "Blackburn Rovers", 14},
        {4, "Bolton", 60},
        {5, "Chelsea", 13},
    }

    // Define a handler function to return the JSON data
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        // Set the content type header to JSON
        w.Header().Set("Content-Type", "application/json")

        // Marshal the data into JSON format
        jsonData, err := json.Marshal(items)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        // Write the JSON data to the response
        w.Write(jsonData)
    })

    // Start the server on port 8080
    fmt.Println("Server starting...")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        fmt.Println(err)
    }
}