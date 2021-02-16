
## sql2csv
SQL(DB) to CSV format output

### Build 

```
make build
```

### Example 

```
$ bin\sql2csv.exe
Usage of bin/sql2csv:
  -     ver. 210216.1
  -d string
        driver name  (mysql, mssql, oracle, adodb)
  -o string
        output filename
  -q string
        query
  -r string
        row term (default "\n")
  -s string
        source
        (e.g mysql user:passwd@tcp(host:3306)/database)
        (e.g mssql server=localhost;uid=dev;pwd=devmember;database=dbname)
        (e.g oracle user/passwd@host:port/sid
        (e.g adodb provider=msdasql;dsn=dnsname;uid=user;pwd=passwd)
  -t string
        field term (default ",")
```

#### mysql
```
$ sql2csv -d "mysql" -s 'user:passwd@tcp(host:3306)/database' -q "select * from tablename" -o output.txt
```

#### mssql 

```
$ sql2csv -d "mssql" -s 'server=host;uid=uid;pwd=passwd;database=dbname' -q "select * from tablename" -o output.txt
$ sql2csv -d "mssql" -s 'sqlserver://user:passwd@host:1433/?database=dbname' -q "select * from tablename" -o output.txt
```

#### oracle 
- windows not test

```
$ sql2csv -d "oracle" -s 'user/passwd@host:port/sid' -q "select * from tablename" -o output.txt
```

#### adodb 
- for windows (slower than other drivers)
  
```
$ sql2csv -d "adodb" -s 'provider=msdasql;dsn=dnsname;uid=user;pwd=passwd' -q "select * from tablename" -o output.txt
```
