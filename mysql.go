package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
)

var database *sql.DB

func main () {
	// Open up our database connection.
	database, _ = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/")

	// Database actions
	createDatabaseIfNotExistCatDatabase()
	useCatDatabase()
	createTableIfNotExistCatOwners()
	//selectAllUsers()

	defer database.Close()
}

func createDatabaseIfNotExistCatDatabase() {
	database.Exec("CREATE DATABASE IF NOT EXISTS cat_database") // Create Database
}

func useCatDatabase() {
	database.Exec("USE cat_database") // Create Database
}

func createTableIfNotExistCatOwners(){
	_, queryError := database.Exec("CREATE TABLE IF NOT EXISTS cat_owners(id integer(11), name varchar(60), age integer(3), PRIMARY KEY(id))")
	if queryError != nil {
		panic(queryError)
	}
}

/*
func selectAllUsers(){
	allUsersData, queryError := database.Query("SELECT * FROM users") // SELECT * query
	if queryError!=nil {                                              // Query Error Handling
		panic(queryError)
	}
	printAllUsers(allUsersData)
	allUsersData.Close() //good habit to close
}

func printAllUsers(allUsersData *sql.Rows){
	var id, age int
	var name string
	fmt.Print("id"); fmt.Print(" ")
	fmt.Print("name"); fmt.Print(" ")
	fmt.Println("age")
	for allUsersData.Next() { //while loop
		allUsersData.Scan(&id, &name, &age)
		fmt.Print(id); fmt.Print(" ")
		fmt.Print(name); fmt.Print(" ")
		fmt.Println(age)
	}
}

func insertUser(){
	// insert
	statement, statementError := database.Prepare("INSERT INTO Users(name, age) values(?,?)")
	if(statementError != nil){
		fmt.Println(statementError.Error())
	} else {
		result, _ := statement.Exec("kage er dejligt", "23")
		id, _ := result.LastInsertId()
		fmt.Println("last insert id: "+ strconv.FormatInt(id, 10))
	}


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

}
*/