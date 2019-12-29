package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	ds "github.com/craftamap/go-luftdaten-api/datastructs"
	flag "github.com/spf13/pflag"
	"log"
	"net/http"
	"os"
)

func logging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.UserAgent(), r.URL.Path)
		f(w, r)
	}
}

func main() {
	port := flag.IntP("port", "p", 8080, "port of web-server")
	addr := flag.String("address", "0.0.0.0", "address of web-server")
	csvFile := flag.StringP("outputfile", "o", "", "file to output to")
	flag.Parse()

	http.HandleFunc("/api/post", logging(
		func(w http.ResponseWriter, r *http.Request) {
			if r.Method == "POST" {
				var sData ds.SensorData
				json.NewDecoder(r.Body).Decode(&sData)

				f, _ := os.OpenFile(*csvFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
				defer f.Close()

				csvW := csv.NewWriter(f)
				defer csvW.Flush()

				flattenedData := sData.FlattenToArray()
				log.Printf("%s", flattenedData)

				csvW.Write(flattenedData)
			} else {
				http.Redirect(w, r, r.URL.Hostname(), 301)
			}
		}),
	)

	http.ListenAndServe(fmt.Sprintf("%s:%d", *addr, *port), nil)
}
