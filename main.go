package main

import (
	
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
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
	url, _ := url.Parse("http://127.0.0.1:5000/")
	return url
}