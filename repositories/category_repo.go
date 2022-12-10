package repositories

import (
	"final_project_4/models"
	"fmt"

	"github.com/jinzhu/gorm"
)

type CategoryRepo struct {
	db *gorm.DB
}

func NewCategoryRepo(db *gorm.DB) CategoryRepo {
	return CategoryRepo{
		db: db,
	}
}

type CategoryRepoApi interface {
	CreateCategory(Categories models.Category) (models.Category, error)
	GetAllCategory(Category []models.Category) ([]models.Category, error)
	GetCategoryById(categoryId int) (models.Category, error)
	UpdateCategory(categoryId int, Category models.Category) (models.Category, error)
	DeleteCategory(categoryId int) error
}

func (cr *CategoryRepo) CreateCategory(Category models.Category) (models.Category, error) {
	err := cr.db.Debug().Create(&Category).Error
	if err != nil {
		fmt.Println(err.Error())
	}

	return Category, err
}

func (cr *CategoryRepo) GetAllCategory(Category []models.Category) ([]models.Category, error) {
	err := cr.db.Model(&models.Category{}).Preload("Product").Find(&Category).Error
	fmt.Println(err)
	return Category, err
}

func (cr *CategoryRepo) GetCategoryById(categoryId int) (models.Category, error) {
	var category models.Category

	err := cr.db.Where("id = ?", categoryId).First(&category).Error

	fmt.Println(err)
	return category, err
}

func (cr *CategoryRepo) UpdateCategory(categoryId int, Category models.Category) (models.Category, error) {
	err := cr.db.Model(&Category).Where("id = ?", categoryId).Updates(models.Category{
		Type: Category.Type,
	}).Error

	return Category, err
}

func (cr *CategoryRepo) DeleteCategory(categoryId int) error {

	err := cr.db.Exec(`
	DELETE Categories
	FROM Categories
	WHERE Categories.id = ?`, categoryId).Error

	return err
}
