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
		System System

		TelegramUser TelegramUser

		Role pgtype.Varchar
	}

	Users []User
)
