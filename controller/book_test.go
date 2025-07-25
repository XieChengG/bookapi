package controller_test

import (
	"context"
	"testing"

	"github.com/XieChengG/bookapi/controller"
	"github.com/XieChengG/bookapi/model"
)

func TestCreateBook(t *testing.T) {
	book := controller.NewBookController()
	isSaled := false
	ins, err := book.CreateBook(context.Background(), &model.BookSpec{
		Title:   "Golang Book",
		Author:  "lucy",
		Price:   15.50,
		IsSaled: &isSaled,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}

func TestGetBook(t *testing.T) {
	book := controller.NewBookController()
	ins, err := book.GetBook(context.Background(), &controller.GetBookRequest{Isbn: 6})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}
