package controller

import (
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/samarthasthan/tanx-task/internal/database"
	"github.com/samarthasthan/tanx-task/internal/database/mysql/sqlc"
	"github.com/samarthasthan/tanx-task/internal/models"
)

func (c *Controller) CreateComodity(ctx echo.Context, com *models.Commodity) (*models.CommodityOut, error) {
	mysql := c.mysql.(*database.MySQL)
	redis := c.redis.(*database.Redis)
	dbCtx := ctx.Request().Context()
	userID := ctx.Get("id").(string)

	commId := uuid.New().String()

	err := mysql.Queries.CreateCommodity(dbCtx, sqlc.CreateCommodityParams{
		Commodityid: commId,
		Userid:      userID,
		Name:        com.Name,
		Description: com.Description,
		Price:       fmt.Sprintf("%f", com.Price),
		Category:    com.Category,
	})

	if err != nil {
		return nil, err
	}

	// Remove cached alerts from Redis
	err = redis.Del(dbCtx, "alerts").Err()
	if err != nil {
		log.Printf("Error removing alerts cache from Redis: %v", err)
		// Optionally, handle this error depending on your use case
	} else {
		log.Printf("Alert created successfully for user %s and cache invalidated", userID)
	}

	res := &models.CommodityOut{
		CommodityId: commId,
		Price:       com.Price,
		CreatedAt:   time.Now().String(),
	}

	return res, nil
}

func (c *Controller) GetCommodities(ctx echo.Context) ([]*models.CommodityResponse, error) {
	mysql := c.mysql.(*database.MySQL)
	dbCtx := ctx.Request().Context()
	// userID := ctx.Get("id").(string)
	// catParam:= ctx.QueryParam("item_category")

	commodities, err := mysql.Queries.GetCommodities(dbCtx)
	if err != nil {
		return nil, err
	}

	var commoditiesOut []*models.CommodityResponse

	for _, commodity := range commodities {
		commoditiesOut = append(commoditiesOut, &models.CommodityResponse{
			CommodityId: commodity.Commodityid,
			Price:       commodity.Price,
			CreatedAt:   commodity.Createdat.String(),
			Category:    commodity.Category,
		})

	}

	return commoditiesOut, nil

}
