package repositories

import (
	"final_project_4/models"
	"fmt"

	"github.com/jinzhu/gorm"
)

type TransactionRepo struct {
	db *gorm.DB
}

func NewTransactionRepo(db *gorm.DB) TransactionRepo {
	return TransactionRepo{
		db: db,
	}
}

type TransactionRepoApi interface {
	GetUserById(userId int) (models.User, error)
	GetProductById(productId int) (models.Product, error)
	GetCategoryById(categoryId int) (models.Category, error)
	CreateTransaction(Transaction models.TransactionHistory) (models.TransactionHistory, error)
	UpdateProduct(productId int, Product models.Product) (models.Product, error)
	UpdateBalanceUser(userId int, User models.User) (models.User, error)
	UpdateSold_Product_Amount(categoryId int, Category models.Category) (models.Category, error)
	GetMyTransaction(userId int, Transaction []models.TransactionHistory) ([]models.TransactionHistory, error)
	GetAllUserTransaction(Transaction []models.TransactionHistory) ([]models.TransactionHistory, error)
}

func (tr *TransactionRepo) GetUserById(userId int) (models.User, error) {
	var user models.User
	err := tr.db.Where("id = ?", userId).First(&user).Error
	return user, err
}

func (tr *TransactionRepo) GetProductById(productId int) (models.Product, error) {
	var product models.Product
	err := tr.db.Where("id = ?", productId).First(&product).Error
	return product, err
}
func (tr *TransactionRepo) GetCategoryById(categoryId int) (models.Category, error) {
	var category models.Category
	err := tr.db.Where("id = ?", categoryId).First(&category).Error
	return category, err
}

func (tr *TransactionRepo) CreateTransaction(Transaction models.TransactionHistory) (models.TransactionHistory, error) {
	err := tr.db.Debug().Create(&Transaction).Error
	if err != nil {
		fmt.Println(err.Error())
	}

	return Transaction, err
}
func (tr *TransactionRepo) UpdateProduct(productId int, Product models.Product) (models.Product, error) {
	err := tr.db.Model(&Product).Where("id = ?", productId).Updates(models.Product{
		Stock: Product.Stock,
	}).Error

	return Product, err
}

func (tr *TransactionRepo) UpdateBalanceUser(userId int, User models.User) (models.User, error) {
	err := tr.db.Model(&User).Where("id = ?", userId).Updates(models.User{
		Balance: User.Balance,
	}).Error

	return User, err
}

func (tr *TransactionRepo) UpdateSold_Product_Amount(categoryId int, Category models.Category) (models.Category, error) {
	err := tr.db.Model(&Category).Where("id = ?", categoryId).Updates(models.Category{
		Sold_product_amount: Category.Sold_product_amount,
	}).Error

	return Category, err
}
func (tr *TransactionRepo) GetMyTransaction(userId int, Transaction []models.TransactionHistory) ([]models.TransactionHistory, error) {
	err := tr.db.Where("user_id = ?", userId).Preload("Product").Find(&Transaction).Error
	if err != nil {
		fmt.Println(err.Error())
	}

	return Transaction, err
}

func (tr *TransactionRepo) GetAllUserTransaction(Transaction []models.TransactionHistory) ([]models.TransactionHistory, error) {
	err := tr.db.Preload("Product").Preload("User").Find(&Transaction).Error
	if err != nil {
		fmt.Println(err.Error())
	}

	return Transaction, err
}
