package main

import (
	"log"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/gorilla/context"

	"gopkg.in/paytm/grace.v1"
)

func main() {
	log.Println("$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$ CHIT STARTED $$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$")
	log.Println("$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$")

	muxRouter := mux.NewRouter()

	muxRouter.HandleFunc("/articles/{category}/", Articles)
	http.Handle("/", muxRouter)

	// err = grace.Serve(":"+cfg.Server.Port, context.ClearHandler(http.DefaultServeMux))
	err := grace.Serve(":9000", context.ClearHandler(http.DefaultServeMux))
	if err != nil {
		log.Println("[ERROR GRACEFUL]", err)
		os.Exit(1)
	}

	os.Exit(0)
}

func Articles(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category: %v\n", vars["category"])
}