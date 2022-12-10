package controllers

import (
	"final_project_4/helpers"
	"final_project_4/models"
	"final_project_4/service"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type TransactionController struct { // implementasi Controller
	csa service.TransactionServiceApi
}

func NewTransactionController(csa service.TransactionServiceApi) *TransactionController {
	return &TransactionController{csa: csa}
}

func (cc *TransactionController) CreateTransactionControllers(c *gin.Context) {
	var input models.TransactionHistory

	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := int(userData["id"].(float64))

	ContentType := helpers.GetContentType(c)
	var err error = nil

	if ContentType == appJSON {
		err = c.ShouldBindJSON(&input)
	} else {
		err = c.ShouldBind(&input)
	}

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"result": "Failed Create User",
		})
		return
	}

	_, _, updateStock, user, product, transaction, err := cc.csa.CreateTransactionService(userID, input)

	if transaction.ProductID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Your product_id is required",
		})
	} else if product.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Product Yang anda pilih tidak ditemukan",
		})
	} else if user.Balance < product.Price {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error":               "Saldo ATM Anda Tidak Cukup Untuk Membeli Barang ini",
			"Sisa Saldo ATM anda": user.Balance,
		})
	} else if transaction.Quantity == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Your quantity is required",
		})
	} else if transaction.Quantity > product.Stock {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":                  "quantity yang dipesan melebihi stock yang tersedia",
			"quantity yang tersedia": product.Stock,
		})
	} else if updateStock.Stock < 5 {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error":                   "quantity yang anda pesan mencapai batas minimum stock",
			"Stock Sekarang":          product.Stock,
			"minimum stock":           5,
			"stock yang bisa dipesan": product.Stock - 5,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message":       "You Have Successfully purchased the product",
			"total_price":   transaction.TotalPrice,
			"quantity":      transaction.Quantity,
			"product_title": product.Title,
			"category_id":   product.CategoryID,
		})
	}
}

func (tc *TransactionController) GetMyTransactionController(c *gin.Context) {
	Transaction := []models.TransactionHistory{}
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := int(userData["id"].(float64))

	var err error = nil

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	transaction, err := tc.csa.GetMyTransactionService(userID, Transaction)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"count":  len(transaction),
			"result": transaction,
		})
	}
}

func (tc *TransactionController) GetAllUserTransactionController(c *gin.Context) {
	Transaction := []models.TransactionHistory{}

	var err error = nil

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	transaction, err := tc.csa.GetAllUserTransactionService(Transaction)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"count":  len(transaction),
			"result": transaction,
		})
	}
}
