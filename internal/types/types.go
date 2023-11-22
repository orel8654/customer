package types

import "time"

type UserCustomer struct {
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
	Result []UserCustomer
}
