package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	pongBaseUrl, hostExists := os.LookupEnv("SERVER_HOST")
	if !hostExists {
		pongBaseUrl = "http://localhost"
		log.Println("SERVER_HOST not set, using default")
	}
	pongPort, portExists := os.LookupEnv("SERVER_PORT")
	if !portExists {
		pongPort = "8080"
		log.Println("SERVER_PORT not set, using default")
	}
	resonse, err := http.Get(pongBaseUrl + ":" + pongPort)

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
