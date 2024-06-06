package main

import (
	"net/http"
	"github.com/bmizerany/pat"
)
func (app *application) routes() http.Handler{
	mux := pat.New()
	mux.Get("/",app.session.Enable(http.HandlerFunc(app.getTask)))
	mux.Post("/addtask",app.session.Enable(http.HandlerFunc(app.addTask)))
	mux.Post("/deletetask",app.session.Enable(http.HandlerFunc(app.deleteTask)))
	mux.Post("/updatetask",app.session.Enable(http.HandlerFunc(app.updateTask)))

	//adding new 5 routes
	mux.Get("/user/signup",app.session.Enable(http.HandlerFunc(app.signupUserForm)))
	mux.Post("/user/signup",app.session.Enable(http.HandlerFunc(app.signupUser)))
	mux.Get("/user/login",app.session.Enable(http.HandlerFunc(app.loginUserForm)))
	mux.Post("/user/login",app.session.Enable(http.HandlerFunc(app.loginUser)))
	mux.Post("/user/logout",app.session.Enable(http.HandlerFunc(app.logoutUser)))
	// Serve static files from the "./ui/static/" directory.
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))

	return app.logRequest(secureHeaders(mux))

}