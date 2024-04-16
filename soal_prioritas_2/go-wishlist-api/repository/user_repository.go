package repository

import (
	"gorm.io/gorm"
)

type UserRepository interface {
	Login() ([]Wishlist, error)
	Register(wishlist Wishlist) (Wishlist, error)
}

type SqliteWishlistRepository struct {
	DB gorm.DB
}

func NewSqliteWishlistRepository(db gorm.DB) *SqliteWishlistRepository {
	return &SqliteWishlistRepository{DB: db}
}

func (repo *SqliteWishlistRepository) GetAll() ([]Wishlist, error) {
	var wishlists []Wishlist
	repo.DB.Find(&wishlists)

	return wishlists, nil
}

func (repo *SqliteWishlistRepository) Create(wishlist Wishlist) (Wishlist, error) {
	repo.DB.Create(&wishlist)

	return wishlist, nil
}
