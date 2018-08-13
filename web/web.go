package web

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/mongodb/mongo-go-driver/mongo"
)

var (
	collection  *mongo.Collection
	campaignApp *Campaign
)

//New creates a new iris application
func New(c *mongo.Collection) *iris.Application {
	collection = c
	campaignApp = &Campaign{collection}

	return initApp()
}

func initApp() *iris.Application {

	app := iris.New()
	app.RegisterView(iris.HTML("./web/client", ".html"))

	app.Get("/", func(ctx iris.Context) {

		ctx.View("index.html")
	})
	mvc.Configure(app.Party("/api/campaign")).Handle(campaignApp)

	app.Run(iris.Addr(":8080"))

	return app
}
