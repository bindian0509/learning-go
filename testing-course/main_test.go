package main

import (
	"testing"
)

/**
Running test
go test -v

Running test with html based coverage
go test -coverprofile=coverage.out && go tool cover -html=coverage.out

*/

// slice of structs
var tests = []struct {
	name     string
	dividend float32
	divisor  float32
	expected float32
	isError  bool
}{
	{"valid-data", 100.00, 10.00, 10.00, false},
	{"valid-5", 100.00, 20.00, 5.00, false},
	{"invalid-data", 100.00, 0.00, 0.00, true},
}

func TestDivison(t *testing.T) {
	for _, tt := range tests {
		got, err := divide(tt.dividend, tt.divisor)
		if tt.isError {
			if err == nil {
				t.Error("Got an error while we shouldn't have got one....")
			}
		} else {
			if err != nil {
				t.Error("Should have got an error but didn't something fishy...", err.Error())
			}
		}
		if got != tt.expected {
			t.Errorf("Expected %f but got %f", tt.expected, got)
		}
	}
}

func TestDivide(t *testing.T) {
	_, err := divide(1.0, 1.0)
	if err != nil {
		t.Error("Got an error while we shouldn't have got one....")
	}
}

func TestBadDivide(t *testing.T) {
	_, err := divide(0, 0)
	if err == nil {
		t.Error("Should have got an error but didn't something fishy...")
	}
}
