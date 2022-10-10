package main

import (
	"database/sql"
	"fmt"
	_ "github.com/vertica/vertica-sql-go"
	"net/url"
	"os"
)

func main() {
	var rawQuery = url.Values{}
	var query = url.URL{
		Scheme:   "vertica",
		User:     url.UserPassword("dbadmin", "gotest"),
		Host:     fmt.Sprintf("%s:%d", "vertica", 5433),
		Path:     "gotest",
		RawQuery: rawQuery.Encode(),
	}
	fmt.Println(query.String())
	conn, err := sql.Open("vertica", query.String())
	if err != nil {
		fmt.Printf(err.Error())
		os.Exit(99)
	}

	rows, err := conn.Query("select * from golearn.test")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(99)
	}

	for rows.Next() {
		var id int
		var name string
		if err = rows.Scan(&id, &name); err != nil {
			fmt.Println(err.Error())
			os.Exit(99)
		}
		fmt.Println(fmt.Sprintf("id: %d, name: %s", id, name))
	}
}
