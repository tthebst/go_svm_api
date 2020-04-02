package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math"
	"net/http"

	"github.com/gorilla/mux"
	"gonum.org/v1/gonum/stat/distuv"
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
	call_price, err := blackscholes(1.0, 300.0, 250.0, 0.03, 0.15)

	fmt.Println(call_price)

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

// blackscholes calculates option prices
// Arguments:
// float time_to_mat -> time to maturity
// float spot -> spot price of underlying asset
// float strike -> strike price
// float risk_free -> risk free rate
// float sig -> volatility of underlying asset
// Returns:
// float c -> call price
func blackscholes(time_to_mat float64, spot float64, strike float64, risk_free float64, sig float64) (float64, error) {

	//check if any negativce values were provided
	if time_to_mat < 0 || spot < 0 || strike < 0 || risk_free < 0 || sig < 0 {
		return -1, errors.New("Provided negaitve number as arsgument to blackscholes")
	}
	//check if risk_free rate and volatility value percentage between 0,1
	if risk_free > 1.0 || sig > 1.0 {
		return -1, errors.New("percentage value not in the range between 0 and 1")
	}

	//define standart normal distribution
	norm_dist := distuv.Normal{
		Mu:    0,
		Sigma: 1,
	}
	// black scholes formulas SEE: https://en.wikipedia.org/wiki/Black%E2%80%93Scholes_model
	d1 := 1 / (sig * math.Sqrt(time_to_mat)) * (math.Log(spot/strike) + (risk_free+math.Pow(sig, 2)/2)*time_to_mat)
	d2 := d1 - sig*math.Sqrt(time_to_mat)
	pv := strike * math.Exp(-risk_free*time_to_mat)
	call_price := norm_dist.CDF(d1)*spot - norm_dist.CDF(d2)*pv

	return call_price, nil

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
