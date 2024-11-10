package database

import (
	"context"
	"gin-bookstore-api/models"

	"gorm.io/gorm"
)

func (c Client) AddBook(ctx context.Context, bookParams models.BookParams) *models.Book {
	var Book models.Book

	Book.Title = bookParams.Title
	Book.Author = bookParams.Author
	Book.Read = bookParams.Read

	result := c.db.WithContext(ctx).Create(&Book)

	if result.Error != nil {
		return nil
	}

	return &Book
}

func (c Client) ListBooks(ctx context.Context) ([]models.Book, error) {

	var books []models.Book

	result := c.db.WithContext(ctx).Find(&books)

	if result.Error != nil {
		return nil, result.Error
	}

	return books, nil
}

func (c Client) UpdateBook(ctx context.Context, UpdateBookParams models.UpdateBookParams, bookId int) (bool, error) {

	c.db.Save(&models.Book{Model: gorm.Model{ID: uint(bookId)}, Title: UpdateBookParams.Title, Author: UpdateBookParams.Author, Read: UpdateBookParams.Read})

	return true, nil
}

func (c Client) DeleteBook(ctx context.Context, bookId int) error {

	result := c.db.WithContext(ctx).Delete(bookId)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
