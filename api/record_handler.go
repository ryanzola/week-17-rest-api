package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
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

func (h *RecordHandler) HandleGetRecords(c *fiber.Ctx) error {
	var params types.RecordQueryParams
	if err := c.BodyParser(&params); err != nil {
		return err
	}

	records, err := h.store.GetRecords(c.Context(), params)
	if err != nil {
		return err
	}

	resp := types.ResourceResponse{
		Code:    0,
		Message: "Success",
		Records: records,
	}

	return c.JSON(resp)
}

func (h *RecordHandler) HandleGetRecordByID(c *fiber.Ctx) error {
	fmt.Printf("GetRecordByID: %s\n", c.Params("id"))
	return nil
}

func (h *RecordHandler) HandleInsertRecord(c *fiber.Ctx) error {
	var params types.Record
	if err := c.BodyParser(&params); err != nil {
		return err // error bad request
	}
	fmt.Printf("InsertRecord: %+v\n", params)

	return nil
}

func (h *RecordHandler) HandleUpdateRecord(c *fiber.Ctx) error {
	fmt.Printf("UpdateRecord: %s\n", c.Body())
	return nil
}

func (h *RecordHandler) HandleDeleteRecord(c *fiber.Ctx) error {
	fmt.Printf("DeleteRecord: %s\n", c.Params("id"))
	return nil
}
