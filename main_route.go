package main

import (
	"food-delivery-service/component"
	restaurantgin "food-delivery-service/module/restaurant/transport/gin"
	"github.com/gin-gonic/gin"
)

func mainRoute(router *gin.Engine, appCtx component.AppContext) {
	v1 := router.Group("/v1")
	{
		restaurants := v1.Group("/restaurants")
		{
			restaurants.POST("", restaurantgin.CreateRestaurantHandler(appCtx))
			restaurants.GET("", restaurantgin.ListRestaurant(appCtx))
			restaurants.GET("/:restaurant_id", restaurantgin.GetRestaurantHandler(appCtx))
			restaurants.PUT("/:restaurant_id", restaurantgin.UpdateRestaurantHandler(appCtx))
			restaurants.DELETE("/:restaurant_id", restaurantgin.DeleteRestaurantHandler(appCtx))
		}
	}
}
