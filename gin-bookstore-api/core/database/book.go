package database

import (
	"context"
	"errors"
	"gin-bookstore-api/models"

	"gorm.io/gorm"
)

func (c Client) AddBook(ctx context.Context, bookParams models.BookParams) (*models.Book, error) {
	var Book models.Book

	Book.Title = bookParams.Title
	Book.Author = bookParams.Author
	Book.Read = bookParams.Read

	result := c.db.WithContext(ctx).Create(&Book)

	if result.Error != nil {
		return nil, result.Error
	}

	return &Book, nil
}

func (c Client) GetBook(ctx context.Context, bookId int64) (*models.Book, error) {
	var book models.Book

	result := c.db.First(&book, bookId)

	if result.Error != nil {
		return nil, result.Error
	}

	return &book, nil

}

func (c Client) ListBooks(ctx context.Context) ([]models.Book, error) {

	var books []models.Book

	result := c.db.WithContext(ctx).Find(&books)

	if result.Error != nil {
		return nil, result.Error
	}

	return books, nil
}

func (c Client) UpdateBook(ctx context.Context, UpdateBookParams models.UpdateBookParams, bookId int64) (bool, error) {

	c.db.Save(&models.Book{Model: gorm.Model{ID: uint(bookId)}, Title: UpdateBookParams.Title, Author: UpdateBookParams.Author, Read: UpdateBookParams.Read})

	return true, nil
}

func (c Client) DeleteBook(ctx context.Context, bookId int64) error {

	var bookInfo models.Book
	bookInfo.ID = uint(bookId) // if Book allows direct assignment like this

	if err := c.db.First(&bookInfo).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("there is no book associated with this ID")
		}
	}
	// Delete Book (Hard delete)
	c.db.Unscoped().Delete(&bookInfo)

	return nil
}
