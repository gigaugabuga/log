package main

import(
	"fmt"
	"log"
	"net/http"
    "math"
    "strconv"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/hello" {
        http.Error(w, "404 not found", http.StatusNotFound)
        return
    }
    if r.Method != "GET" {
        http.Error(w, "method is not supported", http.StatusNotFound)
        return
    }
    fmt.Fprintf(w, "Coded with GO!")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
    if err := r.ParseForm(); err != nil {
        fmt.Fprintf(w, "ParseForm() err: %v", err)
        return
    }
    fmt.Fprintf(w, "POST request successful")
    name := r.FormValue("name")
    address := r.FormValue("address")
    fmt.Fprintf(w, "Name = %s\n", name)
    fmt.Fprintf(w, "Address = %s\n", address)
}

func logHandler(w http.ResponseWriter, r *http.Request) {
    if err := r.ParseForm(); err != nil {
        fmt.Fprintf(w, "ParseForm() err: %v", err)
        return
    }
    sVal := r.FormValue("logvalue")
    initValue, err := strconv.ParseFloat(sVal, 64)
    if err != nil {
        fmt.Fprintf(w, "String to float64 cast err: %v", err)
    }
    result := math.Log(initValue)
    fmt.Fprintf(w, "%v", math.Round(result*10000)/10000)
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
    http.HandleFunc("/hello", helloHandler)
    http.HandleFunc("/log", logHandler)

    fmt.Printf("Starting server at port 5000\n")
    if err := http.ListenAndServe(":5000", nil); err !=nil {
        log.Fatal(err)
    }
}