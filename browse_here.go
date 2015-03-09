package main

// got get github.com/skratchdot/open-golang/open

import (
	"fmt"
	"github.com/skratchdot/open-golang/open"
	"net/http"
	"net/url"
	"time"
)

const port = 8123

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	var metric_type string
	timestamp := int32(time.Now().Unix())
	/*
		// todo: track a metric based on the url
		// (optional) params override:
		// the default metric-type: counter [vs. timer]
		// the default metric-amount: 1  [vs. ...]
		// the date/timestamp: now.to_epoch
		// var newValues url.Values
	*/
	if r.URL != nil {
		// var e error
		//url := *r.URL
		params, e := url.ParseQuery(r.URL.RawQuery)
		metric := r.URL.Path
		if e != nil {
			panic(e)
		}
		if nil == params["type"] {
			metric_type = "counter"
		} else {
			metric_type = params["type"][0]
		}
		// tracker.Track(metric, metric_type, amount, timestamp)
		fmt.Fprintf(w, "Tracked %s (%s) at %d\n", metric, metric_type, timestamp)
	}
}

// https://golang.org/doc/articles/wiki/
func main() {
	url := fmt.Sprintf("http://localhost:%d", port)
	port_string := fmt.Sprintf(":%d", port)
	path := "/"

	//setup server on port
	http.HandleFunc(path, IndexHandler)
	// fmt.Printf("visit %s in your browser\n", url)
	open.Start(url)

	//start server
	http.ListenAndServe(port_string, nil)
}
