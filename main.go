package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/XieChengG/bookapi/config"
	"github.com/gin-gonic/gin"
)

// Book's table structure
type Book struct {
	IsBn uint `json:"isbn" gorm:"primaryKey;column:isbn"`
	BookSpec
}

type BookSpec struct {
	Title   string  `json:"title" gorm:"type:varchar(200);column:title"`
	Author  string  `json:"author" gorm:"type:varchar(200);column:author;index"`
	Price   float64 `json:"price" gorm:"column:price"`
	IsSaled *bool   `json:"is_saled" gorm:"column:is_saled"`
}

// error handler
func Failed(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": err})
}

func main() {
	// load config
	configFilePath := os.Getenv("CONFIG_FILE_PATH")
	if configFilePath == "" {
		configFilePath = "application.yaml"
	}
	if err := config.LoadConfigFromYaml(configFilePath); err != nil {
		fmt.Println("Failed to load config")
		os.Exit(1)
	}
	conf := config.GetConfig()

	// initial db
	db := conf.MySQL.DB()
	db.AutoMigrate(&Book{})

	// create a gin server
	server := gin.Default()

	// api group
	book := server.Group("/api/books")

	// create book
	book.POST("", func(ctx *gin.Context) {
		var book Book
		if err := ctx.ShouldBindJSON(&book); err != nil {
			Failed(ctx, err)
			return
		}

		if err := db.Save(&book).Error; err != nil {
			Failed(ctx, err)
			return
		}

		ctx.JSON(http.StatusOK, book)
	})

	// get all books
	book.GET("", func(ctx *gin.Context) {
		var books []*Book
		if err := db.Find(&books).Error; err != nil {
			Failed(ctx, err)
			return
		}
		ctx.JSON(http.StatusOK, books)
	})

	// get info of a book
	book.GET("/:isbn", func(ctx *gin.Context) {
		var book Book
		id := ctx.Param("isbn")
		if err := db.Where("isbn = ?", id).Find(&book).Error; err != nil {
			Failed(ctx, err)
			return
		}
		ctx.JSON(http.StatusOK, book)
	})

	// update a book
	book.PUT("/:isbn", func(ctx *gin.Context) {
		spec := BookSpec{}
		id := ctx.Param("isbn")

		if err := ctx.ShouldBindJSON(&spec); err != nil {
			Failed(ctx, err)
			return
		}

		if err := db.Where("isbn = ?", id).Model(&Book{}).Updates(spec).Error; err != nil {
			Failed(ctx, err)
			return
		}
	})

	// delete a book
	book.DELETE("/:isbn", func(ctx *gin.Context) {
		var book Book
		id := ctx.Param("isbn")
		if err := db.Where("isbn = ?", id).Delete(&book).Error; err != nil {
			Failed(ctx, err)
			return
		}
	})

	// run server
	if err := server.Run(conf.App.Address()); err != nil {
		panic(err)
	}
}
