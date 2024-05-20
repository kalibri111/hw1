package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	resonse, err := http.Get("http://localhost:8080/statistics")

	if err != nil {
		log.Fatal(err)
		return
	}

	body, errRead := io.ReadAll(resonse.Body)

	if errRead != nil {
		log.Fatal(errRead)
		return
	}

	bodyString := string(body)

	file, errFile := os.OpenFile("stat.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if errFile != nil {
		log.Fatal(errFile)
	}
	defer file.Close()

	file.Write([]byte("requested: " + bodyString + " times\n"))
}
