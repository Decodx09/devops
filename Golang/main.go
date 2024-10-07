package main

import (
    "database/sql"
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    _ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
    var err error
	db, err = sql.Open("mysql", "root:root@tcp(mysql:3306)/testdb")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Ensure the connection is established
    if err := db.Ping(); err != nil {
        log.Fatal(err)
    }

    fmt.Println("Connected to the database!")

    // Set up HTTP routes
    http.HandleFunc("/create", createRecord)
    http.HandleFunc("/read", readRecords)

    log.Fatal(http.ListenAndServe(":8080", nil))
}

func createRecord(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
        var data struct {
            Name string `json:"name"`
        }
        if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        _, err := db.Exec("INSERT INTO your_table (name) VALUES (?)", data.Name)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        w.Write([]byte("Record added!"))
    }
}

func readRecords(w http.ResponseWriter, r *http.Request) {
    rows, err := db.Query("SELECT name FROM your_table")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var names []string
    for rows.Next() {
        var name string
        if err := rows.Scan(&name); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        names = append(names, name)
    }
    json.NewEncoder(w).Encode(names)
}
