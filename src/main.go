package main

import (
    "time"
    "log"
    "os"
    "net/http"
    "architalks/routes"
    "database/sql"
    "github.com/joho/godotenv"
    "github.com/gorilla/mux"
    _ "github.com/go-sql-driver/mysql"
)

func main() {
    godotenv.Load(".env")

    database_driver := os.Getenv("DB_DRIVER")
    database_user := os.Getenv("DB_USER")
    database_password := os.Getenv("DB_PASSWORD")
    database := os.Getenv("DB_NAME")

    var connectionStr string = database_user + ":" + database_password + "@/" + database

    db, err := sql.Open(database_driver, connectionStr)

    if err != nil {
        panic(err)
    }

    // DB initialization
    db.SetConnMaxLifetime(time.Minute * 3)
    db.SetMaxOpenConns(10)
    db.SetMaxIdleConns(10)

    // Routing
    r := mux.NewRouter()
    r.HandleFunc("/", routes.IndexHandler)
    http.Handle("/", r)

    s := &http.Server{
        Addr:   ":8000",
    }

    log.Fatal(s.ListenAndServe())
}
