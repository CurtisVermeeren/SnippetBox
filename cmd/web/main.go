package main

import (
	"crypto/tls"
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"
	"text/template"
	"time"

	"github.com/curtisvermeeren/snippetbox/pkg/models/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golangcollege/sessions"
	"github.com/joho/godotenv"
)

// application structure holds application-wide dependencies
// the application object can inject dependencies in a neater manner than a global variable
type application struct {
	errorLog      *log.Logger
	infoLog       *log.Logger
	session       *sessions.Session
	snippets      *mysql.SnippetModel
	templateCache map[string]*template.Template
	users         *mysql.UserModel
}

func main() {
	// Load config file
	err := godotenv.Load("../../config.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Create and parse flags for runtime parameters
	addr := flag.String("addr", ":4000", "HTTP network address")
	dsn := flag.String("dsn", os.Getenv("DB_CONNECTION"), "MySQL database connection string")
	secret := flag.String("secret", os.Getenv("SESSION_SECRET"), "Session secret key")
	flag.Parse()

	// Create new logs for error or info
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// Create and open a database connection
	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()

	// Create a new cache of template files
	templateCache, err := newTemplateCache("../../ui/html/")
	if err != nil {
		errorLog.Fatal(err)
	}

	// Create a new session manager
	session := sessions.New([]byte(*secret))
	session.Lifetime = 12 * time.Hour

	// Initialize a new application
	app := &application{
		errorLog:      errorLog,
		infoLog:       infoLog,
		session:       session,
		snippets:      &mysql.SnippetModel{DB: db},
		templateCache: templateCache,
		users:         &mysql.UserModel{DB: db},
	}

	// Configure the TLS settings
	tlsConfig := &tls.Config{
		PreferServerCipherSuites: true,
		CurvePreferences:         []tls.CurveID{tls.X25519, tls.CurveP256},
	}

	// Initialize a new server
	srv := &http.Server{
		Addr:         *addr,
		ErrorLog:     errorLog,
		Handler:      app.routes(),
		TLSConfig:    tlsConfig,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	infoLog.Printf("Starting server on %s", *addr)
	err = srv.ListenAndServeTLS("../../tls/cert.pem", "../../tls/key.pem")
	errorLog.Fatal(err)
}

// openDB opens a new database connection
// dsn is a data source name also called a connection string for the database
func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
