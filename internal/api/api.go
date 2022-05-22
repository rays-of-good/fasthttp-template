package api

import (
	idatabase "github.com/rays-of-good/fasthttp-template/internal/database"
)

type (
	API struct {
		Database *idatabase.Database
	}
)

func NewAPI(d *idatabase.Database) (a *API) {
	a = &API{
		Database: d,
	}

	return
}
