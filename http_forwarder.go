package httpforwarder

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/fatih/color"
)

const (
	configFile   = "sites.json"
	errorWarning = 0
	errorPanic   = 99
)

var httpClient *http.Client

type sites []struct {
	Identificator string        `json:"identificator"`
	Forward       []forwardSite `json:"forward"`
}

type forwardSite struct {
	Address        string `json:"address"`
	Method         string `json:"method"`
	Auth           string `json:"auth,omitempty"`
	Username       string `json:"username,omitempty"`
	Password       string `json:"password,omitempty"`
	Retry          int    `json:"retry"`
	ExpectedStatus int    `json:"expected_status"`
	ID             string `json:"id"`
}

func init() {
	httpClient = &http.Client{Timeout: time.Second * 10}
}

func (e sites) String() string {
	var sitesClients string

	for _, site := range e {
		sitesClients += site.Identificator + " "
	}

	return sitesClients
}

/*
  Function new will create new sites struct "object" with loaded sites from JSON file
*/

func New() sites {
	return (sites{}).prepareSites()
}

func (e sites) prepareSites() sites {
	sites := sites{}
	sitesFile, err := os.Open(configFile)
	defer sitesFile.Close()
	CheckErr(err, errorPanic)
	sitesJSON, err := ioutil.ReadAll(sitesFile)
	CheckErr(err, errorPanic)
	err = json.Unmarshal(sitesJSON, &sites)
	CheckErr(err, errorPanic)
	err = nil
	sitesJSON = nil
	return sites
}

func (e sites) appendHeadersToRequest(headers http.Header, originIP string, r *http.Request) {

	for key, value := range headers {
		r.Header.Add(key, value[0])
	}

	r.Header.Add("X-Real-IP", originIP)
	r.Header.Add("X-Forwarder-For", originIP)
	r.Header.Add("X-Forwarded-Host", originIP)
}

func (e sites) setAuthMethodForForwardSite(site *forwardSite, r *http.Request) {
	if site.Auth == "basic" {
		r.SetBasicAuth(site.Username, site.Password)
	}
}

func (e sites) forwardHTTPRequest(headers http.Header, wantedSite string, forwardSites []forwardSite, content []byte, originIP string) {

	for _, forwardSite := range forwardSites {
		go e.executeHTTPRequest(forwardSite, headers, originIP, content)
	}
}

func (e sites) executeHTTPRequest(site forwardSite, headers http.Header, originIP string, content []byte) {

	req, err := http.NewRequest(site.Method, site.Address, bytes.NewBuffer(content))
	CheckErr(err, errorWarning)

	// Request stuff like set auth header and append origin headers and IP
	e.appendHeadersToRequest(headers, originIP, req)
	e.setAuthMethodForForwardSite(&site, req)

	tried := site.Retry

	for tried >= 0 {
		log.Printf("Try to call %s", site.Address)
		res, err := httpClient.Do(req)
		if err != nil || res.StatusCode != site.ExpectedStatus {
			CheckErr(err, errorWarning)
			if err == nil {
				color.Red("Failed. Status Code: %d, try again.", res.StatusCode)
				res.Body.Close()
			}
			log.Printf("Tries left %d", tried)
			tried--
			time.Sleep(time.Second * 15)
		} else {
			res.Body.Close()
			color.Green("Done %s. Status: %d", site.Address, site.ExpectedStatus)
			return
		}

		if tried == 0 {
			color.Yellow("Sentry or webhook")
			return
		}
	}
}
