package handlers

import (
	"app/service"
	"github.com/gin-gonic/gin"
)

func ReverseHandler(c *gin.Context) {

	addressService := service.NewAddressService()
	locations := addressService.Reverse(50.455472, 30.486827)

	c.JSON(200, locations)

}
