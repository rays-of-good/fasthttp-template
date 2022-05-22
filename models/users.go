package models

import (
	"github.com/jackc/pgtype"
)

type (
	TelegramUser struct {
		ID       pgtype.Varchar
		Username pgtype.Varchar
	}

	User struct {
		Database     Database
		TelegramUser TelegramUser

		Role pgtype.EnumType
	}

	Users []User
)
