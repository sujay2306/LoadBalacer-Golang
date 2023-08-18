package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

var (
	serverList = []string{
			"http://127.0.0.1:5000",
	 		"http://127.0.0.1:5001",
			"http://127.0.0.1:5000",
			"http://127.0.0.1:5003",
			"http://127.0.0.1:5004",

		}
	lastServedIndex = 0

)
func main(){
	http.HandleFunc("/", forwardRequest)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
//Todo:
// We need to forward the request to flask server
//we need to create a single host reverse proxy using http util
func forwardRequest(res http.ResponseWriter ,req *http.Request) {
	url := getServer()
	rProxy := httputil.NewSingleHostReverseProxy(url) 	 //this accepts a url
	rProxy.ServeHTTP(res, req) 			//takes reswriter and req as i/p
}

func getServer() *url.URL{
	nextIndex := (lastServedIndex+1)%5
	url, _ := url.Parse(serverList[lastServedIndex]) //make sure range is always 5
	lastServedIndex = nextIndex
	return url
}