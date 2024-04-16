package usecase

import "go-wishlist-api/repository"

type WishlistUseCase struct {
	WishlistRepo repository.WishlistRepository
}

func NewWishlistUseCase(repo repository.WishlistRepository) *WishlistUseCase {
	return &WishlistUseCase{WishlistRepo: repo}
}

func (usecase *WishlistUseCase) GetAll() ([]repository.Wishlist, error) {
	return usecase.WishlistRepo.GetAll()
}

func (usecase *WishlistUseCase) Create(wishlist repository.Wishlist) (repository.Wishlist, error) {
	return usecase.WishlistRepo.Create(wishlist)
}
