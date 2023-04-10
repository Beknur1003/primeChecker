package main

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
)

func main() {
	fmt.Println("Starting server...")

	http.HandleFunc("/check-prime", checkPrimeHandler)
	http.ListenAndServe(":1010", nil)
}
func checkPrimeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received request...")

	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var nums []int
	err := json.NewDecoder(r.Body).Decode(&nums)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	results := make([]bool, len(nums))
	for i, n := range nums {
		results[i] = isPrime(n)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)

	fmt.Println("Sent response...")
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
