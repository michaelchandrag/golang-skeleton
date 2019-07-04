package main

import (
	"log"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/gorilla/context"

	"gopkg.in/paytm/grace.v1"

	// utilhttp "bitbucket.org/michaelchandrag/chit/pkg/util/http"
	// util "bitbucket.org/michalechandrag/chit/pkg/util"
)

func main() {
	log.Println("$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$ CHIT STARTED $$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$")
	log.Println("$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$")

	muxRouter := mux.NewRouter()

	muxRouter.HandleFunc("/", Articles)
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
	// vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello")
	// r.Close = true
	// w.Header().Set("Content-Type", "application/json")
	// w.Header().Set("Access-Control-Allow-Origin", "*")

	/*if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	if err := fn(w, r); err != nil {
		log.Println(err)
		apiObject := ConstructAPIError(http.StatusInternalServerError, ErrGeneral, SysMsgErrGeneral, MsgErrGeneral)
		SendAPIObject(w, apiObject)
		return
	}*/
}