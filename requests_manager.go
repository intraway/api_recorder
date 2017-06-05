package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

type EnhancedRequest struct {
	r             *http.Request
	URL           string      `json:url`
	Method        string      `json:method`
	URI           string      `json:uri`
	Header        http.Header `json:header`
	Proto         string      `json:proto`
	Host          string      `json:host`
	Body          string      `json:body`
	ParsedBody    interface{} `json:parsed_body`
	RemoteAddress string      `json:remote_address`
	ContentType   string      `json:content_type`
	Timestamp     time.Time   `json:timestamp`
}

type RequestsManager struct {
	m        *sync.Mutex
	requests map[string][]*EnhancedRequest
	config   Config
}

func NewRequestsManager(config Config) *RequestsManager {
	rm := &RequestsManager{
		m:        &sync.Mutex{},
		requests: make(map[string][]*EnhancedRequest),
		config:   config,
	}

	return rm
}

func (self *RequestsManager) update(r *http.Request) {
	self.m.Lock()
	defer self.m.Unlock()
	_, ok := self.requests[r.URL.Path]

	er := &EnhancedRequest{}

	if !ok {
		self.requests[r.URL.Path] = make([]*EnhancedRequest, 0)
	}
	er.URI = r.RequestURI
	er.URL = r.URL.Path
	er.Method = r.Method
	er.Proto = r.Proto
	er.Header = r.Header
	er.Host = r.Host
	er.Timestamp = time.Now()
	er.RemoteAddress = r.RemoteAddr
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		er.Body = string(body)
	}

	if content_type, ok := r.Header["Content-Type"]; ok {
		er.ContentType = strings.Join(content_type, ",")
		// Yeah, I know, it's not the correct solution
		if strings.Contains(er.ContentType, "json") {
			decoder := json.NewDecoder(strings.NewReader(er.Body))
			decoder.UseNumber()

			if err := decoder.Decode(&er.ParsedBody); err != nil {
				er.ParsedBody = nil
			}
		}
	}

	if self.config.RecordAll {
		self.requests[r.URL.Path] = append(self.requests[r.URL.Path], er)
	} else {
		self.requests[r.URL.Path][0] = er
	}
}

func (self *RequestsManager) handleAll(w http.ResponseWriter, r *http.Request) {
	self.update(r)
	fmt.Fprintf(w, "{\"message\":\"OK\"}")
}

func (self *RequestsManager) handleShow(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	url := ""
	n := 0
	result := ""

	if url_form, ok := r.Form["url"]; ok && len(url_form) > 0 {
		url = url_form[0]
	}

	if n_form, ok := r.Form["n"]; ok && len(n_form) > 0 {
		if val, err := strconv.Atoi(n_form[0]); err == nil {
			n = val
		}
	}

	self.m.Lock()
	defer self.m.Unlock()
	if url != "" {
		response := make(map[string]interface{})
		if data, ok := self.requests[url]; ok {
			if n == 0 {
				// all elements
				response[url] = data
			} else {
				// return last n elements
				if n > len(data) {
					response[url] = data
				} else {
					response[url] = data[len(data)-n:]
				}
			}
		} else {
			// no results for this url
			response[url] = []string{}
		}
		res, _ := json.MarshalIndent(response, "", "  ")
		result = string(res)
	} else {
		// no filter, return all
		res, _ := json.MarshalIndent(self.requests, "", "  ")
		result = string(res)
	}

	fmt.Fprintf(w, string(result))
}

func (self *RequestsManager) handleReset(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	self.m.Lock()
	defer self.m.Unlock()
	if url_form, ok := r.Form["url"]; ok {
		for _, url := range url_form {
			delete(self.requests, url)
		}
	} else {
		self.requests = make(map[string][]*EnhancedRequest)
	}

	fmt.Fprintf(w, "{\"message\":\"OK\"}")
}

func (self *RequestsManager) Run() {
	log.Info("Registering '/'")
	http.HandleFunc("/", self.handleAll)

	log.Infof("Registering '/%v'", self.config.ShowURL)
	http.HandleFunc(fmt.Sprintf("/%v", self.config.ShowURL), self.handleShow)

	log.Infof("Registering '/%v'", self.config.ResetURL)
	http.HandleFunc(fmt.Sprintf("/%v", self.config.ResetURL), self.handleReset)

	address := fmt.Sprintf("%v:%v", self.config.Host, self.config.Port)

	log.Infof("Listening on %v", address)
	log.Fatal(http.ListenAndServe(address, nil))
}
