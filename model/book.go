package model

import "encoding/json"

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

func (t *Book) TableName() string {
	return "books"
}

func (b *Book) String() string {
	v, _ := json.Marshal(b)
	return string(v)
}

func (s *BookSpec) String() string {
	v, _ := json.Marshal(s)
	return string(v)
}
