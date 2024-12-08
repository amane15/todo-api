package main

import (
	"net/http"
	"time"

	"github.com/amane15/todo-api/internal/data"
)

func (app *application) listTodoHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("list of todos"))
}

func (app *application) showTodohandler(w http.ResponseWriter, r *http.Request) {
	todo := data.Todo{
		ID:        1,
		Task:      "Get api working",
		Status:    data.PENDING,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := app.writeJSON(w, http.StatusOK, envelope{"todo": todo}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
