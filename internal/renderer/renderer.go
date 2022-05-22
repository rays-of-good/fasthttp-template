package renderer

import (
	idatabase "github.com/rays-of-good/fasthttp-template/internal/database"
)

type (
	Renderer struct {
		Database *idatabase.Database
	}
)

func NewRenderer(d *idatabase.Database) (rr *Renderer) {
	rr = &Renderer{
		Database: d,
	}

	return
}
