package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/david-gurley/gopen"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "method is not supported", http.StatusNotFound)
		}
		var quarantinePayload QuarantinePayload
		err := json.NewDecoder(r.Body).Decode(&quarantinePayload)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		c, err := gopen.NewPSMClient(os.Getenv("PSM_USERNAME"), os.Getenv("PSM_PASSWORD"), os.Getenv("PSM"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = gopen.QuarantineWorkload(c, os.Getenv("QUARANTINE_POLICY"), quarantinePayload.IP)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

	})
	fmt.Printf("started server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

type QuarantinePayload struct {
	IP string `json:"ip"`
}
