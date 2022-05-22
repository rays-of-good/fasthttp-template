package api

import (
	"github.com/valyala/fasthttp"

	aheaders "github.com/go-asphyxia/http/headers"
)

func (a *API) CORS(ctx *fasthttp.RequestCtx) {
	headers := &ctx.Response.Header

	headers.Set(aheaders.AccessControlAllowOrigin, "*")
	headers.Set(aheaders.AccessControlAllowHeaders, "Authorization, Accept, Content-Type")
	headers.Set(aheaders.AccessControlAllowMethods, "GET, POST, PUT, DELETE, OPTIONS")
}
