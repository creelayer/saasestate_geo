package handlers

import (
	"app/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

func ReverseHandler(c *gin.Context) {

	addressService := service.NewAddressService()

	lat, err := strconv.ParseFloat(c.Query("lat"), 32)

	if  err != nil {
		c.JSON(200, "Invalid Lat"+err.Error())
		return
	}

	lng, err := strconv.ParseFloat(c.Query("lng"), 32)

	if  err != nil {
		c.JSON(200, "Invalid Lng"+err.Error())
		return
	}

	locations := addressService.Reverse(lat, lng)

	c.JSON(200, locations)

}
