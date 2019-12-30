package helper

import (
	"encoding/csv"
	ds "github.com/craftamap/go-luftdaten-api/datastructs"
	"log"
	"os"
)

func WriteCsv(csvFile *string, sData *ds.SensorData) error {
	f, err := os.OpenFile(*csvFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
	defer f.Close()

	csvW := csv.NewWriter(f)
	defer csvW.Flush()

	flattenedData := sData.FlattenToArray()
	log.Printf("%s", flattenedData)

	err = csvW.Write(flattenedData)
	return err
}
