package main

import (
	"database/sql"
	"flag"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/godror/godror"
)

func PrintRows(rows *sql.Rows) {
	cols, _ := rows.Columns()
	colsize := len(cols)

	for rows.Next() {
		line := ""
		coldata := make([]interface{}, colsize)

		for i := 0; i < colsize; i++ {
			coldata[i] = new(interface{})
		}
		rows.Scan(coldata...)

		for i := 0; i < colsize; i++ {
			v := coldata[i].(*interface{})
			var c string

			switch (*v).(type) {
			case nil:
				c = ""
			case int64:
				c = fmt.Sprintf("%v", *v)
			default:
				c = fmt.Sprintf("%s", *v)
			}

			line += c
			if i != (colsize - 1) {
				line += ","
			}
		}

		fmt.Println(line)
	}
}

type flags struct {
	Driver *string
	Source *string
	Query  *string
}

func getArgs() (flags, bool) {
	args := flags{}

	args.Driver = flag.String("d", "", "driver name  (mysql, mssql, oracle)")
	args.Source = flag.String("s", "",
		`source
(e.g mysql user:passwd@tcp(host:3306)/database) 
(e.g mssql server=localhost;uid=dev;pwd=devmember;database=dbname) 
(e.g oracle user/passwd@host:port/sid`)
	args.Query = flag.String("q", "", "query ")
	flag.Bool("", false, "ver. 210215.0")
	flag.Parse()

	isFlagPassed := func(name string) bool {
		found := false
		flag.Visit(func(f *flag.Flag) {
			if f.Name == name {
				found = true
			}
		})
		return found
	}

	found := isFlagPassed("d")
	found = found && isFlagPassed("s")
	found = found && isFlagPassed("q")

	if !found {
		flag.Usage()
	}
	return args, found
}

func getDriverName() map[string]string {
	driver := map[string]string{
		"mysql":  "mysql",
		"mssql":  "mssql",
		"oracle": "godror",
	}

	return driver
}

func main() {
	args, found := getArgs()
	if !found {
		return
	}

	driver := getDriverName()
	if _, ok := driver[*args.Driver]; !ok {
		fmt.Println("not found driver name")
		return
	}

	driverName := driver[*args.Driver]
	db, err := sql.Open(driverName, *args.Source)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query(*args.Query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	PrintRows(rows)
}
