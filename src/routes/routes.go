package routes

import (
    "fmt"
    "net/http"
    "architalks/models"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Ping!~")
}

func getUser(r *http.Request) models.User {
    name := r.FormValue("name")
    email := r.FormValue("email")
    password := r.FormValue("password")

    return models.User{
        Name:    name,
        Email:   email,
        Password:    password,
    }
}
