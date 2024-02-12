package cart

import (
	"github.com/agusheryanto182/go-online-store-mvp/domain/cart/dto"
	"github.com/agusheryanto182/go-online-store-mvp/entities"
	"github.com/gofiber/fiber/v2"
)

type CartRepositoryInterface interface {
	CreateCart(newCart *entities.Cart) (*entities.Cart, error)
	CreateCartItem(cartItem *entities.CartItem) (*entities.CartItem, error)
	GetCartByUserId(userId int) (*entities.Cart, error)
	GetCartItemByProductID(cartID, productId int) (*entities.CartItem, error)
	UpdateCartItem(cartItem *entities.CartItem) error
	UpdateGrandTotal(cartID, grandTotal int) error
	GetCartItemsByCartID(cartId int) ([]*entities.CartItem, error)
	FindCart(userID int) (*entities.Cart, error)
	RemoveCart(cart *entities.Cart) error
	RemoveProductFromCart(userID, productID int) error
	IsProductInCart(userID, productID int) bool
}

type CartServiceInterface interface {
	AddCartItems(userId int, request *dto.AddCartItemsRequest) (*entities.CartItem, error)
	RecalculateGrandTotal(cart *entities.Cart) error
	GetCart(userID int) (*entities.Cart, error)
	RemoveProductFromCart(userID, productID int) error
	IsProductInCart(userID, productID int) bool
}

type CartHandlerInterface interface {
	AddCartItem(c *fiber.Ctx) error
	GetCart(c *fiber.Ctx) error
	DeleteCartItem(c *fiber.Ctx) error
}
