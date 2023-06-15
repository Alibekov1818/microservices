package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)
	router.HandlerFunc(http.MethodGet, "/phones", app.getPhones)
	router.HandlerFunc(http.MethodPost, "/phones", app.requirePermission("write", app.addPhone))
	router.HandlerFunc(http.MethodGet, "/phones/:id", app.getPhone)
	router.HandlerFunc(http.MethodDelete, "/phones/:id", app.requirePermission("write", app.deletePhone))
	router.HandlerFunc(http.MethodGet, "/computers", app.getComputers)
	router.HandlerFunc(http.MethodPost, "/computers", app.requirePermission("write", app.addComputer))
	router.HandlerFunc(http.MethodGet, "/computers/:id", app.getComputer)
	router.HandlerFunc(http.MethodDelete, "/computers/:id", app.requirePermission("write", app.deleteComputer))
	router.HandlerFunc(http.MethodDelete, "/users/:id", app.requirePermission("write", app.deleteUser))
	router.HandlerFunc(http.MethodPost, "/register", app.registerUser)
	router.HandlerFunc(http.MethodPost, "/authenticate", app.createToken)

	return app.recoverPanic(app.enableCORS(app.rateLimit(app.authenticate(router))))
}
