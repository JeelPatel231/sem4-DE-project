package database

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"

	_ "github.com/lib/pq"
)

// for json marshalling and response
type responseField struct {
	Restaurant_id string  `json:"id"`
	Layout        [][]int `json:"layout"`
}

// function to connect to a URL and return the database pointer
func Connect(connStr string) *sql.DB {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	return db
}

// pass the DB and name of restaurent to register it in db
// returns a string which is uuid of the specific restaurant
func InsertRestaurant(db *sql.DB, name string) string {
	rows, err := db.Query(
		"INSERT INTO uuididentify (name) VALUES ($1) RETURNING id",
		name,
	)
	if err != nil {
		panic(err)
	}
	var response string
	for rows.Next() {
		rows.Scan(&response)
	}
	return response
}

// inserts the restaurant's layout in the 2nd table of db
// referenced uui from 1st table primary key
func InsertLayout(db *sql.DB, uuid string, layout [][]int) {
	array := strings.Join(strings.Fields(fmt.Sprint(layout)), ",")

	_, err := db.Query(
		"INSERT into restaurant (id, layout) VALUES ($1,$2)",
		uuid,
		array,
	)
	if err != nil {
		panic(err)
	}

}

// fetch layout of a restaurant from its unique uuid
func GetLayout(db *sql.DB, id string) [][]int {
	rows, err := db.Query(
		"select layout from restaurant where id=$1",
		id,
	)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var response string
	for rows.Next() {
		rows.Scan(&response)
	}
	var int_arr [][]int
	json.Unmarshal(
		[]byte(response),
		&int_arr,
	)
	return int_arr
}

// query the whole db by using joins and common uuid field
// and return in the responseField struct type
func Query(db *sql.DB, uuid string) []responseField {
	rows, err := db.Query(
		"select * from restaurant where id=$1",
		uuid,
	)
	if err != nil {
		panic(err)
	}

	var responseArray []responseField

	for rows.Next() {
		var response responseField
		var str_layout string
		rows.Scan(&response.Restaurant_id, &str_layout)

		json.Unmarshal(
			[]byte(str_layout),
			&response.Layout,
		)

		responseArray = append(responseArray, response)
	}
	return responseArray
}