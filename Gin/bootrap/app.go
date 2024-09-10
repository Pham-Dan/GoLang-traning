package bootstrap

import (
	"main/models"

	"github.com/uptrace/bun"
)

type Application struct {
	Env   *Env
	DB  *bun.DB
}

func App() Application {
	app := &Application{}
	app.Env = NewEnv()
	app.DB = models.NewConnectDB()
	return *app
}

func (app *Application) CloseDBConnection() {
}
