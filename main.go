package main

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func get(w http.ResponseWriter, r *http.Request) {
	http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
	return
}


//struct form of post object
type data_struct struct {
	X_data []int
	Y_data []int
}

func post(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	r.ParseForm()
	log.Println(r.Form)
	var d data_struct
	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	err := json.NewDecoder(r.Body).Decode(&d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Error if provided x_data array contains no values
	if len(d.X_data) == 0 {
		http.Error(w, "please provide x_data", http.StatusBadRequest)
		return
	}

	// Error if provided y_data array contains no values
	if len(d.Y_data) == 0 {
		http.Error(w, "please provide y_data", http.StatusBadRequest)
		return
	}

	


	//var svm = NewLinearSVC("l2", "l2", true, 1.0, 0.2)(*LinearSVC, error)

	w.Write([]byte(`{"message": "post called"}`))
}

func put(w http.ResponseWriter, r *http.Request) {
	http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
	return
}

func delete(w http.ResponseWriter, r *http.Request) {
	http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
	return
}

func main() {
	r := mux.NewRouter()

	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("", get).Methods(http.MethodGet)
	api.HandleFunc("", post).Methods(http.MethodPost)
	api.HandleFunc("", put).Methods(http.MethodPut)
	api.HandleFunc("", delete).Methods(http.MethodDelete)

	log.Fatal(http.ListenAndServe(":5550", r))
}
