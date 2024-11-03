package transport

import (
	"net/http"
	"social-todo-list/common"
	"social-todo-list/modules/item/buisness"
	"social-todo-list/modules/item/buisness/impl"
	"social-todo-list/modules/item/model"
	"social-todo-list/modules/item/storage"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateItem(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		data := model.TodoItemCreation{}

		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
		store := storage.NewSQLStorage(db)
		var buisnessImpl buisness.ItemInterface = impl.NewItemBuisness(store)
		if err := buisnessImpl.CreateItem(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, common.SimepleSuccessResponse(data.Id))
	}
}
func GetItem(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
			return
		}
		store := storage.NewSQLStorage(db)
		var buisnessImpl buisness.ItemInterface = impl.NewItemBuisness(store)
		item, err := buisnessImpl.GetItem(c.Request.Context(), id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, common.SimepleSuccessResponse(item))
	}
}
func ListItem(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var paging common.Paging
		if err := c.ShouldBindQuery(&paging); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
		var filter model.Filter
		if err := c.ShouldBindQuery(&filter); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
		}

		paging.Offset()

		store := storage.NewSQLStorage(db)
		var buisnessImpl buisness.ItemInterface = impl.NewItemBuisness(store)
		data, err := buisnessImpl.ListItem(c.Request.Context(), &filter, &paging)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, common.NewSuccessResponse(data, &paging, &filter))	
	}
}
func UpdateItem(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
			return
		}
		data := model.ToDoItemUpdate{}
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
		store := storage.NewSQLStorage(db)
		var buisnessImpl buisness.ItemInterface = impl.NewItemBuisness(store)
		if err := buisnessImpl.UpdateItem(c.Request.Context(), id, &data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, common.SimepleSuccessResponse("Updated"))
	}
}
func DeleteItem(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
			return
		}
		store := storage.NewSQLStorage(db)
		var buisnessImpl buisness.ItemInterface = impl.NewItemBuisness(store)
		if err := buisnessImpl.DeleteItem(c.Request.Context(), id); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, common.SimepleSuccessResponse("Deleted"))
	}
}
