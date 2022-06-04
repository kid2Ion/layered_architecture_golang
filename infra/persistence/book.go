package persistence

import (
	"context"

	"github.com/hiroki-kondo-git/chitchat/domain/model"
	"github.com/hiroki-kondo-git/chitchat/domain/repository"
)

type bookPersistence struct{}

func NewBookPersistence() repository.BookRepository {
	return &bookPersistence{}
}

func (bp *bookPersistence) GetAll(context.Context) ([]*model.Book, error) {
	book1 := model.Book{
		Id:     1,
		Title:  "layeredArchitecture",
		Author: "hiro",
	}
	book2 := model.Book{
		Id:     2,
		Title:  "ddd",
		Author: "hiro",
	}
	return []*model.Book{&book1, &book2}, nil
}
