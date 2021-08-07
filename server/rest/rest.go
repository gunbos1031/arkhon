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

type url string

type urlResponse struct {
	URL			url		`json:"url"`
	Method 		string	`json:"method"`
	Description string	`json:"description"`
}

type msgResponse struct {
	Message string
}

func (u url) MarshalText() (text []byte, err error) {
	url := fmt.Sprintf("https://blockminingsite-dqqwx.run.goorm.io%s", u)
	return []byte(url), nil
}

func home(rw http.ResponseWriter, r *http.Request) {
	resp := []urlResponse{
		{
			URL: url("/"),
			Method: "GET",
			Description: "Describes action of each URL",
		},	
		{
			URL: url("/status"),
			Method: "GET",
			Description: "Show blockchain information",
		},
		{
			URL: url("/blocks"),
			Method: "GET",
			Description: "Show blocks of blockchain",
		},
		{
			URL: url("/blocks"),
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
		blockchain.Blockchain().AddBlock(payload.Message)
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