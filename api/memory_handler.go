package api

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/ryanzola/week-17/db"
)

type MemoryHandler struct {
	MemoryDB *db.MemoryDB
}

func NewMemoryHandler() *MemoryHandler {
	return &MemoryHandler{
		MemoryDB: db.NewMemory(),
	}
}

func (h MemoryHandler) HandleGetMemoryRecords(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	if key == "" {
		err := fmt.Errorf("missing key")
		slog.Error("cannot get value in memory db", slog.String("error", err.Error()))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if record, ok := h.MemoryDB.Get(key); ok {
		WriteJSON(w, http.StatusOK, record)
		return
	}

	err := fmt.Errorf("key not found")
	slog.Error("cannot get value in memory db", slog.String("error", err.Error()))
	http.Error(w, err.Error(), http.StatusNotFound)
}

func (h MemoryHandler) HandlePostMemoryRecord(w http.ResponseWriter, r *http.Request) {
	var record db.MemoryRecord
	if err := json.NewDecoder(r.Body).Decode(&record); err != nil {
		slog.Error("failed to get request body", slog.String("error", err.Error()))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	if err := h.MemoryDB.Add(record); err != nil {
		slog.Error("failed to add a new record", slog.String("error", err.Error()))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	WriteJSON(w, http.StatusCreated, record)
}
