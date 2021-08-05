package rest

import (
	"github.com/gorilla/mux"
	"encoding/json"
	"net/http"
	"log"
	"github.com/gunbos1031/arkhon/utils"
)

type urlResponse struct {
	URL			string	`json:"url"`
	Method 		string	`json:"method"`
	Description string	`json:"description"`
}

func home(rw http.ResponseWriter, r *http.Request) {
	resp := []urlResponse{
		{
			URL: "/",
			Method: "GET",
			Description: "Describes action of each URL",
		},
	}
	utils.HandleErr(json.NewEncoder(rw).Encode(resp))
}

func Start() {
	router := mux.NewRouter()
	
	router.HandleFunc("/", home).Methods("GET")
	log.Fatal(http.ListenAndServe(":80", router))
}