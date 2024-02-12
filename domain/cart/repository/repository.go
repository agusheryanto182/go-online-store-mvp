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

func (r *CartRepositoryImpl) GetCartItemByProductID(cartID, productId int) (*entities.CartItem, error) {
	carts := entities.CartItem{}
	if err := r.db.Where("cart_id = ? AND product_id = ?", cartID, productId).First(&carts).Error; err != nil {
		return nil, err
	}
	return &carts, nil
}

func (r *CartRepositoryImpl) UpdateCartItem(cartItem *entities.CartItem) error {
	if err := r.db.Save(cartItem).Error; err != nil {
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

func (r *CartRepositoryImpl) FindCart(userID int) (*entities.Cart, error) {
	carts := &entities.Cart{}
	if err := r.db.
		Preload("CartItems", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, cart_id, product_id, quantity, total_price").
				Preload("Product", func(db *gorm.DB) *gorm.DB {
					return db.Select("id, name, price")
				})
		}).Where("user_id = ?", userID).First(&carts).Error; err != nil {
		return nil, err
	}
	return carts, nil
}

func (r *CartRepositoryImpl) RemoveProductFromCart(userID, productID int) error {
	var carts entities.Cart
	if err := r.db.Where("user_id = ?", userID).Preload("CartItems").First(&carts).Error; err != nil {
		return err
	}

	var cartItem entities.CartItem
	for _, item := range carts.CartItems {
		if item.ProductId == productID {
			cartItem = *item
			break
		}
	}
	if cartItem.Id == 0 {
		return nil
	}

	if err := r.db.Delete(&cartItem).Error; err != nil {
		return err
	}

	return nil
}

func (r *CartRepositoryImpl) RemoveCart(cart *entities.Cart) error {
	if err := r.db.Delete(&cart).Error; err != nil {
		return err
	}

	return nil
}

func (r *CartRepositoryImpl) IsProductInCart(userID, productID int) bool {
	var count int64
	r.db.Model(&entities.CartItem{}).
		Joins("JOIN carts ON cart_items.cart_id = carts.id").
		Where("carts.user_id = ? AND cart_items.product_id = ?", userID, productID).
		Count(&count)
	return count > 0
}
