package api

import (
	"strconv"

	"github.com/XieChengG/bookapi/config"
	"github.com/XieChengG/bookapi/controller"
	"github.com/XieChengG/bookapi/model"
	"github.com/XieChengG/bookapi/response"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BookApiHandler struct {
	db  *gorm.DB
	svc *controller.BookController
}

func NewBookApiHander() *BookApiHandler {
	return &BookApiHandler{
		db:  config.GetConfig().MySQL.DB(),
		svc: controller.NewBookController(),
	}
}

// 提供一个注册的功能
func (h *BookApiHandler) Registry(r *gin.Engine) {
	book := r.Group("/api/books")
	book.POST("", h.CreateBook)
	book.GET("", h.GetBookList)
	book.GET("/:isbn", h.GetBook)
	book.PUT("/:isbn", h.UpdateBook)
	book.DELETE("/:isbn", h.DeleteBook)
}

// 创建书籍
func (h *BookApiHandler) CreateBook(ctx *gin.Context) {
	ins := &model.BookSpec{}
	if err := ctx.ShouldBindJSON(ins); err != nil {
		response.Failed(ctx, err)
		return
	}

	book, err := h.svc.CreateBook(ctx, ins)
	if err != nil {
		response.Failed(ctx, err)
	}

	response.Success(ctx, book)

}

// 获取书籍列表
func (h *BookApiHandler) GetBookList(ctx *gin.Context) {
	var ins []*model.Book
	books, err := h.svc.GetBookList(ctx, ins)
	if err != nil {
		response.Failed(ctx, err)
	}
	response.Success(ctx, books)
}

// 获取单个书籍
func (h *BookApiHandler) GetBook(ctx *gin.Context) {
	idStr := ctx.Param("isbn")
	id, _ := strconv.ParseInt(idStr, 10, 64)

	book, err := h.svc.GetBook(ctx, id)
	if err != nil {
		response.Failed(ctx, err)
	}
	response.Success(ctx, book)
}

// 更新书籍
func (h *BookApiHandler) UpdateBook(ctx *gin.Context) {
	ins := model.BookSpec{}
	idStr := ctx.Param("isbn")

	id, _ := strconv.ParseInt(idStr, 10, 64)

	if err := ctx.ShouldBindJSON(&ins); err != nil {
		response.Failed(ctx, err)
	}

	err := h.svc.UpdateBook(ctx, id, &ins)
	if err != nil {
		response.Failed(ctx, err)
	}
}

// 删除书籍
func (h *BookApiHandler) DeleteBook(ctx *gin.Context) {
	var ins model.Book
	id := ctx.Param("isbn")

	err := h.svc.DeleteBook(ctx, id, &ins)
	if err != nil {
		response.Failed(ctx, err)
	}
}
