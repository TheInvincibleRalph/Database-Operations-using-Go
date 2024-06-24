package utility

import (
	"log"
	"os"
)

func ConnectionString() string {
	connStr, status := os.LookupEnv("CONN_STR")
	if !status {
		log.Fatalln("Missing environment variable CONN_STR")
	}

	return connStr
}

/*
The purpose of this program or the function above is to retrieve the
database connection string from an environment variable named CONN_STR.
If the environment variable is not set, the function logs an error message
and stops the program.

Environment variables are often used to configure applications,
especially for sensitive information like database connection strings,
because they can be set outside the source code and provide a way to
change the configuration without modifying the code.

*/
