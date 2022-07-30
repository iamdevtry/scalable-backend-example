package main

import (
	restaurantgin "food-delivery-service/module/restaurant/transport/gin"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

func main() {
	dsn := os.Getenv("DB_CONN_STR")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln("Cannot connect to MySQL:", err)
	}

	log.Println("Connected:", db)

	router := gin.Default()

	v1 := router.Group("/v1")
	{
		restaurants := v1.Group("/restaurants")
		{
			restaurants.POST("", restaurantgin.CreateRestaurantHandler(db))
			restaurants.GET("", restaurantgin.ListRestaurant(db))
			restaurants.GET("/:restaurant_id", restaurantgin.GetRestaurantHandler(db))
			restaurants.PUT("/:restaurant_id", restaurantgin.UpdateRestaurantHandler(db))
			restaurants.DELETE("/:restaurant_id", restaurantgin.DeleteRestaurantHandler(db))
		}
	}

	router.Run(":3000")
}
