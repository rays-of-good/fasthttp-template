package api

import (
	lmodels "github.com/rays-of-good/fasthttp-template/models"

	"github.com/goccy/go-json"
	"github.com/jackc/pgx/v4"
	"github.com/valyala/fasthttp"

	aheaders "github.com/go-asphyxia/http/headers"
	amime "github.com/go-asphyxia/mime"
)

func (a *API) GetUsers(ctx *fasthttp.RequestCtx) {
	c, err := a.Database.Pool.Acquire(ctx)
	if err != nil {
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
		return
	}

	defer c.Release()

	rows, err := c.Query(ctx, "select id, created_at, updated_at, deleted_at, user_role from users order by created_at desc")
	if err != nil {
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
		return
	}

	us := lmodels.Users{}
	u := lmodels.User{}

	for rows.Next() {
		err = rows.Scan(&u.Database.ID, &u.Database.CreatedAt, &u.Database.UpdatedAt, &u.Database.DeletedAt, &u.Role)
		if err != nil {
			continue
		}

		us = append(us, u)
	}

	b, err := json.Marshal(us)
	if err != nil {
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
		return
	}

	headers := &ctx.Response.Header

	headers.Set(aheaders.ContentType, amime.ApplicationJSON)

	ctx.Write(b)
}

func (a *API) GetUser(ctx *fasthttp.RequestCtx) {
	c, err := a.Database.Pool.Acquire(ctx)
	if err != nil {
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
		return
	}

	defer c.Release()

	u := lmodels.User{}

	err = u.Database.ID.Set(ctx.UserValue("id").(string))
	if err != nil {
		ctx.Error(err.Error(), fasthttp.StatusBadRequest)
		return
	}

	err = c.QueryRow(
		ctx,
		"select created_at, updated_at, deleted_at, user_role from users where id = $1",
		u.Database.ID,
	).Scan(
		&u.Database.CreatedAt,
		&u.Database.UpdatedAt,
		&u.Database.DeletedAt,
		&u.Role,
	)

	switch err {
	case pgx.ErrNoRows:
		ctx.Error(err.Error(), fasthttp.StatusNotFound)
	case nil:
		b, err := json.Marshal(&u)
		if err != nil {
			ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
			return
		}

		headers := &ctx.Response.Header

		headers.Set(aheaders.ContentType, amime.ApplicationJSON)

		ctx.Write(b)
	default:
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
	}
}

func (a *API) UpdateUser(ctx *fasthttp.RequestCtx) {
	u := lmodels.User{}

	err := json.Unmarshal(ctx.Request.Body(), &u)
	if err != nil {
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
		return
	}

	c, err := a.Database.Pool.Acquire(ctx)
	if err != nil {
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
		return
	}

	defer c.Release()

	err = c.QueryRow(ctx, "update users set user_role = $1 where id = $2 returning user_role", u.Role, u.Database.ID).Scan(&u.Role)

	switch err {
	case pgx.ErrNoRows:
		ctx.Error(err.Error(), fasthttp.StatusNotFound)
	case nil:
		b, err := json.Marshal(&u)
		if err != nil {
			ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
			return
		}

		headers := &ctx.Response.Header

		headers.Set(aheaders.ContentType, amime.ApplicationJSON)

		ctx.Write(b)
	default:
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
	}
}

func (a *API) DeleteUser(ctx *fasthttp.RequestCtx) {
	d := lmodels.Database{}

	err := d.ID.Set(ctx.UserValue("id").(string))
	if err != nil {
		ctx.Error(err.Error(), fasthttp.StatusBadRequest)
		return
	}

	c, err := a.Database.Pool.Acquire(ctx)
	if err != nil {
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
		return
	}

	defer c.Release()

	err = c.QueryRow(ctx, "update users set deleted_at = current_timestamp where id = $1 returning created_at, updated_at, deleted_at", d.ID).Scan(&d.CreatedAt, &d.UpdatedAt, &d.DeletedAt)

	switch err {
	case pgx.ErrNoRows:
		ctx.Error(err.Error(), fasthttp.StatusNotFound)
	case nil:
		b, err := json.Marshal(&d)
		if err != nil {
			ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
		}

		headers := &ctx.Response.Header

		headers.Set(aheaders.ContentType, amime.ApplicationJSON)

		ctx.Write(b)
	default:
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
	}
}
