package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)


func health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(map[string]string{"status":"OK"})
	fmt.Fprintln(w, "good boy")
}

type ScanResult struct {
	IP string `json:"ip"`
	Ports []int		`json:"ports"`
}

func portscanner(w http.ResponseWriter, r *http.Request) {
	ip := r.URL.Query().Get("ip")
	start := r.URL.Query().Get("start")
	end := r.URL.Query().Get("end")

	// error handling

	OpenPorts := main1(ip, start, end)
	resulsts := ScanResult{
		IP : ip,
		Ports: OpenPorts,
	}
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(resulsts)
}

func main(){
	http.HandleFunc("/health",health)
	http.HandleFunc("/scan",portscanner)
	http.ListenAndServe(":8080",nil)
}