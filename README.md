## Running the application

Setup `config.env.TEMPLATE` with values for your setup and rename it to `config.env`

Run using the `go run .` command from the cmd/web directory.

## Parameters

`-addr=` can be used to change the port of the application. Example `go run . -addr=":8080"`

`-dsn=` can be used to specifiy a MySql connection string. 

## Database

Database is provided using MySQL

Schema sql files are available at the root of the project to create the database matching specific chapters. 

For development a default MySQL user is created in the schema with username of `web` and password of `pass`

## config.env

This file contains setup for the project