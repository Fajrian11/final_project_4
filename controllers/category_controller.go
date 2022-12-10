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

type CategoryController struct { // implementasi Controller
	csa service.CategoryServiceApi
}

func NewCategoryController(csa service.CategoryServiceApi) *CategoryController {
	return &CategoryController{csa: csa}
}

func (cc *CategoryController) CreateCategoryControllers(c *gin.Context) {
	var input models.Category

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
			"result": "Failed Create Category",
		})
		return
	}
	Category, err := cc.csa.CreateCategoryService(input)
	if Category.Type == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Your type is required",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"Id":                  Category.ID,
			"Type":                Category.Type,
			"sold_product_amount": Category.Sold_product_amount,
			"created_at":          Category.CreatedAt,
		})
	}
}

func (cc *CategoryController) GetAllCategoryControllers(c *gin.Context) {
	var err error = nil

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}
	Category, err := cc.csa.GetAllCategoryService()

	c.JSON(http.StatusOK, gin.H{
		"result": Category,
		"count":  len(Category),
	})
}

func (cc *CategoryController) UpdateCategoryControllers(c *gin.Context) {
	var input models.Category

	categoryId, _ := strconv.Atoi(c.Param("categoryId"))

	contentType := helpers.GetContentType(c)
	var err error = nil
	if contentType == appJSON {
		err = c.ShouldBindJSON(&input)
	} else {
		err = c.ShouldBind(&input)
	}

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"result": err.Error(),
		})
		return
	}

	Category, err := cc.csa.UpdateCategoryService(categoryId, input)
	if Category.Type == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Your type is required",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"Result": "Your Category Has Been Updated",
			"Type":   Category.Type,
		})
	}
}

func (cc *CategoryController) DeleteCategoryControllers(c *gin.Context) {
	categoryId, _ := strconv.Atoi(c.Param("categoryId"))

	var err error = nil

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"result": err.Error(),
		})
		return
	}

	err = cc.csa.DeleteCategoryService(categoryId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"result": "Your Category Has Been Successfully Deleted",
		})
	}
}
