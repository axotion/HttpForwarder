package httpforwarder

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/fatih/color"
	"github.com/gorilla/mux"
)

func (e sites) Run(host string, port string) {
	configSites := e.prepareSites()
	r := mux.NewRouter()
	r.HandleFunc("/forward/{client}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		if r.Method == "POST" && vars["client"] != "" {
			body, err := ioutil.ReadAll(r.Body)
			CheckErr(err, errorWarning)

			for _, client := range configSites {
				log.Println(client.Identificator)
				if client.Identificator == vars["client"] {
					go e.executeHTTPRequests(r.Header, vars["client"], client.Forward, body, r.RemoteAddr)
					w.WriteHeader(http.StatusOK)
					return
				}
			}
			w.WriteHeader(http.StatusNotFound)
		}
		w.WriteHeader(http.StatusTeapot)
	})
	color.Green("Server started at... %s:%s", host, port)
	http.Handle("/api", r)
	http.ListenAndServe(host+":"+port, r)
}
