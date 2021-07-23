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
   // insert
   stmt, err := db.Prepare("INSERT INTO userinfo(username, departname, created) values(?,?,?)")
   checkErr(err)

   res, err := stmt.Exec("astaxie", "研发部门", "2012-12-09")
   checkErr(err)

   id, err := res.LastInsertId()
   checkErr(err)

   fmt.Println(id)
   // update
   stmt, err = db.Prepare("update userinfo set username=? where uid=?")
   checkErr(err)

   res, err = stmt.Exec("astaxieupdate", id)
   checkErr(err)

   affect, err := res.RowsAffected()
   checkErr(err)

   fmt.Println(affect)

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