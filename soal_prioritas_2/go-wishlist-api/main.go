package main

import (
	"go-wishlist-api/controllers"
	"go-wishlist-api/database"
	"go-wishlist-api/repository"
	"go-wishlist-api/usecase"

	"github.com/labstack/echo/v4"
)

func main() {
	database.InitDatabase()
	wishlistRepo := repository.NewSqliteWishlistRepository(*database.DB)

	wishlistUseCase := usecase.NewWishlistUseCase(wishlistRepo)
	wishlistController := controllers.NewWishlistController(wishlistUseCase)

	e := echo.New()

	e.GET("/wishlists", wishlistController.GetAll)
	e.POST("/wishlists", wishlistController.Create)

	e.Logger.Fatal(e.Start(":1323"))
}
