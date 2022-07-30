package restaurantgin

import (
	"context"
	restaurantbiz "food-delivery-service/module/restaurant/biz"
	restaurantmodel "food-delivery-service/module/restaurant/model"
	restaurantstorage "food-delivery-service/module/restaurant/storage"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type mockCreateStore struct{}

func (mockCreateStore) InsertRestaurant(ctx context.Context, data *restaurantmodel.RestaurantCreate) error {
	data.Id = 20
	return nil
}

func CreateRestaurantHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data restaurantmodel.RestaurantCreate

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		storage := restaurantstorage.NewSQLStore(db)
		//storage := &mockCreateStore{}
		biz := restaurantbiz.NewCreateRestaurantBiz(storage)

		if err := biz.CreateRestaurant(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": data.Id})
	}
}
