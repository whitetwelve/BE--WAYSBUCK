package repositories

import (
	"waysbuck/models"

	"gorm.io/gorm"
)

type CartRepository interface {
	FindCarts() ([]models.Cart, error)
	GetCart(ID int) (models.Cart, error)
	CreateCart(cart models.Cart) (models.Cart, error)
	UpdateCart(cart models.Cart) (models.Cart, error)
	DeleteCart(cart models.Cart, ID int) (models.Cart, error)
	FindCartTopings(ID []int) ([]models.Toping, error)
	GetUserCart(ID int) ([]models.Cart, error)
	// FindCartsTransaction(TransactionID int) ([]models.Cart, error)
	// CreateDataTransaction(transaction models.Transaction) (models.Transaction, error)
}

func RepositoryCart(db *gorm.DB) *repository {
	return &repository{db}
}

// GET ALL CARTS
func (r *repository) FindCarts() ([]models.Cart, error) {
	var carts []models.Cart
	err := r.db.Preload("Topping").Preload("Product").Find(&carts).Error

	return carts, err
}

// GET BY ID
func (r *repository) GetCart(ID int) (models.Cart, error) {
	var cart models.Cart
	err := r.db.Preload("Topping").Preload("Product").First(&cart, ID).Error

	return cart, err
}

// CREATE CART
func (r *repository) CreateCart(cart models.Cart) (models.Cart, error) {
	err := r.db.Preload("Product").Preload("Topping").Create(&cart).Error

	return cart, err
}

// UPDATE CART
func (r *repository) UpdateCart(cart models.Cart) (models.Cart, error) {
	err := r.db.Save(&cart).Error

	return cart, err
}

// DELETE
func (r *repository) DeleteCart(cart models.Cart, ID int) (models.Cart, error) {
	err := r.db.Delete(&cart).Error

	return cart, err
}

// CART TOPING
func (r *repository) FindCartTopings(ID []int) ([]models.Toping, error) {
	var topings []models.Toping
	err := r.db.Find(&topings, ID).Error

	return topings, err
}

// GET USER CART
func (r *repository) GetUserCart(UserID int) ([]models.Cart, error) {
	var user []models.Cart
	err := r.db.Preload("Product").Preload("Topping").Find(&user, "user_id  = ?", UserID).Error

	return user, err
}

// FIND CARTS TRANSACTION
// func (r *repository) FindCartsTransaction(idTrans int) ([]models.Cart, error) {
// 	var carts []models.Cart
// 	err := r.db.Preload("Product").Preload("Topping").Find(&carts, "transaction_id = ?", idTrans).Error

// 	return carts, err
// }

// // CREATE ID TRANS
// func (r *repository) CreateDataTransaction(transaction models.Transaction) (models.Transaction, error) {
// 	err := r.db.Create(&transaction).Error

// 	return transaction, err
// }
