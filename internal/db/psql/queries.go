package psql

import (
	"context"
	"customer/internal/types"

	"github.com/jmoiron/sqlx"
)

type Query struct {
	db *sqlx.DB
}

func NewDB(db *sqlx.DB) *Query {
	return &Query{
		db: db,
	}
}

func (q *Query) CreateUser(ctx context.Context, data types.UserCustomer) error {
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

func (q *Query) GetOfficeUsers(ctx context.Context, data types.OfficeID) (types.UsersList, error) {
	query := `
		SELECT uuid, name, office_uuid, office_name, created_at
		FROM customer
		WHERE office_uuid = :office_uuid
	`

	var result types.UsersList

	if err := q.db.SelectContext(ctx, &result, query, data.OfficeUuID); err != nil {
		return result, err
	}

	return result, nil
}
