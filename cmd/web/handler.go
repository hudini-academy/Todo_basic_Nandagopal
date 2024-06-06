package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
	"todo/pkg/models"
)

// getTask handles the home page and displays the tasks.
func (app *application) getTask(w http.ResponseWriter, r *http.Request) {

	// Get all records from the db
	s, err := app.todo.GetAll()
	if err != nil {
		app.errorLog.Println(err.Error())
		// http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	for _, item := range s {
		fmt.Println(item)
	}

	// Define the template files to parse.
	files := []string{
		"./ui/html/home.page.tmpl",
	}

	// Parse the template files.
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.errorLog.Println(err.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	// Execute the template with the list of tasks.
	err = ts.Execute(w, struct{
		Tasks []*models.Todos
		Flash string
	}{
		Tasks: s,
		Flash: app.session.PopString(r,"flash"),
	})
	if err != nil {
		app.errorLog.Println(err.Error())
		// http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	app.infoLog.Println("Displayed home page")
}

// addTask handles adding a new task.
func (app *application) addTask(w http.ResponseWriter, r *http.Request) {

	taskName := r.FormValue("task")

	if len(strings.TrimSpace(taskName)) != 0 {
		// Add the entry to the db
		_, err := app.todo.Insert(taskName)
		if err != nil {
			// app.errorLog("Hello")
			fmt.Println(err)
		}
	app.session.Put(r, "flash", "Task Added successfull!")
	}else{
	app.session.Put(r, "flash", "item cannot be  empty!")

	}
	// Redirect back to the home page.
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// deleteTask handles deleting a task.
func (app *application) deleteTask(w http.ResponseWriter, r *http.Request) {

	// Get the ID of the task to delete from the form data.
	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		app.errorLog.Println("Invalid task ID")
		http.Error(w, "invalid task ID", http.StatusBadRequest)
		return
	}

	errDel := app.todo.Remove(id)
	if errDel != nil {
		log.Println(err)
		return
	}
	app.session.Put(r, "flash", "Task deleted successfull!")

	// Redirect back to the home page.
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

//update task handles user input to side of each task and it gets updated

func (app *application) updateTask(w http.ResponseWriter, r *http.Request) {

	// TODO: First check if the id given is present in the DB, if not, send an Not Found Back

	id, _ := strconv.Atoi(r.FormValue("id"))

	//call the function checkifxists to check if id is present
	doesExist, err := app.todo.CheckIfExists(id)
	if err != nil {
		app.errorLog.Println(err.Error())
		return
	}

	if doesExist {
		taskName:=r.FormValue("updatetask")
		if len(strings.TrimSpace(taskName)) != 0{
		_, err := app.todo.Update(id,taskName)
		if err != nil {
			fmt.Println(err)
		}}
		app.session.Put(r, "flash", "Task Updated successfull!")

	} else {
		http.Error(w, "Data not found", http.StatusNotModified)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}


func (app *application) signupUserForm(w http.ResponseWriter, r *http.Request){
	// fmt.Fprintln(w, "Display the user signup form...")
	files := []string{
		"./ui/html/signup.page.tmpl",
		"./ui/html/base.layout.tmpl",
	}

	// Parse the template files.
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.errorLog.Println(err.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	ts.Execute(w,nil)

}

func (app *application) signupUser(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(w, "Create a new user...")
	userName := r.FormValue("name")
	userEmail := r.FormValue("email")
	userPassword := r.FormValue("password")

	err := app.users.Insert(userName,userEmail,userPassword)
		if err != nil {
			// app.errorLog("Hello")
			fmt.Println(err)
		}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (app *application) loginUserForm(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(w, "Display the user login form...")
	files := []string{
		"./ui/html/login.page.tmpl",
		"./ui/html/base.layout.tmpl",
	}

	// Parse the template files.
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.errorLog.Println(err.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	ts.Execute(w,app.session.PopString(r, "flash"))
}

func (app *application) loginUser(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(w, "Authenticate and login the user...")
	userEmail := r.FormValue("email")
	userPassword:= r.FormValue("password")
	isUser, err := app.users.Authenticate(userEmail,userPassword)
	log.Print(isUser)
	if err != nil {
		app.errorLog.Println(err.Error())
		http.Error(w, "internal server error", 500)
		return
	}
	if isUser{
		app.session.Put(r, "Authenticated", true)
		app.session.Put(r,"flash","Login Successfully")
		http.Redirect(w,r,"/", http.StatusSeeOther)	
	}else{
		app.session.Put(r, "flash", "login failed")
		app.session.Put(r,"Authenticated",true)
		http.Redirect(w,r,"/user/login", http.StatusSeeOther)	
	}
}

func (app *application) logoutUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Logout the user...")
}
