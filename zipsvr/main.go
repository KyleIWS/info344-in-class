package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
	"strings"

	"github.com/kyleiws/info344-in-class/zipsvr/handlers"
	"github.com/kyleiws/info344-in-class/zipsvr/models"
)

const zipsPath = "/zips/"

func helloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	w.Header().Add("Content-type", "text/plain")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	fmt.Fprintf(w, "Hello %s", name)
}

func memoryHandler(w http.ResponseWriter, r *http.Request) {
	runTrash := r.URL.Query().Get("trash")
	if runTrash != "" {
		runtime.GC()
	}
	//runtime.GC() // Run the garbage collector
	stats := &runtime.MemStats{}
	runtime.ReadMemStats(stats)
	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(stats)
}

func main() {
	addr := os.Getenv("ADDR")
	if len(addr) == 0 {
		addr = ":443"
	}

	tlskey := os.Getenv("TLSKEY")
	tlscert := os.Getenv("TLSCERT")
	if len(tlskey) == 0 || len(tlscert) == 0 {
		log.Fatal("Please set TLSKEY and TLSCERT")
	}

	zips, err := models.LoadZips("zips.csv")
	if err != nil {
		log.Fatalf("error reading csv file: %v\n", err)
	}
	log.Printf("loaded %d zips\n", len(zips))
	cityIndex := models.ZipIndex{}
	for _, z := range zips {
		cityLower := strings.ToLower(z.City)
		cityIndex[cityLower] = append(cityIndex[cityLower], z)
	}
	//
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/mem", memoryHandler)
	cityHandler := &handlers.CityHandler{
		Index:      cityIndex,
		PathPrefix: zipsPath,
	}
	mux.Handle(zipsPath, cityHandler)
	fmt.Printf("Server is listening on https://%s...\n", addr)
	log.Fatal(http.ListenAndServeTLS(addr, tlscert, tlskey, mux))
}
