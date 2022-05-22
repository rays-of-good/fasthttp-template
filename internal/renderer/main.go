package renderer

import (
	tindex "github.com/rays-of-good/fasthttp-template/templates/index"

	"github.com/valyala/fasthttp"
)

func (rr *Renderer) Main() (handler fasthttp.RequestHandler) {
	handler = func(ctx *fasthttp.RequestCtx) {
		tindex.WriteCode(ctx, tindex.MainData{Title: "Hello", Text: "World"})
	}

	return
}
