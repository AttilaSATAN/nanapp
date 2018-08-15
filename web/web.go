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
	// app.RegisterView(iris.HTML("./web/client/build", ".html"))

	// app.Get("/", func(ctx iris.Context) {

	// 	ctx.View("index.html")
	// })
	assetHandler := app.StaticHandler("./web/client/build", false, false)
	// as an alternative of SPA you can take a look at the /routing/dynamic-path/root-wildcard
	// example too
	app.SPA(assetHandler)
	mvc.Configure(app.Party("/api/campaign")).Handle(campaignApp)

	app.Run(iris.Addr(":8080"))

	return app
}
