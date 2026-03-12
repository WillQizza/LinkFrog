package services

import (
	"context"
	"database/sql"

	"github.com/willqizza/linkfrog/backend/db"
	"github.com/willqizza/linkfrog/backend/models"
)

func GetLinksByUser(ctx context.Context, user *models.User) ([]*models.Link, error) {
	rows, err := db.DB.QueryContext(ctx, "SELECT id, owner, path, url FROM links WHERE owner = ?", user.ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var links []*models.Link
	for rows.Next() {
		var link models.Link
		err := rows.Scan(&link.ID, &link.Owner, &link.Path, &link.URL)
		if err != nil {
			return nil, err
		}

		links = append(links, &link)
	}

	return links, rows.Err()
}

func CreateLink(ctx context.Context, link *models.Link) error {
	_, err := db.DB.ExecContext(ctx, "INSERT INTO links (owner, path, url) VALUES (?, ?, ?)", link.Owner, link.Path, link.URL)
	return err
}

func DeleteLink(ctx context.Context, code string) error {
	_, err := db.DB.ExecContext(ctx, "DELETE FROM links WHERE path = ?", code)
	return err
}

func DoesCodeExist(ctx context.Context, code string) (bool, error) {
	var count int
	err := db.DB.QueryRowContext(ctx, "SELECT COUNT(*) FROM links WHERE path = ?", code).Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func GetLinkByCode(ctx context.Context, code string) (*models.Link, error) {
	var link models.Link
	err := db.DB.QueryRowContext(ctx, "SELECT id, owner, path, url FROM links WHERE path = ?", code).Scan(&link.ID, &link.Owner, &link.Path, &link.URL)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &link, nil
}
