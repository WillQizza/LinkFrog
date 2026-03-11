package services

import (
	"context"

	"github.com/willqizza/linkfrog/backend/db"
)

func GetTotalUsers(ctx context.Context) (int, error) {
	var count int
	if err := db.DB.QueryRowContext(ctx, "SELECT COUNT(*) FROM users").Scan(&count); err != nil {
		return 0, err
	}

	return count, nil
}

func GetUserIdByEmail(ctx context.Context, email string) (int, error) {
	var userId int
	if err := db.DB.QueryRowContext(ctx, "SELECT id FROM users WHERE email = ?", email).Scan(&userId); err != nil {
		return 0, err
	}

	return int(userId), nil
}

func WhitelistUser(ctx context.Context, email string) (int, error) {
	insertResults, err := db.DB.ExecContext(ctx, "INSERT INTO users (email) VALUES (?)", email)
	if err != nil {
		return 0, err
	}

	insertId, err := insertResults.LastInsertId()
	if err != nil {
		return 0, err
	}

	// Using int here isn't a problem bc the database is 32 bits for the user id
	return int(insertId), nil
}
