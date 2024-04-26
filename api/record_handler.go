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

func (h *RecordHandler) GetRecords(c *fiber.Ctx) error {
	fmt.Println("GetRecords")
	return nil
}

func (h *RecordHandler) GetRecordByID(c *fiber.Ctx) error {
	fmt.Printf("GetRecordByID: %s\n", c.Params("id"))
	return nil
}

func (h *RecordHandler) InsertRecord(c *fiber.Ctx) error {
	var params types.Record
	if err := c.BodyParser(&params); err != nil {
		return err // error bad request
	}
	fmt.Printf("InsertRecord: %v\n", params)

	return nil
}

func (h *RecordHandler) UpdateRecord(c *fiber.Ctx) error {
	fmt.Printf("UpdateRecord: %s\n", c.Body())
	return nil
}

func (h *RecordHandler) DeleteRecord(c *fiber.Ctx) error {
	fmt.Printf("DeleteRecord: %s\n", c.Params("id"))
	return nil
}
