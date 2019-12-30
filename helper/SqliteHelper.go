package helper

import (
	ds "github.com/craftamap/go-luftdaten-api/datastructs"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var db *sqlx.DB

func OpenDb() (*sqlx.DB, error) {
	db, err := sqlx.Open("sqlite3", "data.db")
	return db, err
}

func table_exists(table string) bool {
	_, table_check := db.Query("select * from SensorData;")
	return table_check == nil

}

func CloseDb() error {
	return db.Close()
}

func CreateDatabase() error {
	var err error

	if db == nil {
		db, err = OpenDb()
	}

	if !table_exists("SensorData") {
		log.Print("Table does not exists... creating it")
		sqlStmt := `
		    CREATE TABLE "SensorData" (
		    "date"	INTEGER NOT NULL UNIQUE,
		    "SensorId"	TEXT NOT NULL,
		    "SDS_P1"	NUMERIC,
		    "SDS_P2"	NUMERIC,
		    "temperature"	NUMERIC,
		    "humidity"	NUMERIC,
		    "samples"	NUMERIC,
		    "min_micro"	NUMERIC,
		    "max_micro"	NUMERIC,
		    "signal"	NUMERIC,
		    PRIMARY KEY("date","SensorId")
		    );
		`
		_, err = db.Exec(sqlStmt)
	} else {
		log.Print("Table already exists...")
	}
	return err
}

func InsertData(sData *ds.SensorData) error {
	sqlStmt := `INSERT INTO SensorData (date, SensorId, SDS_P1, SDS_P2, temperature, humidity, samples, min_micro, max_micro, signal) VALUES (:date, :SensorId, :SDS_P1, :SDS_P2, :temperature, :humidity, :samples, :min_micro, :max_micro, :signal)`
	_, err := db.NamedExec(sqlStmt, sData.FlattenToMap())

	return err
}
