package main

import (
	"food-delivery-service/component"
	"food-delivery-service/middleware"
	restaurantgin "food-delivery-service/module/restaurant/transport/gin"
	userstorage "food-delivery-service/module/user/storage"
	usergin "food-delivery-service/module/user/transport/gin"
	"github.com/gin-gonic/gin"
	"net/http"
)

func mainRoute(router *gin.Engine, appCtx component.AppContext) {
	userStore := userstorage.NewSQLStore(appCtx.GetMainDBConnection())

	v1 := router.Group("/v1")
	{
		v1.GET("/admin",
			middleware.RequiredAuth(appCtx, userStore),
			middleware.RequiredRoles(appCtx, "admin", "mod"),
			func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{"data": 1})
			})

		v1.POST("/register", usergin.Register(appCtx))
		v1.POST("/auth", usergin.Login(appCtx))
		v1.GET("/profile", middleware.RequiredAuth(appCtx, userStore), usergin.Profile(appCtx))

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
