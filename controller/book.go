package controller

import (
	"context"

	"github.com/XieChengG/bookapi/config"
	"github.com/XieChengG/bookapi/model"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

type BookController struct {
	db  *gorm.DB
	log *zerolog.Logger
}

type GetBookRequest struct {
	Isbn int64 `json:"isbn"`
}

func NewBookController() *BookController {
	return &BookController{
		db:  config.GetConfig().MySQL.DB(),
		log: config.GetConfig().Log.Logger(),
	}
}

// 创建书籍
func (c *BookController) CreateBook(ctx context.Context, b *model.BookSpec) (*model.Book, error) {

	c.log.Debug().Msgf("create book: %+v", b)

	ins := &model.Book{
		BookSpec: *b,
	}
	if err := c.db.Save(ins).Error; err != nil {
		return nil, err
	}
	return ins, nil
}

// 获取书籍列表
func (c *BookController) GetBookList(ctx context.Context, b []*model.Book) ([]*model.Book, error) {
	if err := c.db.Find(&b).Error; err != nil {
		return nil, err
	}
	return b, nil
}

// 获取单个书籍
func (c *BookController) GetBook(ctx context.Context, req *GetBookRequest) (*model.Book, error) {
	ins := &model.Book{}
	if err := c.db.WithContext(ctx).Where("isbn = ?", req.Isbn).Take(ins).Error; err != nil {
		return nil, err
	}
	return ins, nil
}

// 更新书籍
func (c *BookController) UpdateBook(ctx context.Context, req *GetBookRequest, b *model.BookSpec) error {
	err := c.db.Where("isbn = ?", req).Model(&model.Book{}).Updates(b).Error
	if err != nil {
		return err
	}
	return nil
}

// 删除书籍
func (c *BookController) DeleteBook(ctx context.Context, req *GetBookRequest, b *model.Book) error {
	err := c.db.Where("isbn = ?", req).Delete(b).Error
	if err != nil {
		return err
	}
	return nil
}
