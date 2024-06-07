package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"
	"time"

	"todo/pkg/models/mysql"

	_ "github.com/go-sql-driver/mysql"  // MySQL driver
	"github.com/golangcollege/sessions" // sessions
)

// application struct holds the loggers for logging errors and info.
type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	todo     *mysql.TodosModel
	session  *sessions.Session
	users    *mysql.UserModel
	special  *mysql.SpecialModel             // for special db
}

func main() {
	// Define command-line flags for specifying the HTTP network address and DSN (Data Source Name).
	addr := flag.String("addr", ":4000", "HTTP network address")
	dsn := flag.String("dsn", "root:root@/TODO?parseTime=true", "MySQL data")
	secret := flag.String("secret", "s6Ndh+pPbnzHbS*+9Pk8qGWhTzbpa@ge", "Secret")
	flag.Parse()

	// Open files for logging informational and error messages.
	infoFile, err := os.OpenFile("./info.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer infoFile.Close()

	errorFile, err := os.OpenFile("./error.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer errorFile.Close()

	// Create loggers for informational and error messages.
	infoLog := log.New(infoFile, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(errorFile, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// Open a connection to the database.
	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()

	session := sessions.New([]byte(*secret))
	session.Lifetime = 12 * time.Hour

	// Create an instance of the application struct containing the loggers.
	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		todo:     &mysql.TodosModel{DB: db},
		session:  session, 
		users: 	  &mysql.UserModel{DB: db},
		special:  &mysql.SpecialModel{DB: db},               //special db
		}

	

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	// Log the start of the server and listen for incoming connections.

	err = srv.ListenAndServe()
	if err != nil {
		errorLog.Fatal(err)
	}
	infoLog.Printf("Starting server on %s", srv.Addr)
}

// openDB opens a connection to the database specified by the Data Source Name (DSN).
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
