package main

import (
	"fmt"
	"github.com/josibake/calculator"
	"html/template"
	"log"
	"net/http"
	"os"
)

type Calculation struct {
	Input  string
	Result float64
}

// get the port from the ENV on startup
func getPort() (string, error) {
	port := os.Getenv("PORT")
	if port == "" {
		return "", fmt.Errorf("$PORT not set")
	}
	return ":" + port, nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	result := calculator.Calculate(r.Form["infix"][0])
	calculation := Calculation{r.Form["infix"][0], result}

	t, _ := template.ParseFiles("results.html")
	t.ExecuteTemplate(w, "results.html", calculation)
}

func main() {

	// get port on app startup
	port, err := getPort()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("listening on %s...", port[1:])
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/calculate", handler)
	log.Fatal(http.ListenAndServe(port, nil))
}
