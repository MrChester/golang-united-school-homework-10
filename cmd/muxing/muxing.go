package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

/**
Please note Start functions is a placeholder for you to start your own solution.
Feel free to drop gorilla.mux if you want and use any other solution available.

main function reads host/port from env just for an example, flavor it following your taste
*/

func nameHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if vars["name"] != "" {
		response := fmt.Sprintf("Hello, %s!", vars["name"])
		fmt.Fprint(w, response)
	}
}
func badNameHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("500"))
}

func dataHandler(w http.ResponseWriter, r *http.Request) {
	d, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, "I got message:\n%s", d)
}

func headersPostHandler(w http.ResponseWriter, r *http.Request) {
	var sum int = 0
	for headerName, headerValue := range r.Header {
		if headerName == "A" || headerName == "B" {
			v,_ := strconv.Atoi(headerValue[0])
			sum += v
		}
	}

	sumStr := strconv.Itoa(sum)
	w.Header().Set("a+b", sumStr)
	w.WriteHeader(http.StatusOK)
}

// Start /** Starts the web server listener on given host and port.
func Start(host string, port int) {
	router := mux.NewRouter()
	router.HandleFunc("/name/{name:[a-z A-Z]+}", nameHandler).Methods(http.MethodGet)
	router.HandleFunc("/bad", badNameHandler).Methods(http.MethodGet)
	router.HandleFunc("/data", dataHandler).Methods(http.MethodPost)
	router.HandleFunc("/headers", headersPostHandler).Methods(http.MethodPost)
	http.Handle("/", router)

	log.Println(fmt.Printf("Starting API server on %s:%d\n", host, port))
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), router); err != nil {
		fmt.Println("error")
		log.Fatal(err)
	}
}

//main /** starts program, gets HOST:PORT param and calls Start func.
func main() {
	host := os.Getenv("HOST")
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8181
	}
	Start(host, port)
}
