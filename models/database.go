package models

import (
	"github.com/jackc/pgtype"
)

type (
	System struct {
		ID pgtype.UUID

		CreatedAt pgtype.Timestamptz
		UpdatedAt pgtype.Timestamptz
		DeletedAt pgtype.Timestamptz
	}
)
