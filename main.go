package main

import (
    "database/sql"
    "fmt"
    _ "github.com/lib/pq"
)

func main() {
    connectionString := "postgres://postgres:aronmilu@localhost:5432/user_data"

    db, err := sql.Open("postgres", connectionString)
    if err != nil {
        fmt.Println(err)
        return
    }

    defer db.Close()

    // Get user input
    var name string
    fmt.Print("Enter your name: ")
    fmt.Scanf("%s", &name)

    var email string
    fmt.Print("Enter your email: ")
    fmt.Scanf("%s", &email)

    var password string
    fmt.Print("Enter your password: ")
    fmt.Scanf("%s", &password)

    // Create a new user
    stmt, err := db.Prepare("INSERT INTO users (name, email, password) VALUES ($1, $2, $3)")
    if err != nil {
        fmt.Println(err)
        return
    }

    _, err = stmt.Exec(name, email, password)
    if err != nil {
        fmt.Println(err)
        return
    }

    // Print success message
    fmt.Println("User created successfully")
}
