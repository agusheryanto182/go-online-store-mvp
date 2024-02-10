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
	GetCartItemByProductID(cartId, productId int) (*entities.CartItem, error)
	UpdateCartItem(cartItem *entities.CartItem) error
	UpdateGrandTotal(cartID, grandTotal int) error
	GetCartItemsByCartID(cartId int) ([]*entities.CartItem, error)
}

type CartServiceInterface interface {
	AddCartItems(userId int, request *dto.AddCartItemsRequest) (*entities.CartItem, error)
	RecalculateGrandTotal(cart *entities.Cart) error
}

type CartHandlerInterface interface {
	AddCartItem(c *fiber.Ctx) error
}
