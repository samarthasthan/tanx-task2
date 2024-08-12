package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/samarthasthan/tanx-task/internal/models"
)

func (h *Handlers) handleAddCommodity(c echo.Context) error {

	if c.Get("type") != "lender" {
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Unauthorized2"})
	}
	s := new(models.Commodity)
	if err := c.Bind(s); err != nil {
		return err
	}
	if err := c.Validate(s); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	res, err := h.controller.CreateComodity(c, s)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{"status": "error", "message": "Commodity could not be listed", "payload": nil})
	}

	return c.JSON(200, map[string]any{"status": "success", "message": "Commodity listed successfully", "payload": res})
}

func (h *Handlers) handleGetCommodities(c echo.Context) error {
	res, err := h.controller.GetCommodities(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{"status": "error", "message": "Commodities could not be fetched", "payload": nil})
	}
	return c.JSON(200, map[string]any{"status": "success", "message": "Available commodities fetched successfully", "payload": res})
}
