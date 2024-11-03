package main

import (
	"log"
	"social-todo-list/modules/item/model"
	"social-todo-list/modules/item/transport"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(mysql.Open("root:root@tcp(mysql:3306)/todolist?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	db.AutoMigrate(&model.TodoItem{})
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

	r.Run(":8080")
}
