package session

import (
	"context"
	"cookdroogers/app"
	"cookdroogers/models"
	"github.com/pkg/errors"
)

func CheckAdmin(username, password string, app *app.App) bool {
	return app.Config.Root.Username == username && app.Config.Root.Password == password
}

func LoginAdmin(ctx context.Context, app *app.App) error {

	err := app.Services.UserService.SetRole(ctx, models.AdminUser)
	if err != nil {
		return errors.Wrap(ErrCantSetAdminRole, err.Error())
	}

	return nil
}
