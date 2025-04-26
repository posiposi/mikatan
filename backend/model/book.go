package model

import (
	"fmt"
	"time"

	"github.com/shopspring/decimal"
)

type Book struct {
	ID                 string          `json:"id" gorm:"primary_key"`
	Title              string          `json:"title" gorm:"not null"`
	Genre              string          `json:"genre"`
	TotalPage          int             `json:"totalPage"`
	ProgressPage       int             `json:"progressPage"`
	ProgressPercentage int             `json:"progressPercentage"`
	Price              decimal.Decimal `json:"price"`
	Author             string          `json:"author"`
	Publisher          string          `json:"publisher"`
	PublishedAt        int             `json:"publishedAt"`
	CreatedAt          time.Time       `json:"createdAt"`
	UpdatedAt          time.Time       `json:"updatedAt"`
}

func NewBook(
	id, title, genre, author, publisher string,
	totalPage, progressPage, publishedAt int,
	price decimal.Decimal,
) (_ *Book, err error) {
	if title == "" {
		return nil, fmt.Errorf("title is required")
	}
	return &Book{
		ID:                 id,
		Title:              title,
		Genre:              genre,
		TotalPage:          totalPage,
		ProgressPage:       progressPage,
		ProgressPercentage: calcProgressPercentage(totalPage, progressPage),
		Price:              price,
		Author:             author,
		Publisher:          publisher,
		PublishedAt:        publishedAt,
		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
	}, nil
}

// TODO レスポンスパッケージに移管しても良いかも
type BookResponse struct {
	ID                 string          `json:"id" gorm:"primary_key"`
	Title              string          `json:"title" gorm:"not null"`
	Genre              string          `json:"genre"`
	TotalPage          int             `json:"total_page"`
	ProgressPage       int             `json:"progressPage"`
	ProgressPercentage int             `json:"progressPercentage"`
	Price              decimal.Decimal `json:"price"`
	Author             string          `json:"author"`
	Publisher          string          `json:"publisher"`
	PublishedAt        int             `json:"publishedAt"`
	CreatedAt          time.Time       `json:"createdAt"`
	UpdatedAt          time.Time       `json:"updatedAt"`
}

func calcProgressPercentage(totalPage, progressPage int) int {
	if totalPage == 0 || progressPage == 0 {
		return 0
	}
	percetange := float64(progressPage) / float64(totalPage) * 100
	return int(percetange)
}
