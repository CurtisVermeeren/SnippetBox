# Snippet Box
Snippet Box is a CRUD application written in Go.

## Project setup
Setup `config.env.TEMPLATE` with values for your setup and rename it to `config.env`. See the `config.env` below for details of each variable.

See the `Database` section below for recreating the DB used for this project.

See the `Generating a self-signed TLS certificate` section below for setting up TLS on the Go server.

## Running the project
Run using the `go run .` command from the cmd/web directory.

Then proceed to `https://localhost:4000/`.

You can then view snippets, Register an account, or Login with an existing account on the DB.

## Parameters

`-addr=` can be used to change the port of the application. Example `go run . -addr=":8080"`

`-dsn=` can be used to specifiy a MySql connection string. 

`-secret=` can be used to specifiy a secret token used for creating sessions

## config.env
This file contains setup for the project.

DB_CONNECTION is the connection string for your MySQL database.

SESSION_SECRET is a 32 Byte string for creating sessions.

## Database
Database is provided using MySQL

Schema sql files are available at the root of the project to create the database matching specific chapters. Apply the chapter-3 file before the chapter-11 file.

For development a default MySQL user is created in the schema with username of `web` and password of `pass` which can be used for your connection string in the config.env file. 

## Generating a self-signed TLS certificate
You can generate a self-signed TSL certificate for development using GO itself. 

Create a `tls` folderin the root of the project.

On Windows from the `tls` directory of the project run `go run PATHTOGOSOURCE\src\crypto\tls\generate_cert.go --rsa-bits=2048 --host=localhost` replacing the path to the generate_cert.go file with your own systems path to the Go source.

This wil create the `cert.pem` and `key.pem` files needed for the Go server to use tls.
