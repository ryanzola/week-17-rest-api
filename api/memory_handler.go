package api

import (
	"encoding/json"
	"net/http"

	"github.com/ryanzola/week-17/types"
)

var MemoryDB = map[string]types.Memory{}

func HandleMemoryRecords(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case http.MethodGet:
		return HandleGetMemoryRecords(w, r)
	case http.MethodPost:
		return HandlePostMemoryRecord(w, r)
	default:
		return HandleMethodNotAllowed(w, r)
	}
}

func HandleGetMemoryRecords(w http.ResponseWriter, r *http.Request) error {
	key := r.URL.Query().Get("key")
	if key == "" {
		resp := types.ErrorResponse{
			Code:    4,
			Message: "Key is required",
		}

		return WriteJSON(w, http.StatusBadRequest, resp)
	}

	memory, ok := MemoryDB[key]
	if !ok {
		resp := types.ErrorResponse{
			Code:    3,
			Message: "Key not found",
		}

		return WriteJSON(w, http.StatusNotFound, resp)
	}

	resp := &types.MemoryResponse{
		Key:   memory.Key,
		Value: memory.Value,
	}

	return WriteJSON(w, http.StatusOK, resp)
}

func HandlePostMemoryRecord(w http.ResponseWriter, r *http.Request) error {
	var params types.MemoryRequestParams
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		resp := types.ErrorResponse{
			Code:    2,
			Message: err.Error(),
		}

		return WriteJSON(w, http.StatusBadRequest, resp)
	}

	if params.Key == "" {
		resp := types.ErrorResponse{
			Code:    3,
			Message: "Key is required",
		}

		return WriteJSON(w, http.StatusBadRequest, resp)
	}

	if params.Value == "" {
		resp := types.ErrorResponse{
			Code:    3,
			Message: "Value is required",
		}

		return WriteJSON(w, http.StatusBadRequest, resp)
	}

	memory := types.Memory(params)

	MemoryDB[params.Key] = memory

	resp := &types.MemoryResponse{
		Key:   params.Key,
		Value: params.Value,
	}

	return WriteJSON(w, http.StatusCreated, resp)
}

func HandleMethodNotAllowed(w http.ResponseWriter, r *http.Request) error {
	resp := types.ErrorResponse{
		Code:    1,
		Message: "Method not allowed",
	}

	return WriteJSON(w, http.StatusMethodNotAllowed, resp)
}
