package main

import (
	"encoding/json"
	"fmt"
	ds "github.com/craftamap/go-luftdaten-api/datastructs"
	helper "github.com/craftamap/go-luftdaten-api/helper"
	flag "github.com/spf13/pflag"
	"log"
	"net/http"
	"time"
)

func logging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.UserAgent(), r.URL.Path)
		f(w, r)
	}
}

func main() {
	helper.OpenDb()
	defer helper.CloseDb()

	helper.CreateDatabase()

	port := flag.IntP("port", "p", 8080, "port of web-server")
	addr := flag.String("address", "0.0.0.0", "address of web-server")
	csvEnabled := flag.Bool("csv", true, "Enables or disables csv writing")
	csvFile := flag.StringP("outputfile", "o", "", "file to output to")
	flag.Parse()

	http.HandleFunc("/api/post", logging(
		func(w http.ResponseWriter, r *http.Request) {
			if r.Method == "POST" {
				var sData ds.SensorData
				json.NewDecoder(r.Body).Decode(&sData)
				sData.Date = time.Now().Unix()
				if *csvEnabled {
					helper.WriteCsv(csvFile, &sData)
				}
				helper.InsertData(&sData)

			} else {
				http.Redirect(w, r, r.URL.Hostname(), 301)
			}
		}),
	)

	http.ListenAndServe(fmt.Sprintf("%s:%d", *addr, *port), nil)
}
