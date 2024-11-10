package abstract

import (
	"context"
	"gin-bookstore-api/models"
)

type Book interface {
	AddBook(ctx context.Context, bookParmas models.BookParams) (*models.BookParams, error)
	ListBooks(ctx context.Context) ([]models.Book, error)
	UpdateBook(ctx context.Context, UpdateBookParams models.UpdateBookParams, bookId int) (bool, error)
	DeleteBook(ctx context.Context, bookId int) error
}
