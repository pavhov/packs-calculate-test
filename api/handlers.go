package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/pavhov/packs-calculate-test/models"
	"github.com/pavhov/packs-calculate-test/utils"
)

func CalculateHandler(c *gin.Context) {
	var order models.Order
	if err := c.ShouldBind(&order); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	packCounts := utils.CalculatePacks(&order)
	c.JSON(http.StatusOK, packCounts)
}
