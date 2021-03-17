## Running the application

Setup `config.env.TEMPLATE` with values for your setup and rename it to `config.env`

Run using the `go run .` command from the cmd/web directory.

## Parameters

`-addr=` can be used to change the port of the application. Example `go run . -addr=":8080"`

`-dsn=` can be used to specifiy a MySql connection string. 

`-secret=` can be used to specifiy a secret token used for creating sessions

## Database

Database is provided using MySQL

Schema sql files are available at the root of the project to create the database matching specific chapters. 

For development a default MySQL user is created in the schema with username of `web` and password of `pass`

## config.env

This file contains setup for the project.

DB_CONNECTION is the connection string for your MySQL database.

SESSION_SECRET is a 32 Byte string for creating sessions.

## Generating a self-signed TLS certificate

You can generate a self-signed TSL certificate for development using GO itself. 

On Windows from the tls directory of the project run `go run PATHTOGOSOURCE\src\crypto\tls\generate_cert.go --rsa-bits=2048 --host=localhost` replacing the path to the generate_cert.go file with your own systems path to the Go source.