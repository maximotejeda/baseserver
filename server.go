package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	conf := preRun()
	if conf["address"] == "" {
		fmt.Printf("Server running on addr -> %s:%s\n", "localhost", conf["port"])
	} else {
		fmt.Printf("Server running on addr -> %s:%s\n", conf["address"], conf["port"])
	}

	// serve static files from static folder on the root of the project
	http.HandleFunc("/static/", ServeStatic)

	// Metrics to collect from Prometheus service
	http.Handle("/metrics", promhttp.Handler())

	log.Fatal(http.ListenAndServe(conf["address"]+":"+conf["port"], nil))
}

// preRun check all necesary env to be present
func preRun() map[string]string {
	config := make(map[string]string)

	config["address"] = os.Getenv("ADDR")
	config["port"] = os.Getenv("PORT")

	if config["port"] == "" {
		log.Fatal("Port variable not set")
	}
	return config
}

// ServeStatic serve static failes on the service 
// separated from main func in case of special cases
// SPA manage index.js and index.html 
func ServeStatic(w http.ResponseWriter, r *http.Request) {
	staticD := http.FileServer(http.Dir("static"))
	si := http.StripPrefix("/static/", staticD)

	fileURI := r.URL.Path
	file := strings.Replace(fileURI, "/", "", 1)

	si.ServeHTTP(w, r)
	_, err := os.Stat(file)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			log.Printf("file not found -> %s", file)
		}
	}else{
		log.Println("serving file -> ", file)
	}
}
