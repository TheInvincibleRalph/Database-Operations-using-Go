package example

import (
	"database/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
	"log"
	"woojiahao.com/gda/internal/utility"
)

func Connect() {
	connStr := utility.ConnectionString()
	db, err := sql.Open("pgx", connStr) //This line attempts to open a connection to the database using the sql.Open function. The first argument, "pgx", specifies the database driver to use (in this case, the pgx driver for PostgreSQL). The second argument, connStr, is the connection string obtained earlier.
	if err != nil {
		log.Fatalf("Unable to connect to database because %s", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("Cannot ping database because %s", err)
	}

	log.Println("Successfully connected to database and pinged it")
}
