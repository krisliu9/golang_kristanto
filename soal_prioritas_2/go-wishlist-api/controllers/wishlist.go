package controllers

import (
	"go-wishlist-api/repository"
	"go-wishlist-api/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type WishlistController struct {
	GetAllWishlistsUseCase *usecase.WishlistUseCase
}

func NewWishlistController(getAllWishlists *usecase.WishlistUseCase) *WishlistController {
	return &WishlistController{
		GetAllWishlistsUseCase: getAllWishlists,
	}
}

func (controller *WishlistController) GetAll(c echo.Context) error {
	wishlists, err := controller.GetAllWishlistsUseCase.GetAll()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get wishlists")
	}

	return c.JSON(http.StatusOK, echo.Map{
		"data": wishlists,
	})
}

func (controller *WishlistController) Create(c echo.Context) error {
	var input repository.Wishlist

	c.Bind(&input)

	controller.GetAllWishlistsUseCase.Create(input)

	return c.JSON(http.StatusCreated, echo.Map{
		"data": input,
	})
}
