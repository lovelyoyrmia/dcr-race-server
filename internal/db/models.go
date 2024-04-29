// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"database/sql"
)

type UsersLocation struct {
	ID        int32           `json:"id"`
	Uid       string          `json:"uid"`
	Category  string          `json:"category"`
	Fullname  string          `json:"fullname"`
	Latitude  sql.NullFloat64 `json:"latitude"`
	Longitude sql.NullFloat64 `json:"longitude"`
	Altitude  sql.NullFloat64 `json:"altitude"`
	Timestamp sql.NullTime    `json:"timestamp"`
	Email     sql.NullString  `json:"email"`
}
