## Running the application

Run using the `go run .` command from the cmd/web directory.

## Parameters

`-addr=` can be used to change the port of the application. Example `go run . -addr=":8080"`

## Database

Database is provided using MySQL

Schema sql files are available at the root of the project to create the database mathcing specific chapters. 

For development a default MySQL user is created in the schema with username of `web` and password of `pass`