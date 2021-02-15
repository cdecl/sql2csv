
## sql2csv
sql(db) to csv format output

### Build 

```
make build
```

### Example 

```
$ bin\sql2csv.exe
Usage of bin/sql2csv:
  -     ver. 210215.0
  -d string
        driver name  (mysql, mssql, oracle)
  -q string
        query
  -s string
        source
        (e.g mysql user:passwd@tcp(host:3306)/database)
        (e.g mssql server=localhost;uid=dev;pwd=devmember;database=dbname)
        (e.g oracle user/passwd@host:port/sid
```

- mysql

```
$ sql2csv -d "mysql" -s "user:passwd@tcp(host:3306)/database" -q "select * from tablename"
```

- mssql 

```
$ sql2csv -d "mssql" -s "server=localhost;uid=dev;pwd=devmember;database=dbname" -q "select * from tablename"
$ sql2csv -d "mssql" -s "sqlserver://user:passwd@localhost:1433/?database=dbname" -q "select * from tablename"
```

- oracle 

```
$ sql2csv -d "oracle" -s 'user/passwd@host:port/sid' -q "select * from tablename"
```
