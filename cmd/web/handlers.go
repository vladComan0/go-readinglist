package main

import (
	"fmt"
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "The home page")
}

func (app *application) bookView(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "The view book page")
}

func (app *application) bookCreate(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "The create book page")
}
