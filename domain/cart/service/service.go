package service

import (
	"errors"
	"strconv"

	"github.com/agusheryanto182/go-online-store-mvp/domain/cart"
	"github.com/agusheryanto182/go-online-store-mvp/domain/cart/dto"
	"github.com/agusheryanto182/go-online-store-mvp/domain/product"
	"github.com/agusheryanto182/go-online-store-mvp/entities"
)

type CartServiceImpl struct {
	cartRepository cart.CartRepositoryInterface
	productService product.ProductServiceInterface
}

func NewCartService(cartRepository cart.CartRepositoryInterface, productService product.ProductServiceInterface) cart.CartServiceInterface {
	return &CartServiceImpl{cartRepository: cartRepository, productService: productService}
}

func (s *CartServiceImpl) AddCartItems(userId int, request *dto.AddCartItemsRequest) (*entities.CartItem, error) {
	carts, err := s.cartRepository.GetCartByUserId(userId)
	if err != nil {
		if carts == nil {
			newCart := &entities.Cart{
				UserId: userId,
			}
			_, err := s.cartRepository.CreateCart(newCart)
			if err != nil {
				return nil, errors.New("failed create new cart")
			}
			carts = newCart
		}
	}

	existingCartItem, err := s.cartRepository.GetCartItemByProductID(carts.Id, request.ProductID)
	if err == nil && existingCartItem != nil {
		existingCartItem.Quantity += request.Quantity
		existingCartItem.TotalPrice = existingCartItem.Quantity * existingCartItem.Price

		err := s.cartRepository.UpdateCartItem(existingCartItem)
		if err != nil {
			return nil, errors.New("failed to change the number of products in the cart")
		}
		err = s.RecalculateGrandTotal(carts)
		if err != nil {
			return nil, errors.New("failed to recalculate the grand total")
		}
		return existingCartItem, nil
	}

	getProductByID, err := s.productService.GetProductByID(request.ProductID)
	if err != nil {
		return nil, errors.New("product not found")
	}

	cartItem := &entities.CartItem{
		CartId:     carts.Id,
		ProductId:  request.ProductID,
		Quantity:   request.Quantity,
		Price:      getProductByID.Price,
		TotalPrice: getProductByID.Price * request.Quantity,
	}
	result, err := s.cartRepository.CreateCartItem(cartItem)
	if err != nil {
		return nil, errors.New("failed to add product to cart")
	}
	err = s.RecalculateGrandTotal(carts)
	if err != nil {
		return nil, errors.New("failed to recalculate the grand total")
	}
	return result, nil

}

func (s *CartServiceImpl) RecalculateGrandTotal(cart *entities.Cart) error {
	cartItems, err := s.cartRepository.GetCartItemsByCartID(cart.Id)
	if err != nil {
		return err
	}
	var grandTotal int
	for _, item := range cartItems {
		grandTotal += item.TotalPrice
	}

	cart.GrandTotal = grandTotal

	err = s.cartRepository.UpdateGrandTotal(cart.Id, grandTotal)
	if err != nil {
		return err
	}
	return nil
}

func (s *CartServiceImpl) GetCart(userID int) (*entities.Cart, error) {
	carts, err := s.cartRepository.FindCart(userID)
	if err != nil {
		return &entities.Cart{}, errors.New("cart not found")
	}
	return carts, nil
}

func (s *CartServiceImpl) RemoveProductFromCart(userID, productID int) error {
	isProductInCart := s.cartRepository.IsProductInCart(userID, productID)
	if !isProductInCart {
		return errors.New("the product with ID " + strconv.Itoa(userID) + " is not on user")
	}

	getCart, _ := s.cartRepository.GetCartByUserId(userID)

	getCartItem, _ := s.cartRepository.GetCartItemByProductID(userID, productID)

	newGrandTotal := getCart.GrandTotal - getCartItem.TotalPrice

	if err := s.cartRepository.UpdateGrandTotal(getCart.Id, newGrandTotal); err != nil {
		return errors.New("failed to update grand total")
	}

	err := s.cartRepository.RemoveProductFromCart(userID, productID)
	if err != nil {
		return errors.New("failed to remove product")
	}

	return nil
}

func (s *CartServiceImpl) IsProductInCart(userID, productID int) bool {
	isInCart := s.cartRepository.IsProductInCart(userID, productID)
	return isInCart
}
