package handler

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Order struct{}

func (o *Order) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create an order")
}

func (o *Order) List(w http.ResponseWriter, r *http.Request) {
	fmt.Println("List all order")
}

func (o *Order) GetByID(w http.ResponseWriter, r *http.Request) {
	ID := chi.URLParam(r, "id")
	fmt.Println("Get an order by ID", ID)
}

func (o *Order) UpdateByID(w http.ResponseWriter, r *http.Request) {
	ID := chi.URLParam(r, "id")
	fmt.Println("Update an order by ID: ", ID)
}

func (o *Order) DeleteByID(w http.ResponseWriter, r *http.Request) {
	ID := chi.URLParam(r, "id")
	fmt.Println("Delete an order by ID", ID)
}

