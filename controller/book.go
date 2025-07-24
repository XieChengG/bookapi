package controller

import (
	"context"

	"github.com/XieChengG/bookapi/config"
	"github.com/XieChengG/bookapi/model"
	"gorm.io/gorm"
)

type BookController struct {
	db *gorm.DB
}

func NewBookController() *BookController {
	return &BookController{
		db: config.GetConfig().MySQL.DB(),
	}
}

// 创建书籍
func (c *BookController) CreateBook(ctx context.Context, b *model.BookSpec) (*model.Book, error) {
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
func (c *BookController) GetBook(ctx context.Context, id int64) (*model.Book, error) {
	ins := &model.Book{}
	if err := c.db.Where("isbn = ?", id).Take(ins).Error; err != nil {
		return nil, err
	}
	return ins, nil
}

// 更新书籍
func (c *BookController) UpdateBook(ctx context.Context, id int64, b *model.BookSpec) error {
	err := c.db.Where("isbn = ?", id).Model(&model.Book{}).Updates(b).Error
	if err != nil {
		return err
	}
	return nil
}

// 删除书籍
func (c *BookController) DeleteBook(ctx context.Context, id string, b *model.Book) error {
	err := c.db.Where("isbn = ?", id).Delete(b).Error
	if err != nil {
		return err
	}
	return nil
}
