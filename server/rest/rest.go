package rest

import (
	"github.com/gorilla/mux"
	"encoding/json"
	"net/http"
	"log"
	"fmt"
	"github.com/gunbos1031/arkhon/utils"
	"github.com/gunbos1031/arkhon/blockchain"
)

type urlResponse struct {
	URL			string	`json:"url"`
	Method 		string	`json:"method"`
	Description string	`json:"description"`
}

type msgResponse struct {
	Message string
}

func home(rw http.ResponseWriter, r *http.Request) {
	resp := []urlResponse{
		{
			URL: "/",
			Method: "GET",
			Description: "Describes action of each URL",
		},	
		{
			URL: "/status",
			Method: "GET",
			Description: "Show blockchain information",
		},
		{
			URL: "/blocks",
			Method: "GET",
			Description: "Show blocks of blockchain",
		},
		{
			URL: "/blocks",
			Method: "POST",
			Description: "Add blocks to blockchain",
		},
	}
	utils.HandleErr(json.NewEncoder(rw).Encode(resp))
}

func status(rw http.ResponseWriter, r *http.Request) {
	utils.HandleErr(json.NewEncoder(rw).Encode(blockchain.Blockchain()))
}

func blocks(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case "GET":
		blocks := blockchain.Blocks(blockchain.Blockchain())
		utils.HandleErr(json.NewEncoder(rw).Encode(blocks))
		case "POST":
		var payload msgResponse
		utils.HandleErr(json.NewDecoder(r.Body).Decode(&payload))
		blockchain.AddBlock(payload.Message)
	}
}

func Start() {
	router := mux.NewRouter()
	
	router.HandleFunc("/", home).Methods("GET")
	router.HandleFunc("/status", status).Methods("GET")
	router.HandleFunc("/blocks", blocks).Methods("GET", "POST")
	fmt.Println("localhost:80 starts")
	log.Fatal(http.ListenAndServe(":80", router))
}