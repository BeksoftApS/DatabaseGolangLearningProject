package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, _ := sql.Open("sqlite3", "./DatabaseGolangLearningProject/sqliteData.db")

	// query
	rows, err := db.Query("SELECT * FROM Users")

	if(err!=nil){
		panic(err)
	}

	var id int
	var name string
	var age int

	fmt.Print("id"); fmt.Print(" ")
	fmt.Print("name"); fmt.Print(" ")
	fmt.Println("age")
	for rows.Next() {
		rows.Scan(&id, &name, &age)
		fmt.Print(id); fmt.Print(" ")
		fmt.Print(name); fmt.Print(" ")
		fmt.Println(age)
	}

	rows.Close() //good habit to close
/*
	// delete
	stmt, err = db.Prepare("delete from userinfo where uid=?")
	checkErr(err)

	res, err = stmt.Exec(id)
	checkErr(err)

	affect, err = res.RowsAffected()
	checkErr(err)
*/

	db.Close()



}