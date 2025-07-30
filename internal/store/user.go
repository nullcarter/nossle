package store

import (
	"context"
	"database/sql"
	"fmt"
)

type User struct {
	id          int
	username    string
	displayName string
}

type UserStore struct {
	db *sql.DB
}

func (us *UserStore) Initialize() error {
	_, err := us.db.ExecContext(context.Background(),
		`
		CREATE TABLE IF NOT EXISTS users (
  			id INTEGER PRIMARY KEY AUTOINCREMENT,
     		username TEXT UNIQUE NOT NULL,
       		display_name TEXT,
         	avatar_url TEXT,
          	created_at DATETIME DEFAULT CURRENT_TIMESTAMP
        );
		`)

	if err != nil {
		fmt.Printf("user table: %s", err)
		return err
	}

	fmt.Println("users table initialized.")

	return nil
}

func (us *UserStore) Create(ctx context.Context) error {
	return nil
}

func (us *UserStore) Get(ctx context.Context) error {
	return nil
}

func (us *UserStore) Update(ctx context.Context) error {
	return nil
}

func (us *UserStore) Delete(ctx context.Context) error {
	return nil
}
