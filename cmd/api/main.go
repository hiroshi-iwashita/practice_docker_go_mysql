package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/hiroshi-iwashita/practice_docker_go_mysql/internal/data"
	"github.com/hiroshi-iwashita/practice_docker_go_mysql/internal/driver"

	_ "github.com/go-sql-driver/mysql"
)

// config is the type for all application configuration
type config struct {
	port int //what port do we want the web server to listen on
}

// application is the type for all data we want to share with the
// various parts of our application. We will share this information
// in most cases by using this type as the receiver for functions.
type application struct {
	config      config
	infoLog     *log.Logger
	errorLog    *log.Logger
	models      data.Models
	environment string
}

func open(path string, count uint) *sql.DB {
	db, err := sql.Open("mysql", path)
	if err != nil {
		log.Fatal("open error:", err)
	}

	if err = db.Ping(); err != nil {
		time.Sleep(time.Second * 2)
		count--
		fmt.Printf("retry... count:%v\n", count)
		return open(path, count)
	}

	fmt.Println("db connected!!")
	return db
}

func connectDB() *sql.DB {
	var path string = fmt.Sprintf(
		"%s:%s@tcp(db:3306)/%s?charset=utf8&parseTime=true",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_DATABASE"))

	return open(path, 100)
}

func main() {
	// port := os.Getenv("PORT")
	var cfg config
	cfg.port = 8080

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	var path string = fmt.Sprintf(
		"%s:%s@tcp(db:3306)/%s?charset=utf8&parseTime=true",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_DATABASE"))
	environment := os.Getenv("ENV")

	db, err := driver.ConnectDb(path)
	if err != nil {
		log.Fatal("Cannot connect to database")
	}
	defer db.SQL.Close()

	app := &application{
		config:      cfg,
		infoLog:     infoLog,
		errorLog:    errorLog,
		models:      data.New(db.SQL),
		environment: environment,
	}

	err = app.serve()
	if err != nil {
		log.Fatal(err)
	}
}

// serve starts the web server
func (app *application) serve() error {
	app.infoLog.Println("API listening on port", app.config.port)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", app.config.port),
		Handler: app.routes(),
	}

	return srv.ListenAndServe()
}
