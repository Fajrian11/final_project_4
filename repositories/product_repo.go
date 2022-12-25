package repositories

import (
	"final_project_4/models"
	"fmt"

	"github.com/jinzhu/gorm"
)

type ProductRepo struct {
	db *gorm.DB
}

func NewProductRepo(db *gorm.DB) ProductRepo {
	return ProductRepo{
		db: db,
	}
}

type ProductRepoApi interface {
	ProductGetCategoryById(categoryId int) (models.Category, error)
	CreateProduct(Product models.Product) (models.Product, error)
	GetAllProduct(Product []models.Product) ([]models.Product, error)
	UpdateProduct(productId int, Product models.Product) (models.Product, error)
	DeleteProduct(productId int) error
}

func (pr *ProductRepo) ProductGetCategoryById(categoryId int) (models.Category, error) {
	var category models.Category

	err := pr.db.Where("id = ?", categoryId).First(&category).Error

	fmt.Println(err)
	fmt.Println(category)
	return category, err
}

func (pr *ProductRepo) CreateProduct(Product models.Product) (models.Product, error) {
	err := pr.db.Debug().Create(&Product).Error
	if err != nil {
		fmt.Println(err.Error())
	}

	return Product, err
}

func (pr *ProductRepo) GetAllProduct(Product []models.Product) ([]models.Product, error) {
	err := pr.db.Find(&Product).Error
	fmt.Println(err)

	return Product, err
}

func (pr *ProductRepo) UpdateProduct(productId int, Product models.Product) (models.Product, error) {
	err := pr.db.Model(&Product).Where("id = ?", productId).Updates(models.Product{
		Title:      Product.Title,
		Price:      Product.Price,
		Stock:      Product.Stock,
		CategoryID: Product.CategoryID,
	}).Error

	return Product, err
}

func (pr *ProductRepo) DeleteProduct(productId int) error {
	err := pr.db.Exec(`
	DELETE products
	FROM products
	WHERE products.id = ?`, productId).Error

	return err
}
