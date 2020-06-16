package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3" // import sqlite driver
)

//GlobalDB - variable that allows everything to access the database
var GlobalDB *sql.DB

//CreateTable if it doesnt exist
func CreateTable() {
	storageTable := `
	CREATE TABLE IF NOT EXISTS log_table(
		Id INTEGER NOT NULL PRIMARY KEY,
		Storage TEXT,
		InsertedDatetime DATETIME
	);
	`

	_, err := GlobalDB.Exec(storageTable)
	if err != nil {
		panic(err)
	}
}

//GetShowList .
func GetShowList() []ShowList {
	sqlReadRequest := `
	SELECT Storage, InsertedDatetime
	FROM log_table
	ORDER BY Id DESC
	LIMIT 75
	`
	rows, err := GlobalDB.Query(sqlReadRequest)
	if err != nil {
		panic(err)
	}
	outputList := []ShowList{}
	for rows.Next() {
		//
		tempList := ShowList{}
		rows.Scan(&tempList.Data, &tempList.Date)
		outputList = append(outputList, tempList)
	}
	// fmt.Print(outputList)
	return outputList
}

//AddRequests Add json to database
func AddRequests(request string) {
	sqlAddRequest := `
	INSERT OR REPLACE INTO log_table(
		Storage,
		InsertedDatetime
	) values(?, CURRENT_TIMESTAMP)
	`

	statement, err := GlobalDB.Prepare(sqlAddRequest)
	if err != nil {
		panic(err)
	}
	defer statement.Close()

	_, err2 := statement.Exec(request)
	if err2 != nil {
		panic(err2)
	}

}
