package main

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	data, err := readjson()
	if err != nil {
		log.Fatal(err)
	}

	db, err := dbconnection()
	if err != nil {
		log.Fatal(err)
	}
	inserttodb(data, db)
}

func readjson() (data JsonStructure, err error) {
	file, err := ioutil.ReadFile("./students.json")
	if err != nil {
		fmt.Println("Unable to read json file because ", err)
		return
	}

	err = json.Unmarshal(file, &data)
	if err != nil {
		fmt.Println("Unable to convert file data from json into struct because ", err)
		return
	}
	fmt.Println("File read successfully")
	return
}

func inserttodb(data JsonStructure, db *sql.DB) {
	for _, student := range data.Student {
		query := "insert into new_student(first_name, last_name, age, gender, eye_color, disability, house, stream) values (?,?,?,?,?,?,?,?)"
		row, err := db.Exec(query, student.FirstName, student.LastName, student.Age, student.Gender, student.EyeColor, student.Disability, data.House, data.Stream)
		if err != nil {
			fmt.Println("unable to insert because", err)
			continue
		}
		id, _ := row.LastInsertId()
		fmt.Println(id, "Record inserted")
	}
}

func dbconnection() (connection *sql.DB, err error) {
	dburi := "maxine:daddy@tcp(localhost:3306)/school_system"

	connection, err = sql.Open("mysql", dburi)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("DB Connection successful")
	return
}
