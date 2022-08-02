package restaurantgin

import (
	"food-delivery-service/common"
	"food-delivery-service/component"
	restaurantbiz "food-delivery-service/module/restaurant/biz"
	restaurantmodel "food-delivery-service/module/restaurant/model"
	restaurantstorage "food-delivery-service/module/restaurant/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UpdateRestaurantHandler(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data restaurantmodel.RestaurantUpdate

		//id, err := strconv.Atoi(c.Param("restaurant_id"))
		uid, err := common.FromBase58(c.Param("restaurant_id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		storage := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := restaurantbiz.NewUpdateRestaurantBiz(storage)

		if err := biz.UpdateRestaurantById(c.Request.Context(), int(uid.GetLocalID()), &data); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
