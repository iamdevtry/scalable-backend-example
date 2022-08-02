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

func ListRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			panic(err)
		}

		var filter restaurantmodel.Filter

		if err := c.ShouldBind(&filter); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		_ = paging.Validate()

		store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := restaurantbiz.NewListRestaurantBiz(store)

		result, err := biz.ListRestaurant(c.Request.Context(), &filter, &paging)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		for i := range result {
			result[i].Mask(true)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}
