package main

import (
	"bufio"
	"database/sql"
	"flag"
	"fmt"
	"os"
	"strings"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/godror/godror"
	_ "github.com/mattn/go-adodb"
)

// PrintRows Rows
func PrintRows(rows *sql.Rows, filename string, fs string, rs string) {
	cols, _ := rows.Columns()
	colsize := len(cols)
	colsArr := []string{}
	fin, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer fin.Close()

	w := bufio.NewWriter(fin)
	idx := 0

	for rows.Next() {
		colmeta := make([]interface{}, colsize)

		for i := 0; i < colsize; i++ {
			colmeta[i] = new(interface{})
		}
		rows.Scan(colmeta...)

		for i := 0; i < colsize; i++ {
			v := colmeta[i].(*interface{})
			var c string

			switch (*v).(type) {
			case nil:
				c = ""
			case float64, float32:
				c = fmt.Sprintf("%f", *v)
			case int64, int32, int16:
				c = fmt.Sprintf("%v", *v)
			default:
				c = fmt.Sprintf("%s", *v)
			}

			colsArr = append(colsArr, c)
		}

		line := strings.Join(colsArr, fs)
		colsArr = colsArr[:0]

		w.WriteString(line)
		w.WriteString(rs)

		if (idx % 10000) == 0 {
			w.Flush()
			fmt.Printf("> %10d rows flush     \r", idx)
		}
		idx++
	}

	w.Flush()
	fmt.Printf("> %10d rows flush     \n", idx)
}

type flags struct {
	Driver    *string
	Source    *string
	Output    *string
	FieldTerm *string
	RowTerm   *string
	Query     *string
}

func getArgs() (flags, bool) {
	args := flags{}

	args.Driver = flag.String("d", "", "driver name  (mysql, mssql, oracle, adodb)")
	args.FieldTerm = flag.String("t", ",", "field term")
	args.RowTerm = flag.String("r", "\n", "row term")
	args.Output = flag.String("o", "", "output filename")
	args.Query = flag.String("q", "", "query ")
	args.Source = flag.String("s", "",
		`source
(e.g mysql user:passwd@tcp(host:3306)/database) 
(e.g mssql server=localhost;uid=dev;pwd=devmember;database=dbname) 
(e.g oracle user/passwd@host:port/sid
(e.g adodb provider=msdasql;dsn=dnsname;uid=user;pwd=passwd) `)

	flag.Bool("", false, "ver. 210216.1")
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
	found = found && isFlagPassed("o")

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
		"adodb":  "adodb",
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

	PrintRows(rows, *args.Output, *args.FieldTerm, *args.RowTerm)
}
