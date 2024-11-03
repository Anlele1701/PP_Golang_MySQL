package main

import (
	"log"
	"os"
	"social-todo-list/modules/item/transport"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	godotenv.Load(".env")
	dsn := os.Getenv("DB_CONNECTION_STRING")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("missing MySQL connection string.")
	}

	r := gin.Default()
	v1 := r.Group("/v1")
	{
		items := v1.Group("/items")
		{
			items.POST("", transport.CreateItem(db))
			items.GET("", transport.ListItem(db))
			items.GET("/:id", transport.GetItem(db))
			items.PATCH("/:id", transport.UpdateItem(db))
			items.DELETE("/:id", transport.DeleteItem(db))
		}
	}

	r.Run(":3000")
}
