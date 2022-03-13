package main

import (
    "time"
    "log"
    "net/http"
    "architalks/routes"
    "database/sql"
    "github.com/gorilla/mux"
    _ "github.com/go-sql-driver/mysql"
)

func main() {
    db, err := sql.Open("mysql", "user:password@/database")

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
