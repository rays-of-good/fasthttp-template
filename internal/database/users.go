package database

import (
	"context"

	"github.com/rays-of-good/fasthttp-template/models"
)

func (database *Database) InsertUser(ctx context.Context, user *models.User) (err error) {
	c, err := database.Pool.Acquire(ctx)
	if err != nil {
		return
	}

	defer c.Release()

	err = c.QueryRow(
		ctx, "insert into users (telegram_id, telegram_username, role) values ($1, $2, $3) returning id, created_at, updated_at, deleted_at",
		user.TelegramUser.ID, user.TelegramUser.Username, user.Role,
	).Scan(
		&user.System.ID, &user.System.CreatedAt, &user.System.UpdatedAt, &user.System.DeletedAt,
	)

	return
}

func (database *Database) UpdateUser(ctx context.Context, user *models.User) (err error) {
	c, err := database.Pool.Acquire(ctx)
	if err != nil {
		return
	}

	defer c.Release()

	err = c.QueryRow(
		ctx, "update languages set telegram_id = $1, telegram_username = $2, role = $3 where id = $4 returning created_at, updated_at, deleted_at",
		user.TelegramUser.ID, user.TelegramUser.Username, user.Role, user.System.ID,
	).Scan(
		&user.System.CreatedAt, &user.System.UpdatedAt, &user.System.DeletedAt,
	)

	return
}

func (database *Database) SelectUsers(ctx context.Context) (users models.Users, err error) {
	c, err := database.Pool.Acquire(ctx)
	if err != nil {
		return
	}

	defer c.Release()

	rows, err := c.Query(ctx, "select id, created_at, updated_at, deleted_at, telegram_id, telegram_username, role from users order by updated_at desc")
	if err != nil {
		return
	}

	user := models.User{}

	for rows.Next() {
		err = rows.Scan(&user.System.ID, &user.System.CreatedAt, &user.System.UpdatedAt, &user.System.DeletedAt, &user.TelegramUser.ID, &user.TelegramUser.Username, &user.Role)
		if err != nil {
			continue
		}

		users = append(users, user)
	}

	return
}

func (database *Database) SelectUser(ctx context.Context, id string) (user models.User, err error) {
	c, err := database.Pool.Acquire(ctx)
	if err != nil {
		return
	}

	defer c.Release()

	err = c.QueryRow(
		ctx, "select id, created_at, updated_at, deleted_at, telegram_id, telegram_username, role from users where id = $1",
		id,
	).Scan(
		&user.System.ID, &user.System.CreatedAt, &user.System.UpdatedAt, &user.System.DeletedAt, &user.TelegramUser.ID, &user.TelegramUser.Username, &user.Role,
	)

	return
}

func (database *Database) DeleteUser(ctx context.Context, id string) (system models.System, err error) {
	c, err := database.Pool.Acquire(ctx)
	if err != nil {
		return
	}

	defer c.Release()

	err = c.QueryRow(
		ctx, "update users set deleted_at = current_timestamp where id = $1 returning id, created_at, updated_at, deleted_at",
		id,
	).Scan(
		&system.ID, &system.CreatedAt, &system.UpdatedAt, &system.DeletedAt,
	)

	return
}
