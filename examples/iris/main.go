package main

import (
	"github.com/kataras/iris"
	_ "github.com/chenhg5/go-admin/adapter/iris"
	"github.com/chenhg5/go-admin/engine"
	"github.com/chenhg5/go-admin/examples/datamodel"
	"github.com/chenhg5/go-admin/modules/config"
	"github.com/chenhg5/go-admin/plugins/admin"
	"github.com/chenhg5/go-admin/plugins/example"
)

func main() {
	app := iris.Default()

	eng := engine.Default()

	cfg := config.Config{
		DATABASE: []config.Database{
			{
				HOST:         "127.0.0.1",
				PORT:         "3306",
				USER:         "root",
				PWD:          "root",
				NAME:         "godmin",
				MAX_IDLE_CON: 50,
				MAX_OPEN_CON: 150,
				DRIVER:       "mysql",
			},
		},
		DOMAIN: "localhost",
		PREFIX: "admin",
		INDEX:  "/",
	}

	adminPlugin := admin.NewAdmin(datamodel.Generators)
	examplePlugin := example.NewExample()

	if err := eng.AddConfig(cfg).AddPlugins(adminPlugin, examplePlugin).Use(app); err != nil {
		panic(err)
	}

	app.Run(iris.Addr(":8099"))
}
