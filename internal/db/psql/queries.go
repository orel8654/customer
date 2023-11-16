package psql

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
)

type UserDB struct {
	UuID       string    `db:"uuid"`
	Name       string    `db:"name"`
	OfficeUuID string    `db:"office_uuid"`
	OfficeName string    `db:"office_name"`
	CreatedAt  time.Time `db:"created_at"`
}

type OfficeID struct {
	OfficeUuID string `db:"office_uuid"`
}

type UsersList struct {
	result []*UserDB
}

type Query struct {
	db *sqlx.DB
}

func NewDB(db *sqlx.DB) *Query {
	return &Query{
		db: db,
	}
}

func (q *Query) CreateUser(ctx context.Context, data UserDB) error {
	query := `
		INSERT INTO customer (uuid, name, office_uuid, office_name, created_at)
		VALUES(:uuid, :name, :office_uuid, :office_name, :created_at)
	`
	_, err := q.db.NamedExecContext(ctx, query, data)
	if err != nil {
		return err
	}
	return nil
}

func (q *Query) GetOfficeUsers(ctx context.Context, data OfficeID) (UsersList, error) {
	query := `
		SELECT uuid, name, office_uuid, office_name, created_at
		FROM customer
		WHERE office_uuid = :office_uuid
	`

	var result UsersList

	if err := q.db.SelectContext(ctx, &result, query, data.OfficeUuID); err != nil {
		return result, err
	}

	return result, nil
}
