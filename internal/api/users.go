package api

import (
	"github.com/rays-of-good/fasthttp-template/models"

	"github.com/goccy/go-json"
	"github.com/valyala/fasthttp"
)

func (api *API) AddUser() (handler fasthttp.RequestHandler) {
	handler = func(ctx *fasthttp.RequestCtx) {
		user := models.User{}

		err := json.NewDecoder(ctx.RequestBodyStream()).Decode(&user)
		if err != nil {
			ctx.Error(err.Error(), fasthttp.StatusBadRequest)
			return
		}

		err = api.Database.InsertUser(ctx, &user)
		if err != nil {
			ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
			return
		}

		err = json.NewEncoder(ctx.Response.BodyWriter()).Encode(&user)
		if err != nil {
			ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
			return
		}
	}

	return
}

func (api *API) UpdateUser() (handler fasthttp.RequestHandler) {
	handler = func(ctx *fasthttp.RequestCtx) {
		user := models.User{}

		err := json.NewDecoder(ctx.RequestBodyStream()).Decode(&user)
		if err != nil {
			ctx.Error(err.Error(), fasthttp.StatusBadRequest)
			return
		}

		err = api.Database.UpdateUser(ctx, &user)
		if err != nil {
			ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
			return
		}

		err = json.NewEncoder(ctx.Response.BodyWriter()).Encode(&user)
		if err != nil {
			ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
			return
		}
	}

	return
}

func (api *API) GetUsers() (handler fasthttp.RequestHandler) {
	handler = func(ctx *fasthttp.RequestCtx) {
		coaches, err := api.Database.SelectUsers(ctx)
		if err != nil {
			ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
			return
		}

		err = json.NewEncoder(ctx.Response.BodyWriter()).Encode(coaches)
		if err != nil {
			ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
			return
		}
	}

	return
}

func (api *API) GetUser() (handler fasthttp.RequestHandler) {
	handler = func(ctx *fasthttp.RequestCtx) {
		coach, err := api.Database.SelectUser(ctx, ctx.UserValue("id").(string))
		if err != nil {
			ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
			return
		}

		err = json.NewEncoder(ctx.Response.BodyWriter()).Encode(&coach)
		if err != nil {
			ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
			return
		}
	}

	return
}

func (api *API) DeleteUser() (handler fasthttp.RequestHandler) {
	handler = func(ctx *fasthttp.RequestCtx) {
		system, err := api.Database.DeleteUser(ctx, ctx.UserValue("id").(string))
		if err != nil {
			ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
			return
		}

		err = json.NewEncoder(ctx.Response.BodyWriter()).Encode(&system)
		if err != nil {
			ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
			return
		}
	}

	return
}
