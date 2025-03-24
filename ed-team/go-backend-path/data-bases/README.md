# Databases in Go

In Go exists a package named: `database-sql` but that package works with: `driver`.

We must create a specific driver that implements the interfaces from `driver`

this is the [list of recommended drivers from Go](https://go.dev/wiki/SQLDrivers)

## Open Database

In this sample we will be connecting to a Postgres Database via `pq` package:

```Go
package main

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("drivername", "dataSourceName")
	if err != nil {
		panic(err)
    }
	defer db.close()
	
	if err := db.Ping(); err != nil {
		panic(err)
    }
}
```

## SQL Commands

The db struct that is returned by the sql.Open method showed before, has the following methods to make SQL Commands:

### SELECT (DQL | DML | DDL)

- Query() - execute a query
- QueryContext() - Get query context
- QueryRow() - Get a single row

We can execute queries but is recommended only to SELECT

In those methods is not recommended to use the `_` operator when calling those methods, because return the connection from the db.

### (DDL | DML)

- Exec()
- ExecContext()

In those methods is recommended to use the `_` operator when executing.

## Prepared statements

We can tell to the DB that prepare an statement and then we only have to pass the parameters to execute the instruction.

```Go
stmt, err := db.Prepare("INSERT INTO products(name, price) VALUES ($1, $2)")
stmt.Exec(val1, val2)
```

This will help us with SQL Injection attacks. `stmt` on each call of `Exec()` he validates the pool connection, so it's possible that with many petitions he will search a new pool connection to do the prepared statement.

## Transactions

to do transactions we can do:

```Go
tx, err := db.Begin()
if err != nil {
    // handle error
    return
}

stmtInvoice, err := tx.Prepare("INSERT INTO invoices(client) VALUES(?)")
if err != nil {
    tx.Rollback()
    return
}
defer stmtInvoice.Close()

invRes, err := stmtInvoice.Exec("Alejandro")
if err != nil {
    tx.Rollback()
    return
}

invID, err := invRes.LastInsertId()
if err != nil {
    tx.Rollback()
    return
}

stmtItem, err := tx.Prepare("INSERT INTO invoice_items(invoice_id, product, price) VALUES(?, ?, ?)")
if err != nil {
    tx.Rollback()
    return
}
defer stmtItem.Close()

_, err = stmtItem.Exec(invID, "Curso Go", 50)
if err != nil {
    tx.Rollback()
    return
}

err = tx.Commit()
if err != nil {
    // handle commit error
    return
}
```

In the previous code will be necessary on each `rollback` to quit totally from that execution code, also the method returns an error that needs to be controlled and `Commit`too.

## Null Data

The zero value is important when we use databases in Go. We can work via interface `Scanner` and `Valuer`:

```Go
type Product struct { Name string}
for rows.Next() {
  var nameNull sql.NullString
  p := Product{}
  err := rows.Scan(&nameNull)

  p.Name = nameNull.String // code reduces omitting the validation, it contains zero value
}
```

We can also work via pointers without the need of an intermediate structure:

```Go
type Product struct { Name string }
    for rows.Next() {
        var name *string
	    p := Product{}
        err := rows.Scan(&name)
    
        // name contains a value if name is different to nil
        if name != nil { // We need to ensure the validation
            p.Name = *name
        }   
    }
```
