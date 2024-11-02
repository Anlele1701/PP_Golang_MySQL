package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type TodoItem struct {
	Id          int        `json:"id" gorm:"column:id;"`
	Titles      string     `json:"titles" gorm:"column:titles;"`
	Description string     `json:"description" gorm:"column:description;"`
	Status      string     `json:"status" gorm:"column:status;"`
	CreatedAt   *time.Time `json:"createdAt" gorm:"column:createdAt;"`
	UpdatedAt   *time.Time `json:"updatedAt,omitempty" gorm:"column:updatedAt;"`
}
func (TodoItem) TableName() string { return "tasks" }

type TodoItemCreation struct {
	Id          int    `json:"-" gorm:"column:id;"`
	Title       string `json:"title" gorm:"column:titles;"`
	Description string `json:"description" gorm:"column:description;"`
}
func (TodoItemCreation) TableName() string { return TodoItem{}.TableName() }

type ToDoItemUpdate struct {
	Title       *string `json:"title" gorm:"column:titles;"`
	Description *string `json:"description" gorm:"column:description;"`
	Status      *string `json:"status" gorm:"column:status;"`
}
func (ToDoItemUpdate) TableName() string { return TodoItem{}.TableName() }

type Paging struct {
	Page int `json:"page" form:"page"` 
	Size int `json:"size" form:"size"`
	Total int64 `json:"total" form:"-"`
}
func (p *Paging) Offset(){
	if p.Page <= 0{
		p.Page = 1
	}
	if p.Size <=0 || p.Size >= 100{
		p.Size = 10
	}
}
func main() {
	godotenv.Load(".env")
	dsn := os.Getenv("DB_CONNECTION_STRING")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	r := gin.Default()
	v1 := r.Group("/v1")
	{
		items := v1.Group("/items")
		{
			items.POST("", CreateItem(db))
			items.GET("", ListItem(db))
			items.GET("/:id", GetItem(db))
			items.PATCH("/:id", UpdateItem(db))
			items.DELETE("/:id", DeleteItem(db))
		}
	}

	r.Run(":3000")
}
func CreateItem(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		data := TodoItemCreation{}

		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
		if err := db.Create(&data).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": data.Id,
		})
	}
}

func GetItem(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		data := TodoItem{}
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		if err := db.Where("id = ?", id).First(&data).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	}
}
func ListItem(db *gorm.DB) func(*gin.Context) {
    return func(c *gin.Context) {
		var paging Paging
        var result []TodoItem
		if err := c.ShouldBindQuery(&paging); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
		paging.Offset()
		db = db.Where("status <> ?", "Deleted")
		if err:= db.Table(TodoItem{}.TableName()).Count(&paging.Total).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
        if err := db.Order("id asc").
		Offset((paging.Page -1)*paging.Size).
		Limit(paging.Size).Find(&result).Error; err != nil {
            c.JSON(http.StatusBadRequest, gin.H{
                "message": err.Error(),
            })
			return
        } else {
            c.JSON(http.StatusOK, gin.H{
				"paging": paging,
                "data": result,
            })
        }
    }
}
func UpdateItem(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		data := ToDoItemUpdate{}
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
		if err := db.Where("id = ?", id).Updates(&data).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": true,
		})
	}
}
func DeleteItem(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
		if err := db.Table(TodoItem{}.TableName()).Where("id=?", id).Update("status", "Deleted").Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"data": true,
		})
	}
}
