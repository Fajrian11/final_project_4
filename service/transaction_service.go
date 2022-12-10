package service

import (
	"final_project_4/models"
	"final_project_4/repositories"
)

type TransactionService struct {
	rr repositories.TransactionRepoApi
}

func NewTransactionService(rr repositories.TransactionRepoApi) *TransactionService { //provie service
	return &TransactionService{rr: rr}
}

type TransactionServiceApi interface {
	CreateTransactionService(userId int, input models.TransactionHistory) (models.Category, models.User, models.Product, models.User, models.Product, models.TransactionHistory, error)
	GetMyTransactionService(userId int, Transaction []models.TransactionHistory) ([]models.TransactionHistory, error)
	GetAllUserTransactionService(Transaction []models.TransactionHistory) ([]models.TransactionHistory, error)
}

func (ts TransactionService) CreateTransactionService(userId int, input models.TransactionHistory) (models.Category, models.User, models.Product, models.User, models.Product, models.TransactionHistory, error) {
	transaction := models.TransactionHistory{}
	editStock := models.Product{}
	editBalance := models.User{}
	editSPA := models.Category{}

	transaction.UserID = userId
	transaction.ProductID = input.ProductID
	transaction.Quantity = input.Quantity

	// GET PRODUCT
	Product, err := ts.rr.GetProductById(input.ProductID)
	User, err := ts.rr.GetUserById(userId)
	if Product.ID == 0 {
		return editSPA, editBalance, editStock, User, Product, transaction, err
	} else if input.Quantity > Product.Stock {
		return editSPA, editBalance, editStock, User, Product, transaction, err
	} else if User.Balance < Product.Price {
		return editSPA, editBalance, editStock, User, Product, transaction, err

	}
	transaction.TotalPrice = Product.Price * input.Quantity

	// GET CATEGORY
	Category, err := ts.rr.GetCategoryById(int(Product.CategoryID))
	if Category.ID == 0 {
		return editSPA, editBalance, editStock, User, Product, transaction, err
	}

	// EDIT STOCK
	editStock.Stock = Product.Stock - input.Quantity
	UpdateStock, err := ts.rr.UpdateProduct(input.ProductID, editStock)
	if editStock.Stock < 5 {
		return editSPA, editBalance, UpdateStock, User, Product, transaction, err
	}

	// EDIT BALANCE
	editBalance.Balance = User.Balance - Product.Price
	UpdateBalance, err := ts.rr.UpdateBalanceUser(userId, editBalance)
	if editBalance.Balance < 5 {
		return editSPA, UpdateBalance, UpdateStock, User, Product, transaction, err
	}

	// EDIT SOLD_PRODUCT_AMOUNT
	editSPA.Sold_product_amount = Category.Sold_product_amount + input.Quantity
	UpdateSPA, err := ts.rr.UpdateSold_Product_Amount(int(Product.CategoryID), editSPA)
	if err != nil {
		return UpdateSPA, UpdateBalance, UpdateStock, User, Product, transaction, err
	}

	// CATAT TRANSAKSI
	Transaction, err := ts.rr.CreateTransaction(transaction)
	if err != nil {
		return UpdateSPA, UpdateBalance, UpdateStock, User, Product, Transaction, err
	}

	return UpdateSPA, UpdateBalance, UpdateStock, User, Product, Transaction, nil
}

func (ts TransactionService) GetMyTransactionService(userId int, Transaction []models.TransactionHistory) ([]models.TransactionHistory, error) {
	Transaction, err := ts.rr.GetMyTransaction(userId, Transaction)
	if err != nil {
		return Transaction, err
	}

	return Transaction, nil
}

func (ts TransactionService) GetAllUserTransactionService(Transaction []models.TransactionHistory) ([]models.TransactionHistory, error) {
	Transaction, err := ts.rr.GetAllUserTransaction(Transaction)
	if err != nil {
		return Transaction, err
	}

	return Transaction, nil
}
