package main

import (
    "fmt"
    "net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/home" {
        http.NotFound(w, r)
        return
    }
    fmt.Fprintln(w, "Welcome to devopsified project.\n")
    fmt.Fprintln(w, "In this project, we will devopsify the simple Go application by:")
    fmt.Fprintln(w, "- Creating a Docker image from a Dockerfile")
    fmt.Fprintln(w, "- Deploying it to Kubernetes")
    fmt.Fprintln(w, "- Implementing CI/CD using GitHub Actions")
    fmt.Fprintln(w, "- Using Argo CD for continuous deployment")
}

func main() {
    // Register the handler function for the /home path
    http.HandleFunc("/home", homeHandler)

    // Start the server on port 8080
    fmt.Println("Server is listening on port 8080...")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Printf("Error starting server: %s\n", err)
    }
}