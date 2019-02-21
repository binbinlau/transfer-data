package main

import (
	"database/sql"
	"fmt"
	"github.com/binsix/transfer-data/src/main/utils"
	_ "github.com/go-sql-driver/mysql"
)

func GetSession() (*sql.DB, error) {
	return Connec(utils.Conf.Mysql.User, utils.Conf.Mysql.Password, utils.Conf.Mysql.Collection)
}

func Connec(user string, passwd string, collection string) (*sql.DB, error) {
	db, err := sql.Open("mysql", user+":"+passwd+"@/"+collection)
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	return db, err
}

func GetList() {
	db, err := GetSession()
	defer db.Close()
	rows, err := db.Query("SELECT * FROM class")
	if err != nil {
		panic(err.Error())
	}
	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error())
	}
	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error())
		}
		var value string
		for i, col := range values {
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			fmt.Println(columns[i], ": ", value)
		}
		fmt.Println("-----------------------------------")
	}
	if err = rows.Err(); err != nil {
		panic(err.Error())
	}
}

// func GetOne(id int) (interface{}) {

// }

func main() {
	// db, err := sql.Open("mysql", "root:123456@/test")
	// if err != nil {
	// 	panic(err.Error())  // Just for example purpose. You should use proper error handling instead of panic
	// }
	// defer db.Close()
	db, err := GetSession()
	defer db.Close()
	// Prepare statement for inserting data
	// stmtIns, err := db.Prepare("INSERT INTO class VALUES( ?, ? )") // ? = placeholder
	// if err != nil {
	// 	panic(err.Error()) // proper error handling instead of panic in your app
	// }
	// defer stmtIns.Close() // Close the statement when we leave main() / the program terminates

	// Prepare statement for reading data
	stmtOut, err := db.Prepare("SELECT name FROM class WHERE id = ?")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtOut.Close()

	// Insert square numbers for 0-24 in the database
	// for i := 0; i < 25; i++ {
	// 	_, err = stmtIns.Exec(i, (i * i)) // Insert tuples (i, i^2)
	// 	if err != nil {
	// 		panic(err.Error()) // proper error handling instead of panic in your app
	// 	}
	// }

	var squareNum int // we "scan" the result in here

	// Query the square-number of 13
	err = stmtOut.QueryRow(13).Scan(&squareNum) // WHERE number = 13
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	fmt.Printf("The square number of 13 is: %d", squareNum)

	// Query another number.. 1 maybe?
	err = stmtOut.QueryRow(1).Scan(&squareNum) // WHERE number = 1
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	fmt.Printf("The square number of 1 is: %d", squareNum)
	GetList()
}
