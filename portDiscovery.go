package main

import (
	"fmt"
	"log"
	"net/http"
)

func startPortDiscoveryServer() {
	http.HandleFunc("/udp", getUDPPort)
	log.Fatalln(http.ListenAndServe(discoveryPort, nil))
}

func getUDPPort(w http.ResponseWriter, r *http.Request) {
	log.Println(r.RemoteAddr, "GET udp port 200")
	w.WriteHeader(200)
	fmt.Fprint(w, udpPort)
}
