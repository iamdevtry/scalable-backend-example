package restaurantgin

import (
	"food-delivery-service/common"
	"food-delivery-service/component"
	restaurantbiz "food-delivery-service/module/restaurant/biz"
	restaurantstorage "food-delivery-service/module/restaurant/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetRestaurantHandler(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		//id, err := strconv.Atoi(c.Param("restaurant_id"))

		uid, err := common.FromBase58(c.Param("restaurant_id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		storage := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := restaurantbiz.NewFindRestaurantBiz(storage)

		data, err := biz.FindRestaurantById(c.Request.Context(), int(uid.GetLocalID()))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		data.Mask(true)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
