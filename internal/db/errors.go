package db

import (
	"context"
	"errors"
)

var (
	ErrNil = errors.New("no matching record found in redis database")
	Ctx    = context.TODO()
)
