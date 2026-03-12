package services

import (
	"context"

	"github.com/willqizza/linkfrog/backend/db"
	"github.com/willqizza/linkfrog/backend/models"
)

func GetTotalUsers(ctx context.Context) (int, error) {
	var count int
	err := db.DB.QueryRowContext(ctx, "SELECT COUNT(*) FROM users").Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func GetUserByID(ctx context.Context, id int) (*models.User, error) {
	var user models.User
	err := db.DB.QueryRowContext(ctx, "SELECT id, email FROM users WHERE id = ?", id).Scan(&user.ID, &user.Email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserIdByEmail(ctx context.Context, email string) (int, error) {
	var userId int
	err := db.DB.QueryRowContext(ctx, "SELECT id FROM users WHERE email = ?", email).Scan(&userId)
	if err != nil {
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
