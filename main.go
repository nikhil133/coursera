package main

import (
	crConfig "coursera/config"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//JSON writes json to server
func JSON(c crConfig.Config, str string) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, str)
	}
}

func main() {

	pg := crConfig.InitPG()
	config := crConfig.Config{
		PG:   pg,
		Port: "8080",
	}
	defer pg.Close()
	coursera := mux.NewRouter()
	coursera.HandleFunc("/coursera/push/courses", Create(config)).Methods("POST")
	coursera.HandleFunc("/coursera/get/courses", FetchData(config)).Methods("GET")
	http.Handle("/", coursera)
	log.Println("listening on :", config.Port)
	log.Fatal(http.ListenAndServe(":"+config.Port, nil))

}
