package oppgave8

import (
	"database/sql"
	"fmt"

	"github.com/fingann/swat/pkg/flags"

	_ "modernc.org/sqlite"
)

// This task is a SQL oppgave8 task,it contains a in-memory sqlite database and a login page.
// The database should contain two tables one for users with two users, admin and user, and another table "geodata" containing a flag.

func SetupDb() (*sql.DB, error) {
	db, err := sql.Open("sqlite", "file::memory:?cache=shared")
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	if err = CreateTables(db); err != nil {
		return nil, fmt.Errorf("failed to create tables: %w", err)
	}

	return db, nil
}

func CreateTables(db *sql.DB) error {
	_, err := db.Exec("CREATE TABLE users (id INTEGER PRIMARY KEY, username TEXT, password TEXT, message TEXT)")
	if err != nil {
		return fmt.Errorf("failed to create users table: %w", err)
	}

	_, err = db.Exec(fmt.Sprintf("INSERT INTO users (username, password, message) VALUES ('admin', 'why is this not hash?', '%s')", flags.Oppgave8Flag))
	if err != nil {
		return fmt.Errorf("failed to insert admin user: %w", err)
	}

	_, err = db.Exec("INSERT INTO users (username, password, message) VALUES ('user', 'user', 'Det er desverre ikke noe flagg her, kansje admin-brukeren har det?')")
	if err != nil {
		return fmt.Errorf("failed to insert user user: %w", err)
	}

	return nil
}

// Login is a function that is vulnerable to sql oppgave8 and returns the result
func Login(db *sql.DB, username, password string) (string, error) {
	var message string
	sql := fmt.Sprintf("SELECT message FROM users WHERE username='%s' AND password='%s'", username, password)
	fmt.Println("SQL: " + sql)
	err := db.QueryRow(sql).Scan(&message)
	if err != nil {
		return message, fmt.Errorf("failed to query database: %w", err)
	}

	return message, nil
}
