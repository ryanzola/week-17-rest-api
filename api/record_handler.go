package api

import (
	"encoding/json"
	"net/http"

	"github.com/ryanzola/week-17/db"
	"github.com/ryanzola/week-17/types"
)

type RecordHandler struct {
	store db.RecordStore
}

func NewRecordHandler(store db.RecordStore) *RecordHandler {
	return &RecordHandler{
		store: store,
	}
}

func (h *RecordHandler) HandleGetRecords(w http.ResponseWriter, r *http.Request) error {
	var params types.FilterParamsRequest
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		resp := types.ErrorResponse{
			Code:    2,
			Message: err.Error(),
		}

		return WriteJSON(w, http.StatusBadRequest, resp)
	}

	if err := ValidateParams(params); err != nil {
		resp := types.ErrorResponse{
			Code:    3,
			Message: err.Error(),
		}

		return WriteJSON(w, http.StatusBadRequest, resp)
	}

	records, err := h.store.GetRecords(r.Context(), params)
	if err != nil {
		resp := types.ErrorResponse{
			Code:    4,
			Message: err.Error(),
		}

		return WriteJSON(w, http.StatusInternalServerError, resp)
	}

	resp := types.ResourceResponse{
		Code:    0,
		Message: "Success",
		Records: records,
	}

	return WriteJSON(w, http.StatusOK, resp)
}
