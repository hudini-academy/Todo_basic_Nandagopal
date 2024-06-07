package main

import (
	"net/http"
	"github.com/bmizerany/pat"
)
func (app *application) routes() http.Handler{
	mux := pat.New()
	//basic cred application routes
	mux.Get("/",app.session.Enable(app.AuthenticateMiddleware( http.HandlerFunc(app.getTask))))
	mux.Post("/addtask",app.session.Enable(app.AuthenticateMiddleware(http.HandlerFunc(app.addTask))))
	mux.Post("/deletetask",app.session.Enable(app.AuthenticateMiddleware(http.HandlerFunc(app.deleteTask))))
	mux.Post("/updatetask",app.session.Enable(app.AuthenticateMiddleware(http.HandlerFunc(app.updateTask))))

	

	//adding new 5 routes
	mux.Get("/user/signup",app.session.Enable(http.HandlerFunc(app.signupUserForm)))
	mux.Post("/user/signup",app.session.Enable(http.HandlerFunc(app.signupUser)))
	mux.Get("/user/login",app.session.Enable(http.HandlerFunc(app.loginUserForm)))
	mux.Post("/user/login",app.session.Enable(http.HandlerFunc(app.loginUser)))
	mux.Post("/user/logout",app.session.Enable(http.HandlerFunc(app.logoutUser)))

	//add a new route for special tasks
	mux.Get("/user/special",app.session.Enable(http.HandlerFunc(app.specialUser)))
	mux.Post("/user/deletetask",app.session.Enable(app.AuthenticateMiddleware(http.HandlerFunc(app.deleteSpecial))))
	
	// Serve static files from the "./ui/static/" directory.
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))

	return app.logRequest(secureHeaders(mux))

}