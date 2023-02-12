package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"task4/cmd/web/repository"
	"task4/dto"
)

type Handler struct {
	Repo repository.Repository
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)

	if err := json.NewEncoder(w).Encode(dto.NewCommonResponse(data)); err != nil {
		panic(err)
	}
}

func (h *Handler) GetItems(w http.ResponseWriter, r *http.Request) {
	log.Println("Request Received For Get Items")

	items, err := h.Repo.GetItems()
	if err != nil {
		writeResponse(w, 404, err.Error())
		return
	}
	writeResponse(w, 200, items)
}
func (h *Handler) GetItemById(w http.ResponseWriter, r *http.Request) {
	log.Println("Request Received For Get Item By Id")
	vars := mux.Vars(r)
	itemId := vars["id"]

	items, err := h.Repo.GetItemById(itemId)
	if err != nil {
		writeResponse(w, 404, err.Error())
		return
	}
	writeResponse(w, 200, items)
}
func (h *Handler) UpdateItem(w http.ResponseWriter, r *http.Request) {
	log.Println("Request Received For Update Item")
	vars := mux.Vars(r)
	itemId := vars["id"]

	var itemRequest dto.ItemRequest
	err := json.NewDecoder(r.Body).Decode(&itemRequest)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	err = h.Repo.UpdateItem(itemId, itemRequest.ToModel())
	if err != nil {
		writeResponse(w, 404, err.Error())
		return
	}
	writeResponse(w, 200, "Updated")
}

func (h *Handler) DeleteItem(w http.ResponseWriter, r *http.Request) {
	log.Println("Request Received For Delete Item")
	vars := mux.Vars(r)
	itemId := vars["id"]
	err := h.Repo.DeleteItem(itemId)
	if err != nil {
		writeResponse(w, 404, err.Error())
		return
	}
	writeResponse(w, 200, "Success")
}
