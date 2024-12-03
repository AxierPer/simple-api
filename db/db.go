package db

import (
  "database/sql"
  "log"
  "simple-api/types"

  _ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func ConnectDB() {
    var err error
    DB, err = sql.Open("sqlite3", "./users.db")
    if err != nil {
        log.Fatal(err)
    }
}

func InitDB() {
    _, err := DB.Exec(`CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        first_name TEXT,
        last_name TEXT,
        email TEXT
    )`)
    if err != nil {
        log.Fatal(err)
    }
}

func InsertUser(firstName, lastName, email string) (int64, error) {
    stmt, err := DB.Prepare("INSERT INTO users(first_name, last_name, email) VALUES (?, ?, ?)")
    if err != nil {
        return 0, err
    }
    defer stmt.Close()

    res, err := stmt.Exec(firstName, lastName, email)
    if err != nil {
        return 0, err
    }

    return res.LastInsertId()
}

func GetUsers() ([]types.User, error) {
    rows, err := DB.Query("SELECT id, first_name, last_name, email FROM users")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var users []types.User
    for rows.Next() {
        var user types.User
        err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email)
        if err != nil {
            return nil, err
        }
        users = append(users, user)
    }

    return users, nil
}

func GetUser(id int) (types.User, error) {
    var user types.User
    err := DB.QueryRow("SELECT id, first_name, last_name, email FROM users WHERE id = ?", id).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email)
    if err != nil {
        return user, err
    }
    return user, nil
}

func UpdateUser(id int, firstName, lastName, email string) error {
    stmt, err := DB.Prepare("UPDATE users SET first_name = ?, last_name = ?, email = ? WHERE id = ?")
    if err != nil {
        return err
    }
    defer stmt.Close()

    _, err = stmt.Exec(firstName, lastName, email, id)
    return err
}

func DeleteUser(id int) error {
    stmt, err := DB.Prepare("DELETE FROM users WHERE id = ?")
    if err != nil {
        return err
    }
    defer stmt.Close()

    _, err = stmt.Exec(id)
    return err
}
