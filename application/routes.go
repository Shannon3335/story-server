package application

import (
	aigen "github.com/shannon3335/story-server/services/aiGen"
	"github.com/shannon3335/story-server/services/auth"
)

func (app *App) LoadRoutes() {
	// any new routes should be added here, with business logic in handler files
	// app.loadTestRoutes()
	app.loadAuthRoutes()
	app.loadAIgenRoutes()

}

// func (app *App) loadTestRoutes() {
// 	t := &Test{}
// 	app.e.GET("/", t.HelloHandler)
// 	app.e.GET("/user/:id", t.GetHandler)
// 	app.e.POST("/make", t.MakeString)
// 	app.e.POST("/create", t.HandlePost)
// }

func (app *App) loadAuthRoutes() {
	authService := auth.NewAuthService(app.DB)
	a := auth.NewAuthHandler(authService)

	app.e.POST("/signup", a.Signup)
	app.e.POST("/login", a.Login)
}

func (app *App) loadAIgenRoutes() {
	// Pass api key here? and AI generator instance?

	aiService := aigen.NewAiService("testing123")
	a := aigen.NewAiHandler(aiService)

	app.e.POST("/generate", a.TestHello)

}
