package service

import (
	"final_project_4/models"
	"final_project_4/repositories"
)

type ProductService struct {
	rr repositories.ProductRepoApi
}

func NewProductService(rr repositories.ProductRepoApi) *ProductService { //provie service
	return &ProductService{rr: rr}
}

type ProductServiceApi interface {
	// ProductGetCategoryByIdService(input models.Product) (models.Category, error)
	CreateProductService(input models.Product) (models.Product, models.Category, error)
	GetAllProductService(Product []models.Product) ([]models.Product, error)
	UpdateProductService(productId int, input models.Product) (models.Product, models.Category, error)
	DeleteProductService(productId int) error
}

// func (ps ProductService) ProductGetCategoryByIdService(input models.Product) (models.Category, error) {
// 	Category, _ := ps.rr.ProductGetCategoryById(int(input.CategoryID))
// 	if Category.ID == 0 {
// 		return Category, errors.New("Category Yang Anda Pilih Tidak Tersedia")
// 	}
// 	return Category, nil
// }

func (ps ProductService) CreateProductService(input models.Product) (models.Product, models.Category, error) {

	product := models.Product{}
	product.Title = input.Title
	product.Price = input.Price
	product.Stock = input.Stock
	product.CategoryID = input.CategoryID

	Category, err := ps.rr.ProductGetCategoryById(int(input.CategoryID))
	if Category.ID == 0 {
		return product, Category, err //harus beda return product
	}

	Product, err := ps.rr.CreateProduct(product)
	if err != nil {
		return Product, Category, err
	}

	return Product, Category, nil
}

func (ps ProductService) GetAllProductService(Product []models.Product) ([]models.Product, error) {
	Product, err := ps.rr.GetAllProduct(Product)
	if err != nil {
		return Product, err
	}
	return Product, nil

}

func (ps ProductService) UpdateProductService(productId int, input models.Product) (models.Product, models.Category, error) {
	Category, err := ps.rr.ProductGetCategoryById(int(input.CategoryID))
	if Category.ID == 0 {
		return input, Category, err //harus beda return product
	}

	Product, err := ps.rr.UpdateProduct(productId, input)
	if err != nil {
		return Product, Category, err
	}

	return Product, Category, nil
}

func (ps ProductService) DeleteProductService(productId int) error {
	err := ps.rr.DeleteProduct(productId)
	if err != nil {
		return err
	}
	return err
}
