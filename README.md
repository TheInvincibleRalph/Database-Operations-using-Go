# How to Work with SQL Databases in Go

Demo repository for
[article on SQL database access in Go](https://betterstack.com/community/guides/scaling-go/sql-databases-in-go/).

## 🟢 Prerequisites

- The latest version of Go and PostgreSQL installed

## 📦 Getting started

Clone the repository:

```bash
git clone https://github.com/betterstack-community/go-database-access.git
cd go-database-access/
```

Build the project:

```bash
go build
```

Rename `.env.example` to `.env` and modify the PostgreSQL connection string as
needed.

[Create a new database](https://www.tutorialspoint.com/postgresql/postgresql_create_database.htm)
through the `psql` console named "gda".

Populate the `gda` database:

```bash
./gda setup
```

Run code examples:

```bash
./gda example [connect|single|multi|parameterised|null|insert|transaction|struct|return|prepared|conn|timeout]
```

## ⚖ License

The code used in this project and in the linked tutorial are licensed under the
[Apache License, Version 2.0](LICENSE).
