package main

import (

    "log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"

)

func main() {
	cors := handlers.CORS(
		handlers.AllowedHeaders([]string{"content-type"}),
		handlers.AllowedOrigins([]string{"*"}),
	)
	r := mux.NewRouter()
    r.HandleFunc("/createAsset", createAsset)
    r.HandleFunc("/createEscrow", createEscrow)
	r.HandleFunc("/lookup", lookupAssets)
	r.HandleFunc("/getSellings", lookupEscrowAssets)
	r.HandleFunc("/buy", buy)
	r.HandleFunc("/lookupSellings", lookupSellings)
	r.HandleFunc("/withdrawAlgos", withdrawAlgos)
	r.HandleFunc("/withdrawAssets", withdrawAssets)
	r.Use(cors)
	log.Fatal(http.ListenAndServe(":8081", (r)))
}

