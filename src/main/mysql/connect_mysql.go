package mysql

import (
	"database/sql"
	"fmt"
	"github.com/binsix/transfer-data/src/main/utils"
	"reflect"
)

func getSession() (*sql.DB, error) {
	return Connec(utils.Conf.Mysql.User, utils.Conf.Mysql.Password, utils.Conf.Mysql.Database)
}

func Connec(user string, passwd string, collection string) (*sql.DB, error) {
	db, err := sql.Open("mysql", user+":"+passwd+"@/"+collection)
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	return db, err
}

func test() {
	Connec("1", "2", "3")
}

func GetList(table string) {
	db, err := getSession()
	defer db.Close()
	rows, err := db.Query("SELECT * FROM" + table)
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

func GetOne(id string, obj interface{}, table string) interface{} {
	getType := reflect.TypeOf(obj)
	fmt.Println("get Type is :", getType.Name())
	getValue := reflect.ValueOf(obj)
	fmt.Println("get all Fields is:", getValue)
	for i := 0; i < getType.NumField(); i++ {
		field := getType.Field(i)
		value := getValue.Field(i).Interface()
		fmt.Printf("%s: %v = %v\n", field.Name, field.Type, value)
	}

	db, err := getSession()
	stmtOut, err := db.Prepare("SELECT * FROM" + table + "WHERE id = ?")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtOut.Close()
	var squareNum string
	err = stmtOut.QueryRow(id).Scan(&squareNum) // WHERE number = 13
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	return obj
}
