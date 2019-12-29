package main

import (
	"encoding/csv"
	"encoding/json"
	ds "github.com/craftamap/go-luftdaten-api/datastructs"
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
	http.HandleFunc("/api/post", logging(
		func(w http.ResponseWriter, r *http.Request) {
			var sData ds.SensorData
			json.NewDecoder(r.Body).Decode(&sData)

			f, _ := os.OpenFile("data.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
			defer f.Close()

			csvW := csv.NewWriter(f)
			defer csvW.Flush()

			flattenedData := sData.FlattenToArray()
			log.Printf("%s", flattenedData)

			csvW.Write(flattenedData)
		}),
	)

	http.ListenAndServe(":8080", nil)
}
