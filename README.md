
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
  -     ver. 210929.0
  -d string
        driver name  (mysql, mssql, oracle, adodb)
  -f string
        query file (- stdin)
  -o string
        output filename
  -r string
        row term (default "\n")
  -s string
        source
        (e.g mysql user:passwd@tcp(host:3306)/database)
        (e.g mssql server=localhost;user id=dev;password=devmember;database=dbname)
        (e.g oracle user/passwd@host:port/sid
        (e.g adodb provider=msdasql;dsn=dnsname;user id=user;password=passwd)
  -t string
        field term (default ",")
```

#### mysql
```sh
$ cat q.txt 
select * from tablename

# query file path 
$ sql2csv -d "mysql" -s 'user:passwd@tcp(host:3306)/database' -f q.txt -o output.txt

# query stdin
$ cat q.txt | sql2csv -d "mysql" -s 'user:passwd@tcp(host:3306)/database' -o output.txt
```

#### mssql 

```sh
$ sql2csv -d "mssql" -s 'server=host;user id=uid;password=passwd;database=dbname' -f q.txt -o output.txt
$ sql2csv -d "mssql" -s 'sqlserver://user:passwd@host:1433/?database=dbname' -f q.txt -o output.txt
```

#### oracle 
- windows not test

```sh
$ sql2csv -d "oracle" -s 'user/passwd@host:port/sid' -f q.txt -o output.txt
```

#### adodb 
- for windows (slower than other drivers)
  
```sh
$ sql2csv -d "adodb" -s 'provider=msdasql;dsn=dnsname;uid=user;pwd=passwd' -f q.txt -o output.txt
```
