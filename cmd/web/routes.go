package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func (app *Config) routes() http.Handler {
	//create router
	mux := chi.NewRouter()
	//set up middleware
	mux.Use(middleware.Recoverer)
	mux.Use(app.SessionLoad)
	//define applications routes
	mux.Get("/", app.HomePage)

	mux.Get("/login", app.LoginPage)
	mux.Post("/login", app.PostLoginPage)
	mux.Get("/logout", app.Logout)
	mux.Get("/register", app.PostRegisterPage)
	mux.Get("/activate", app.ActivateAccount)
	mux.Get("/plans", app.ChooseSubscription)
	mux.Get("/subscribe", app.SubscribeToPlan)
	mux.Mount("/members", app.authRouter())

	mux.Get("/test-email", func(writer http.ResponseWriter, request *http.Request) {
		m := Mail{
			Domain:      "localhost",
			Host:        "localhost",
			Port:        1025,
			Encryption:  "none",
			FromAddress: "info@mycompany.com",
			FromName:    "info",
			ErrorChan:   make(chan error),
		}

		msg := Message{
			To:      "me@here.com",
			Subject: "Test email",
			Data:    "Hello, World",
		}
		m.sendMail(msg, make(chan error))
	})
	return mux
}

func (app *Config) authRouter() http.Handler {
	mux := chi.NewRouter()
	mux.Use(app.Auth)

	mux.Get("/plans", app.ChooseSubscription)
	mux.Get("/subscribe", app.SubscribeToPlan)

	return mux
}
