package main

import (
	"fmt"
	"net/http"
)

var isHealthy bool

func checkhealthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("health checked...\n")
	if isHealthy {
		fmt.Fprintf(w, "Healthy\n")
	} else {
		http.Error(w, "application is unhealthy", 500)
	}
}

func killHealthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("killing health ...\n")
	isHealthy = false
	fmt.Fprintf(w, "Applicaiton will now appear down\n")
}

func restoreHealthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("restoring health...\n")
	isHealthy = true
	fmt.Fprintf(w, "Application will now appear healthy\n")
}

func main() {
	isHealthy = true

	fmt.Printf("Starting test web app...\n")
	http.HandleFunc("/checkhealth", checkhealthHandler)
	http.HandleFunc("/killhealth", killHealthHandler)
	http.HandleFunc("/restorehealth", restoreHealthHandler)
	http.ListenAndServe(":8080", nil)
}
