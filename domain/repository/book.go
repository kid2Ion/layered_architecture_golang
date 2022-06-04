package repository

import (
	"context"

	"github.com/hiroki-kondo-git/chitchat/domain/model"
)

// やることはCRUD　技術的関心はinfraなので、interface→全ての書籍を受けとるgetallを定義

type BookRepository interface {
	GetAll(context.Context) ([]*model.Book, error)
}

//このinterfaceは依存性を逆転させるためのもの（本来はここに具体的な処理を書く）
