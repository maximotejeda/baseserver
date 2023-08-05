package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	conf := preRun()
	if conf["address"] == "" {
		fmt.Printf("Server running on addr -> %s:%s\n", "localhost", conf["port"])
	} else {
		fmt.Printf("Server running on addr -> %s:%s\n", conf["address"], conf["port"])

	}
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
