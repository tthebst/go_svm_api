package main

import (
	"math"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetMethod(t *testing.T) {
	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	req, err := http.NewRequest("GET", "/api/v1", nil)
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(get)
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusMethodNotAllowed)
	}

}
func TestBlackScholes(t *testing.T) {
	got, err := blackscholes(1.0, 300.0, 250.0, 0.03, 0.15)
	if err != nil {
		t.Errorf("Unexpected Error when calling blackscholes function:  %v", err)
	}
	expected := 58.82
	diff := got - expected
	if math.Abs(diff) < 1e-9 {
		t.Errorf("Flalse blackschole calculation. Diffrence: %v", diff)
	}
}
func TestBlackScholesError(t *testing.T) {
	_, err := blackscholes(1.0, 300.0, -250.0, 0.03, 0.15)
	if err == nil {
		t.Errorf("Should raise error")
	}

}
