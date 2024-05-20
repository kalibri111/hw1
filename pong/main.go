package main

import (
	"net/http"
	"strconv"
	"time"
)

var requested = 0

func timeNow(w http.ResponseWriter, r *http.Request) {
	t := time.Now()
	requested += 1
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(t.Format(time.RFC850)))
}

func stat(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(strconv.Itoa(requested)))
}

func main() {
	http.HandleFunc("/time", timeNow)
	http.HandleFunc("/statistics", stat)

	http.ListenAndServe(":8080", nil)
}
