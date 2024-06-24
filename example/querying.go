package example

import (
	"context"
	"database/sql"
	"log"
	"woojiahao.com/gda/internal/utility"
)

// Single Row Query
func SingleRowQuery() {
	//connecting to the database through the pgx driver
	connStr := utility.ConnectionString()
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatalf("Unable to connect to database because %s", err)
	}

	//querying the databasse
	var johnDoeId string
	row := db.QueryRowContext(context.TODO(), `SELECT id FROM customer WHERE name = 'John Doe';`)
	err = row.Scan(&johnDoeId)
	switch {
	case err == sql.ErrNoRows:
		log.Fatalf("Unable to retrieve anyone called 'John Doe'")
	case err != nil:
		log.Fatalf("Database query failed because %s", err)
	default:
		log.Printf("John Doe has an ID of %s", johnDoeId)
	}
}

// Miltiple Row Query
func MultiRowQuery() {
	connStr := utility.ConnectionString()
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatalf("Unable to connect to database because %s", err)
	}

	orderQuantities := make(map[string]int)
	rows, err := db.QueryContext(context.TODO(), `SELECT food, sum(quantity) FROM "order" GROUP BY food;`)
	if err != nil {
		log.Fatalf("Database query failed because %s", err)
	}

	for rows.Next() {
		var food string
		var totalQuantity int
		err = rows.Scan(&food, &totalQuantity)
		if err != nil {
			log.Fatalf("Failed to retrieve row because %s", err)
		}
		orderQuantities[food] = totalQuantity
	}
	log.Printf("Total order quantity per food %v", orderQuantities)
}

// Parameterised Query
func ParameterisedQuery(target string) {
	connStr := utility.ConnectionString()
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatalf("Unable to connect to database because %s", err)
	}

	var id string
	row := db.QueryRowContext(context.TODO(), `SELECT id FROM customer WHERE name = $1;`, target)
	err = row.Scan(&id)
	switch {
	case err == sql.ErrNoRows:
		log.Fatalf("Unable to retrieve anyone called %s", target)
	case err != nil:
		log.Fatalf("Database query failed because %s", err)
	default:
		log.Printf("%s has an ID of %s", target, id)
	}
}

// Null Type Query
func NullTypeQuery() {
	connStr := utility.ConnectionString()
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatalf("Unable to connect to database because %s", err)
	}

	var allergies []sql.NullString
	rows, err := db.QueryContext(context.TODO(), `SELECT allergy FROM customer;`)
	if err != nil {
		log.Fatalf("Unable to retrieve customer allergies because %s", err)
	}

	for rows.Next() {
		var allergy sql.NullString
		err = rows.Scan(&allergy)
		if err != nil {
			log.Fatalf("Failed to scan for row because %s", err)
		}
		allergies = append(allergies, allergy)
	}
	log.Printf("Customer allergies are %v", allergies)
}

/*
QueryRowContext()

The QueryRowContext() method returns a *sql.Row type,
which has a Scan() method to map the values returned from the database
to variables. In our case, we are mapping the result of the SELECT
statement to the johnDoeId variable since we expect a string result.

When we successfully scan the row returned from the database,
the variable johnDoeId will hold the id returned from the database.
Notice that we passed a pointer reference of johnDoeId to Scan().
If a pointer reference is unused, there will be an error.




SPECIFYING CONTEXT

In Go, contexts carry deadline and cancellation signals (among others)
across API boundaries and between processes so that you can control
how long a task is allowed to take.

When working with databases, this context can be used to inform the
database service to cancel a query if too much time has elapsed to
prevent performance degradation.

In the codes above, the context.TODO() method is used, which returns
an empty context that allows the query to run for as long as it needs.


NB: rows.Next() is a method of the Rows type, which is part of the
database/sql package.


    <------var allergies []sql.NullString------>

This line declares a variable allergies which is a slice of sql.NullString.
The sql.NullString type is used to handle nullable string values from
the database. It contains a String field to hold the string value and
a Valid field to indicate if the value is non-null.
*/
