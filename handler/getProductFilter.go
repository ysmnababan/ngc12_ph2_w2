package handler

import (
	"net/http"
	"ngc12/logger"
	"ngc12/model"

	"github.com/labstack/echo/v4"
)

func (r *Repo) GetProductsFilter(c echo.Context) error {

	//get queryparam
	key := c.QueryParam("filter")
	if key == "" {
		logger.Logging(c).Error("empty query")
		return c.JSON(http.StatusBadRequest, "filter is empty")
	}
	// filter := "%" + key
	filter := key + "%"
	var products []model.Product
	res := r.DB.Where("name LIKE ?", filter).Find(&products)
	if res.Error != nil {
		logger.Logging(c).Error(res.Error)
		return c.JSON(http.StatusInternalServerError, "Internal Server Error")
	}

	return c.JSON(http.StatusOK, products)
}
