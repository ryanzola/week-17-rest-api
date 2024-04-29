package api

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	"github.com/ryanzola/week-17/types"
)

func ValidateParams(params types.RecordQueryParams) error {
	var err error
	if params.StartDate == "" {
		err = errors.New("missing start date")
	}

	if params.EndDate == "" {
		err = errors.New("missing end date")
	}

	if params.MinCount == 0 {
		err = errors.New("missing min count")
	}

	if params.MaxCount == 0 {
		err = errors.New("missing max count")
	}

	return err
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

func Make(h func(http.ResponseWriter, *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			slog.Error("internal server error", "err", err, "path", r.URL.Path)
		}
	}
}
