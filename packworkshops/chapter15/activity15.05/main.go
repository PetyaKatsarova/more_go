package main

import (
	"log"
	"net/http"
)
/*
Get-NetTCPConnection | Where-Object { $_.LocalPort -eq 8088 }
Stop-Process -Id 12345 -Force

*/
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static-index.html")
	})
	log.Fatal(http.ListenAndServe(":8088", nil))
}