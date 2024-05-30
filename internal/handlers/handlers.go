package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"GOncurrently-Calculator/internal/storage"
)

var (
	id = 1
)

func AddExpression(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "something went wrong... (resolving body)", http.StatusInternalServerError)
		return
	}

	exp := &storage.Expression{}
	err = json.Unmarshal(body, exp)
	if err != nil {
		http.Error(w, "something went wrong... (parsing body to json)", http.StatusInternalServerError)
		return
	}

	if exp.RemoveSpaces().IsInvalid() {
		http.Error(w, "invalid expression", http.StatusUnprocessableEntity)
		return
	}

	response := fmt.Sprintf("{\"id\": %d}", id)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(response))

	exp.Id = id
	id++

	storage.DB = append(storage.DB, exp)
}
