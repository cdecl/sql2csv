
## sql2csv
sql(db) to csv format output

### Build 

```
make build
```

### Example 

```
$ bin\sql2csv.exe
Usage of bin\sql2csv.exe:
  -d string
        driver name  (e.g mysql)
  -q string
        query
  -s string
        source (e.g user:passwd@tcp(host:3306)/database )

```

- mysql

```
$ sql2csv -d "mysql" -s "user:passwd@tcp(host:3306)/database" -q "select * from tablename"
```

- mssql 

```
$ sql2csv -d "mssql" -s "sqlserver://user:passwd@localhost:1433/?database=glass" -q "select * from tablename"

// OR 

$ sql2csv -d "mssql" -s "server=localhost;uid=dev;pwd=devmember;database=glass" -q "select * from tablename"
```

