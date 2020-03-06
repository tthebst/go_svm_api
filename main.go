package main

import (
    "fmt"
    "log"
    "net/http"
	"encoding/json"

    "github.com/gorilla/mux"
)

func get(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(`{"message": "get called"}`))
}


type Data struct {
    x_data string
    y_data  string
}

func post(w http.ResponseWriter, r *http.Request) {
    
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)

    var d Data

    // Try to decode the request body into the struct. If there is an error,
    // respond to the client with the error message and a 400 status code.
    err := json.NewDecoder(r.Body).Decode(&d)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Do something with the Person struct...
    fmt.Fprintf(w, "Person: %+v", d)
    fmt.Printf(d.x_data)



    for k, v := range r.URL.Query() {
        fmt.Printf("%s: %s\n", k, v)
    }
    


    w.Write([]byte(`{"message": "post called"}`))
}

func put(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusAccepted)
    w.Write([]byte(`{"message": "put called"}`))
}

func delete(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(`{"message": "delete called"}`))
}


func main() {
    r := mux.NewRouter()

    api := r.PathPrefix("/api/v1").Subrouter()
    api.HandleFunc("", get).Methods(http.MethodGet)
    api.HandleFunc("", post).Methods(http.MethodPost)
    api.HandleFunc("", put).Methods(http.MethodPut)
    api.HandleFunc("", delete).Methods(http.MethodDelete)

    log.Fatal(http.ListenAndServe(":8080", r))
}
