package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"app/dto"
	"app/service"
)

func ReverseHandler(c *gin.Context) {

	var query struct {
		Lat float64 `form:"lat" json:"lat" binding:"required"`
		Lng float64 `form:"lng" json:"lng" binding:"required"`
	}

	if err := c.ShouldBind(&query); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	addressService := service.NewAddressService()

	locations := addressService.Reverse(query.Lat, query.Lng)

	c.JSON(200, dto.NewLocationResponse(locations))

}
