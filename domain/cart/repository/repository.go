package repository

import (
	"github.com/agusheryanto182/go-online-store-mvp/domain/cart"
	"github.com/agusheryanto182/go-online-store-mvp/entities"
	"gorm.io/gorm"
)

type CartRepositoryImpl struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) cart.CartRepositoryInterface {
	return &CartRepositoryImpl{db}
}

func (r *CartRepositoryImpl) CreateCart(newCart *entities.Cart) (*entities.Cart, error) {
	err := r.db.Create(newCart).Error
	if err != nil {
		return nil, err
	}
	return newCart, nil
}

func (r *CartRepositoryImpl) CreateCartItem(cartItem *entities.CartItem) (*entities.CartItem, error) {
	err := r.db.Create(cartItem).Error
	if err != nil {
		return nil, err
	}
	return cartItem, nil
}

func (r *CartRepositoryImpl) GetCartByUserId(userId int) (*entities.Cart, error) {
	carts := &entities.Cart{}
	if err := r.db.Where("user_id = ?", userId).First(carts).Error; err != nil {
		return nil, err
	}
	return carts, nil
}

func (r *CartRepositoryImpl) GetCartItemByProductID(cartId, productId int) (*entities.CartItem, error) {
	carts := entities.CartItem{}
	if err := r.db.Where("cart_id = ? AND product_id = ?", cartId, productId).First(&carts).Error; err != nil {
		return nil, err
	}
	return &carts, nil
}

func (r *CartRepositoryImpl) UpdateCartItem(cartItem *entities.CartItem) error {
	if err := r.db.Save(&cartItem).Error; err != nil {
		return err
	}
	return nil
}

func (r *CartRepositoryImpl) UpdateGrandTotal(cartID, grandTotal int) error {
	var carts *entities.Cart
	result := r.db.Model(&carts).Where("id = ?", cartID).Update("grand_total", grandTotal)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *CartRepositoryImpl) GetCartItemsByCartID(cartId int) ([]*entities.CartItem, error) {
	var cartItems []*entities.CartItem
	if err := r.db.Where("cart_id = ?", cartId).Find(&cartItems).Error; err != nil {
		return nil, err
	}
	return cartItems, nil
}
