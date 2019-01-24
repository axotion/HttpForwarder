package main

import (
	"net/http"

	"github.com/fatih/color"
	"github.com/gorilla/mux"
)

const (
	host = "127.0.0.1"
	port = "8050"
)

func main() {
	configSites := PrepareSites()
	router := mux.NewRouter()
	mappedIdentifcators := make(map[string]*site)

	for _, client := range configSites {
		mappedIdentifcators[client.Identificator] = client
	}

	router.HandleFunc("/forward/{client}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		if r.Method != "POST" && vars["client"] == "" {
			w.WriteHeader(http.StatusTeapot)
			return
		}

		client, exist := mappedIdentifcators[vars["client"]]

		if exist {
			client.forwardHTTPRequest(r, client)
			w.WriteHeader(http.StatusOK)
			return
		}

		w.WriteHeader(http.StatusNotFound)
		return

	})
	color.Green("Server started at... %s:%s", host, port)
	http.Handle("/api", router)
	http.ListenAndServe(host+":"+port, router)
}
