package controllers

import (
	"final_project_4/helpers"
	"final_project_4/models"
	"final_project_4/service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductController struct { // implementasi Controller
	csa service.ProductServiceApi
}

func NewProductController(csa service.ProductServiceApi) *ProductController {
	return &ProductController{csa: csa}
}

func (cc *ProductController) CreateProductControllers(c *gin.Context) {
	var input models.Product

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
	// category, err := cc.csa.ProductGetCategoryByIdService(input)

	product, category, err := cc.csa.CreateProductService(input)

	if product.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Your title is required",
		})
	} else if product.Price == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Your price is required",
		})
	} else if product.Price > 50000000 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Price Tidak Boleh Melebihi angka 50.000.000",
		})
	} else if product.Stock == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Stock Tidak Boleh Kurang Dari angka 5",
		})
	} else if product.Stock < 5 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Stock Tidak Boleh Kurang Dari angka 5",
		})
	} else if product.CategoryID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Your category_id is required",
		})
	} else if category.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "category_id yang ada pilih tidak tersedia",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"Id":          product.ID,
			"Title":       product.Title,
			"price":       product.Price,
			"stock":       product.Stock,
			"category_id": product.CategoryID,
			"created_at":  product.CreatedAt,
		})
	}
}

func (cc *ProductController) GetAllProductControllers(c *gin.Context) {
	Product := []models.Product{}
	var err error = nil

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	product, err := cc.csa.GetAllProductService(Product)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"result": product,
			"count":  len(product),
		})
	}
}

func (cc *ProductController) UpdateProductControllers(c *gin.Context) {
	var input models.Product

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
	productId, _ := strconv.Atoi(c.Param("productId"))

	product, category, err := cc.csa.UpdateProductService(productId, input)

	if product.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Your title is required",
		})
	} else if product.Price == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Your price is required",
		})
	} else if product.Price > 50000000 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Price Tidak Boleh Melebihi angka 50.000.000",
		})
	} else if product.Stock == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Stock Tidak Boleh Kurang Dari angka 5",
		})
	} else if product.Stock < 5 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Stock Tidak Boleh Kurang Dari angka 5",
		})
	} else if product.CategoryID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Your category_id is required",
		})
	} else if category.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "category_id yang ada pilih tidak tersedia",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"RESULT":      "Your Product Has Been Updated",
			"Id":          product.ID,
			"Title":       product.Title,
			"price":       product.Price,
			"stock":       product.Stock,
			"category_id": product.CategoryID,
			"created_at":  product.CreatedAt,
		})
	}
}

func (cc *ProductController) DeleteProductControllers(c *gin.Context) {
	var err error = nil

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}
	productId, _ := strconv.Atoi(c.Param("productId"))

	err = cc.csa.DeleteProductService(productId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Gagal menghapus data",
			"message": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"Success": "Your Product has been successfully deleted",
		})
	}
}
