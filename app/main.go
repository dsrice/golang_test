package main

import (
	"fmt"
	"net/url"
	"os"
)

import (
	"database/sql"
	_ "github.com/vertica/vertica-sql-go"
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
	defer conn.Close()
	if err != nil {
		fmt.Printf(err.Error())
		os.Exit(99)
	}

	rows, err := conn.Query("select table_name as name from tables")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(99)
	}

	for rows.Next() {
		var name string
		if err = rows.Scan(&name); err != nil {
			fmt.Println(err.Error())
			os.Exit(99)
		}
		fmt.Println(fmt.Sprintf("name: %s", name))
	}
}
