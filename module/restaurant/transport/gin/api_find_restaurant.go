package restaurantgin

import (
	restaurantbiz "food-delivery-service/module/restaurant/biz"
	restaurantstorage "food-delivery-service/module/restaurant/storage"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func GetRestaurantHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("restaurant_id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		storage := restaurantstorage.NewSQLStore(db)
		biz := restaurantbiz.NewFindRestaurantBiz(storage)

		data, err := biz.FindRestaurantById(c.Request.Context(), id)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": data})
	}
}
